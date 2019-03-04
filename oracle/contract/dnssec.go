// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// DNSSECABI is the input ABI used to generate the binding from.
const DNSSECABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"dnstype\",\"type\":\"uint16\"},{\"name\":\"name\",\"type\":\"bytes\"}],\"name\":\"rrdata\",\"outputs\":[{\"name\":\"inception\",\"type\":\"uint32\"},{\"name\":\"inserted\",\"type\":\"uint64\"},{\"name\":\"hash\",\"type\":\"bytes20\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"input\",\"type\":\"bytes\"},{\"name\":\"sig\",\"type\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"submitRRSet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"submitRRSets\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"anchors\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"deletetype\",\"type\":\"uint16\"},{\"name\":\"deletename\",\"type\":\"bytes\"},{\"name\":\"nsec\",\"type\":\"bytes\"},{\"name\":\"sig\",\"type\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"deleteRRSet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"name\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"rrset\",\"type\":\"bytes\"}],\"name\":\"RRSetUpdated\",\"type\":\"event\"}]"

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
	return address, tx, &DNSSEC{DNSSECCaller: DNSSECCaller{contract: contract}, DNSSECTransactor: DNSSECTransactor{contract: contract}, DNSSECFilterer: DNSSECFilterer{contract: contract}}, nil
}

// DNSSEC is an auto generated Go binding around an Ethereum contract.
type DNSSEC struct {
	DNSSECCaller     // Read-only binding to the contract
	DNSSECTransactor // Write-only binding to the contract
	DNSSECFilterer   // Log filterer for contract events
}

// DNSSECCaller is an auto generated read-only Go binding around an Ethereum contract.
type DNSSECCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DNSSECTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DNSSECTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DNSSECFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DNSSECFilterer struct {
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
	contract, err := bindDNSSEC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DNSSEC{DNSSECCaller: DNSSECCaller{contract: contract}, DNSSECTransactor: DNSSECTransactor{contract: contract}, DNSSECFilterer: DNSSECFilterer{contract: contract}}, nil
}

// NewDNSSECCaller creates a new read-only instance of DNSSEC, bound to a specific deployed contract.
func NewDNSSECCaller(address common.Address, caller bind.ContractCaller) (*DNSSECCaller, error) {
	contract, err := bindDNSSEC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DNSSECCaller{contract: contract}, nil
}

// NewDNSSECTransactor creates a new write-only instance of DNSSEC, bound to a specific deployed contract.
func NewDNSSECTransactor(address common.Address, transactor bind.ContractTransactor) (*DNSSECTransactor, error) {
	contract, err := bindDNSSEC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DNSSECTransactor{contract: contract}, nil
}

// NewDNSSECFilterer creates a new log filterer instance of DNSSEC, bound to a specific deployed contract.
func NewDNSSECFilterer(address common.Address, filterer bind.ContractFilterer) (*DNSSECFilterer, error) {
	contract, err := bindDNSSEC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DNSSECFilterer{contract: contract}, nil
}

// bindDNSSEC binds a generic wrapper to an already deployed contract.
func bindDNSSEC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DNSSECABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// Anchors is a free data retrieval call binding the contract method 0x98d35f20.
//
// Solidity: function anchors() constant returns(bytes)
func (_DNSSEC *DNSSECCaller) Anchors(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _DNSSEC.contract.Call(opts, out, "anchors")
	return *ret0, err
}

// Anchors is a free data retrieval call binding the contract method 0x98d35f20.
//
// Solidity: function anchors() constant returns(bytes)
func (_DNSSEC *DNSSECSession) Anchors() ([]byte, error) {
	return _DNSSEC.Contract.Anchors(&_DNSSEC.CallOpts)
}

// Anchors is a free data retrieval call binding the contract method 0x98d35f20.
//
// Solidity: function anchors() constant returns(bytes)
func (_DNSSEC *DNSSECCallerSession) Anchors() ([]byte, error) {
	return _DNSSEC.Contract.Anchors(&_DNSSEC.CallOpts)
}

