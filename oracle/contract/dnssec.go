// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// DNSSECABI is the input ABI used to generate the binding from.
const DNSSECABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"class\",\"type\":\"uint16\"},{\"name\":\"name\",\"type\":\"bytes\"},{\"name\":\"input\",\"type\":\"bytes\"},{\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"submitRRSet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"class\",\"type\":\"uint16\"},{\"name\":\"dnstype\",\"type\":\"uint16\"},{\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"rrset\",\"outputs\":[{\"name\":\"inception\",\"type\":\"uint32\"},{\"name\":\"expiration\",\"type\":\"uint32\"},{\"name\":\"inserted\",\"type\":\"uint64\"},{\"name\":\"rrs\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"nameHash\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"RRSetUpdated\",\"type\":\"event\"}]"

// DNSSECBin is the compiled bytecode used for deploying new contracts.
const DNSSECBin = `0x`

// DeployDNSSEC deploys a new Ethereum contract, binding an instance of DNSSEC to it.
func DeployDNSSEC(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DNSSEC, error) {
	parsed, err := abi.JSON(strings.NewReader(DNSSECABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DNSSECBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DNSSEC{DNSSECCaller: DNSSECCaller{contract: contract}, DNSSECTransactor: DNSSECTransactor{contract: contract}}, nil
}

// DNSSEC is an auto generated Go binding around an Ethereum contract.
type DNSSEC struct {
	DNSSECCaller     // Read-only binding to the contract
	DNSSECTransactor // Write-only binding to the contract
}

// DNSSECCaller is an auto generated read-only Go binding around an Ethereum contract.
type DNSSECCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DNSSECTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DNSSECTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DNSSECSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DNSSECSession struct {
	Contract     *DNSSEC           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DNSSECCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DNSSECCallerSession struct {
	Contract *DNSSECCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DNSSECTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DNSSECTransactorSession struct {
	Contract     *DNSSECTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DNSSECRaw is an auto generated low-level Go binding around an Ethereum contract.
type DNSSECRaw struct {
	Contract *DNSSEC // Generic contract binding to access the raw methods on
}

// DNSSECCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DNSSECCallerRaw struct {
	Contract *DNSSECCaller // Generic read-only contract binding to access the raw methods on
}

// DNSSECTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DNSSECTransactorRaw struct {
	Contract *DNSSECTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDNSSEC creates a new instance of DNSSEC, bound to a specific deployed contract.
func NewDNSSEC(address common.Address, backend bind.ContractBackend) (*DNSSEC, error) {
	contract, err := bindDNSSEC(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DNSSEC{DNSSECCaller: DNSSECCaller{contract: contract}, DNSSECTransactor: DNSSECTransactor{contract: contract}}, nil
}

// NewDNSSECCaller creates a new read-only instance of DNSSEC, bound to a specific deployed contract.
func NewDNSSECCaller(address common.Address, caller bind.ContractCaller) (*DNSSECCaller, error) {
	contract, err := bindDNSSEC(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &DNSSECCaller{contract: contract}, nil
}

// NewDNSSECTransactor creates a new write-only instance of DNSSEC, bound to a specific deployed contract.
func NewDNSSECTransactor(address common.Address, transactor bind.ContractTransactor) (*DNSSECTransactor, error) {
	contract, err := bindDNSSEC(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &DNSSECTransactor{contract: contract}, nil
}

// bindDNSSEC binds a generic wrapper to an already deployed contract.
func bindDNSSEC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DNSSECABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, nil), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DNSSEC *DNSSECRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DNSSEC.Contract.DNSSECCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DNSSEC *DNSSECRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DNSSEC.Contract.DNSSECTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DNSSEC *DNSSECRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DNSSEC.Contract.DNSSECTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DNSSEC *DNSSECCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DNSSEC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DNSSEC *DNSSECTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DNSSEC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DNSSEC *DNSSECTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DNSSEC.Contract.contract.Transact(opts, method, params...)
}

// Rrset is a free data retrieval call binding the contract method 0x5f2b77ce.
//
// Solidity: function rrset(class uint16, dnstype uint16, name bytes) constant returns(inception uint32, expiration uint32, inserted uint64, rrs bytes)
func (_DNSSEC *DNSSECCaller) Rrset(opts *bind.CallOpts, class uint16, dnstype uint16, name []byte) (struct {
	Inception  uint32
	Expiration uint32
	Inserted   uint64
	Rrs        []byte
}, error) {
	ret := new(struct {
		Inception  uint32
		Expiration uint32
		Inserted   uint64
		Rrs        []byte
	})
	out := ret
	err := _DNSSEC.contract.Call(opts, out, "rrset", class, dnstype, name)
	return *ret, err
}

// Rrset is a free data retrieval call binding the contract method 0x5f2b77ce.
//
// Solidity: function rrset(class uint16, dnstype uint16, name bytes) constant returns(inception uint32, expiration uint32, inserted uint64, rrs bytes)
func (_DNSSEC *DNSSECSession) Rrset(class uint16, dnstype uint16, name []byte) (struct {
	Inception  uint32
	Expiration uint32
	Inserted   uint64
	Rrs        []byte
}, error) {
	return _DNSSEC.Contract.Rrset(&_DNSSEC.CallOpts, class, dnstype, name)
}

// Rrset is a free data retrieval call binding the contract method 0x5f2b77ce.
//
// Solidity: function rrset(class uint16, dnstype uint16, name bytes) constant returns(inception uint32, expiration uint32, inserted uint64, rrs bytes)
func (_DNSSEC *DNSSECCallerSession) Rrset(class uint16, dnstype uint16, name []byte) (struct {
	Inception  uint32
	Expiration uint32
	Inserted   uint64
	Rrs        []byte
}, error) {
	return _DNSSEC.Contract.Rrset(&_DNSSEC.CallOpts, class, dnstype, name)
}

// SubmitRRSet is a paid mutator transaction binding the contract method 0x11267632.
//
// Solidity: function submitRRSet(class uint16, name bytes, input bytes, sig bytes) returns()
func (_DNSSEC *DNSSECTransactor) SubmitRRSet(opts *bind.TransactOpts, class uint16, name []byte, input []byte, sig []byte) (*types.Transaction, error) {
	return _DNSSEC.contract.Transact(opts, "submitRRSet", class, name, input, sig)
}

// SubmitRRSet is a paid mutator transaction binding the contract method 0x11267632.
//
// Solidity: function submitRRSet(class uint16, name bytes, input bytes, sig bytes) returns()
func (_DNSSEC *DNSSECSession) SubmitRRSet(class uint16, name []byte, input []byte, sig []byte) (*types.Transaction, error) {
	return _DNSSEC.Contract.SubmitRRSet(&_DNSSEC.TransactOpts, class, name, input, sig)
}

// SubmitRRSet is a paid mutator transaction binding the contract method 0x11267632.
//
// Solidity: function submitRRSet(class uint16, name bytes, input bytes, sig bytes) returns()
func (_DNSSEC *DNSSECTransactorSession) SubmitRRSet(class uint16, name []byte, input []byte, sig []byte) (*types.Transaction, error) {
	return _DNSSEC.Contract.SubmitRRSet(&_DNSSEC.TransactOpts, class, name, input, sig)
}
