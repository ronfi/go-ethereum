// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ronswapv3fe

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

// Ronswapv3feMetaData contains all meta data concerning the Ronswapv3fe contract.
var Ronswapv3feMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_executors\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_charger\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"charger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalOk\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"failed\",\"type\":\"bool\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_executors\",\"type\":\"address[]\"}],\"name\":\"addExecutor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_executors\",\"type\":\"address[]\"}],\"name\":\"delExecutor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"pancakeCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"pancakeV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenPairsAndFee\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"ronSwapV3\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"uniswapV2Call\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// Ronswapv3feABI is the input ABI used to generate the binding from.
// Deprecated: Use Ronswapv3feMetaData.ABI instead.
var Ronswapv3feABI = Ronswapv3feMetaData.ABI

// Ronswapv3fe is an auto generated Go binding around an Ethereum contract.
type Ronswapv3fe struct {
	Ronswapv3feCaller     // Read-only binding to the contract
	Ronswapv3feTransactor // Write-only binding to the contract
	Ronswapv3feFilterer   // Log filterer for contract events
}

// Ronswapv3feCaller is an auto generated read-only Go binding around an Ethereum contract.
type Ronswapv3feCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ronswapv3feTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Ronswapv3feTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ronswapv3feFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Ronswapv3feFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ronswapv3feSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Ronswapv3feSession struct {
	Contract     *Ronswapv3fe      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Ronswapv3feCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Ronswapv3feCallerSession struct {
	Contract *Ronswapv3feCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// Ronswapv3feTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Ronswapv3feTransactorSession struct {
	Contract     *Ronswapv3feTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// Ronswapv3feRaw is an auto generated low-level Go binding around an Ethereum contract.
type Ronswapv3feRaw struct {
	Contract *Ronswapv3fe // Generic contract binding to access the raw methods on
}

// Ronswapv3feCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Ronswapv3feCallerRaw struct {
	Contract *Ronswapv3feCaller // Generic read-only contract binding to access the raw methods on
}

// Ronswapv3feTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Ronswapv3feTransactorRaw struct {
	Contract *Ronswapv3feTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRonswapv3fe creates a new instance of Ronswapv3fe, bound to a specific deployed contract.
func NewRonswapv3fe(address common.Address, backend bind.ContractBackend) (*Ronswapv3fe, error) {
	contract, err := bindRonswapv3fe(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ronswapv3fe{Ronswapv3feCaller: Ronswapv3feCaller{contract: contract}, Ronswapv3feTransactor: Ronswapv3feTransactor{contract: contract}, Ronswapv3feFilterer: Ronswapv3feFilterer{contract: contract}}, nil
}

// NewRonswapv3feCaller creates a new read-only instance of Ronswapv3fe, bound to a specific deployed contract.
func NewRonswapv3feCaller(address common.Address, caller bind.ContractCaller) (*Ronswapv3feCaller, error) {
	contract, err := bindRonswapv3fe(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Ronswapv3feCaller{contract: contract}, nil
}

// NewRonswapv3feTransactor creates a new write-only instance of Ronswapv3fe, bound to a specific deployed contract.
func NewRonswapv3feTransactor(address common.Address, transactor bind.ContractTransactor) (*Ronswapv3feTransactor, error) {
	contract, err := bindRonswapv3fe(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Ronswapv3feTransactor{contract: contract}, nil
}

// NewRonswapv3feFilterer creates a new log filterer instance of Ronswapv3fe, bound to a specific deployed contract.
func NewRonswapv3feFilterer(address common.Address, filterer bind.ContractFilterer) (*Ronswapv3feFilterer, error) {
	contract, err := bindRonswapv3fe(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Ronswapv3feFilterer{contract: contract}, nil
}

// bindRonswapv3fe binds a generic wrapper to an already deployed contract.
func bindRonswapv3fe(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Ronswapv3feABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ronswapv3fe *Ronswapv3feRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ronswapv3fe.Contract.Ronswapv3feCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ronswapv3fe *Ronswapv3feRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ronswapv3fe.Contract.Ronswapv3feTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ronswapv3fe *Ronswapv3feRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ronswapv3fe.Contract.Ronswapv3feTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ronswapv3fe *Ronswapv3feCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ronswapv3fe.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ronswapv3fe *Ronswapv3feTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ronswapv3fe.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ronswapv3fe *Ronswapv3feTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ronswapv3fe.Contract.contract.Transact(opts, method, params...)
}

// AddExecutor is a paid mutator transaction binding the contract method 0xee26468a.
//
// Solidity: function addExecutor(address[] _executors) returns()
func (_Ronswapv3fe *Ronswapv3feTransactor) AddExecutor(opts *bind.TransactOpts, _executors []common.Address) (*types.Transaction, error) {
	return _Ronswapv3fe.contract.Transact(opts, "addExecutor", _executors)
}

// AddExecutor is a paid mutator transaction binding the contract method 0xee26468a.
//
// Solidity: function addExecutor(address[] _executors) returns()
func (_Ronswapv3fe *Ronswapv3feSession) AddExecutor(_executors []common.Address) (*types.Transaction, error) {
	return _Ronswapv3fe.Contract.AddExecutor(&_Ronswapv3fe.TransactOpts, _executors)
}

// AddExecutor is a paid mutator transaction binding the contract method 0xee26468a.
//
// Solidity: function addExecutor(address[] _executors) returns()
func (_Ronswapv3fe *Ronswapv3feTransactorSession) AddExecutor(_executors []common.Address) (*types.Transaction, error) {
	return _Ronswapv3fe.Contract.AddExecutor(&_Ronswapv3fe.TransactOpts, _executors)
}

// DelExecutor is a paid mutator transaction binding the contract method 0x30f9006a.
//
// Solidity: function delExecutor(address[] _executors) returns()
func (_Ronswapv3fe *Ronswapv3feTransactor) DelExecutor(opts *bind.TransactOpts, _executors []common.Address) (*types.Transaction, error) {
	return _Ronswapv3fe.contract.Transact(opts, "delExecutor", _executors)
}

// DelExecutor is a paid mutator transaction binding the contract method 0x30f9006a.
//
// Solidity: function delExecutor(address[] _executors) returns()
func (_Ronswapv3fe *Ronswapv3feSession) DelExecutor(_executors []common.Address) (*types.Transaction, error) {
	return _Ronswapv3fe.Contract.DelExecutor(&_Ronswapv3fe.TransactOpts, _executors)
}

// DelExecutor is a paid mutator transaction binding the contract method 0x30f9006a.
//
// Solidity: function delExecutor(address[] _executors) returns()
func (_Ronswapv3fe *Ronswapv3feTransactorSession) DelExecutor(_executors []common.Address) (*types.Transaction, error) {
	return _Ronswapv3fe.Contract.DelExecutor(&_Ronswapv3fe.TransactOpts, _executors)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) payable returns()
func (_Ronswapv3fe *Ronswapv3feTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Ronswapv3fe.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) payable returns()
func (_Ronswapv3fe *Ronswapv3feSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Ronswapv3fe.Contract.Deposit(&_Ronswapv3fe.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) payable returns()
func (_Ronswapv3fe *Ronswapv3feTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Ronswapv3fe.Contract.Deposit(&_Ronswapv3fe.TransactOpts, amount)
}

// PancakeCall is a paid mutator transaction binding the contract method 0x84800812.
//
// Solidity: function pancakeCall(address sender, uint256 amount0, uint256 amount1, bytes _data) returns()
func (_Ronswapv3fe *Ronswapv3feTransactor) PancakeCall(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _Ronswapv3fe.contract.Transact(opts, "pancakeCall", sender, amount0, amount1, _data)
}

// PancakeCall is a paid mutator transaction binding the contract method 0x84800812.
//
// Solidity: function pancakeCall(address sender, uint256 amount0, uint256 amount1, bytes _data) returns()
func (_Ronswapv3fe *Ronswapv3feSession) PancakeCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _Ronswapv3fe.Contract.PancakeCall(&_Ronswapv3fe.TransactOpts, sender, amount0, amount1, _data)
}

// PancakeCall is a paid mutator transaction binding the contract method 0x84800812.
//
// Solidity: function pancakeCall(address sender, uint256 amount0, uint256 amount1, bytes _data) returns()
func (_Ronswapv3fe *Ronswapv3feTransactorSession) PancakeCall(sender common.Address, amount0 *big.Int, amount1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _Ronswapv3fe.Contract.PancakeCall(&_Ronswapv3fe.TransactOpts, sender, amount0, amount1, _data)
}

// PancakeV3SwapCallback is a paid mutator transaction binding the contract method 0x23a69e75.
//
// Solidity: function pancakeV3SwapCallback(int256 amount0, int256 amount1, bytes _data) returns()
func (_Ronswapv3fe *Ronswapv3feTransactor) PancakeV3SwapCallback(opts *bind.TransactOpts, amount0 *big.Int, amount1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _Ronswapv3fe.contract.Transact(opts, "pancakeV3SwapCallback", amount0, amount1, _data)
}

// PancakeV3SwapCallback is a paid mutator transaction binding the contract method 0x23a69e75.
//
// Solidity: function pancakeV3SwapCallback(int256 amount0, int256 amount1, bytes _data) returns()
func (_Ronswapv3fe *Ronswapv3feSession) PancakeV3SwapCallback(amount0 *big.Int, amount1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _Ronswapv3fe.Contract.PancakeV3SwapCallback(&_Ronswapv3fe.TransactOpts, amount0, amount1, _data)
}

// PancakeV3SwapCallback is a paid mutator transaction binding the contract method 0x23a69e75.
//
// Solidity: function pancakeV3SwapCallback(int256 amount0, int256 amount1, bytes _data) returns()
func (_Ronswapv3fe *Ronswapv3feTransactorSession) PancakeV3SwapCallback(amount0 *big.Int, amount1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _Ronswapv3fe.Contract.PancakeV3SwapCallback(&_Ronswapv3fe.TransactOpts, amount0, amount1, _data)
}

// RonSwapV3 is a paid mutator transaction binding the contract method 0x69bf3560.
//
// Solidity: function ronSwapV3(uint256[] tokenPairsAndFee, uint256 amountIn) returns(uint256 amountOut)
func (_Ronswapv3fe *Ronswapv3feTransactor) RonSwapV3(opts *bind.TransactOpts, tokenPairsAndFee []*big.Int, amountIn *big.Int) (*types.Transaction, error) {
	return _Ronswapv3fe.contract.Transact(opts, "ronSwapV3", tokenPairsAndFee, amountIn)
}

// RonSwapV3 is a paid mutator transaction binding the contract method 0x69bf3560.
//
// Solidity: function ronSwapV3(uint256[] tokenPairsAndFee, uint256 amountIn) returns(uint256 amountOut)
func (_Ronswapv3fe *Ronswapv3feSession) RonSwapV3(tokenPairsAndFee []*big.Int, amountIn *big.Int) (*types.Transaction, error) {
	return _Ronswapv3fe.Contract.RonSwapV3(&_Ronswapv3fe.TransactOpts, tokenPairsAndFee, amountIn)
}

// RonSwapV3 is a paid mutator transaction binding the contract method 0x69bf3560.
//
// Solidity: function ronSwapV3(uint256[] tokenPairsAndFee, uint256 amountIn) returns(uint256 amountOut)
func (_Ronswapv3fe *Ronswapv3feTransactorSession) RonSwapV3(tokenPairsAndFee []*big.Int, amountIn *big.Int) (*types.Transaction, error) {
	return _Ronswapv3fe.Contract.RonSwapV3(&_Ronswapv3fe.TransactOpts, tokenPairsAndFee, amountIn)
}

// UniswapV2Call is a paid mutator transaction binding the contract method 0x10d1e85c.
//
// Solidity: function uniswapV2Call(address sender, uint256 amount0, uint256 amount1, bytes _data) returns()
func (_Ronswapv3fe *Ronswapv3feTransactor) UniswapV2Call(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _Ronswapv3fe.contract.Transact(opts, "uniswapV2Call", sender, amount0, amount1, _data)
}

// UniswapV2Call is a paid mutator transaction binding the contract method 0x10d1e85c.
//
// Solidity: function uniswapV2Call(address sender, uint256 amount0, uint256 amount1, bytes _data) returns()
func (_Ronswapv3fe *Ronswapv3feSession) UniswapV2Call(sender common.Address, amount0 *big.Int, amount1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _Ronswapv3fe.Contract.UniswapV2Call(&_Ronswapv3fe.TransactOpts, sender, amount0, amount1, _data)
}

// UniswapV2Call is a paid mutator transaction binding the contract method 0x10d1e85c.
//
// Solidity: function uniswapV2Call(address sender, uint256 amount0, uint256 amount1, bytes _data) returns()
func (_Ronswapv3fe *Ronswapv3feTransactorSession) UniswapV2Call(sender common.Address, amount0 *big.Int, amount1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _Ronswapv3fe.Contract.UniswapV2Call(&_Ronswapv3fe.TransactOpts, sender, amount0, amount1, _data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0, int256 amount1, bytes _data) returns()
func (_Ronswapv3fe *Ronswapv3feTransactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0 *big.Int, amount1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _Ronswapv3fe.contract.Transact(opts, "uniswapV3SwapCallback", amount0, amount1, _data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0, int256 amount1, bytes _data) returns()
func (_Ronswapv3fe *Ronswapv3feSession) UniswapV3SwapCallback(amount0 *big.Int, amount1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _Ronswapv3fe.Contract.UniswapV3SwapCallback(&_Ronswapv3fe.TransactOpts, amount0, amount1, _data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0, int256 amount1, bytes _data) returns()
func (_Ronswapv3fe *Ronswapv3feTransactorSession) UniswapV3SwapCallback(amount0 *big.Int, amount1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _Ronswapv3fe.Contract.UniswapV3SwapCallback(&_Ronswapv3fe.TransactOpts, amount0, amount1, _data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 amount, address token) returns()
func (_Ronswapv3fe *Ronswapv3feTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int, token common.Address) (*types.Transaction, error) {
	return _Ronswapv3fe.contract.Transact(opts, "withdraw", amount, token)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 amount, address token) returns()
func (_Ronswapv3fe *Ronswapv3feSession) Withdraw(amount *big.Int, token common.Address) (*types.Transaction, error) {
	return _Ronswapv3fe.Contract.Withdraw(&_Ronswapv3fe.TransactOpts, amount, token)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 amount, address token) returns()
func (_Ronswapv3fe *Ronswapv3feTransactorSession) Withdraw(amount *big.Int, token common.Address) (*types.Transaction, error) {
	return _Ronswapv3fe.Contract.Withdraw(&_Ronswapv3fe.TransactOpts, amount, token)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ronswapv3fe *Ronswapv3feTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ronswapv3fe.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ronswapv3fe *Ronswapv3feSession) Receive() (*types.Transaction, error) {
	return _Ronswapv3fe.Contract.Receive(&_Ronswapv3fe.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ronswapv3fe *Ronswapv3feTransactorSession) Receive() (*types.Transaction, error) {
	return _Ronswapv3fe.Contract.Receive(&_Ronswapv3fe.TransactOpts)
}

// Ronswapv3feDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Ronswapv3fe contract.
type Ronswapv3feDepositedIterator struct {
	Event *Ronswapv3feDeposited // Event containing the contract specifics and raw log

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
func (it *Ronswapv3feDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Ronswapv3feDeposited)
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
		it.Event = new(Ronswapv3feDeposited)
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
func (it *Ronswapv3feDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Ronswapv3feDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Ronswapv3feDeposited represents a Deposited event raised by the Ronswapv3fe contract.
type Ronswapv3feDeposited struct {
	Charger     common.Address
	TotalOk     *big.Int
	TotalAmount *big.Int
	Failed      bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x65e6f89b6907d6277741ee4ccbc4ae260163a17f16bbb55b5344dff064867c40.
//
// Solidity: event Deposited(address indexed charger, uint256 totalOk, uint256 totalAmount, bool failed)
func (_Ronswapv3fe *Ronswapv3feFilterer) FilterDeposited(opts *bind.FilterOpts, charger []common.Address) (*Ronswapv3feDepositedIterator, error) {

	var chargerRule []interface{}
	for _, chargerItem := range charger {
		chargerRule = append(chargerRule, chargerItem)
	}

	logs, sub, err := _Ronswapv3fe.contract.FilterLogs(opts, "Deposited", chargerRule)
	if err != nil {
		return nil, err
	}
	return &Ronswapv3feDepositedIterator{contract: _Ronswapv3fe.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x65e6f89b6907d6277741ee4ccbc4ae260163a17f16bbb55b5344dff064867c40.
//
// Solidity: event Deposited(address indexed charger, uint256 totalOk, uint256 totalAmount, bool failed)
func (_Ronswapv3fe *Ronswapv3feFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *Ronswapv3feDeposited, charger []common.Address) (event.Subscription, error) {

	var chargerRule []interface{}
	for _, chargerItem := range charger {
		chargerRule = append(chargerRule, chargerItem)
	}

	logs, sub, err := _Ronswapv3fe.contract.WatchLogs(opts, "Deposited", chargerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Ronswapv3feDeposited)
				if err := _Ronswapv3fe.contract.UnpackLog(event, "Deposited", log); err != nil {
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
func (_Ronswapv3fe *Ronswapv3feFilterer) ParseDeposited(log types.Log) (*Ronswapv3feDeposited, error) {
	event := new(Ronswapv3feDeposited)
	if err := _Ronswapv3fe.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