// Rrdata is a free data retrieval call binding the contract method 0x087991bc.
//
// Solidity: function rrdata(uint16 dnstype, bytes name) constant returns(uint32 inception, uint64 inserted, bytes20 hash)
func (_DNSSEC *DNSSECCaller) Rrdata(opts *bind.CallOpts, dnstype uint16, name []byte) (struct {
	Inception uint32
	Inserted  uint64
	Hash      [20]byte
}, error) {
	ret := new(struct {
		Inception uint32
		Inserted  uint64
		Hash      [20]byte
	})
	out := ret
	err := _DNSSEC.contract.Call(opts, out, "rrdata", dnstype, name)
	return *ret, err
}

// Rrdata is a free data retrieval call binding the contract method 0x087991bc.
//
// Solidity: function rrdata(uint16 dnstype, bytes name) constant returns(uint32 inception, uint64 inserted, bytes20 hash)
func (_DNSSEC *DNSSECSession) Rrdata(dnstype uint16, name []byte) (struct {
	Inception uint32
	Inserted  uint64
	Hash      [20]byte
}, error) {
	return _DNSSEC.Contract.Rrdata(&_DNSSEC.CallOpts, dnstype, name)
}

// Rrdata is a free data retrieval call binding the contract method 0x087991bc.
//
// Solidity: function rrdata(uint16 dnstype, bytes name) constant returns(uint32 inception, uint64 inserted, bytes20 hash)
func (_DNSSEC *DNSSECCallerSession) Rrdata(dnstype uint16, name []byte) (struct {
	Inception uint32
	Inserted  uint64
	Hash      [20]byte
}, error) {
	return _DNSSEC.Contract.Rrdata(&_DNSSEC.CallOpts, dnstype, name)
}

// DeleteRRSet is a paid mutator transaction binding the contract method 0xe60b202f.
//
// Solidity: function deleteRRSet(uint16 deletetype, bytes deletename, bytes nsec, bytes sig, bytes proof) returns()
func (_DNSSEC *DNSSECTransactor) DeleteRRSet(opts *bind.TransactOpts, deletetype uint16, deletename []byte, nsec []byte, sig []byte, proof []byte) (*types.Transaction, error) {
	return _DNSSEC.contract.Transact(opts, "deleteRRSet", deletetype, deletename, nsec, sig, proof)
}

// DeleteRRSet is a paid mutator transaction binding the contract method 0xe60b202f.
//
// Solidity: function deleteRRSet(uint16 deletetype, bytes deletename, bytes nsec, bytes sig, bytes proof) returns()
func (_DNSSEC *DNSSECSession) DeleteRRSet(deletetype uint16, deletename []byte, nsec []byte, sig []byte, proof []byte) (*types.Transaction, error) {
	return _DNSSEC.Contract.DeleteRRSet(&_DNSSEC.TransactOpts, deletetype, deletename, nsec, sig, proof)
}

// DeleteRRSet is a paid mutator transaction binding the contract method 0xe60b202f.
//
// Solidity: function deleteRRSet(uint16 deletetype, bytes deletename, bytes nsec, bytes sig, bytes proof) returns()
func (_DNSSEC *DNSSECTransactorSession) DeleteRRSet(deletetype uint16, deletename []byte, nsec []byte, sig []byte, proof []byte) (*types.Transaction, error) {
	return _DNSSEC.Contract.DeleteRRSet(&_DNSSEC.TransactOpts, deletetype, deletename, nsec, sig, proof)
}

// SubmitRRSet is a paid mutator transaction binding the contract method 0x4d46d581.
//
// Solidity: function submitRRSet(bytes input, bytes sig, bytes proof) returns()
func (_DNSSEC *DNSSECTransactor) SubmitRRSet(opts *bind.TransactOpts, input []byte, sig []byte, proof []byte) (*types.Transaction, error) {
	return _DNSSEC.contract.Transact(opts, "submitRRSet", input, sig, proof)
}

