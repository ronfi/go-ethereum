// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dmm

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

// DmmMetaData contains all meta data concerning the Dmm contract.
var DmmMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeInPrecision\",\"type\":\"uint256\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vReserve0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vReserve1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reserve0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reserve1\",\"type\":\"uint256\"}],\"name\":\"Sync\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shortEMA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"longEMA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"lastBlockVolume\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"skipBlock\",\"type\":\"uint256\"}],\"name\":\"UpdateEMA\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MINIMUM_LIQUIDITY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ampBps\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"contractIDMMFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"_reserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"_reserve1\",\"type\":\"uint112\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTradeInfo\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"_reserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"_reserve1\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"_vReserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"_vReserve1\",\"type\":\"uint112\"},{\"internalType\":\"uint256\",\"name\":\"feeInPrecision\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVolumeTrendData\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"_shortEMA\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_longEMA\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_currentBlockVolume\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_lastTradeBlock\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_token0\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_token1\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_ampBps\",\"type\":\"uint32\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"skim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callbackData\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sync\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DmmABI is the input ABI used to generate the binding from.
// Deprecated: Use DmmMetaData.ABI instead.
var DmmABI = DmmMetaData.ABI

// Dmm is an auto generated Go binding around an Ethereum contract.
type Dmm struct {
	DmmCaller     // Read-only binding to the contract
	DmmTransactor // Write-only binding to the contract
	DmmFilterer   // Log filterer for contract events
}

