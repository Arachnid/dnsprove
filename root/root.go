// Copyright 2019 Nick Johnson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package root

//go:generate abigen --sol ../contracts/root.sol --pkg contracts --out ../contracts/root.go

import (
	"errors"
	"math/big"

	"github.com/arachnid/dnsprove/contracts"
	"github.com/arachnid/dnsprove/oracle"
	"github.com/arachnid/dnsprove/proofs"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/inconshreveable/log15"
	"github.com/miekg/dns"
)

var DNSSEC_ROOT_CLAIM_INTERFACE_ID = [4]byte{0xc7, 0xfe, 0x16, 0xbf}
var InterfaceNotSupportedError = errors.New("Interface not supported")

type Root struct {
	r       *contracts.Root
	backend bind.ContractBackend
}

func New(addr common.Address, backend bind.ContractBackend) (*Root, error) {
	root, err := contracts.NewRoot(addr, backend)
	if err != nil {
		return nil, err
	}

	// Check it really implements the interface we need.
	if ok, err := root.SupportsInterface(nil, DNSSEC_ROOT_CLAIM_INTERFACE_ID); !ok || err != nil {
		return nil, InterfaceNotSupportedError
	}

	return &Root{
		root,
		backend,
	}, nil
}

func (r *Root) GetOracle() (*oracle.Oracle, error) {
	addr, err := r.r.Oracle(nil)
	if err != nil {
		return nil, err
	}

	return oracle.New(addr, r.backend)
}

func (r *Root) Claim(opts *bind.TransactOpts, name string, sets []proofs.SignedSet) (*types.Transaction, error) {
	dnsname, err := oracle.PackName(name)
	if err != nil {
		return nil, err
	}

	o, err := r.GetOracle()
	if err != nil {
		return nil, err
	}

	// If the RRset already matches, we just need to claim, not prove.
	matches, err := o.RecordMatches(sets[len(sets)-1])
	if err != nil {
		return nil, err
	}

	if matches {
		proof, err := sets[len(sets)-1].PackRRSet()
		if err != nil {
			return nil, err
		}

		log.Info("Transaction to registerTLD()", "name", name, "proof", hexutil.Encode(proof))
		return r.r.RegisterTLD(opts, dnsname, proof)
	} else {
		known, err := o.FindFirstUnknownProof(sets)
		if err != nil {
			return nil, err
		}

		data, proof, err := o.SerializeProofs(sets, known)
		if err != nil {
			return nil, err
		}

		log.Info("Transaction to proveAndRegisterTLD()", "name", name, "data", hexutil.Encode(data), "lastProof", hexutil.Encode(proof))
		return r.r.ProveAndRegisterTLD(opts, dnsname, data, proof)
	}
}

func (r *Root) ClaimDefault(opts *bind.TransactOpts, name string, nsecsets, dssets []proofs.SignedSet) ([]*types.Transaction, error) {
	var txs []*types.Transaction

	o, err := r.GetOracle()
	if err != nil {
		return nil, err
	}

	// Check if the TXT record is there and so requires deleting.
	_, _, hash, err := o.Rrdata(dns.TypeTXT, "_ens.nic."+name)
	if err != nil {
		return nil, err
	}
	if hash != [20]byte{} {
    nsec, nsecsets := nsecsets[len(nsecsets)-1], nsecsets[:len(nsecsets)-1]

  	known, err := o.FindFirstUnknownProof(nsecsets)
  	if err != nil {
  		return nil, err
  	}

		var proof []byte
		// Update proofs so the NSEC can be verified.
		if known < len(nsecsets) {
			log.Info("Sending transaction to update proofs", "name", "_ens."+name, "count", len(nsecsets)-known)
			tx, err := o.SendProofs(opts, nsecsets, known)
			if err != nil {
				return nil, err
			}
			txs = append(txs, tx)
			opts.Nonce = opts.Nonce.Add(opts.Nonce, big.NewInt(1))
		}

		// Use the NSEC's signing record as proof of its validity
		proof, err = nsecsets[len(nsecsets)-2].PackRRSet()
		if err != nil {
			return txs, err
		}

		// Send the NSEC to delete the TXT records
		log.Info("Sending transaction to delete RRSet", "type", "TXT", "name", "_ens."+name)
		deletetx, err := o.DeleteRRSet(opts, dns.TypeTXT, "_ens."+name, nsec, proof)
		if err != nil {
			return txs, err
		}
		opts.Nonce = opts.Nonce.Add(opts.Nonce, big.NewInt(1))
		txs = append(txs, deletetx)
	}

  known, err := o.FindFirstUnknownProof(dssets)
  if err != nil {
    return nil, err
  }

  dnsname, err := oracle.PackName(name)
	if err != nil {
		return nil, err
	}

  if known < len(dssets) {
    data, proof, err := o.SerializeProofs(dssets, known)
		if err != nil {
			return nil, err
		}

		log.Info("Transaction to proveAndRegisterDefaultTLD()", "name", name, "data", hexutil.Encode(data), "lastProof", hexutil.Encode(proof))
		tx, err := r.r.ProveAndRegisterDefaultTLD(opts, dnsname, data, proof)
    if err != nil {
      return nil, err
    }
    txs = append(txs, tx)
  } else {
  	log.Info("Sending transaction to set name to default registrar", "name", name)
  	tx, err := r.r.RegisterTLD(opts, dnsname, []byte{})
  	if err != nil {
  		return txs, err
  	}
  	txs = append(txs, tx)
  }

	return txs, nil
}
