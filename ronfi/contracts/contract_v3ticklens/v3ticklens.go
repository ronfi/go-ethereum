// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package v3TickLens

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

// ITickLensPopulatedTick is an auto generated low-level Go binding around an user-defined struct.
type ITickLensPopulatedTick struct {
	Tick           *big.Int
	LiquidityNet   *big.Int
	LiquidityGross *big.Int
}

// V3TickLensMetaData contains all meta data concerning the V3TickLens contract.
var V3TickLensMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"int16\",\"name\":\"tickBitmapIndex\",\"type\":\"int16\"}],\"name\":\"getPopulatedTicksInWord\",\"outputs\":[{\"components\":[{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"internalType\":\"int128\",\"name\":\"liquidityNet\",\"type\":\"int128\"},{\"internalType\":\"uint128\",\"name\":\"liquidityGross\",\"type\":\"uint128\"}],\"internalType\":\"structITickLens.PopulatedTick[]\",\"name\":\"populatedTicks\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// V3TickLensABI is the input ABI used to generate the binding from.
// Deprecated: Use V3TickLensMetaData.ABI instead.
var V3TickLensABI = V3TickLensMetaData.ABI

// V3TickLens is an auto generated Go binding around an Ethereum contract.
type V3TickLens struct {
	V3TickLensCaller     // Read-only binding to the contract
	V3TickLensTransactor // Write-only binding to the contract
	V3TickLensFilterer   // Log filterer for contract events
}

// V3TickLensCaller is an auto generated read-only Go binding around an Ethereum contract.
type V3TickLensCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V3TickLensTransactor is an auto generated write-only Go binding around an Ethereum contract.
type V3TickLensTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V3TickLensFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type V3TickLensFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V3TickLensSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type V3TickLensSession struct {
	Contract     *V3TickLens       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// V3TickLensCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type V3TickLensCallerSession struct {
	Contract *V3TickLensCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// V3TickLensTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type V3TickLensTransactorSession struct {
	Contract     *V3TickLensTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// V3TickLensRaw is an auto generated low-level Go binding around an Ethereum contract.
type V3TickLensRaw struct {
	Contract *V3TickLens // Generic contract binding to access the raw methods on
}

// V3TickLensCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type V3TickLensCallerRaw struct {
	Contract *V3TickLensCaller // Generic read-only contract binding to access the raw methods on
}

// V3TickLensTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type V3TickLensTransactorRaw struct {
	Contract *V3TickLensTransactor // Generic write-only contract binding to access the raw methods on
}

// NewV3TickLens creates a new instance of V3TickLens, bound to a specific deployed contract.
func NewV3TickLens(address common.Address, backend bind.ContractBackend) (*V3TickLens, error) {
	contract, err := bindV3TickLens(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &V3TickLens{V3TickLensCaller: V3TickLensCaller{contract: contract}, V3TickLensTransactor: V3TickLensTransactor{contract: contract}, V3TickLensFilterer: V3TickLensFilterer{contract: contract}}, nil
}

// NewV3TickLensCaller creates a new read-only instance of V3TickLens, bound to a specific deployed contract.
func NewV3TickLensCaller(address common.Address, caller bind.ContractCaller) (*V3TickLensCaller, error) {
	contract, err := bindV3TickLens(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &V3TickLensCaller{contract: contract}, nil
}

// NewV3TickLensTransactor creates a new write-only instance of V3TickLens, bound to a specific deployed contract.
func NewV3TickLensTransactor(address common.Address, transactor bind.ContractTransactor) (*V3TickLensTransactor, error) {
	contract, err := bindV3TickLens(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &V3TickLensTransactor{contract: contract}, nil
}

// NewV3TickLensFilterer creates a new log filterer instance of V3TickLens, bound to a specific deployed contract.
func NewV3TickLensFilterer(address common.Address, filterer bind.ContractFilterer) (*V3TickLensFilterer, error) {
	contract, err := bindV3TickLens(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &V3TickLensFilterer{contract: contract}, nil
}

// bindV3TickLens binds a generic wrapper to an already deployed contract.
func bindV3TickLens(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(V3TickLensABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V3TickLens *V3TickLensRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V3TickLens.Contract.V3TickLensCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V3TickLens *V3TickLensRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V3TickLens.Contract.V3TickLensTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V3TickLens *V3TickLensRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V3TickLens.Contract.V3TickLensTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V3TickLens *V3TickLensCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V3TickLens.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V3TickLens *V3TickLensTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V3TickLens.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V3TickLens *V3TickLensTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V3TickLens.Contract.contract.Transact(opts, method, params...)
}

// GetPopulatedTicksInWord is a free data retrieval call binding the contract method 0x351fb478.
//
// Solidity: function getPopulatedTicksInWord(address pool, int16 tickBitmapIndex) view returns((int24,int128,uint128)[] populatedTicks)
func (_V3TickLens *V3TickLensCaller) GetPopulatedTicksInWord(opts *bind.CallOpts, pool common.Address, tickBitmapIndex int16) ([]ITickLensPopulatedTick, error) {
	var out []interface{}
	err := _V3TickLens.contract.Call(opts, &out, "getPopulatedTicksInWord", pool, tickBitmapIndex)

	if err != nil {
		return *new([]ITickLensPopulatedTick), err
	}

	out0 := *abi.ConvertType(out[0], new([]ITickLensPopulatedTick)).(*[]ITickLensPopulatedTick)

	return out0, err

}

// GetPopulatedTicksInWord is a free data retrieval call binding the contract method 0x351fb478.
//
// Solidity: function getPopulatedTicksInWord(address pool, int16 tickBitmapIndex) view returns((int24,int128,uint128)[] populatedTicks)
func (_V3TickLens *V3TickLensSession) GetPopulatedTicksInWord(pool common.Address, tickBitmapIndex int16) ([]ITickLensPopulatedTick, error) {
	return _V3TickLens.Contract.GetPopulatedTicksInWord(&_V3TickLens.CallOpts, pool, tickBitmapIndex)
}

// GetPopulatedTicksInWord is a free data retrieval call binding the contract method 0x351fb478.
//
// Solidity: function getPopulatedTicksInWord(address pool, int16 tickBitmapIndex) view returns((int24,int128,uint128)[] populatedTicks)
func (_V3TickLens *V3TickLensCallerSession) GetPopulatedTicksInWord(pool common.Address, tickBitmapIndex int16) ([]ITickLensPopulatedTick, error) {
	return _V3TickLens.Contract.GetPopulatedTicksInWord(&_V3TickLens.CallOpts, pool, tickBitmapIndex)
}
