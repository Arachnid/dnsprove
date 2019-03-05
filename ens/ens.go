// Copyright 2019 Nick Johnson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package ens

//go:generate abigen --sol ../contracts/ens.sol --pkg contracts --out ../contracts/ens.go
//go:generate abigen --sol ../contracts/resolver.sol --pkg contracts --out ../contracts/resolver.go

import (
	"strings"

	"github.com/arachnid/dnsprove/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"golang.org/x/crypto/sha3"
)

type Resolver struct {
	resolver *contracts.Resolver
	backend  bind.ContractBackend
	node     common.Hash
}

func (r *Resolver) Addr() (common.Address, error) {
	return r.resolver.Addr(nil, r.node)
}

type ENS struct {
	ens     *contracts.ENS
	backend bind.ContractBackend
}

func Namehash(name string) common.Hash {
	if name == "" {
		return common.Hash{}
	}

	parts := strings.SplitN(name, ".", 2)

	h := sha3.NewLegacyKeccak256()
	h.Write([]byte(parts[0]))
	label := h.Sum(nil)

	parent := common.Hash{}
	if len(parts) > 1 {
		parent = Namehash(parts[1])
	}

	h = sha3.NewLegacyKeccak256()
	h.Write(parent[:])
	h.Write(label[:])
	return common.BytesToHash(h.Sum(nil))
}

func New(addr common.Address, backend bind.ContractBackend) (*ENS, error) {
	contract, err := contracts.NewENS(addr, backend)
	if err != nil {
		return nil, err
	}

	return &ENS{
		contract,
		backend,
	}, nil
}

func (e *ENS) Owner(name string) (common.Address, error) {
	h := Namehash(name)
	return e.ens.Owner(nil, h)
}

func (e *ENS) Resolver(name string) (*Resolver, error) {
	h := Namehash(name)
	addr, err := e.ens.Resolver(nil, h)
	if err != nil {
		return nil, err
	}

	contract, err := contracts.NewResolver(addr, e.backend)
	if err != nil {
		return nil, err
	}

	return &Resolver{
		contract,
		e.backend,
		h,
	}, nil
}
