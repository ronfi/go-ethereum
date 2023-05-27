// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ronswapv3

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
)

// Ronswapv3MetaData contains all meta data concerning the Ronswapv3 contract.
var Ronswapv3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_executors\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_charger\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"charger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalOk\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"failed\",\"type\":\"bool\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_executors\",\"type\":\"address[]\"}],\"name\":\"addExecutor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_executors\",\"type\":\"address[]\"}],\"name\":\"delExecutor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"pancakeV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenPairsAndFee\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"ronSwapV3\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_mainContract\",\"type\":\"address\"}],\"name\":\"setMain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// Ronswapv3ABI is the input ABI used to generate the binding from.
// Deprecated: Use Ronswapv3MetaData.ABI instead.
var Ronswapv3ABI = Ronswapv3MetaData.ABI

// Ronswapv3 is an auto generated Go binding around an Ethereum contract.
type Ronswapv3 struct {
	Ronswapv3Caller     // Read-only binding to the contract
	Ronswapv3Transactor // Write-only binding to the contract
	Ronswapv3Filterer   // Log filterer for contract events
}

// Ronswapv3Caller is an auto generated read-only Go binding around an Ethereum contract.
type Ronswapv3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ronswapv3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Ronswapv3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ronswapv3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Ronswapv3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ronswapv3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Ronswapv3Session struct {
	Contract     *Ronswapv3        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Ronswapv3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Ronswapv3CallerSession struct {
	Contract *Ronswapv3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// Ronswapv3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Ronswapv3TransactorSession struct {
	Contract     *Ronswapv3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// Ronswapv3Raw is an auto generated low-level Go binding around an Ethereum contract.
type Ronswapv3Raw struct {
	Contract *Ronswapv3 // Generic contract binding to access the raw methods on
}

// Ronswapv3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Ronswapv3CallerRaw struct {
	Contract *Ronswapv3Caller // Generic read-only contract binding to access the raw methods on
}

// Ronswapv3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Ronswapv3TransactorRaw struct {
	Contract *Ronswapv3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewRonswapv3 creates a new instance of Ronswapv3, bound to a specific deployed contract.
func NewRonswapv3(address common.Address, backend bind.ContractBackend) (*Ronswapv3, error) {
	contract, err := bindRonswapv3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ronswapv3{Ronswapv3Caller: Ronswapv3Caller{contract: contract}, Ronswapv3Transactor: Ronswapv3Transactor{contract: contract}, Ronswapv3Filterer: Ronswapv3Filterer{contract: contract}}, nil
}

// NewRonswapv3Caller creates a new read-only instance of Ronswapv3, bound to a specific deployed contract.
func NewRonswapv3Caller(address common.Address, caller bind.ContractCaller) (*Ronswapv3Caller, error) {
	contract, err := bindRonswapv3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Ronswapv3Caller{contract: contract}, nil
}

// NewRonswapv3Transactor creates a new write-only instance of Ronswapv3, bound to a specific deployed contract.
func NewRonswapv3Transactor(address common.Address, transactor bind.ContractTransactor) (*Ronswapv3Transactor, error) {
	contract, err := bindRonswapv3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Ronswapv3Transactor{contract: contract}, nil
}

// NewRonswapv3Filterer creates a new log filterer instance of Ronswapv3, bound to a specific deployed contract.
func NewRonswapv3Filterer(address common.Address, filterer bind.ContractFilterer) (*Ronswapv3Filterer, error) {
	contract, err := bindRonswapv3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Ronswapv3Filterer{contract: contract}, nil
}

// bindRonswapv3 binds a generic wrapper to an already deployed contract.
func bindRonswapv3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Ronswapv3ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ronswapv3 *Ronswapv3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ronswapv3.Contract.Ronswapv3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ronswapv3 *Ronswapv3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ronswapv3.Contract.Ronswapv3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ronswapv3 *Ronswapv3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ronswapv3.Contract.Ronswapv3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ronswapv3 *Ronswapv3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ronswapv3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ronswapv3 *Ronswapv3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ronswapv3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ronswapv3 *Ronswapv3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ronswapv3.Contract.contract.Transact(opts, method, params...)
}

// AddExecutor is a paid mutator transaction binding the contract method 0xee26468a.
//
// Solidity: function addExecutor(address[] _executors) returns()
func (_Ronswapv3 *Ronswapv3Transactor) AddExecutor(opts *bind.TransactOpts, _executors []common.Address) (*types.Transaction, error) {
	return _Ronswapv3.contract.Transact(opts, "addExecutor", _executors)
}

// AddExecutor is a paid mutator transaction binding the contract method 0xee26468a.
//
// Solidity: function addExecutor(address[] _executors) returns()
func (_Ronswapv3 *Ronswapv3Session) AddExecutor(_executors []common.Address) (*types.Transaction, error) {
	return _Ronswapv3.Contract.AddExecutor(&_Ronswapv3.TransactOpts, _executors)
}

