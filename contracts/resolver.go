// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ResolverABI is the input ABI used to generate the binding from.
const ResolverABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"addr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ResolverBin is the compiled bytecode used for deploying new contracts.
const ResolverBin = `0x`

// DeployResolver deploys a new Ethereum contract, binding an instance of Resolver to it.
func DeployResolver(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Resolver, error) {
	parsed, err := abi.JSON(strings.NewReader(ResolverABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ResolverBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Resolver{ResolverCaller: ResolverCaller{contract: contract}, ResolverTransactor: ResolverTransactor{contract: contract}, ResolverFilterer: ResolverFilterer{contract: contract}}, nil
}

// Resolver is an auto generated Go binding around an Ethereum contract.
type Resolver struct {
	ResolverCaller     // Read-only binding to the contract
	ResolverTransactor // Write-only binding to the contract
	ResolverFilterer   // Log filterer for contract events
}

// ResolverCaller is an auto generated read-only Go binding around an Ethereum contract.
type ResolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ResolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ResolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ResolverSession struct {
	Contract     *Resolver         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ResolverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ResolverCallerSession struct {
	Contract *ResolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ResolverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ResolverTransactorSession struct {
	Contract     *ResolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ResolverRaw is an auto generated low-level Go binding around an Ethereum contract.
type ResolverRaw struct {
	Contract *Resolver // Generic contract binding to access the raw methods on
}

// ResolverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ResolverCallerRaw struct {
	Contract *ResolverCaller // Generic read-only contract binding to access the raw methods on
}

// ResolverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ResolverTransactorRaw struct {
	Contract *ResolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewResolver creates a new instance of Resolver, bound to a specific deployed contract.
func NewResolver(address common.Address, backend bind.ContractBackend) (*Resolver, error) {
	contract, err := bindResolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Resolver{ResolverCaller: ResolverCaller{contract: contract}, ResolverTransactor: ResolverTransactor{contract: contract}, ResolverFilterer: ResolverFilterer{contract: contract}}, nil
}

// NewResolverCaller creates a new read-only instance of Resolver, bound to a specific deployed contract.
func NewResolverCaller(address common.Address, caller bind.ContractCaller) (*ResolverCaller, error) {
	contract, err := bindResolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ResolverCaller{contract: contract}, nil
}

// NewResolverTransactor creates a new write-only instance of Resolver, bound to a specific deployed contract.
func NewResolverTransactor(address common.Address, transactor bind.ContractTransactor) (*ResolverTransactor, error) {
	contract, err := bindResolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ResolverTransactor{contract: contract}, nil
}

// NewResolverFilterer creates a new log filterer instance of Resolver, bound to a specific deployed contract.
func NewResolverFilterer(address common.Address, filterer bind.ContractFilterer) (*ResolverFilterer, error) {
	contract, err := bindResolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ResolverFilterer{contract: contract}, nil
}

// bindResolver binds a generic wrapper to an already deployed contract.
func bindResolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ResolverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Resolver *ResolverRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Resolver.Contract.ResolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Resolver *ResolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Resolver.Contract.ResolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Resolver *ResolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Resolver.Contract.ResolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Resolver *ResolverCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Resolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Resolver *ResolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Resolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Resolver *ResolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Resolver.Contract.contract.Transact(opts, method, params...)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) constant returns(address)
func (_Resolver *ResolverCaller) Addr(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Resolver.contract.Call(opts, out, "addr", node)
	return *ret0, err
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) constant returns(address)
func (_Resolver *ResolverSession) Addr(node [32]byte) (common.Address, error) {
	return _Resolver.Contract.Addr(&_Resolver.CallOpts, node)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) constant returns(address)
func (_Resolver *ResolverCallerSession) Addr(node [32]byte) (common.Address, error) {
	return _Resolver.Contract.Addr(&_Resolver.CallOpts, node)
}
