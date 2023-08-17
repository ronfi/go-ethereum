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

// RonExecutorEPayload is an auto generated low-level Go binding around an user-defined struct.
type RonExecutorEPayload struct {
	Target common.Address
	Data   []byte
	Value  *big.Int
}

// Ronswapv3feMetaData contains all meta data concerning the Ronswapv3fe contract.
var Ronswapv3feMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_executors\",\"type\":\"address[]\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_executors\",\"type\":\"address[]\"}],\"name\":\"addExecutor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_executors\",\"type\":\"address[]\"}],\"name\":\"delExecutor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structRonExecutorE.Payload[]\",\"name\":\"payloads\",\"type\":\"tuple[]\"}],\"name\":\"executePayloads\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
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

// ExecutePayloads is a paid mutator transaction binding the contract method 0x2685cff5.
//
// Solidity: function executePayloads((address,bytes,uint256)[] payloads) payable returns()
func (_Ronswapv3fe *Ronswapv3feTransactor) ExecutePayloads(opts *bind.TransactOpts, payloads []RonExecutorEPayload) (*types.Transaction, error) {
	return _Ronswapv3fe.contract.Transact(opts, "executePayloads", payloads)
}

// ExecutePayloads is a paid mutator transaction binding the contract method 0x2685cff5.
//
// Solidity: function executePayloads((address,bytes,uint256)[] payloads) payable returns()
func (_Ronswapv3fe *Ronswapv3feSession) ExecutePayloads(payloads []RonExecutorEPayload) (*types.Transaction, error) {
	return _Ronswapv3fe.Contract.ExecutePayloads(&_Ronswapv3fe.TransactOpts, payloads)
}

// ExecutePayloads is a paid mutator transaction binding the contract method 0x2685cff5.
//
// Solidity: function executePayloads((address,bytes,uint256)[] payloads) payable returns()
func (_Ronswapv3fe *Ronswapv3feTransactorSession) ExecutePayloads(payloads []RonExecutorEPayload) (*types.Transaction, error) {
	return _Ronswapv3fe.Contract.ExecutePayloads(&_Ronswapv3fe.TransactOpts, payloads)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0, int256 amount1, bytes ) returns()
func (_Ronswapv3fe *Ronswapv3feTransactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0 *big.Int, amount1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Ronswapv3fe.contract.Transact(opts, "uniswapV3SwapCallback", amount0, amount1, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0, int256 amount1, bytes ) returns()
func (_Ronswapv3fe *Ronswapv3feSession) UniswapV3SwapCallback(amount0 *big.Int, amount1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Ronswapv3fe.Contract.UniswapV3SwapCallback(&_Ronswapv3fe.TransactOpts, amount0, amount1, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0, int256 amount1, bytes ) returns()
func (_Ronswapv3fe *Ronswapv3feTransactorSession) UniswapV3SwapCallback(amount0 *big.Int, amount1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Ronswapv3fe.Contract.UniswapV3SwapCallback(&_Ronswapv3fe.TransactOpts, amount0, amount1, arg2)
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