// AddExecutor is a paid mutator transaction binding the contract method 0xee26468a.
//
// Solidity: function addExecutor(address[] _executors) returns()
func (_Ronswapv3 *Ronswapv3TransactorSession) AddExecutor(_executors []common.Address) (*types.Transaction, error) {
	return _Ronswapv3.Contract.AddExecutor(&_Ronswapv3.TransactOpts, _executors)
}

// DelExecutor is a paid mutator transaction binding the contract method 0x30f9006a.
//
// Solidity: function delExecutor(address[] _executors) returns()
func (_Ronswapv3 *Ronswapv3Transactor) DelExecutor(opts *bind.TransactOpts, _executors []common.Address) (*types.Transaction, error) {
	return _Ronswapv3.contract.Transact(opts, "delExecutor", _executors)
}

// DelExecutor is a paid mutator transaction binding the contract method 0x30f9006a.
//
// Solidity: function delExecutor(address[] _executors) returns()
func (_Ronswapv3 *Ronswapv3Session) DelExecutor(_executors []common.Address) (*types.Transaction, error) {
	return _Ronswapv3.Contract.DelExecutor(&_Ronswapv3.TransactOpts, _executors)
}

// DelExecutor is a paid mutator transaction binding the contract method 0x30f9006a.
//
// Solidity: function delExecutor(address[] _executors) returns()
func (_Ronswapv3 *Ronswapv3TransactorSession) DelExecutor(_executors []common.Address) (*types.Transaction, error) {
	return _Ronswapv3.Contract.DelExecutor(&_Ronswapv3.TransactOpts, _executors)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) payable returns()
func (_Ronswapv3 *Ronswapv3Transactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Ronswapv3.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) payable returns()
func (_Ronswapv3 *Ronswapv3Session) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Ronswapv3.Contract.Deposit(&_Ronswapv3.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) payable returns()
func (_Ronswapv3 *Ronswapv3TransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Ronswapv3.Contract.Deposit(&_Ronswapv3.TransactOpts, amount)
}

// PancakeV3SwapCallback is a paid mutator transaction binding the contract method 0x23a69e75.
//
// Solidity: function pancakeV3SwapCallback(int256 amount0, int256 amount1, bytes _data) returns()
func (_Ronswapv3 *Ronswapv3Transactor) PancakeV3SwapCallback(opts *bind.TransactOpts, amount0 *big.Int, amount1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _Ronswapv3.contract.Transact(opts, "pancakeV3SwapCallback", amount0, amount1, _data)
}

// PancakeV3SwapCallback is a paid mutator transaction binding the contract method 0x23a69e75.
//
// Solidity: function pancakeV3SwapCallback(int256 amount0, int256 amount1, bytes _data) returns()
func (_Ronswapv3 *Ronswapv3Session) PancakeV3SwapCallback(amount0 *big.Int, amount1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _Ronswapv3.Contract.PancakeV3SwapCallback(&_Ronswapv3.TransactOpts, amount0, amount1, _data)
}

// PancakeV3SwapCallback is a paid mutator transaction binding the contract method 0x23a69e75.
//
// Solidity: function pancakeV3SwapCallback(int256 amount0, int256 amount1, bytes _data) returns()
func (_Ronswapv3 *Ronswapv3TransactorSession) PancakeV3SwapCallback(amount0 *big.Int, amount1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _Ronswapv3.Contract.PancakeV3SwapCallback(&_Ronswapv3.TransactOpts, amount0, amount1, _data)
}

// RonSwapV3 is a paid mutator transaction binding the contract method 0x69bf3560.
//
// Solidity: function ronSwapV3(uint256[] tokenPairsAndFee, uint256 amountIn) returns(uint256 amountOut)
func (_Ronswapv3 *Ronswapv3Transactor) RonSwapV3(opts *bind.TransactOpts, tokenPairsAndFee []*big.Int, amountIn *big.Int) (*types.Transaction, error) {
	return _Ronswapv3.contract.Transact(opts, "ronSwapV3", tokenPairsAndFee, amountIn)
}

// RonSwapV3 is a paid mutator transaction binding the contract method 0x69bf3560.
//
// Solidity: function ronSwapV3(uint256[] tokenPairsAndFee, uint256 amountIn) returns(uint256 amountOut)
func (_Ronswapv3 *Ronswapv3Session) RonSwapV3(tokenPairsAndFee []*big.Int, amountIn *big.Int) (*types.Transaction, error) {
	return _Ronswapv3.Contract.RonSwapV3(&_Ronswapv3.TransactOpts, tokenPairsAndFee, amountIn)
}

// RonSwapV3 is a paid mutator transaction binding the contract method 0x69bf3560.
//
// Solidity: function ronSwapV3(uint256[] tokenPairsAndFee, uint256 amountIn) returns(uint256 amountOut)
func (_Ronswapv3 *Ronswapv3TransactorSession) RonSwapV3(tokenPairsAndFee []*big.Int, amountIn *big.Int) (*types.Transaction, error) {
	return _Ronswapv3.Contract.RonSwapV3(&_Ronswapv3.TransactOpts, tokenPairsAndFee, amountIn)
}

