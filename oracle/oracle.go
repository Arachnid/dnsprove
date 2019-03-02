// Copyright 2017 Nick Johnson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package oracle

//go:generate abigen --sol contract/dnssec.sol --pkg contract --out contract/dnssec.go

import (
    "bytes"
    "fmt"
    "math/big"

    "github.com/miekg/dns"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/common/hexutil"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/arachnid/dnsprove/oracle/contract"
    log "github.com/inconshreveable/log15"
    "github.com/arachnid/dnsprove/proofs"
    "golang.org/x/crypto/sha3"
)

type Oracle struct {
    o *contract.DNSSEC
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

func (o *Oracle) FindFirstUnknownProof(p []proofs.SignedSet, found bool) (int, error) {
    for i, set := range p {
        if matches, err := o.RecordMatches(set); err != nil || !matches {
            return i, err
        }
    }
    return len(p), nil
}

func (o *Oracle) Rrdata(rrtype uint16, name string) (uint32, uint64, [20]byte, error) {
    packed, err := packName(name)
    if err != nil {
        return 0, 0, [20]byte{}, err
    }

    result, err := o.o.Rrdata(nil, rrtype, packed)
    return result.Inception, result.Inserted, result.Hash, err
}

func (o *Oracle) RecordMatches(set proofs.SignedSet) (bool, error) {
    header := set.Rrs[0].Header()

    inception, _, hash, err := o.Rrdata(header.Rrtype, header.Name)
    if err != nil {
        return false, err
    }

    rrset, err := set.PackRRSet()
    if err != nil {
        return false, err
    }
    h := sha3.NewLegacyKeccak256()
    h.Write(rrset)
    ourhash := h.Sum(nil)

    if inception == 0 {
        log.Info("RRSET does not exist", "name", header.Name, "type", dns.TypeToString[header.Rrtype])
        return false, nil
    } else if inception <= set.Sig.Inception && !bytes.Equal(hash[:], ourhash[:20]) {
        log.Info("RRSET exists but is out of date", "name", header.Name, "type", dns.TypeToString[header.Rrtype], "current", inception, "new", set.Sig.Inception, "oldhash", hexutil.Encode(hash[:]), "newhash", hexutil.Encode(ourhash[:20]))
        return false, nil
    } else if inception > set.Sig.Inception {
        return false, fmt.Errorf("Oracle's RRSET has inception after our record's inception: name=%s, type=%s, oracleInception=%d, inception=%d", header.Name, dns.TypeToString[header.Rrtype], inception, set.Sig.Inception)
    }

    log.Info("RRSET already exists", "name", header.Name, "type", dns.TypeToString[header.Rrtype])
    return true, nil
}

func (o *Oracle) SendProofs(opts *bind.TransactOpts, p []proofs.SignedSet, known int, found bool) ([]*types.Transaction, []byte, error) {
    ret := make([]*types.Transaction, 0, known)

    // Get the trust anchors as initial proof
    proof, err := o.o.Anchors(nil)
    if err != nil {
        return nil, nil, err
    }

    for i, set := range p {
        if i >= known {
            header := set.Rrs[0].Header()

            data, err := set.Pack()
            if err != nil {
                return nil, nil, err
            }

            sig, err := set.PackSignature()
            if err != nil {
                return nil, nil, err
            }

            log.Info("Submitting transaction", "name", header.Name, "type", dns.TypeToString[header.Rrtype])
            log.Debug("Signature info", "data", hexutil.Encode(data), "sig", hexutil.Encode(sig), "proof", hexutil.Encode(proof))
            tx, err := o.o.SubmitRRSet(opts, data, sig, proof)
            if err != nil {
                return nil, nil, err
            }
            ret = append(ret, tx)
            opts.Nonce = opts.Nonce.Add(opts.Nonce, big.NewInt(1))
        }

        proof, err = set.PackRRSet()
        if err != nil {
            return nil, nil, err
        }
    }


    return ret, proof, nil
}

func (o *Oracle) DeleteRRSet(opts *bind.TransactOpts, dnsType uint16, name string, nsec proofs.SignedSet, proof []byte) (*types.Transaction, error) {
    log.Info("Deleting RRSet", "type", dns.TypeToString[dnsType], "name", name, "nsec", nsec.Rrs)
    packedName, err := packName(name)
    if err != nil {
        return nil, err
    }

    data, err := nsec.Pack()
    if err != nil {
        return nil, err
    }

    sig, err := nsec.PackSignature()
    if err != nil {
        return nil, err
    }

    opts.GasLimit = 150000
    tx, err := o.o.DeleteRRSet(opts, dnsType, packedName, data, sig, proof)
    opts.Nonce = opts.Nonce.Add(opts.Nonce, big.NewInt(1))
    opts.GasLimit = 0

    return tx, err
}
