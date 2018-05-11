// Copyright 2017 Nick Johnson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package oracle

//go:generate abigen --sol contract/dnssec.sol --pkg contract --out contract/dnssec.go

import (
  "bytes"
  "context"
  "fmt"
  "math/big"

  "github.com/miekg/dns"
  "github.com/ethereum/go-ethereum/accounts/abi/bind"
  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/common/hexutil"
  "github.com/ethereum/go-ethereum/core/types"
  "github.com/ethereum/go-ethereum/crypto/sha3"
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

func (o *Oracle) FindFirstUnknownProof(p []proofs.SignedSet) (int, error) {
  for i, set := range p {
    header := set.Rrs[0].Header()

    name, err := packName(header.Name)
    if err != nil {
      return -1, err
    }

    result, err := o.Rrdata(nil, header.Rrtype, name)
    if err != nil {
      return -1, err
    }

    rrset, err := set.PackRRSet()
    if err != nil {
        return -1, err
    }
    h := sha3.NewKeccak256()
    h.Write(rrset)
    hash := h.Sum(nil)

    if result.Inception == 0 {
      log.Info("RRSET does not exist", "name", header.Name, "type", dns.TypeToString[header.Rrtype])
      return i, nil
    } else if result.Inception <= set.Sig.Inception && !bytes.Equal(result.Hash[:], hash[:20]) {
      log.Info("RRSET exists but is out of date", "name", header.Name, "type", dns.TypeToString[header.Rrtype], "current", result.Inception, "new", set.Sig.Inception, "oldhash", hexutil.Encode(result.Hash[:]), "newhash", hexutil.Encode(hash[:20]))
      return i, nil
    } else if result.Inception > set.Sig.Inception {
      return -1, fmt.Errorf("Oracle's RRSET has inception after our record's inception", "name", name, "oracleInception", result.Inception, "inception", set.Sig.Inception)
    } else {
      log.Info("RRSET already exists", "name", header.Name, "type", dns.TypeToString[header.Rrtype])
    }
  }

  return len(p), nil
}

func (o *Oracle) SendProofs(opts *bind.TransactOpts, p []proofs.SignedSet, known int) ([]*types.Transaction, error) {
  ret := make([]*types.Transaction, 0, known)

  startnonce, err := o.backend.PendingNonceAt(context.TODO(), opts.From)
  if err != nil {
    return nil, err
  }

  // Get the trust anchors as initial proof
  proof, err := o.Anchors(nil)
  if err != nil {
      return nil, err
  }

  for i, set := range p {
    if i >= known {
        header := set.Rrs[0].Header()

        name, err := packName(header.Name)
        if err != nil {
          return nil, err
        }

        data, err := set.Pack()
        if err != nil {
          return nil, err
        }

        sig, err := set.PackSignature()
        if err != nil {
            return nil, err
        }

        log.Info("Submitting transaction", "name", header.Name, "type", dns.TypeToString[header.Rrtype])
        opts.Nonce = big.NewInt(int64(startnonce) + int64(i))
        tx, err := o.SubmitRRSet(opts, name, data, sig, proof)
        if err != nil {
          return nil, err
        }
        ret = append(ret, tx)
    }

    proof, err = set.PackRRSet()
    if err != nil {
        return nil, err
    }
  }
  return ret, nil
}
