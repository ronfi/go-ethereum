// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ronswapz

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// RonswapzABI is the input ABI used to generate the binding from.
const RonswapzABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_executors\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_executors\",\"type\":\"address[]\"}],\"name\":\"addExecutor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_executors\",\"type\":\"address[]\"}],\"name\":\"delExecutor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenPairsAndFee\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"profitMin\",\"type\":\"uint256\"}],\"name\":\"ronSwapAuto\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenPairsAndFee\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"profitMin\",\"type\":\"uint256\"}],\"name\":\"ronSwapD\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// Ronswapz is an auto generated Go binding around an Ethereum contract.
type Ronswapz struct {
	RonswapzCaller     // Read-only binding to the contract
	RonswapzTransactor // Write-only binding to the contract
	RonswapzFilterer   // Log filterer for contract events
}

// RonswapzCaller is an auto generated read-only Go binding around an Ethereum contract.
type RonswapzCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RonswapzTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RonswapzTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RonswapzFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RonswapzFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RonswapzSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RonswapzSession struct {
	Contract     *Ronswapz         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RonswapzCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RonswapzCallerSession struct {
	Contract *RonswapzCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RonswapzTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RonswapzTransactorSession struct {
	Contract     *RonswapzTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RonswapzRaw is an auto generated low-level Go binding around an Ethereum contract.
type RonswapzRaw struct {
	Contract *Ronswapz // Generic contract binding to access the raw methods on
}

// RonswapzCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RonswapzCallerRaw struct {
	Contract *RonswapzCaller // Generic read-only contract binding to access the raw methods on
}

// RonswapzTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RonswapzTransactorRaw struct {
	Contract *RonswapzTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRonswapz creates a new instance of Ronswapz, bound to a specific deployed contract.
func NewRonswapz(address common.Address, backend bind.ContractBackend) (*Ronswapz, error) {
	contract, err := bindRonswapz(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ronswapz{RonswapzCaller: RonswapzCaller{contract: contract}, RonswapzTransactor: RonswapzTransactor{contract: contract}, RonswapzFilterer: RonswapzFilterer{contract: contract}}, nil
}

// NewRonswapzCaller creates a new read-only instance of Ronswapz, bound to a specific deployed contract.
func NewRonswapzCaller(address common.Address, caller bind.ContractCaller) (*RonswapzCaller, error) {
	contract, err := bindRonswapz(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RonswapzCaller{contract: contract}, nil
}

// NewRonswapzTransactor creates a new write-only instance of Ronswapz, bound to a specific deployed contract.
func NewRonswapzTransactor(address common.Address, transactor bind.ContractTransactor) (*RonswapzTransactor, error) {
	contract, err := bindRonswapz(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RonswapzTransactor{contract: contract}, nil
}

// NewRonswapzFilterer creates a new log filterer instance of Ronswapz, bound to a specific deployed contract.
func NewRonswapzFilterer(address common.Address, filterer bind.ContractFilterer) (*RonswapzFilterer, error) {
	contract, err := bindRonswapz(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RonswapzFilterer{contract: contract}, nil
}

// bindRonswapz binds a generic wrapper to an already deployed contract.
func bindRonswapz(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RonswapzABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ronswapz *RonswapzRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ronswapz.Contract.RonswapzCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ronswapz *RonswapzRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ronswapz.Contract.RonswapzTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ronswapz *RonswapzRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ronswapz.Contract.RonswapzTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ronswapz *RonswapzCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ronswapz.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ronswapz *RonswapzTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ronswapz.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ronswapz *RonswapzTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ronswapz.Contract.contract.Transact(opts, method, params...)
}

// AddExecutor is a paid mutator transaction binding the contract method 0xee26468a.
//
// Solidity: function addExecutor(address[] _executors) returns()
func (_Ronswapz *RonswapzTransactor) AddExecutor(opts *bind.TransactOpts, _executors []common.Address) (*types.Transaction, error) {
	return _Ronswapz.contract.Transact(opts, "addExecutor", _executors)
}

// AddExecutor is a paid mutator transaction binding the contract method 0xee26468a.
//
// Solidity: function addExecutor(address[] _executors) returns()
func (_Ronswapz *RonswapzSession) AddExecutor(_executors []common.Address) (*types.Transaction, error) {
	return _Ronswapz.Contract.AddExecutor(&_Ronswapz.TransactOpts, _executors)
}

// AddExecutor is a paid mutator transaction binding the contract method 0xee26468a.
//
// Solidity: function addExecutor(address[] _executors) returns()
func (_Ronswapz *RonswapzTransactorSession) AddExecutor(_executors []common.Address) (*types.Transaction, error) {
	return _Ronswapz.Contract.AddExecutor(&_Ronswapz.TransactOpts, _executors)
}

// DelExecutor is a paid mutator transaction binding the contract method 0x30f9006a.
//
// Solidity: function delExecutor(address[] _executors) returns()
func (_Ronswapz *RonswapzTransactor) DelExecutor(opts *bind.TransactOpts, _executors []common.Address) (*types.Transaction, error) {
	return _Ronswapz.contract.Transact(opts, "delExecutor", _executors)
}

// DelExecutor is a paid mutator transaction binding the contract method 0x30f9006a.
//
// Solidity: function delExecutor(address[] _executors) returns()
func (_Ronswapz *RonswapzSession) DelExecutor(_executors []common.Address) (*types.Transaction, error) {
	return _Ronswapz.Contract.DelExecutor(&_Ronswapz.TransactOpts, _executors)
}

// DelExecutor is a paid mutator transaction binding the contract method 0x30f9006a.
//
// Solidity: function delExecutor(address[] _executors) returns()
func (_Ronswapz *RonswapzTransactorSession) DelExecutor(_executors []common.Address) (*types.Transaction, error) {
	return _Ronswapz.Contract.DelExecutor(&_Ronswapz.TransactOpts, _executors)
}

// RonSwapAuto is a paid mutator transaction binding the contract method 0xfe5856bb.
//
// Solidity: function ronSwapAuto(uint256[] tokenPairsAndFee, uint256 profitMin) returns()
func (_Ronswapz *RonswapzTransactor) RonSwapAuto(opts *bind.TransactOpts, tokenPairsAndFee []*big.Int, profitMin *big.Int) (*types.Transaction, error) {
	return _Ronswapz.contract.Transact(opts, "ronSwapAuto", tokenPairsAndFee, profitMin)
}

// RonSwapAuto is a paid mutator transaction binding the contract method 0xfe5856bb.
//
// Solidity: function ronSwapAuto(uint256[] tokenPairsAndFee, uint256 profitMin) returns()
func (_Ronswapz *RonswapzSession) RonSwapAuto(tokenPairsAndFee []*big.Int, profitMin *big.Int) (*types.Transaction, error) {
	return _Ronswapz.Contract.RonSwapAuto(&_Ronswapz.TransactOpts, tokenPairsAndFee, profitMin)
}

// RonSwapAuto is a paid mutator transaction binding the contract method 0xfe5856bb.
//
// Solidity: function ronSwapAuto(uint256[] tokenPairsAndFee, uint256 profitMin) returns()
func (_Ronswapz *RonswapzTransactorSession) RonSwapAuto(tokenPairsAndFee []*big.Int, profitMin *big.Int) (*types.Transaction, error) {
	return _Ronswapz.Contract.RonSwapAuto(&_Ronswapz.TransactOpts, tokenPairsAndFee, profitMin)
}

// RonSwapD is a paid mutator transaction binding the contract method 0x6a610127.
//
// Solidity: function ronSwapD(uint256[] tokenPairsAndFee, uint256 profitMin) returns()
func (_Ronswapz *RonswapzTransactor) RonSwapD(opts *bind.TransactOpts, tokenPairsAndFee []*big.Int, profitMin *big.Int) (*types.Transaction, error) {
	return _Ronswapz.contract.Transact(opts, "ronSwapD", tokenPairsAndFee, profitMin)
}

// RonSwapD is a paid mutator transaction binding the contract method 0x6a610127.
//
// Solidity: function ronSwapD(uint256[] tokenPairsAndFee, uint256 profitMin) returns()
func (_Ronswapz *RonswapzSession) RonSwapD(tokenPairsAndFee []*big.Int, profitMin *big.Int) (*types.Transaction, error) {
	return _Ronswapz.Contract.RonSwapD(&_Ronswapz.TransactOpts, tokenPairsAndFee, profitMin)
}

// RonSwapD is a paid mutator transaction binding the contract method 0x6a610127.
//
// Solidity: function ronSwapD(uint256[] tokenPairsAndFee, uint256 profitMin) returns()
func (_Ronswapz *RonswapzTransactorSession) RonSwapD(tokenPairsAndFee []*big.Int, profitMin *big.Int) (*types.Transaction, error) {
	return _Ronswapz.Contract.RonSwapD(&_Ronswapz.TransactOpts, tokenPairsAndFee, profitMin)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 amount, address token) returns()
func (_Ronswapz *RonswapzTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int, token common.Address) (*types.Transaction, error) {
	return _Ronswapz.contract.Transact(opts, "withdraw", amount, token)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 amount, address token) returns()
func (_Ronswapz *RonswapzSession) Withdraw(amount *big.Int, token common.Address) (*types.Transaction, error) {
	return _Ronswapz.Contract.Withdraw(&_Ronswapz.TransactOpts, amount, token)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 amount, address token) returns()
func (_Ronswapz *RonswapzTransactorSession) Withdraw(amount *big.Int, token common.Address) (*types.Transaction, error) {
	return _Ronswapz.Contract.Withdraw(&_Ronswapz.TransactOpts, amount, token)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ronswapz *RonswapzTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ronswapz.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ronswapz *RonswapzSession) Receive() (*types.Transaction, error) {
	return _Ronswapz.Contract.Receive(&_Ronswapz.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ronswapz *RonswapzTransactorSession) Receive() (*types.Transaction, error) {
	return _Ronswapz.Contract.Receive(&_Ronswapz.TransactOpts)
}
