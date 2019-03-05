// Copyright 2019 Nick Johnson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package registrar

//go:generate abigen --sol ../contracts/dnsregistrar.sol --pkg contracts --out ../contracts/dnsregistrar.go

import (
  "github.com/arachnid/dnsprove/contracts"
  "github.com/arachnid/dnsprove/oracle"
  "github.com/arachnid/dnsprove/proofs"
  "github.com/ethereum/go-ethereum/accounts/abi/bind"
  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/common/hexutil"
  "github.com/ethereum/go-ethereum/core/types"
  log "github.com/inconshreveable/log15"
)

type DNSRegistrar struct {
    r *contracts.DNSRegistrar
    backend bind.ContractBackend
}

func New(addr common.Address, backend bind.ContractBackend) (*DNSRegistrar, error) {
    registrar, err := contracts.NewDNSRegistrar(addr, backend)
    if err != nil {
        return nil, err
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
  matches, err := o.RecordMatches(sets[len(sets) - 1])
  if err != nil {
    return nil, err
  }

  if matches {
    proof, err := sets[len(sets) - 1].PackRRSet()
    if err != nil {
      return nil, err
    }

    log.Info("Transaction to claim()", "name", name, "proof", hexutil.Encode(proof))
    return r.r.Claim(opts, dnsname, proof)
  } else {
    known, err := o.FindFirstUnknownProof(sets, true)
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
  return nil, nil
}