// DmmCaller is an auto generated read-only Go binding around an Ethereum contract.
type DmmCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DmmTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DmmTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DmmFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DmmFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DmmSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DmmSession struct {
	Contract     *Dmm              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DmmCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DmmCallerSession struct {
	Contract *DmmCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DmmTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DmmTransactorSession struct {
	Contract     *DmmTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DmmRaw is an auto generated low-level Go binding around an Ethereum contract.
type DmmRaw struct {
	Contract *Dmm // Generic contract binding to access the raw methods on
}

// DmmCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DmmCallerRaw struct {
	Contract *DmmCaller // Generic read-only contract binding to access the raw methods on
}

// DmmTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DmmTransactorRaw struct {
	Contract *DmmTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDmm creates a new instance of Dmm, bound to a specific deployed contract.
func NewDmm(address common.Address, backend bind.ContractBackend) (*Dmm, error) {
	contract, err := bindDmm(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dmm{DmmCaller: DmmCaller{contract: contract}, DmmTransactor: DmmTransactor{contract: contract}, DmmFilterer: DmmFilterer{contract: contract}}, nil
}

// NewDmmCaller creates a new read-only instance of Dmm, bound to a specific deployed contract.
func NewDmmCaller(address common.Address, caller bind.ContractCaller) (*DmmCaller, error) {
	contract, err := bindDmm(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DmmCaller{contract: contract}, nil
}

// NewDmmTransactor creates a new write-only instance of Dmm, bound to a specific deployed contract.
func NewDmmTransactor(address common.Address, transactor bind.ContractTransactor) (*DmmTransactor, error) {
	contract, err := bindDmm(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DmmTransactor{contract: contract}, nil
}

// NewDmmFilterer creates a new log filterer instance of Dmm, bound to a specific deployed contract.
func NewDmmFilterer(address common.Address, filterer bind.ContractFilterer) (*DmmFilterer, error) {
	contract, err := bindDmm(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DmmFilterer{contract: contract}, nil
}

// bindDmm binds a generic wrapper to an already deployed contract.
func bindDmm(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DmmABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dmm *DmmRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dmm.Contract.DmmCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dmm *DmmRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dmm.Contract.DmmTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dmm *DmmRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dmm.Contract.DmmTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dmm *DmmCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dmm.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dmm *DmmTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dmm.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dmm *DmmTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dmm.Contract.contract.Transact(opts, method, params...)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_Dmm *DmmCaller) MINIMUMLIQUIDITY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dmm.contract.Call(opts, &out, "MINIMUM_LIQUIDITY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_Dmm *DmmSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _Dmm.Contract.MINIMUMLIQUIDITY(&_Dmm.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_Dmm *DmmCallerSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _Dmm.Contract.MINIMUMLIQUIDITY(&_Dmm.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Dmm *DmmCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Dmm.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Dmm *DmmSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Dmm.Contract.PERMITTYPEHASH(&_Dmm.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Dmm *DmmCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Dmm.Contract.PERMITTYPEHASH(&_Dmm.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Dmm *DmmCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Dmm.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Dmm *DmmSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Dmm.Contract.Allowance(&_Dmm.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Dmm *DmmCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Dmm.Contract.Allowance(&_Dmm.CallOpts, owner, spender)
}

// AmpBps is a free data retrieval call binding the contract method 0x49386b16.
//
// Solidity: function ampBps() view returns(uint32)
func (_Dmm *DmmCaller) AmpBps(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Dmm.contract.Call(opts, &out, "ampBps")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// AmpBps is a free data retrieval call binding the contract method 0x49386b16.
//
// Solidity: function ampBps() view returns(uint32)
func (_Dmm *DmmSession) AmpBps() (uint32, error) {
	return _Dmm.Contract.AmpBps(&_Dmm.CallOpts)
}

// AmpBps is a free data retrieval call binding the contract method 0x49386b16.
//
// Solidity: function ampBps() view returns(uint32)
func (_Dmm *DmmCallerSession) AmpBps() (uint32, error) {
	return _Dmm.Contract.AmpBps(&_Dmm.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Dmm *DmmCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Dmm.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Dmm *DmmSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Dmm.Contract.BalanceOf(&_Dmm.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Dmm *DmmCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Dmm.Contract.BalanceOf(&_Dmm.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Dmm *DmmCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Dmm.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Dmm *DmmSession) Decimals() (uint8, error) {
	return _Dmm.Contract.Decimals(&_Dmm.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Dmm *DmmCallerSession) Decimals() (uint8, error) {
	return _Dmm.Contract.Decimals(&_Dmm.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_Dmm *DmmCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Dmm.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_Dmm *DmmSession) DomainSeparator() ([32]byte, error) {
	return _Dmm.Contract.DomainSeparator(&_Dmm.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_Dmm *DmmCallerSession) DomainSeparator() ([32]byte, error) {
	return _Dmm.Contract.DomainSeparator(&_Dmm.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Dmm *DmmCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dmm.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Dmm *DmmSession) Factory() (common.Address, error) {
	return _Dmm.Contract.Factory(&_Dmm.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Dmm *DmmCallerSession) Factory() (common.Address, error) {
	return _Dmm.Contract.Factory(&_Dmm.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1)
func (_Dmm *DmmCaller) GetReserves(opts *bind.CallOpts) (struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
}, error) {
	var out []interface{}
	err := _Dmm.contract.Call(opts, &out, "getReserves")

	outstruct := new(struct {
		Reserve0 *big.Int
		Reserve1 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1)
func (_Dmm *DmmSession) GetReserves() (struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
}, error) {
	return _Dmm.Contract.GetReserves(&_Dmm.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1)
func (_Dmm *DmmCallerSession) GetReserves() (struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
}, error) {
	return _Dmm.Contract.GetReserves(&_Dmm.CallOpts)
}

// GetTradeInfo is a free data retrieval call binding the contract method 0xd6694027.
//
// Solidity: function getTradeInfo() view returns(uint112 _reserve0, uint112 _reserve1, uint112 _vReserve0, uint112 _vReserve1, uint256 feeInPrecision)
func (_Dmm *DmmCaller) GetTradeInfo(opts *bind.CallOpts) (struct {
	Reserve0       *big.Int
	Reserve1       *big.Int
	VReserve0      *big.Int
	VReserve1      *big.Int
	FeeInPrecision *big.Int
}, error) {
	var out []interface{}
	err := _Dmm.contract.Call(opts, &out, "getTradeInfo")

	outstruct := new(struct {
		Reserve0       *big.Int
		Reserve1       *big.Int
		VReserve0      *big.Int
		VReserve1      *big.Int
		FeeInPrecision *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.VReserve0 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.VReserve1 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.FeeInPrecision = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetTradeInfo is a free data retrieval call binding the contract method 0xd6694027.
//
// Solidity: function getTradeInfo() view returns(uint112 _reserve0, uint112 _reserve1, uint112 _vReserve0, uint112 _vReserve1, uint256 feeInPrecision)
func (_Dmm *DmmSession) GetTradeInfo() (struct {
	Reserve0       *big.Int
	Reserve1       *big.Int
	VReserve0      *big.Int
	VReserve1      *big.Int
	FeeInPrecision *big.Int
}, error) {
	return _Dmm.Contract.GetTradeInfo(&_Dmm.CallOpts)
}

// GetTradeInfo is a free data retrieval call binding the contract method 0xd6694027.
//
// Solidity: function getTradeInfo() view returns(uint112 _reserve0, uint112 _reserve1, uint112 _vReserve0, uint112 _vReserve1, uint256 feeInPrecision)
func (_Dmm *DmmCallerSession) GetTradeInfo() (struct {
	Reserve0       *big.Int
	Reserve1       *big.Int
	VReserve0      *big.Int
	VReserve1      *big.Int
	FeeInPrecision *big.Int
}, error) {
	return _Dmm.Contract.GetTradeInfo(&_Dmm.CallOpts)
}

// GetVolumeTrendData is a free data retrieval call binding the contract method 0x0d94d50b.
//
// Solidity: function getVolumeTrendData() view returns(uint128 _shortEMA, uint128 _longEMA, uint128 _currentBlockVolume, uint128 _lastTradeBlock)
func (_Dmm *DmmCaller) GetVolumeTrendData(opts *bind.CallOpts) (struct {
	ShortEMA           *big.Int
	LongEMA            *big.Int
	CurrentBlockVolume *big.Int
	LastTradeBlock     *big.Int
}, error) {
	var out []interface{}
	err := _Dmm.contract.Call(opts, &out, "getVolumeTrendData")

	outstruct := new(struct {
		ShortEMA           *big.Int
		LongEMA            *big.Int
		CurrentBlockVolume *big.Int
		LastTradeBlock     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ShortEMA = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LongEMA = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CurrentBlockVolume = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LastTradeBlock = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetVolumeTrendData is a free data retrieval call binding the contract method 0x0d94d50b.
//
// Solidity: function getVolumeTrendData() view returns(uint128 _shortEMA, uint128 _longEMA, uint128 _currentBlockVolume, uint128 _lastTradeBlock)
func (_Dmm *DmmSession) GetVolumeTrendData() (struct {
	ShortEMA           *big.Int
	LongEMA            *big.Int
	CurrentBlockVolume *big.Int
	LastTradeBlock     *big.Int
}, error) {
	return _Dmm.Contract.GetVolumeTrendData(&_Dmm.CallOpts)
}

// GetVolumeTrendData is a free data retrieval call binding the contract method 0x0d94d50b.
//
// Solidity: function getVolumeTrendData() view returns(uint128 _shortEMA, uint128 _longEMA, uint128 _currentBlockVolume, uint128 _lastTradeBlock)
func (_Dmm *DmmCallerSession) GetVolumeTrendData() (struct {
	ShortEMA           *big.Int
	LongEMA            *big.Int
	CurrentBlockVolume *big.Int
	LastTradeBlock     *big.Int
}, error) {
	return _Dmm.Contract.GetVolumeTrendData(&_Dmm.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_Dmm *DmmCaller) KLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dmm.contract.Call(opts, &out, "kLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_Dmm *DmmSession) KLast() (*big.Int, error) {
	return _Dmm.Contract.KLast(&_Dmm.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_Dmm *DmmCallerSession) KLast() (*big.Int, error) {
	return _Dmm.Contract.KLast(&_Dmm.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Dmm *DmmCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Dmm.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Dmm *DmmSession) Name() (string, error) {
	return _Dmm.Contract.Name(&_Dmm.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Dmm *DmmCallerSession) Name() (string, error) {
	return _Dmm.Contract.Name(&_Dmm.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Dmm *DmmCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Dmm.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Dmm *DmmSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Dmm.Contract.Nonces(&_Dmm.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Dmm *DmmCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Dmm.Contract.Nonces(&_Dmm.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Dmm *DmmCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Dmm.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Dmm *DmmSession) Symbol() (string, error) {
	return _Dmm.Contract.Symbol(&_Dmm.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Dmm *DmmCallerSession) Symbol() (string, error) {
	return _Dmm.Contract.Symbol(&_Dmm.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Dmm *DmmCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dmm.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Dmm *DmmSession) Token0() (common.Address, error) {
	return _Dmm.Contract.Token0(&_Dmm.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Dmm *DmmCallerSession) Token0() (common.Address, error) {
	return _Dmm.Contract.Token0(&_Dmm.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Dmm *DmmCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dmm.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Dmm *DmmSession) Token1() (common.Address, error) {
	return _Dmm.Contract.Token1(&_Dmm.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Dmm *DmmCallerSession) Token1() (common.Address, error) {
	return _Dmm.Contract.Token1(&_Dmm.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Dmm *DmmCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dmm.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Dmm *DmmSession) TotalSupply() (*big.Int, error) {
	return _Dmm.Contract.TotalSupply(&_Dmm.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Dmm *DmmCallerSession) TotalSupply() (*big.Int, error) {
	return _Dmm.Contract.TotalSupply(&_Dmm.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Dmm *DmmTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dmm.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Dmm *DmmSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dmm.Contract.Approve(&_Dmm.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Dmm *DmmTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dmm.Contract.Approve(&_Dmm.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_Dmm *DmmTransactor) Burn(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Dmm.contract.Transact(opts, "burn", to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_Dmm *DmmSession) Burn(to common.Address) (*types.Transaction, error) {
	return _Dmm.Contract.Burn(&_Dmm.TransactOpts, to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_Dmm *DmmTransactorSession) Burn(to common.Address) (*types.Transaction, error) {
	return _Dmm.Contract.Burn(&_Dmm.TransactOpts, to)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Dmm *DmmTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Dmm.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Dmm *DmmSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Dmm.Contract.DecreaseAllowance(&_Dmm.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Dmm *DmmTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Dmm.Contract.DecreaseAllowance(&_Dmm.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Dmm *DmmTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Dmm.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Dmm *DmmSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Dmm.Contract.IncreaseAllowance(&_Dmm.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Dmm *DmmTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Dmm.Contract.IncreaseAllowance(&_Dmm.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x6ecf2b22.
//
// Solidity: function initialize(address _token0, address _token1, uint32 _ampBps) returns()
func (_Dmm *DmmTransactor) Initialize(opts *bind.TransactOpts, _token0 common.Address, _token1 common.Address, _ampBps uint32) (*types.Transaction, error) {
	return _Dmm.contract.Transact(opts, "initialize", _token0, _token1, _ampBps)
}

// Initialize is a paid mutator transaction binding the contract method 0x6ecf2b22.
//
// Solidity: function initialize(address _token0, address _token1, uint32 _ampBps) returns()
func (_Dmm *DmmSession) Initialize(_token0 common.Address, _token1 common.Address, _ampBps uint32) (*types.Transaction, error) {
	return _Dmm.Contract.Initialize(&_Dmm.TransactOpts, _token0, _token1, _ampBps)
}

// Initialize is a paid mutator transaction binding the contract method 0x6ecf2b22.
//
// Solidity: function initialize(address _token0, address _token1, uint32 _ampBps) returns()
func (_Dmm *DmmTransactorSession) Initialize(_token0 common.Address, _token1 common.Address, _ampBps uint32) (*types.Transaction, error) {
	return _Dmm.Contract.Initialize(&_Dmm.TransactOpts, _token0, _token1, _ampBps)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_Dmm *DmmTransactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Dmm.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_Dmm *DmmSession) Mint(to common.Address) (*types.Transaction, error) {
	return _Dmm.Contract.Mint(&_Dmm.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_Dmm *DmmTransactorSession) Mint(to common.Address) (*types.Transaction, error) {
	return _Dmm.Contract.Mint(&_Dmm.TransactOpts, to)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Dmm *DmmTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Dmm.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Dmm *DmmSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Dmm.Contract.Permit(&_Dmm.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Dmm *DmmTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Dmm.Contract.Permit(&_Dmm.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_Dmm *DmmTransactor) Skim(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Dmm.contract.Transact(opts, "skim", to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_Dmm *DmmSession) Skim(to common.Address) (*types.Transaction, error) {
	return _Dmm.Contract.Skim(&_Dmm.TransactOpts, to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_Dmm *DmmTransactorSession) Skim(to common.Address) (*types.Transaction, error) {
	return _Dmm.Contract.Skim(&_Dmm.TransactOpts, to)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes callbackData) returns()
func (_Dmm *DmmTransactor) Swap(opts *bind.TransactOpts, amount0Out *big.Int, amount1Out *big.Int, to common.Address, callbackData []byte) (*types.Transaction, error) {
	return _Dmm.contract.Transact(opts, "swap", amount0Out, amount1Out, to, callbackData)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes callbackData) returns()
func (_Dmm *DmmSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, callbackData []byte) (*types.Transaction, error) {
	return _Dmm.Contract.Swap(&_Dmm.TransactOpts, amount0Out, amount1Out, to, callbackData)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes callbackData) returns()
func (_Dmm *DmmTransactorSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, callbackData []byte) (*types.Transaction, error) {
	return _Dmm.Contract.Swap(&_Dmm.TransactOpts, amount0Out, amount1Out, to, callbackData)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_Dmm *DmmTransactor) Sync(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dmm.contract.Transact(opts, "sync")
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_Dmm *DmmSession) Sync() (*types.Transaction, error) {
	return _Dmm.Contract.Sync(&_Dmm.TransactOpts)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_Dmm *DmmTransactorSession) Sync() (*types.Transaction, error) {
	return _Dmm.Contract.Sync(&_Dmm.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Dmm *DmmTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dmm.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Dmm *DmmSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dmm.Contract.Transfer(&_Dmm.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Dmm *DmmTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dmm.Contract.Transfer(&_Dmm.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Dmm *DmmTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dmm.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Dmm *DmmSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dmm.Contract.TransferFrom(&_Dmm.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Dmm *DmmTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dmm.Contract.TransferFrom(&_Dmm.TransactOpts, sender, recipient, amount)
}

// DmmApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Dmm contract.
type DmmApprovalIterator struct {
	Event *DmmApproval // Event containing the contract specifics and raw log

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
func (it *DmmApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DmmApproval)
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
		it.Event = new(DmmApproval)
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
func (it *DmmApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DmmApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DmmApproval represents a Approval event raised by the Dmm contract.
type DmmApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Dmm *DmmFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*DmmApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Dmm.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &DmmApprovalIterator{contract: _Dmm.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Dmm *DmmFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DmmApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Dmm.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DmmApproval)
				if err := _Dmm.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Dmm *DmmFilterer) ParseApproval(log types.Log) (*DmmApproval, error) {
	event := new(DmmApproval)
	if err := _Dmm.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DmmBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Dmm contract.
type DmmBurnIterator struct {
	Event *DmmBurn // Event containing the contract specifics and raw log

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
func (it *DmmBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DmmBurn)
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
		it.Event = new(DmmBurn)
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
func (it *DmmBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DmmBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DmmBurn represents a Burn event raised by the Dmm contract.
type DmmBurn struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	To      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_Dmm *DmmFilterer) FilterBurn(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*DmmBurnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Dmm.contract.FilterLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DmmBurnIterator{contract: _Dmm.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_Dmm *DmmFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *DmmBurn, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Dmm.contract.WatchLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DmmBurn)
				if err := _Dmm.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_Dmm *DmmFilterer) ParseBurn(log types.Log) (*DmmBurn, error) {
	event := new(DmmBurn)
	if err := _Dmm.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DmmMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Dmm contract.
type DmmMintIterator struct {
	Event *DmmMint // Event containing the contract specifics and raw log

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
func (it *DmmMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DmmMint)
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
		it.Event = new(DmmMint)
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
func (it *DmmMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DmmMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DmmMint represents a Mint event raised by the Dmm contract.
type DmmMint struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_Dmm *DmmFilterer) FilterMint(opts *bind.FilterOpts, sender []common.Address) (*DmmMintIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Dmm.contract.FilterLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return &DmmMintIterator{contract: _Dmm.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_Dmm *DmmFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *DmmMint, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Dmm.contract.WatchLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DmmMint)
				if err := _Dmm.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_Dmm *DmmFilterer) ParseMint(log types.Log) (*DmmMint, error) {
	event := new(DmmMint)
	if err := _Dmm.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DmmSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Dmm contract.
type DmmSwapIterator struct {
	Event *DmmSwap // Event containing the contract specifics and raw log

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
func (it *DmmSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DmmSwap)
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
		it.Event = new(DmmSwap)
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
func (it *DmmSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DmmSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DmmSwap represents a Swap event raised by the Dmm contract.
type DmmSwap struct {
	Sender         common.Address
	Amount0In      *big.Int
	Amount1In      *big.Int
	Amount0Out     *big.Int
	Amount1Out     *big.Int
	To             common.Address
	FeeInPrecision *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0x606ecd02b3e3b4778f8e97b2e03351de14224efaa5fa64e62200afc9395c2499.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to, uint256 feeInPrecision)
func (_Dmm *DmmFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*DmmSwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Dmm.contract.FilterLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DmmSwapIterator{contract: _Dmm.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0x606ecd02b3e3b4778f8e97b2e03351de14224efaa5fa64e62200afc9395c2499.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to, uint256 feeInPrecision)
func (_Dmm *DmmFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *DmmSwap, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Dmm.contract.WatchLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DmmSwap)
				if err := _Dmm.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0x606ecd02b3e3b4778f8e97b2e03351de14224efaa5fa64e62200afc9395c2499.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to, uint256 feeInPrecision)
func (_Dmm *DmmFilterer) ParseSwap(log types.Log) (*DmmSwap, error) {
	event := new(DmmSwap)
	if err := _Dmm.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DmmSyncIterator is returned from FilterSync and is used to iterate over the raw logs and unpacked data for Sync events raised by the Dmm contract.
type DmmSyncIterator struct {
	Event *DmmSync // Event containing the contract specifics and raw log

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
func (it *DmmSyncIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DmmSync)
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
		it.Event = new(DmmSync)
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
func (it *DmmSyncIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DmmSyncIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DmmSync represents a Sync event raised by the Dmm contract.
type DmmSync struct {
	VReserve0 *big.Int
	VReserve1 *big.Int
	Reserve0  *big.Int
	Reserve1  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSync is a free log retrieval operation binding the contract event 0x2f9d55abfefdfd4c3a83e00a1b419b3c2fe4b83100c559f0e2213e57f6e0bba9.
//
// Solidity: event Sync(uint256 vReserve0, uint256 vReserve1, uint256 reserve0, uint256 reserve1)
func (_Dmm *DmmFilterer) FilterSync(opts *bind.FilterOpts) (*DmmSyncIterator, error) {

	logs, sub, err := _Dmm.contract.FilterLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return &DmmSyncIterator{contract: _Dmm.contract, event: "Sync", logs: logs, sub: sub}, nil
}

// WatchSync is a free log subscription operation binding the contract event 0x2f9d55abfefdfd4c3a83e00a1b419b3c2fe4b83100c559f0e2213e57f6e0bba9.
//
// Solidity: event Sync(uint256 vReserve0, uint256 vReserve1, uint256 reserve0, uint256 reserve1)
func (_Dmm *DmmFilterer) WatchSync(opts *bind.WatchOpts, sink chan<- *DmmSync) (event.Subscription, error) {

	logs, sub, err := _Dmm.contract.WatchLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DmmSync)
				if err := _Dmm.contract.UnpackLog(event, "Sync", log); err != nil {
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

// ParseSync is a log parse operation binding the contract event 0x2f9d55abfefdfd4c3a83e00a1b419b3c2fe4b83100c559f0e2213e57f6e0bba9.
//
// Solidity: event Sync(uint256 vReserve0, uint256 vReserve1, uint256 reserve0, uint256 reserve1)
func (_Dmm *DmmFilterer) ParseSync(log types.Log) (*DmmSync, error) {
	event := new(DmmSync)
	if err := _Dmm.contract.UnpackLog(event, "Sync", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DmmTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Dmm contract.
type DmmTransferIterator struct {
	Event *DmmTransfer // Event containing the contract specifics and raw log

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
func (it *DmmTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DmmTransfer)
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
		it.Event = new(DmmTransfer)
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
func (it *DmmTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DmmTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DmmTransfer represents a Transfer event raised by the Dmm contract.
type DmmTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Dmm *DmmFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DmmTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Dmm.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DmmTransferIterator{contract: _Dmm.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Dmm *DmmFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DmmTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Dmm.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DmmTransfer)
				if err := _Dmm.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Dmm *DmmFilterer) ParseTransfer(log types.Log) (*DmmTransfer, error) {
	event := new(DmmTransfer)
	if err := _Dmm.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DmmUpdateEMAIterator is returned from FilterUpdateEMA and is used to iterate over the raw logs and unpacked data for UpdateEMA events raised by the Dmm contract.
type DmmUpdateEMAIterator struct {
	Event *DmmUpdateEMA // Event containing the contract specifics and raw log

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
func (it *DmmUpdateEMAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DmmUpdateEMA)
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
		it.Event = new(DmmUpdateEMA)
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
func (it *DmmUpdateEMAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DmmUpdateEMAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DmmUpdateEMA represents a UpdateEMA event raised by the Dmm contract.
type DmmUpdateEMA struct {
	ShortEMA        *big.Int
	LongEMA         *big.Int
	LastBlockVolume *big.Int
	SkipBlock       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateEMA is a free log retrieval operation binding the contract event 0x96e2c334d3c0fa98c8b728ee84471864ffe5b28e05f46e52f8a469d0ab3a8b8b.
//
// Solidity: event UpdateEMA(uint256 shortEMA, uint256 longEMA, uint128 lastBlockVolume, uint256 skipBlock)
func (_Dmm *DmmFilterer) FilterUpdateEMA(opts *bind.FilterOpts) (*DmmUpdateEMAIterator, error) {

	logs, sub, err := _Dmm.contract.FilterLogs(opts, "UpdateEMA")
	if err != nil {
		return nil, err
	}
	return &DmmUpdateEMAIterator{contract: _Dmm.contract, event: "UpdateEMA", logs: logs, sub: sub}, nil
}

// WatchUpdateEMA is a free log subscription operation binding the contract event 0x96e2c334d3c0fa98c8b728ee84471864ffe5b28e05f46e52f8a469d0ab3a8b8b.
//
// Solidity: event UpdateEMA(uint256 shortEMA, uint256 longEMA, uint128 lastBlockVolume, uint256 skipBlock)
func (_Dmm *DmmFilterer) WatchUpdateEMA(opts *bind.WatchOpts, sink chan<- *DmmUpdateEMA) (event.Subscription, error) {

	logs, sub, err := _Dmm.contract.WatchLogs(opts, "UpdateEMA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DmmUpdateEMA)
				if err := _Dmm.contract.UnpackLog(event, "UpdateEMA", log); err != nil {
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

// ParseUpdateEMA is a log parse operation binding the contract event 0x96e2c334d3c0fa98c8b728ee84471864ffe5b28e05f46e52f8a469d0ab3a8b8b.
//
// Solidity: event UpdateEMA(uint256 shortEMA, uint256 longEMA, uint128 lastBlockVolume, uint256 skipBlock)
func (_Dmm *DmmFilterer) ParseUpdateEMA(log types.Log) (*DmmUpdateEMA, error) {
	event := new(DmmUpdateEMA)
	if err := _Dmm.contract.UnpackLog(event, "UpdateEMA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
