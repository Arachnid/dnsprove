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

// RootABI is the input ABI used to generate the binding from.
const RootABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"proveAndRegisterDefaultTLD\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"proveAndRegisterTLD\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"registerTLD\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RootBin is the compiled bytecode used for deploying new contracts.
const RootBin = `0x`

// DeployRoot deploys a new Ethereum contract, binding an instance of Root to it.
func DeployRoot(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Root, error) {
	parsed, err := abi.JSON(strings.NewReader(RootABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RootBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Root{RootCaller: RootCaller{contract: contract}, RootTransactor: RootTransactor{contract: contract}, RootFilterer: RootFilterer{contract: contract}}, nil
}

// Root is an auto generated Go binding around an Ethereum contract.
type Root struct {
	RootCaller     // Read-only binding to the contract
	RootTransactor // Write-only binding to the contract
	RootFilterer   // Log filterer for contract events
}

// RootCaller is an auto generated read-only Go binding around an Ethereum contract.
type RootCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RootTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RootFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RootSession struct {
	Contract     *Root             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RootCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RootCallerSession struct {
	Contract *RootCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RootTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RootTransactorSession struct {
	Contract     *RootTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RootRaw is an auto generated low-level Go binding around an Ethereum contract.
type RootRaw struct {
	Contract *Root // Generic contract binding to access the raw methods on
}

// RootCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RootCallerRaw struct {
	Contract *RootCaller // Generic read-only contract binding to access the raw methods on
}

// RootTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RootTransactorRaw struct {
	Contract *RootTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRoot creates a new instance of Root, bound to a specific deployed contract.
func NewRoot(address common.Address, backend bind.ContractBackend) (*Root, error) {
	contract, err := bindRoot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Root{RootCaller: RootCaller{contract: contract}, RootTransactor: RootTransactor{contract: contract}, RootFilterer: RootFilterer{contract: contract}}, nil
}

// NewRootCaller creates a new read-only instance of Root, bound to a specific deployed contract.
func NewRootCaller(address common.Address, caller bind.ContractCaller) (*RootCaller, error) {
	contract, err := bindRoot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RootCaller{contract: contract}, nil
}

// NewRootTransactor creates a new write-only instance of Root, bound to a specific deployed contract.
func NewRootTransactor(address common.Address, transactor bind.ContractTransactor) (*RootTransactor, error) {
	contract, err := bindRoot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RootTransactor{contract: contract}, nil
}

// NewRootFilterer creates a new log filterer instance of Root, bound to a specific deployed contract.
func NewRootFilterer(address common.Address, filterer bind.ContractFilterer) (*RootFilterer, error) {
	contract, err := bindRoot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RootFilterer{contract: contract}, nil
}

// bindRoot binds a generic wrapper to an already deployed contract.
func bindRoot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RootABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Root *RootRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Root.Contract.RootCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Root *RootRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Root.Contract.RootTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Root *RootRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Root.Contract.RootTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Root *RootCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Root.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Root *RootTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Root.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Root *RootTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Root.Contract.contract.Transact(opts, method, params...)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() constant returns(address)
func (_Root *RootCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Root.contract.Call(opts, out, "oracle")
	return *ret0, err
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() constant returns(address)
func (_Root *RootSession) Oracle() (common.Address, error) {
	return _Root.Contract.Oracle(&_Root.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() constant returns(address)
func (_Root *RootCallerSession) Oracle() (common.Address, error) {
	return _Root.Contract.Oracle(&_Root.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) constant returns(bool)
func (_Root *RootCaller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Root.contract.Call(opts, out, "supportsInterface", interfaceID)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) constant returns(bool)
func (_Root *RootSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _Root.Contract.SupportsInterface(&_Root.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) constant returns(bool)
func (_Root *RootCallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _Root.Contract.SupportsInterface(&_Root.CallOpts, interfaceID)
}

// ProveAndRegisterDefaultTLD is a paid mutator transaction binding the contract method 0x19f5b1e2.
//
// Solidity: function proveAndRegisterDefaultTLD(bytes name, bytes input, bytes proof) returns()
func (_Root *RootTransactor) ProveAndRegisterDefaultTLD(opts *bind.TransactOpts, name []byte, input []byte, proof []byte) (*types.Transaction, error) {
	return _Root.contract.Transact(opts, "proveAndRegisterDefaultTLD", name, input, proof)
}

// ProveAndRegisterDefaultTLD is a paid mutator transaction binding the contract method 0x19f5b1e2.
//
// Solidity: function proveAndRegisterDefaultTLD(bytes name, bytes input, bytes proof) returns()
func (_Root *RootSession) ProveAndRegisterDefaultTLD(name []byte, input []byte, proof []byte) (*types.Transaction, error) {
	return _Root.Contract.ProveAndRegisterDefaultTLD(&_Root.TransactOpts, name, input, proof)
}

// ProveAndRegisterDefaultTLD is a paid mutator transaction binding the contract method 0x19f5b1e2.
//
// Solidity: function proveAndRegisterDefaultTLD(bytes name, bytes input, bytes proof) returns()
func (_Root *RootTransactorSession) ProveAndRegisterDefaultTLD(name []byte, input []byte, proof []byte) (*types.Transaction, error) {
	return _Root.Contract.ProveAndRegisterDefaultTLD(&_Root.TransactOpts, name, input, proof)
}

// ProveAndRegisterTLD is a paid mutator transaction binding the contract method 0x245b79ad.
//
// Solidity: function proveAndRegisterTLD(bytes name, bytes input, bytes proof) returns()
func (_Root *RootTransactor) ProveAndRegisterTLD(opts *bind.TransactOpts, name []byte, input []byte, proof []byte) (*types.Transaction, error) {
	return _Root.contract.Transact(opts, "proveAndRegisterTLD", name, input, proof)
}

// ProveAndRegisterTLD is a paid mutator transaction binding the contract method 0x245b79ad.
//
// Solidity: function proveAndRegisterTLD(bytes name, bytes input, bytes proof) returns()
func (_Root *RootSession) ProveAndRegisterTLD(name []byte, input []byte, proof []byte) (*types.Transaction, error) {
	return _Root.Contract.ProveAndRegisterTLD(&_Root.TransactOpts, name, input, proof)
}

// ProveAndRegisterTLD is a paid mutator transaction binding the contract method 0x245b79ad.
//
// Solidity: function proveAndRegisterTLD(bytes name, bytes input, bytes proof) returns()
func (_Root *RootTransactorSession) ProveAndRegisterTLD(name []byte, input []byte, proof []byte) (*types.Transaction, error) {
	return _Root.Contract.ProveAndRegisterTLD(&_Root.TransactOpts, name, input, proof)
}

// RegisterTLD is a paid mutator transaction binding the contract method 0x87900f20.
//
// Solidity: function registerTLD(bytes name, bytes proof) returns()
func (_Root *RootTransactor) RegisterTLD(opts *bind.TransactOpts, name []byte, proof []byte) (*types.Transaction, error) {
	return _Root.contract.Transact(opts, "registerTLD", name, proof)
}

// RegisterTLD is a paid mutator transaction binding the contract method 0x87900f20.
//
// Solidity: function registerTLD(bytes name, bytes proof) returns()
func (_Root *RootSession) RegisterTLD(name []byte, proof []byte) (*types.Transaction, error) {
	return _Root.Contract.RegisterTLD(&_Root.TransactOpts, name, proof)
}

// RegisterTLD is a paid mutator transaction binding the contract method 0x87900f20.
//
// Solidity: function registerTLD(bytes name, bytes proof) returns()
func (_Root *RootTransactorSession) RegisterTLD(name []byte, proof []byte) (*types.Transaction, error) {
	return _Root.Contract.RegisterTLD(&_Root.TransactOpts, name, proof)
}