// SetMain is a paid mutator transaction binding the contract method 0x801161b0.
//
// Solidity: function setMain(address _mainContract) returns()
func (_Ronswapv3 *Ronswapv3Transactor) SetMain(opts *bind.TransactOpts, _mainContract common.Address) (*types.Transaction, error) {
	return _Ronswapv3.contract.Transact(opts, "setMain", _mainContract)
}

// SetMain is a paid mutator transaction binding the contract method 0x801161b0.
//
// Solidity: function setMain(address _mainContract) returns()
func (_Ronswapv3 *Ronswapv3Session) SetMain(_mainContract common.Address) (*types.Transaction, error) {
	return _Ronswapv3.Contract.SetMain(&_Ronswapv3.TransactOpts, _mainContract)
}

// SetMain is a paid mutator transaction binding the contract method 0x801161b0.
//
// Solidity: function setMain(address _mainContract) returns()
func (_Ronswapv3 *Ronswapv3TransactorSession) SetMain(_mainContract common.Address) (*types.Transaction, error) {
	return _Ronswapv3.Contract.SetMain(&_Ronswapv3.TransactOpts, _mainContract)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 amount, address token) returns()
func (_Ronswapv3 *Ronswapv3Transactor) Withdraw(opts *bind.TransactOpts, amount *big.Int, token common.Address) (*types.Transaction, error) {
	return _Ronswapv3.contract.Transact(opts, "withdraw", amount, token)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 amount, address token) returns()
func (_Ronswapv3 *Ronswapv3Session) Withdraw(amount *big.Int, token common.Address) (*types.Transaction, error) {
	return _Ronswapv3.Contract.Withdraw(&_Ronswapv3.TransactOpts, amount, token)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 amount, address token) returns()
func (_Ronswapv3 *Ronswapv3TransactorSession) Withdraw(amount *big.Int, token common.Address) (*types.Transaction, error) {
	return _Ronswapv3.Contract.Withdraw(&_Ronswapv3.TransactOpts, amount, token)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ronswapv3 *Ronswapv3Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ronswapv3.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ronswapv3 *Ronswapv3Session) Receive() (*types.Transaction, error) {
	return _Ronswapv3.Contract.Receive(&_Ronswapv3.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ronswapv3 *Ronswapv3TransactorSession) Receive() (*types.Transaction, error) {
	return _Ronswapv3.Contract.Receive(&_Ronswapv3.TransactOpts)
}

// Ronswapv3DepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Ronswapv3 contract.
type Ronswapv3DepositedIterator struct {
	Event *Ronswapv3Deposited // Event containing the contract specifics and raw log

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
func (it *Ronswapv3DepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Ronswapv3Deposited)
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
		it.Event = new(Ronswapv3Deposited)
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
func (it *Ronswapv3DepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Ronswapv3DepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Ronswapv3Deposited represents a Deposited event raised by the Ronswapv3 contract.
type Ronswapv3Deposited struct {
	Charger     common.Address
	TotalOk     *big.Int
	TotalAmount *big.Int
	Failed      bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x65e6f89b6907d6277741ee4ccbc4ae260163a17f16bbb55b5344dff064867c40.
//
// Solidity: event Deposited(address indexed charger, uint256 totalOk, uint256 totalAmount, bool failed)
func (_Ronswapv3 *Ronswapv3Filterer) FilterDeposited(opts *bind.FilterOpts, charger []common.Address) (*Ronswapv3DepositedIterator, error) {

	var chargerRule []interface{}
	for _, chargerItem := range charger {
		chargerRule = append(chargerRule, chargerItem)
	}

	logs, sub, err := _Ronswapv3.contract.FilterLogs(opts, "Deposited", chargerRule)
	if err != nil {
		return nil, err
	}
	return &Ronswapv3DepositedIterator{contract: _Ronswapv3.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x65e6f89b6907d6277741ee4ccbc4ae260163a17f16bbb55b5344dff064867c40.
//
// Solidity: event Deposited(address indexed charger, uint256 totalOk, uint256 totalAmount, bool failed)
func (_Ronswapv3 *Ronswapv3Filterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *Ronswapv3Deposited, charger []common.Address) (event.Subscription, error) {

	var chargerRule []interface{}
	for _, chargerItem := range charger {
		chargerRule = append(chargerRule, chargerItem)
	}

	logs, sub, err := _Ronswapv3.contract.WatchLogs(opts, "Deposited", chargerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Ronswapv3Deposited)
				if err := _Ronswapv3.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x65e6f89b6907d6277741ee4ccbc4ae260163a17f16bbb55b5344dff064867c40.
//
// Solidity: event Deposited(address indexed charger, uint256 totalOk, uint256 totalAmount, bool failed)
func (_Ronswapv3 *Ronswapv3Filterer) ParseDeposited(log types.Log) (*Ronswapv3Deposited, error) {
	event := new(Ronswapv3Deposited)
	if err := _Ronswapv3.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
