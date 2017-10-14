// Copyright 2017 Nick Johnson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package oracle

//go:generate abigen --sol contract/dnssec.sol --pkg contract --out contract/dnssec.go

import (
  "context"
  "math/big"

  "github.com/miekg/dns"
  "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/core/types"
  "github.com/arachnid/dnsprove/oracle/contract"
  log "github.com/inconshreveable/log15"
  "github.com/arachnid/dnsprove/proofs"
)

type Oracle struct {
  *contract.DNSSEC
  backend bind.ContractBackend
}

func NewOracle(addr common.Address, backend bind.ContractBackend) (*Oracle, error) {
    oracle, err := contract.NewDNSSEC(addr, backend)
    if err != nil {
      return nil, err
    }

    return &Oracle{
      oracle,
      backend,
    }, nil
}

func packName(name string) ([]byte, error) {
  ret := make([]byte, len(name) + 1)
  pos, err := dns.PackDomainName(name, ret, 0, nil, false)
  if err != nil {
    return nil, err
  }
  return ret[:pos], nil
}

func (o *Oracle) FilterProofs(p []proofs.SignedSet) ([]proofs.SignedSet, error) {
  ret := make([]proofs.SignedSet, 0, len(p))
  for _, proof := range p {
    header := proof.Rrs[0].Header()

    name, err := packName(header.Name)
    if err != nil {
      return nil, err
    }

    result, err := o.Rrset(nil, header.Class, header.Rrtype, name)
    if err != nil {
      return nil, err
    }

    if result.Inception == 0 {
      log.Info("RRSET does not exist", "name", header.Name, "class", dns.ClassToString[header.Class], "type", dns.TypeToString[header.Rrtype])
      ret = append(ret, proof)
    } else if result.Inception < proof.Sig.Inception {
      log.Info("RRSET exists but is out of date", "name", header.Name, "class", dns.ClassToString[header.Class], "type", dns.TypeToString[header.Rrtype], "current", result.Inception, "new", proof.Sig.Inception)
      ret = append(ret, proof)
    } else {
      log.Info("RRSET already exists", "name", header.Name, "class", dns.ClassToString[header.Class], "type", dns.TypeToString[header.Rrtype])
    }
  }

  return ret, nil
}

func (o *Oracle) SendProofs(opts *bind.TransactOpts, p []proofs.SignedSet) ([]*types.Transaction, error) {
  ret := make([]*types.Transaction, 0, len(p))

  startnonce, err := o.backend.PendingNonceAt(context.TODO(), opts.From)
  if err != nil {
    return nil, err
  }

  for i, proof := range p {
    header := proof.Rrs[0].Header()

    name, err := packName(header.Name)
    if err != nil {
      return nil, err
    }

    data, sig, err := proof.Pack()
    if err != nil {
      return nil, err
    }

    log.Info("Submitting transaction", "name", header.Name, "class", dns.ClassToString[header.Class], "type", dns.TypeToString[header.Rrtype])
    opts.Nonce = big.NewInt(int64(startnonce) + int64(i))
    tx, err := o.SubmitRRSet(opts, header.Class, name, data, sig)
    if err != nil {
      return nil, err
    }
    ret = append(ret, tx)
  }
  return ret, nil
}
