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

// DNSRegistrarABI is the input ABI used to generate the binding from.
const DNSRegistrarABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"claim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"proveAndClaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DNSRegistrarBin is the compiled bytecode used for deploying new contracts.
const DNSRegistrarBin = `0x`

// DeployDNSRegistrar deploys a new Ethereum contract, binding an instance of DNSRegistrar to it.
func DeployDNSRegistrar(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DNSRegistrar, error) {
	parsed, err := abi.JSON(strings.NewReader(DNSRegistrarABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DNSRegistrarBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DNSRegistrar{DNSRegistrarCaller: DNSRegistrarCaller{contract: contract}, DNSRegistrarTransactor: DNSRegistrarTransactor{contract: contract}, DNSRegistrarFilterer: DNSRegistrarFilterer{contract: contract}}, nil
}

// DNSRegistrar is an auto generated Go binding around an Ethereum contract.
type DNSRegistrar struct {
	DNSRegistrarCaller     // Read-only binding to the contract
	DNSRegistrarTransactor // Write-only binding to the contract
	DNSRegistrarFilterer   // Log filterer for contract events
}

// DNSRegistrarCaller is an auto generated read-only Go binding around an Ethereum contract.
type DNSRegistrarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DNSRegistrarTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DNSRegistrarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DNSRegistrarFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DNSRegistrarFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DNSRegistrarSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DNSRegistrarSession struct {
	Contract     *DNSRegistrar     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DNSRegistrarCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DNSRegistrarCallerSession struct {
	Contract *DNSRegistrarCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// DNSRegistrarTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DNSRegistrarTransactorSession struct {
	Contract     *DNSRegistrarTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DNSRegistrarRaw is an auto generated low-level Go binding around an Ethereum contract.
type DNSRegistrarRaw struct {
	Contract *DNSRegistrar // Generic contract binding to access the raw methods on
}

// DNSRegistrarCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DNSRegistrarCallerRaw struct {
	Contract *DNSRegistrarCaller // Generic read-only contract binding to access the raw methods on
}

// DNSRegistrarTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DNSRegistrarTransactorRaw struct {
	Contract *DNSRegistrarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDNSRegistrar creates a new instance of DNSRegistrar, bound to a specific deployed contract.
func NewDNSRegistrar(address common.Address, backend bind.ContractBackend) (*DNSRegistrar, error) {
	contract, err := bindDNSRegistrar(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DNSRegistrar{DNSRegistrarCaller: DNSRegistrarCaller{contract: contract}, DNSRegistrarTransactor: DNSRegistrarTransactor{contract: contract}, DNSRegistrarFilterer: DNSRegistrarFilterer{contract: contract}}, nil
}

// NewDNSRegistrarCaller creates a new read-only instance of DNSRegistrar, bound to a specific deployed contract.
func NewDNSRegistrarCaller(address common.Address, caller bind.ContractCaller) (*DNSRegistrarCaller, error) {
	contract, err := bindDNSRegistrar(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DNSRegistrarCaller{contract: contract}, nil
}

// NewDNSRegistrarTransactor creates a new write-only instance of DNSRegistrar, bound to a specific deployed contract.
func NewDNSRegistrarTransactor(address common.Address, transactor bind.ContractTransactor) (*DNSRegistrarTransactor, error) {
	contract, err := bindDNSRegistrar(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DNSRegistrarTransactor{contract: contract}, nil
}

// NewDNSRegistrarFilterer creates a new log filterer instance of DNSRegistrar, bound to a specific deployed contract.
func NewDNSRegistrarFilterer(address common.Address, filterer bind.ContractFilterer) (*DNSRegistrarFilterer, error) {
	contract, err := bindDNSRegistrar(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DNSRegistrarFilterer{contract: contract}, nil
}

// bindDNSRegistrar binds a generic wrapper to an already deployed contract.
func bindDNSRegistrar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DNSRegistrarABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DNSRegistrar *DNSRegistrarRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DNSRegistrar.Contract.DNSRegistrarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DNSRegistrar *DNSRegistrarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DNSRegistrar.Contract.DNSRegistrarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DNSRegistrar *DNSRegistrarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DNSRegistrar.Contract.DNSRegistrarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DNSRegistrar *DNSRegistrarCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DNSRegistrar.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DNSRegistrar *DNSRegistrarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DNSRegistrar.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DNSRegistrar *DNSRegistrarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DNSRegistrar.Contract.contract.Transact(opts, method, params...)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() constant returns(address)
func (_DNSRegistrar *DNSRegistrarCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DNSRegistrar.contract.Call(opts, out, "oracle")
	return *ret0, err
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() constant returns(address)
func (_DNSRegistrar *DNSRegistrarSession) Oracle() (common.Address, error) {
	return _DNSRegistrar.Contract.Oracle(&_DNSRegistrar.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() constant returns(address)
func (_DNSRegistrar *DNSRegistrarCallerSession) Oracle() (common.Address, error) {
	return _DNSRegistrar.Contract.Oracle(&_DNSRegistrar.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) constant returns(bool)
func (_DNSRegistrar *DNSRegistrarCaller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DNSRegistrar.contract.Call(opts, out, "supportsInterface", interfaceID)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) constant returns(bool)
func (_DNSRegistrar *DNSRegistrarSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _DNSRegistrar.Contract.SupportsInterface(&_DNSRegistrar.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) constant returns(bool)
func (_DNSRegistrar *DNSRegistrarCallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _DNSRegistrar.Contract.SupportsInterface(&_DNSRegistrar.CallOpts, interfaceID)
}

// Claim is a paid mutator transaction binding the contract method 0xbe27b22c.
//
// Solidity: function claim(bytes name, bytes proof) returns()
func (_DNSRegistrar *DNSRegistrarTransactor) Claim(opts *bind.TransactOpts, name []byte, proof []byte) (*types.Transaction, error) {
	return _DNSRegistrar.contract.Transact(opts, "claim", name, proof)
}

// Claim is a paid mutator transaction binding the contract method 0xbe27b22c.
//
// Solidity: function claim(bytes name, bytes proof) returns()
func (_DNSRegistrar *DNSRegistrarSession) Claim(name []byte, proof []byte) (*types.Transaction, error) {
	return _DNSRegistrar.Contract.Claim(&_DNSRegistrar.TransactOpts, name, proof)
}

// Claim is a paid mutator transaction binding the contract method 0xbe27b22c.
//
// Solidity: function claim(bytes name, bytes proof) returns()
func (_DNSRegistrar *DNSRegistrarTransactorSession) Claim(name []byte, proof []byte) (*types.Transaction, error) {
	return _DNSRegistrar.Contract.Claim(&_DNSRegistrar.TransactOpts, name, proof)
}

// ProveAndClaim is a paid mutator transaction binding the contract method 0xd94585bd.
//
// Solidity: function proveAndClaim(bytes name, bytes input, bytes proof) returns()
func (_DNSRegistrar *DNSRegistrarTransactor) ProveAndClaim(opts *bind.TransactOpts, name []byte, input []byte, proof []byte) (*types.Transaction, error) {
	return _DNSRegistrar.contract.Transact(opts, "proveAndClaim", name, input, proof)
}

// ProveAndClaim is a paid mutator transaction binding the contract method 0xd94585bd.
//
// Solidity: function proveAndClaim(bytes name, bytes input, bytes proof) returns()
func (_DNSRegistrar *DNSRegistrarSession) ProveAndClaim(name []byte, input []byte, proof []byte) (*types.Transaction, error) {
	return _DNSRegistrar.Contract.ProveAndClaim(&_DNSRegistrar.TransactOpts, name, input, proof)
}

// ProveAndClaim is a paid mutator transaction binding the contract method 0xd94585bd.
//
// Solidity: function proveAndClaim(bytes name, bytes input, bytes proof) returns()
func (_DNSRegistrar *DNSRegistrarTransactorSession) ProveAndClaim(name []byte, input []byte, proof []byte) (*types.Transaction, error) {
	return _DNSRegistrar.Contract.ProveAndClaim(&_DNSRegistrar.TransactOpts, name, input, proof)
}
