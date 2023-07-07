// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Registration

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// RegistrationMetaData contains all meta data concerning the Registration contract.
var RegistrationMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"}],\"name\":\"ParticipantRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"}],\"name\":\"ParticipantResigned\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"participants\",\"outputs\":[{\"internalType\":\"enumIRegistration.LifecycleStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resign\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RegistrationABI is the input ABI used to generate the binding from.
// Deprecated: Use RegistrationMetaData.ABI instead.
var RegistrationABI = RegistrationMetaData.ABI

// Registration is an auto generated Go binding around an Ethereum contract.
type Registration struct {
	RegistrationCaller     // Read-only binding to the contract
	RegistrationTransactor // Write-only binding to the contract
	RegistrationFilterer   // Log filterer for contract events
}

// RegistrationCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistrationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistrationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistrationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrationSession struct {
	Contract     *Registration     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistrationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistrationCallerSession struct {
	Contract *RegistrationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RegistrationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistrationTransactorSession struct {
	Contract     *RegistrationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RegistrationRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistrationRaw struct {
	Contract *Registration // Generic contract binding to access the raw methods on
}

// RegistrationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistrationCallerRaw struct {
	Contract *RegistrationCaller // Generic read-only contract binding to access the raw methods on
}

// RegistrationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistrationTransactorRaw struct {
	Contract *RegistrationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistration creates a new instance of Registration, bound to a specific deployed contract.
func NewRegistration(address common.Address, backend bind.ContractBackend) (*Registration, error) {
	contract, err := bindRegistration(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registration{RegistrationCaller: RegistrationCaller{contract: contract}, RegistrationTransactor: RegistrationTransactor{contract: contract}, RegistrationFilterer: RegistrationFilterer{contract: contract}}, nil
}

// NewRegistrationCaller creates a new read-only instance of Registration, bound to a specific deployed contract.
func NewRegistrationCaller(address common.Address, caller bind.ContractCaller) (*RegistrationCaller, error) {
	contract, err := bindRegistration(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistrationCaller{contract: contract}, nil
}

// NewRegistrationTransactor creates a new write-only instance of Registration, bound to a specific deployed contract.
func NewRegistrationTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistrationTransactor, error) {
	contract, err := bindRegistration(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistrationTransactor{contract: contract}, nil
}

// NewRegistrationFilterer creates a new log filterer instance of Registration, bound to a specific deployed contract.
func NewRegistrationFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistrationFilterer, error) {
	contract, err := bindRegistration(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistrationFilterer{contract: contract}, nil
}

// bindRegistration binds a generic wrapper to an already deployed contract.
func bindRegistration(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RegistrationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registration *RegistrationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registration.Contract.RegistrationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registration *RegistrationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registration.Contract.RegistrationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registration *RegistrationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registration.Contract.RegistrationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registration *RegistrationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registration.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registration *RegistrationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registration.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registration *RegistrationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registration.Contract.contract.Transact(opts, method, params...)
}

// Participants is a free data retrieval call binding the contract method 0x09e69ede.
//
// Solidity: function participants(address ) view returns(uint8)
func (_Registration *RegistrationCaller) Participants(opts *bind.CallOpts, arg0 common.Address) (uint8, error) {
	var out []interface{}
	err := _Registration.contract.Call(opts, &out, "participants", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Participants is a free data retrieval call binding the contract method 0x09e69ede.
//
// Solidity: function participants(address ) view returns(uint8)
func (_Registration *RegistrationSession) Participants(arg0 common.Address) (uint8, error) {
	return _Registration.Contract.Participants(&_Registration.CallOpts, arg0)
}

// Participants is a free data retrieval call binding the contract method 0x09e69ede.
//
// Solidity: function participants(address ) view returns(uint8)
func (_Registration *RegistrationCallerSession) Participants(arg0 common.Address) (uint8, error) {
	return _Registration.Contract.Participants(&_Registration.CallOpts, arg0)
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_Registration *RegistrationTransactor) Register(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registration.contract.Transact(opts, "register")
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_Registration *RegistrationSession) Register() (*types.Transaction, error) {
	return _Registration.Contract.Register(&_Registration.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_Registration *RegistrationTransactorSession) Register() (*types.Transaction, error) {
	return _Registration.Contract.Register(&_Registration.TransactOpts)
}

// Resign is a paid mutator transaction binding the contract method 0x69652fcf.
//
// Solidity: function resign() returns()
func (_Registration *RegistrationTransactor) Resign(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registration.contract.Transact(opts, "resign")
}

// Resign is a paid mutator transaction binding the contract method 0x69652fcf.
//
// Solidity: function resign() returns()
func (_Registration *RegistrationSession) Resign() (*types.Transaction, error) {
	return _Registration.Contract.Resign(&_Registration.TransactOpts)
}

// Resign is a paid mutator transaction binding the contract method 0x69652fcf.
//
// Solidity: function resign() returns()
func (_Registration *RegistrationTransactorSession) Resign() (*types.Transaction, error) {
	return _Registration.Contract.Resign(&_Registration.TransactOpts)
}

// RegistrationParticipantRegisteredIterator is returned from FilterParticipantRegistered and is used to iterate over the raw logs and unpacked data for ParticipantRegistered events raised by the Registration contract.
type RegistrationParticipantRegisteredIterator struct {
	Event *RegistrationParticipantRegistered // Event containing the contract specifics and raw log

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
func (it *RegistrationParticipantRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrationParticipantRegistered)
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
		it.Event = new(RegistrationParticipantRegistered)
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
func (it *RegistrationParticipantRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrationParticipantRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrationParticipantRegistered represents a ParticipantRegistered event raised by the Registration contract.
type RegistrationParticipantRegistered struct {
	Participant common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterParticipantRegistered is a free log retrieval operation binding the contract event 0xe11711cd714e06fbbbea301a8e90822f2f2ea4808e37e3adf06038f33c53ff27.
//
// Solidity: event ParticipantRegistered(address indexed participant)
func (_Registration *RegistrationFilterer) FilterParticipantRegistered(opts *bind.FilterOpts, participant []common.Address) (*RegistrationParticipantRegisteredIterator, error) {

	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _Registration.contract.FilterLogs(opts, "ParticipantRegistered", participantRule)
	if err != nil {
		return nil, err
	}
	return &RegistrationParticipantRegisteredIterator{contract: _Registration.contract, event: "ParticipantRegistered", logs: logs, sub: sub}, nil
}

// WatchParticipantRegistered is a free log subscription operation binding the contract event 0xe11711cd714e06fbbbea301a8e90822f2f2ea4808e37e3adf06038f33c53ff27.
//
// Solidity: event ParticipantRegistered(address indexed participant)
func (_Registration *RegistrationFilterer) WatchParticipantRegistered(opts *bind.WatchOpts, sink chan<- *RegistrationParticipantRegistered, participant []common.Address) (event.Subscription, error) {

	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _Registration.contract.WatchLogs(opts, "ParticipantRegistered", participantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrationParticipantRegistered)
				if err := _Registration.contract.UnpackLog(event, "ParticipantRegistered", log); err != nil {
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

// ParseParticipantRegistered is a log parse operation binding the contract event 0xe11711cd714e06fbbbea301a8e90822f2f2ea4808e37e3adf06038f33c53ff27.
//
// Solidity: event ParticipantRegistered(address indexed participant)
func (_Registration *RegistrationFilterer) ParseParticipantRegistered(log types.Log) (*RegistrationParticipantRegistered, error) {
	event := new(RegistrationParticipantRegistered)
	if err := _Registration.contract.UnpackLog(event, "ParticipantRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrationParticipantResignedIterator is returned from FilterParticipantResigned and is used to iterate over the raw logs and unpacked data for ParticipantResigned events raised by the Registration contract.
type RegistrationParticipantResignedIterator struct {
	Event *RegistrationParticipantResigned // Event containing the contract specifics and raw log

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
func (it *RegistrationParticipantResignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrationParticipantResigned)
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
		it.Event = new(RegistrationParticipantResigned)
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
func (it *RegistrationParticipantResignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrationParticipantResignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrationParticipantResigned represents a ParticipantResigned event raised by the Registration contract.
type RegistrationParticipantResigned struct {
	Participant common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterParticipantResigned is a free log retrieval operation binding the contract event 0x2ef5d7853e7253f46168e32337524fae1de5ca1e34b8438d2df597e50e18f39f.
//
// Solidity: event ParticipantResigned(address indexed participant)
func (_Registration *RegistrationFilterer) FilterParticipantResigned(opts *bind.FilterOpts, participant []common.Address) (*RegistrationParticipantResignedIterator, error) {

	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _Registration.contract.FilterLogs(opts, "ParticipantResigned", participantRule)
	if err != nil {
		return nil, err
	}
	return &RegistrationParticipantResignedIterator{contract: _Registration.contract, event: "ParticipantResigned", logs: logs, sub: sub}, nil
}

// WatchParticipantResigned is a free log subscription operation binding the contract event 0x2ef5d7853e7253f46168e32337524fae1de5ca1e34b8438d2df597e50e18f39f.
//
// Solidity: event ParticipantResigned(address indexed participant)
func (_Registration *RegistrationFilterer) WatchParticipantResigned(opts *bind.WatchOpts, sink chan<- *RegistrationParticipantResigned, participant []common.Address) (event.Subscription, error) {

	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _Registration.contract.WatchLogs(opts, "ParticipantResigned", participantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrationParticipantResigned)
				if err := _Registration.contract.UnpackLog(event, "ParticipantResigned", log); err != nil {
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

// ParseParticipantResigned is a log parse operation binding the contract event 0x2ef5d7853e7253f46168e32337524fae1de5ca1e34b8438d2df597e50e18f39f.
//
// Solidity: event ParticipantResigned(address indexed participant)
func (_Registration *RegistrationFilterer) ParseParticipantResigned(log types.Log) (*RegistrationParticipantResigned, error) {
	event := new(RegistrationParticipantResigned)
	if err := _Registration.contract.UnpackLog(event, "ParticipantResigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
