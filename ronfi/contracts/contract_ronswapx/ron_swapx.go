// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ronswapx

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

// RonswapxABI is the input ABI used to generate the binding from.
const RonswapxABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_executors\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_executors\",\"type\":\"address[]\"}],\"name\":\"addExecutor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_executors\",\"type\":\"address[]\"}],\"name\":\"delExecutor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"reserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"reserve1\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"blockTimestampLast\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"name\":\"getReservesEn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reserves\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"pairs\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"poolFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"}],\"name\":\"ronSwapD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"pairs\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"poolFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"}],\"name\":\"ronSwapL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"pairs\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"poolFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"}],\"name\":\"ronSwapX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// Ronswapx is an auto generated Go binding around an Ethereum contract.
type Ronswapx struct {
	RonswapxCaller     // Read-only binding to the contract
	RonswapxTransactor // Write-only binding to the contract
	RonswapxFilterer   // Log filterer for contract events
}

// RonswapxCaller is an auto generated read-only Go binding around an Ethereum contract.
type RonswapxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RonswapxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RonswapxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RonswapxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RonswapxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RonswapxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RonswapxSession struct {
	Contract     *Ronswapx         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RonswapxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RonswapxCallerSession struct {
	Contract *RonswapxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RonswapxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RonswapxTransactorSession struct {
	Contract     *RonswapxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RonswapxRaw is an auto generated low-level Go binding around an Ethereum contract.
type RonswapxRaw struct {
	Contract *Ronswapx // Generic contract binding to access the raw methods on
}

// RonswapxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RonswapxCallerRaw struct {
	Contract *RonswapxCaller // Generic read-only contract binding to access the raw methods on
}

// RonswapxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RonswapxTransactorRaw struct {
	Contract *RonswapxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRonswapx creates a new instance of Ronswapx, bound to a specific deployed contract.
func NewRonswapx(address common.Address, backend bind.ContractBackend) (*Ronswapx, error) {
	contract, err := bindRonswapx(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ronswapx{RonswapxCaller: RonswapxCaller{contract: contract}, RonswapxTransactor: RonswapxTransactor{contract: contract}, RonswapxFilterer: RonswapxFilterer{contract: contract}}, nil
}

// NewRonswapxCaller creates a new read-only instance of Ronswapx, bound to a specific deployed contract.
func NewRonswapxCaller(address common.Address, caller bind.ContractCaller) (*RonswapxCaller, error) {
	contract, err := bindRonswapx(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RonswapxCaller{contract: contract}, nil
}

// NewRonswapxTransactor creates a new write-only instance of Ronswapx, bound to a specific deployed contract.
func NewRonswapxTransactor(address common.Address, transactor bind.ContractTransactor) (*RonswapxTransactor, error) {
	contract, err := bindRonswapx(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RonswapxTransactor{contract: contract}, nil
}

// NewRonswapxFilterer creates a new log filterer instance of Ronswapx, bound to a specific deployed contract.
func NewRonswapxFilterer(address common.Address, filterer bind.ContractFilterer) (*RonswapxFilterer, error) {
	contract, err := bindRonswapx(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RonswapxFilterer{contract: contract}, nil
}

// bindRonswapx binds a generic wrapper to an already deployed contract.
func bindRonswapx(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RonswapxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ronswapx *RonswapxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ronswapx.Contract.RonswapxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ronswapx *RonswapxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ronswapx.Contract.RonswapxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ronswapx *RonswapxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ronswapx.Contract.RonswapxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ronswapx *RonswapxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ronswapx.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ronswapx *RonswapxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ronswapx.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ronswapx *RonswapxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ronswapx.Contract.contract.Transact(opts, method, params...)
}

// GetReserves is a free data retrieval call binding the contract method 0x3e99c1e4.
//
// Solidity: function getReserves(address pair) view returns(uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast)
func (_Ronswapx *RonswapxCaller) GetReserves(opts *bind.CallOpts, pair common.Address) (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	var out []interface{}
	err := _Ronswapx.contract.Call(opts, &out, "getReserves", pair)

	outstruct := new(struct {
		Reserve0           *big.Int
		Reserve1           *big.Int
		BlockTimestampLast uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BlockTimestampLast = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x3e99c1e4.
//
// Solidity: function getReserves(address pair) view returns(uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast)
func (_Ronswapx *RonswapxSession) GetReserves(pair common.Address) (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _Ronswapx.Contract.GetReserves(&_Ronswapx.CallOpts, pair)
}

// GetReserves is a free data retrieval call binding the contract method 0x3e99c1e4.
//
// Solidity: function getReserves(address pair) view returns(uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast)
func (_Ronswapx *RonswapxCallerSession) GetReserves(pair common.Address) (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _Ronswapx.Contract.GetReserves(&_Ronswapx.CallOpts, pair)
}

// GetReservesEn is a free data retrieval call binding the contract method 0xcd62a188.
//
// Solidity: function getReservesEn(address pair) view returns(uint256 reserves)
func (_Ronswapx *RonswapxCaller) GetReservesEn(opts *bind.CallOpts, pair common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ronswapx.contract.Call(opts, &out, "getReservesEn", pair)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReservesEn is a free data retrieval call binding the contract method 0xcd62a188.
//
// Solidity: function getReservesEn(address pair) view returns(uint256 reserves)
func (_Ronswapx *RonswapxSession) GetReservesEn(pair common.Address) (*big.Int, error) {
	return _Ronswapx.Contract.GetReservesEn(&_Ronswapx.CallOpts, pair)
}

// GetReservesEn is a free data retrieval call binding the contract method 0xcd62a188.
//
// Solidity: function getReservesEn(address pair) view returns(uint256 reserves)
func (_Ronswapx *RonswapxCallerSession) GetReservesEn(pair common.Address) (*big.Int, error) {
	return _Ronswapx.Contract.GetReservesEn(&_Ronswapx.CallOpts, pair)
}

// AddExecutor is a paid mutator transaction binding the contract method 0xee26468a.
//
// Solidity: function addExecutor(address[] _executors) returns()
func (_Ronswapx *RonswapxTransactor) AddExecutor(opts *bind.TransactOpts, _executors []common.Address) (*types.Transaction, error) {
	return _Ronswapx.contract.Transact(opts, "addExecutor", _executors)
}

// AddExecutor is a paid mutator transaction binding the contract method 0xee26468a.
//
// Solidity: function addExecutor(address[] _executors) returns()
func (_Ronswapx *RonswapxSession) AddExecutor(_executors []common.Address) (*types.Transaction, error) {
	return _Ronswapx.Contract.AddExecutor(&_Ronswapx.TransactOpts, _executors)
}

// AddExecutor is a paid mutator transaction binding the contract method 0xee26468a.
//
// Solidity: function addExecutor(address[] _executors) returns()
func (_Ronswapx *RonswapxTransactorSession) AddExecutor(_executors []common.Address) (*types.Transaction, error) {
	return _Ronswapx.Contract.AddExecutor(&_Ronswapx.TransactOpts, _executors)
}

// DelExecutor is a paid mutator transaction binding the contract method 0x30f9006a.
//
// Solidity: function delExecutor(address[] _executors) returns()
func (_Ronswapx *RonswapxTransactor) DelExecutor(opts *bind.TransactOpts, _executors []common.Address) (*types.Transaction, error) {
	return _Ronswapx.contract.Transact(opts, "delExecutor", _executors)
}

// DelExecutor is a paid mutator transaction binding the contract method 0x30f9006a.
//
// Solidity: function delExecutor(address[] _executors) returns()
func (_Ronswapx *RonswapxSession) DelExecutor(_executors []common.Address) (*types.Transaction, error) {
	return _Ronswapx.Contract.DelExecutor(&_Ronswapx.TransactOpts, _executors)
}

// DelExecutor is a paid mutator transaction binding the contract method 0x30f9006a.
//
// Solidity: function delExecutor(address[] _executors) returns()
func (_Ronswapx *RonswapxTransactorSession) DelExecutor(_executors []common.Address) (*types.Transaction, error) {
	return _Ronswapx.Contract.DelExecutor(&_Ronswapx.TransactOpts, _executors)
}

// RonSwapD is a paid mutator transaction binding the contract method 0x34ec0942.
//
// Solidity: function ronSwapD(address[] pairs, uint256 poolFees, uint256 amountIn, uint256 amountOutMin) returns(uint256 amountOut)
func (_Ronswapx *RonswapxTransactor) RonSwapD(opts *bind.TransactOpts, pairs []common.Address, poolFees *big.Int, amountIn *big.Int, amountOutMin *big.Int) (*types.Transaction, error) {
	return _Ronswapx.contract.Transact(opts, "ronSwapD", pairs, poolFees, amountIn, amountOutMin)
}

// RonSwapD is a paid mutator transaction binding the contract method 0x34ec0942.
//
// Solidity: function ronSwapD(address[] pairs, uint256 poolFees, uint256 amountIn, uint256 amountOutMin) returns(uint256 amountOut)
func (_Ronswapx *RonswapxSession) RonSwapD(pairs []common.Address, poolFees *big.Int, amountIn *big.Int, amountOutMin *big.Int) (*types.Transaction, error) {
	return _Ronswapx.Contract.RonSwapD(&_Ronswapx.TransactOpts, pairs, poolFees, amountIn, amountOutMin)
}

// RonSwapD is a paid mutator transaction binding the contract method 0x34ec0942.
//
// Solidity: function ronSwapD(address[] pairs, uint256 poolFees, uint256 amountIn, uint256 amountOutMin) returns(uint256 amountOut)
func (_Ronswapx *RonswapxTransactorSession) RonSwapD(pairs []common.Address, poolFees *big.Int, amountIn *big.Int, amountOutMin *big.Int) (*types.Transaction, error) {
	return _Ronswapx.Contract.RonSwapD(&_Ronswapx.TransactOpts, pairs, poolFees, amountIn, amountOutMin)
}

// RonSwapL is a paid mutator transaction binding the contract method 0x0b1407a8.
//
// Solidity: function ronSwapL(address[] pairs, uint256 poolFees, uint256 amountIn, uint256 amountOutMin) returns(uint256 amountOut)
func (_Ronswapx *RonswapxTransactor) RonSwapL(opts *bind.TransactOpts, pairs []common.Address, poolFees *big.Int, amountIn *big.Int, amountOutMin *big.Int) (*types.Transaction, error) {
	return _Ronswapx.contract.Transact(opts, "ronSwapL", pairs, poolFees, amountIn, amountOutMin)
}

// RonSwapL is a paid mutator transaction binding the contract method 0x0b1407a8.
//
// Solidity: function ronSwapL(address[] pairs, uint256 poolFees, uint256 amountIn, uint256 amountOutMin) returns(uint256 amountOut)
func (_Ronswapx *RonswapxSession) RonSwapL(pairs []common.Address, poolFees *big.Int, amountIn *big.Int, amountOutMin *big.Int) (*types.Transaction, error) {
	return _Ronswapx.Contract.RonSwapL(&_Ronswapx.TransactOpts, pairs, poolFees, amountIn, amountOutMin)
}

// RonSwapL is a paid mutator transaction binding the contract method 0x0b1407a8.
//
// Solidity: function ronSwapL(address[] pairs, uint256 poolFees, uint256 amountIn, uint256 amountOutMin) returns(uint256 amountOut)
func (_Ronswapx *RonswapxTransactorSession) RonSwapL(pairs []common.Address, poolFees *big.Int, amountIn *big.Int, amountOutMin *big.Int) (*types.Transaction, error) {
	return _Ronswapx.Contract.RonSwapL(&_Ronswapx.TransactOpts, pairs, poolFees, amountIn, amountOutMin)
}

// RonSwapX is a paid mutator transaction binding the contract method 0x6326d217.
//
// Solidity: function ronSwapX(address[] pairs, uint256 poolFees, uint256 amountIn, uint256 amountOutMin) returns(uint256 amountOut)
func (_Ronswapx *RonswapxTransactor) RonSwapX(opts *bind.TransactOpts, pairs []common.Address, poolFees *big.Int, amountIn *big.Int, amountOutMin *big.Int) (*types.Transaction, error) {
	return _Ronswapx.contract.Transact(opts, "ronSwapX", pairs, poolFees, amountIn, amountOutMin)
}

// RonSwapX is a paid mutator transaction binding the contract method 0x6326d217.
//
// Solidity: function ronSwapX(address[] pairs, uint256 poolFees, uint256 amountIn, uint256 amountOutMin) returns(uint256 amountOut)
func (_Ronswapx *RonswapxSession) RonSwapX(pairs []common.Address, poolFees *big.Int, amountIn *big.Int, amountOutMin *big.Int) (*types.Transaction, error) {
	return _Ronswapx.Contract.RonSwapX(&_Ronswapx.TransactOpts, pairs, poolFees, amountIn, amountOutMin)
}

// RonSwapX is a paid mutator transaction binding the contract method 0x6326d217.
//
// Solidity: function ronSwapX(address[] pairs, uint256 poolFees, uint256 amountIn, uint256 amountOutMin) returns(uint256 amountOut)
func (_Ronswapx *RonswapxTransactorSession) RonSwapX(pairs []common.Address, poolFees *big.Int, amountIn *big.Int, amountOutMin *big.Int) (*types.Transaction, error) {
	return _Ronswapx.Contract.RonSwapX(&_Ronswapx.TransactOpts, pairs, poolFees, amountIn, amountOutMin)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 amount, address token) returns()
func (_Ronswapx *RonswapxTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int, token common.Address) (*types.Transaction, error) {
	return _Ronswapx.contract.Transact(opts, "withdraw", amount, token)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 amount, address token) returns()
func (_Ronswapx *RonswapxSession) Withdraw(amount *big.Int, token common.Address) (*types.Transaction, error) {
	return _Ronswapx.Contract.Withdraw(&_Ronswapx.TransactOpts, amount, token)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 amount, address token) returns()
func (_Ronswapx *RonswapxTransactorSession) Withdraw(amount *big.Int, token common.Address) (*types.Transaction, error) {
	return _Ronswapx.Contract.Withdraw(&_Ronswapx.TransactOpts, amount, token)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ronswapx *RonswapxTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ronswapx.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ronswapx *RonswapxSession) Receive() (*types.Transaction, error) {
	return _Ronswapx.Contract.Receive(&_Ronswapx.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ronswapx *RonswapxTransactorSession) Receive() (*types.Transaction, error) {
	return _Ronswapx.Contract.Receive(&_Ronswapx.TransactOpts)
}
