// Copyright 2019 Nick Johnson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package registrar

//go:generate abigen --sol ../contracts/dnsregistrar.sol --pkg contracts --out ../contracts/dnsregistrar.go

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

var DNSSEC_CLAIM_INTERFACE_ID = [4]byte{0x1a, 0xa2, 0xe6, 0x41}
var InterfaceNotSupportedError = errors.New("Interface not supported")

type DNSRegistrar struct {
	r       *contracts.DNSRegistrar
	backend bind.ContractBackend
}

func New(addr common.Address, backend bind.ContractBackend) (*DNSRegistrar, error) {
	registrar, err := contracts.NewDNSRegistrar(addr, backend)
	if err != nil {
		return nil, err
	}

	// Check it really implements the interface we need.
	if ok, err := registrar.SupportsInterface(nil, DNSSEC_CLAIM_INTERFACE_ID); !ok || err != nil {
		return nil, InterfaceNotSupportedError
	}

	return &DNSRegistrar{
		registrar,
		backend,
	}, nil
}

func (r *DNSRegistrar) GetOracle() (*oracle.Oracle, error) {
	addr, err := r.r.Oracle(nil)
	if err != nil {
		return nil, err
	}

	return oracle.New(addr, r.backend)
}

func (r *DNSRegistrar) Claim(opts *bind.TransactOpts, name string, sets []proofs.SignedSet) (*types.Transaction, error) {
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

		log.Info("Transaction to claim()", "name", name, "proof", hexutil.Encode(proof))
		return r.r.Claim(opts, dnsname, proof)
	} else {
		known, err := o.FindFirstUnknownProof(sets)
		if err != nil {
			return nil, err
		}

		data, proof, err := o.SerializeProofs(sets, known)
		if err != nil {
			return nil, err
		}

		log.Info("Transaction to proveAndClaim()", "name", name, "data", hexutil.Encode(data), "lastProof", hexutil.Encode(proof))
		return r.r.ProveAndClaim(opts, dnsname, data, proof)
	}
}

func (r *DNSRegistrar) Unclaim(opts *bind.TransactOpts, name string, sets []proofs.SignedSet) ([]*types.Transaction, error) {
	var txs []*types.Transaction

	o, err := r.GetOracle()
	if err != nil {
		return nil, err
	}

	nsec, sets := sets[len(sets)-1], sets[:len(sets)-1]

	known, err := o.FindFirstUnknownProof(sets)
	if err != nil {
		return nil, err
	}

	// We're deleting a domain. If it's not there, there's nothing to do.
	_, _, hash, err := o.Rrdata(dns.TypeTXT, "_ens."+name)
	if err != nil {
		return nil, err
	}
	if hash != [20]byte{} {
		var proof []byte
		// Update proofs so the NSEC can be verified.
		if known < len(sets) {
			log.Info("Sending transaction to update proofs", "name", "_ens."+name, "count", len(sets)-known)
			tx, err := o.SendProofs(opts, sets, known)
			if err != nil {
				return nil, err
			}
			txs = append(txs, tx)
			opts.Nonce = opts.Nonce.Add(opts.Nonce, big.NewInt(1))
		}

		// Use the NSEC's signing record as proof of its validity
		proof, err = sets[len(sets)-2].PackRRSet()
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

	dnsname, err := oracle.PackName(name)
	if err != nil {
		return nil, err
	}

	// Unclaim the name
	log.Info("Sending transaction to unclaim name", "name", name)
	tx, err := r.r.Claim(opts, dnsname, []byte{})
	if err != nil {
		return txs, err
	}
	txs = append(txs, tx)

	return txs, nil
}