// SubmitRRSet is a paid mutator transaction binding the contract method 0x4d46d581.
//
// Solidity: function submitRRSet(bytes input, bytes sig, bytes proof) returns()
func (_DNSSEC *DNSSECSession) SubmitRRSet(input []byte, sig []byte, proof []byte) (*types.Transaction, error) {
	return _DNSSEC.Contract.SubmitRRSet(&_DNSSEC.TransactOpts, input, sig, proof)
}

// SubmitRRSet is a paid mutator transaction binding the contract method 0x4d46d581.
//
// Solidity: function submitRRSet(bytes input, bytes sig, bytes proof) returns()
func (_DNSSEC *DNSSECTransactorSession) SubmitRRSet(input []byte, sig []byte, proof []byte) (*types.Transaction, error) {
	return _DNSSEC.Contract.SubmitRRSet(&_DNSSEC.TransactOpts, input, sig, proof)
}

// SubmitRRSets is a paid mutator transaction binding the contract method 0x76a14d1d.
//
// Solidity: function submitRRSets(bytes data, bytes _proof) returns(bytes)
func (_DNSSEC *DNSSECTransactor) SubmitRRSets(opts *bind.TransactOpts, data []byte, _proof []byte) (*types.Transaction, error) {
	return _DNSSEC.contract.Transact(opts, "submitRRSets", data, _proof)
}

// SubmitRRSets is a paid mutator transaction binding the contract method 0x76a14d1d.
//
// Solidity: function submitRRSets(bytes data, bytes _proof) returns(bytes)
func (_DNSSEC *DNSSECSession) SubmitRRSets(data []byte, _proof []byte) (*types.Transaction, error) {
	return _DNSSEC.Contract.SubmitRRSets(&_DNSSEC.TransactOpts, data, _proof)
}

// SubmitRRSets is a paid mutator transaction binding the contract method 0x76a14d1d.
//
// Solidity: function submitRRSets(bytes data, bytes _proof) returns(bytes)
func (_DNSSEC *DNSSECTransactorSession) SubmitRRSets(data []byte, _proof []byte) (*types.Transaction, error) {
	return _DNSSEC.Contract.SubmitRRSets(&_DNSSEC.TransactOpts, data, _proof)
}

// DNSSECRRSetUpdatedIterator is returned from FilterRRSetUpdated and is used to iterate over the raw logs and unpacked data for RRSetUpdated events raised by the DNSSEC contract.
type DNSSECRRSetUpdatedIterator struct {
	Event *DNSSECRRSetUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DNSSECRRSetUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DNSSECRRSetUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DNSSECRRSetUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DNSSECRRSetUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DNSSECRRSetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DNSSECRRSetUpdated represents a RRSetUpdated event raised by the DNSSEC contract.
type DNSSECRRSetUpdated struct {
	Name  []byte
	Rrset []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRRSetUpdated is a free log retrieval operation binding the contract event 0x55ced933cdd5a34dd03eb5d4bef19ec6ebb251dcd7a988eee0c1b9a13baaa88b.
//
// Solidity: event RRSetUpdated(bytes name, bytes rrset)
func (_DNSSEC *DNSSECFilterer) FilterRRSetUpdated(opts *bind.FilterOpts) (*DNSSECRRSetUpdatedIterator, error) {

	logs, sub, err := _DNSSEC.contract.FilterLogs(opts, "RRSetUpdated")
	if err != nil {
		return nil, err
	}
	return &DNSSECRRSetUpdatedIterator{contract: _DNSSEC.contract, event: "RRSetUpdated", logs: logs, sub: sub}, nil
}

// WatchRRSetUpdated is a free log subscription operation binding the contract event 0x55ced933cdd5a34dd03eb5d4bef19ec6ebb251dcd7a988eee0c1b9a13baaa88b.
//
// Solidity: event RRSetUpdated(bytes name, bytes rrset)
func (_DNSSEC *DNSSECFilterer) WatchRRSetUpdated(opts *bind.WatchOpts, sink chan<- *DNSSECRRSetUpdated) (event.Subscription, error) {

	logs, sub, err := _DNSSEC.contract.WatchLogs(opts, "RRSetUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DNSSECRRSetUpdated)
				if err := _DNSSEC.contract.UnpackLog(event, "RRSetUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
