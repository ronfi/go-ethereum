// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package algebrapool

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

// AlgebrapoolMetaData contains all meta data concerning the Algebrapool contract.
var AlgebrapoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"bottomTick\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"topTick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidityAmount\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"bottomTick\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"topTick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"name\":\"Collect\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"communityFee0New\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"communityFee1New\",\"type\":\"uint16\"}],\"name\":\"CommunityFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"fee\",\"type\":\"uint16\"}],\"name\":\"Fee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paid0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paid1\",\"type\":\"uint256\"}],\"name\":\"Flash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"virtualPoolAddress\",\"type\":\"address\"}],\"name\":\"Incentive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"price\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"Initialize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"liquidityCooldown\",\"type\":\"uint32\"}],\"name\":\"LiquidityCooldown\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"bottomTick\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"topTick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidityAmount\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount0\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount1\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"price\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"activeIncentive\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"bottomTick\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"topTick\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"bottomTick\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"topTick\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"amount0Requested\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1Requested\",\"type\":\"uint128\"}],\"name\":\"collect\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dataStorageOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"flash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"bottomTick\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"topTick\",\"type\":\"int24\"}],\"name\":\"getInnerCumulatives\",\"outputs\":[{\"internalType\":\"int56\",\"name\":\"innerTickCumulative\",\"type\":\"int56\"},{\"internalType\":\"uint160\",\"name\":\"innerSecondsSpentPerLiquidity\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"innerSecondsSpent\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"secondsAgos\",\"type\":\"uint32[]\"}],\"name\":\"getTimepoints\",\"outputs\":[{\"internalType\":\"int56[]\",\"name\":\"tickCumulatives\",\"type\":\"int56[]\"},{\"internalType\":\"uint160[]\",\"name\":\"secondsPerLiquidityCumulatives\",\"type\":\"uint160[]\"},{\"internalType\":\"uint112[]\",\"name\":\"volatilityCumulatives\",\"type\":\"uint112[]\"},{\"internalType\":\"uint256[]\",\"name\":\"volumePerAvgLiquiditys\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalState\",\"outputs\":[{\"internalType\":\"uint160\",\"name\":\"price\",\"type\":\"uint160\"},{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"internalType\":\"uint16\",\"name\":\"fee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"timepointIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"communityFeeToken0\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"communityFeeToken1\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"unlocked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint160\",\"name\":\"initialPrice\",\"type\":\"uint160\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidity\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidityCooldown\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxLiquidityPerTick\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"bottomTick\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"topTick\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"liquidityDesired\",\"type\":\"uint128\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"liquidityActual\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"positions\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"lastLiquidityAddTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"innerFeeGrowth0Token\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"innerFeeGrowth1Token\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"fees0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"fees1\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"communityFee0\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"communityFee1\",\"type\":\"uint16\"}],\"name\":\"setCommunityFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"virtualPoolAddress\",\"type\":\"address\"}],\"name\":\"setIncentive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newLiquidityCooldown\",\"type\":\"uint32\"}],\"name\":\"setLiquidityCooldown\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"zeroToOne\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"amountRequired\",\"type\":\"int256\"},{\"internalType\":\"uint160\",\"name\":\"limitSqrtPrice\",\"type\":\"uint160\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"amount0\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"zeroToOne\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"amountRequired\",\"type\":\"int256\"},{\"internalType\":\"uint160\",\"name\":\"limitSqrtPrice\",\"type\":\"uint160\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swapSupportingFeeOnInputTokens\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"amount0\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tickSpacing\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int16\",\"name\":\"\",\"type\":\"int16\"}],\"name\":\"tickTable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"name\":\"ticks\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidityTotal\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"liquidityDelta\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"outerFeeGrowth0Token\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outerFeeGrowth1Token\",\"type\":\"uint256\"},{\"internalType\":\"int56\",\"name\":\"outerTickCumulative\",\"type\":\"int56\"},{\"internalType\":\"uint160\",\"name\":\"outerSecondsPerLiquidity\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"outerSecondsSpent\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"timepoints\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"blockTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"int56\",\"name\":\"tickCumulative\",\"type\":\"int56\"},{\"internalType\":\"uint160\",\"name\":\"secondsPerLiquidityCumulative\",\"type\":\"uint160\"},{\"internalType\":\"uint88\",\"name\":\"volatilityCumulative\",\"type\":\"uint88\"},{\"internalType\":\"int24\",\"name\":\"averageTick\",\"type\":\"int24\"},{\"internalType\":\"uint144\",\"name\":\"volumePerLiquidityCumulative\",\"type\":\"uint144\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalFeeGrowth0Token\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalFeeGrowth1Token\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AlgebrapoolABI is the input ABI used to generate the binding from.
// Deprecated: Use AlgebrapoolMetaData.ABI instead.
var AlgebrapoolABI = AlgebrapoolMetaData.ABI

// Algebrapool is an auto generated Go binding around an Ethereum contract.
type Algebrapool struct {
	AlgebrapoolCaller     // Read-only binding to the contract
	AlgebrapoolTransactor // Write-only binding to the contract
	AlgebrapoolFilterer   // Log filterer for contract events
}

// AlgebrapoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type AlgebrapoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AlgebrapoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AlgebrapoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AlgebrapoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AlgebrapoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AlgebrapoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AlgebrapoolSession struct {
	Contract     *Algebrapool      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AlgebrapoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AlgebrapoolCallerSession struct {
	Contract *AlgebrapoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AlgebrapoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AlgebrapoolTransactorSession struct {
	Contract     *AlgebrapoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AlgebrapoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type AlgebrapoolRaw struct {
	Contract *Algebrapool // Generic contract binding to access the raw methods on
}

// AlgebrapoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AlgebrapoolCallerRaw struct {
	Contract *AlgebrapoolCaller // Generic read-only contract binding to access the raw methods on
}

// AlgebrapoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AlgebrapoolTransactorRaw struct {
	Contract *AlgebrapoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAlgebrapool creates a new instance of Algebrapool, bound to a specific deployed contract.
func NewAlgebrapool(address common.Address, backend bind.ContractBackend) (*Algebrapool, error) {
	contract, err := bindAlgebrapool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Algebrapool{AlgebrapoolCaller: AlgebrapoolCaller{contract: contract}, AlgebrapoolTransactor: AlgebrapoolTransactor{contract: contract}, AlgebrapoolFilterer: AlgebrapoolFilterer{contract: contract}}, nil
}

// NewAlgebrapoolCaller creates a new read-only instance of Algebrapool, bound to a specific deployed contract.
func NewAlgebrapoolCaller(address common.Address, caller bind.ContractCaller) (*AlgebrapoolCaller, error) {
	contract, err := bindAlgebrapool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AlgebrapoolCaller{contract: contract}, nil
}

// NewAlgebrapoolTransactor creates a new write-only instance of Algebrapool, bound to a specific deployed contract.
func NewAlgebrapoolTransactor(address common.Address, transactor bind.ContractTransactor) (*AlgebrapoolTransactor, error) {
	contract, err := bindAlgebrapool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AlgebrapoolTransactor{contract: contract}, nil
}

// NewAlgebrapoolFilterer creates a new log filterer instance of Algebrapool, bound to a specific deployed contract.
func NewAlgebrapoolFilterer(address common.Address, filterer bind.ContractFilterer) (*AlgebrapoolFilterer, error) {
	contract, err := bindAlgebrapool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AlgebrapoolFilterer{contract: contract}, nil
}

// bindAlgebrapool binds a generic wrapper to an already deployed contract.
func bindAlgebrapool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AlgebrapoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Algebrapool *AlgebrapoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Algebrapool.Contract.AlgebrapoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Algebrapool *AlgebrapoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Algebrapool.Contract.AlgebrapoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Algebrapool *AlgebrapoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Algebrapool.Contract.AlgebrapoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Algebrapool *AlgebrapoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Algebrapool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Algebrapool *AlgebrapoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Algebrapool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Algebrapool *AlgebrapoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Algebrapool.Contract.contract.Transact(opts, method, params...)
}

// ActiveIncentive is a free data retrieval call binding the contract method 0xfacb0eb1.
//
// Solidity: function activeIncentive() view returns(address)
func (_Algebrapool *AlgebrapoolCaller) ActiveIncentive(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Algebrapool.contract.Call(opts, &out, "activeIncentive")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ActiveIncentive is a free data retrieval call binding the contract method 0xfacb0eb1.
//
// Solidity: function activeIncentive() view returns(address)
func (_Algebrapool *AlgebrapoolSession) ActiveIncentive() (common.Address, error) {
	return _Algebrapool.Contract.ActiveIncentive(&_Algebrapool.CallOpts)
}

// ActiveIncentive is a free data retrieval call binding the contract method 0xfacb0eb1.
//
// Solidity: function activeIncentive() view returns(address)
func (_Algebrapool *AlgebrapoolCallerSession) ActiveIncentive() (common.Address, error) {
	return _Algebrapool.Contract.ActiveIncentive(&_Algebrapool.CallOpts)
}

// DataStorageOperator is a free data retrieval call binding the contract method 0x29047dfa.
//
// Solidity: function dataStorageOperator() view returns(address)
func (_Algebrapool *AlgebrapoolCaller) DataStorageOperator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Algebrapool.contract.Call(opts, &out, "dataStorageOperator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DataStorageOperator is a free data retrieval call binding the contract method 0x29047dfa.
//
// Solidity: function dataStorageOperator() view returns(address)
func (_Algebrapool *AlgebrapoolSession) DataStorageOperator() (common.Address, error) {
	return _Algebrapool.Contract.DataStorageOperator(&_Algebrapool.CallOpts)
}

// DataStorageOperator is a free data retrieval call binding the contract method 0x29047dfa.
//
// Solidity: function dataStorageOperator() view returns(address)
func (_Algebrapool *AlgebrapoolCallerSession) DataStorageOperator() (common.Address, error) {
	return _Algebrapool.Contract.DataStorageOperator(&_Algebrapool.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Algebrapool *AlgebrapoolCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Algebrapool.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Algebrapool *AlgebrapoolSession) Factory() (common.Address, error) {
	return _Algebrapool.Contract.Factory(&_Algebrapool.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Algebrapool *AlgebrapoolCallerSession) Factory() (common.Address, error) {
	return _Algebrapool.Contract.Factory(&_Algebrapool.CallOpts)
}

// GetInnerCumulatives is a free data retrieval call binding the contract method 0x920c34e5.
//
// Solidity: function getInnerCumulatives(int24 bottomTick, int24 topTick) view returns(int56 innerTickCumulative, uint160 innerSecondsSpentPerLiquidity, uint32 innerSecondsSpent)
func (_Algebrapool *AlgebrapoolCaller) GetInnerCumulatives(opts *bind.CallOpts, bottomTick *big.Int, topTick *big.Int) (struct {
	InnerTickCumulative           *big.Int
	InnerSecondsSpentPerLiquidity *big.Int
	InnerSecondsSpent             uint32
}, error) {
	var out []interface{}
	err := _Algebrapool.contract.Call(opts, &out, "getInnerCumulatives", bottomTick, topTick)

	outstruct := new(struct {
		InnerTickCumulative           *big.Int
		InnerSecondsSpentPerLiquidity *big.Int
		InnerSecondsSpent             uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.InnerTickCumulative = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.InnerSecondsSpentPerLiquidity = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.InnerSecondsSpent = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetInnerCumulatives is a free data retrieval call binding the contract method 0x920c34e5.
//
// Solidity: function getInnerCumulatives(int24 bottomTick, int24 topTick) view returns(int56 innerTickCumulative, uint160 innerSecondsSpentPerLiquidity, uint32 innerSecondsSpent)
func (_Algebrapool *AlgebrapoolSession) GetInnerCumulatives(bottomTick *big.Int, topTick *big.Int) (struct {
	InnerTickCumulative           *big.Int
	InnerSecondsSpentPerLiquidity *big.Int
	InnerSecondsSpent             uint32
}, error) {
	return _Algebrapool.Contract.GetInnerCumulatives(&_Algebrapool.CallOpts, bottomTick, topTick)
}

// GetInnerCumulatives is a free data retrieval call binding the contract method 0x920c34e5.
//
// Solidity: function getInnerCumulatives(int24 bottomTick, int24 topTick) view returns(int56 innerTickCumulative, uint160 innerSecondsSpentPerLiquidity, uint32 innerSecondsSpent)
func (_Algebrapool *AlgebrapoolCallerSession) GetInnerCumulatives(bottomTick *big.Int, topTick *big.Int) (struct {
	InnerTickCumulative           *big.Int
	InnerSecondsSpentPerLiquidity *big.Int
	InnerSecondsSpent             uint32
}, error) {
	return _Algebrapool.Contract.GetInnerCumulatives(&_Algebrapool.CallOpts, bottomTick, topTick)
}

// GetTimepoints is a free data retrieval call binding the contract method 0x9d3a5241.
//
// Solidity: function getTimepoints(uint32[] secondsAgos) view returns(int56[] tickCumulatives, uint160[] secondsPerLiquidityCumulatives, uint112[] volatilityCumulatives, uint256[] volumePerAvgLiquiditys)
func (_Algebrapool *AlgebrapoolCaller) GetTimepoints(opts *bind.CallOpts, secondsAgos []uint32) (struct {
	TickCumulatives                []*big.Int
	SecondsPerLiquidityCumulatives []*big.Int
	VolatilityCumulatives          []*big.Int
	VolumePerAvgLiquiditys         []*big.Int
}, error) {
	var out []interface{}
	err := _Algebrapool.contract.Call(opts, &out, "getTimepoints", secondsAgos)

	outstruct := new(struct {
		TickCumulatives                []*big.Int
		SecondsPerLiquidityCumulatives []*big.Int
		VolatilityCumulatives          []*big.Int
		VolumePerAvgLiquiditys         []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TickCumulatives = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.SecondsPerLiquidityCumulatives = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.VolatilityCumulatives = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)
	outstruct.VolumePerAvgLiquiditys = *abi.ConvertType(out[3], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetTimepoints is a free data retrieval call binding the contract method 0x9d3a5241.
//
// Solidity: function getTimepoints(uint32[] secondsAgos) view returns(int56[] tickCumulatives, uint160[] secondsPerLiquidityCumulatives, uint112[] volatilityCumulatives, uint256[] volumePerAvgLiquiditys)
func (_Algebrapool *AlgebrapoolSession) GetTimepoints(secondsAgos []uint32) (struct {
	TickCumulatives                []*big.Int
	SecondsPerLiquidityCumulatives []*big.Int
	VolatilityCumulatives          []*big.Int
	VolumePerAvgLiquiditys         []*big.Int
}, error) {
	return _Algebrapool.Contract.GetTimepoints(&_Algebrapool.CallOpts, secondsAgos)
}

// GetTimepoints is a free data retrieval call binding the contract method 0x9d3a5241.
//
// Solidity: function getTimepoints(uint32[] secondsAgos) view returns(int56[] tickCumulatives, uint160[] secondsPerLiquidityCumulatives, uint112[] volatilityCumulatives, uint256[] volumePerAvgLiquiditys)
func (_Algebrapool *AlgebrapoolCallerSession) GetTimepoints(secondsAgos []uint32) (struct {
	TickCumulatives                []*big.Int
	SecondsPerLiquidityCumulatives []*big.Int
	VolatilityCumulatives          []*big.Int
	VolumePerAvgLiquiditys         []*big.Int
}, error) {
	return _Algebrapool.Contract.GetTimepoints(&_Algebrapool.CallOpts, secondsAgos)
}

// GlobalState is a free data retrieval call binding the contract method 0xe76c01e4.
//
// Solidity: function globalState() view returns(uint160 price, int24 tick, uint16 fee, uint16 timepointIndex, uint16 communityFeeToken0, uint16 communityFeeToken1, bool unlocked)
func (_Algebrapool *AlgebrapoolCaller) GlobalState(opts *bind.CallOpts) (struct {
	Price              *big.Int
	Tick               *big.Int
	Fee                uint16
	TimepointIndex     uint16
	CommunityFeeToken0 uint16
	CommunityFeeToken1 uint16
	Unlocked           bool
}, error) {
	var out []interface{}
	err := _Algebrapool.contract.Call(opts, &out, "globalState")

	outstruct := new(struct {
		Price              *big.Int
		Tick               *big.Int
		Fee                uint16
		TimepointIndex     uint16
		CommunityFeeToken0 uint16
		CommunityFeeToken1 uint16
		Unlocked           bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Tick = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Fee = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.TimepointIndex = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.CommunityFeeToken0 = *abi.ConvertType(out[4], new(uint16)).(*uint16)
	outstruct.CommunityFeeToken1 = *abi.ConvertType(out[5], new(uint16)).(*uint16)
	outstruct.Unlocked = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// GlobalState is a free data retrieval call binding the contract method 0xe76c01e4.
//
// Solidity: function globalState() view returns(uint160 price, int24 tick, uint16 fee, uint16 timepointIndex, uint16 communityFeeToken0, uint16 communityFeeToken1, bool unlocked)
func (_Algebrapool *AlgebrapoolSession) GlobalState() (struct {
	Price              *big.Int
	Tick               *big.Int
	Fee                uint16
	TimepointIndex     uint16
	CommunityFeeToken0 uint16
	CommunityFeeToken1 uint16
	Unlocked           bool
}, error) {
	return _Algebrapool.Contract.GlobalState(&_Algebrapool.CallOpts)
}

// GlobalState is a free data retrieval call binding the contract method 0xe76c01e4.
//
// Solidity: function globalState() view returns(uint160 price, int24 tick, uint16 fee, uint16 timepointIndex, uint16 communityFeeToken0, uint16 communityFeeToken1, bool unlocked)
func (_Algebrapool *AlgebrapoolCallerSession) GlobalState() (struct {
	Price              *big.Int
	Tick               *big.Int
	Fee                uint16
	TimepointIndex     uint16
	CommunityFeeToken0 uint16
	CommunityFeeToken1 uint16
	Unlocked           bool
}, error) {
	return _Algebrapool.Contract.GlobalState(&_Algebrapool.CallOpts)
}

// Liquidity is a free data retrieval call binding the contract method 0x1a686502.
//
// Solidity: function liquidity() view returns(uint128)
func (_Algebrapool *AlgebrapoolCaller) Liquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Algebrapool.contract.Call(opts, &out, "liquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Liquidity is a free data retrieval call binding the contract method 0x1a686502.
//
// Solidity: function liquidity() view returns(uint128)
func (_Algebrapool *AlgebrapoolSession) Liquidity() (*big.Int, error) {
	return _Algebrapool.Contract.Liquidity(&_Algebrapool.CallOpts)
}

// Liquidity is a free data retrieval call binding the contract method 0x1a686502.
//
// Solidity: function liquidity() view returns(uint128)
func (_Algebrapool *AlgebrapoolCallerSession) Liquidity() (*big.Int, error) {
	return _Algebrapool.Contract.Liquidity(&_Algebrapool.CallOpts)
}

// LiquidityCooldown is a free data retrieval call binding the contract method 0x17e25b3c.
//
// Solidity: function liquidityCooldown() view returns(uint32)
func (_Algebrapool *AlgebrapoolCaller) LiquidityCooldown(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Algebrapool.contract.Call(opts, &out, "liquidityCooldown")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LiquidityCooldown is a free data retrieval call binding the contract method 0x17e25b3c.
//
// Solidity: function liquidityCooldown() view returns(uint32)
func (_Algebrapool *AlgebrapoolSession) LiquidityCooldown() (uint32, error) {
	return _Algebrapool.Contract.LiquidityCooldown(&_Algebrapool.CallOpts)
}

// LiquidityCooldown is a free data retrieval call binding the contract method 0x17e25b3c.
//
// Solidity: function liquidityCooldown() view returns(uint32)
func (_Algebrapool *AlgebrapoolCallerSession) LiquidityCooldown() (uint32, error) {
	return _Algebrapool.Contract.LiquidityCooldown(&_Algebrapool.CallOpts)
}

// MaxLiquidityPerTick is a free data retrieval call binding the contract method 0x70cf754a.
//
// Solidity: function maxLiquidityPerTick() pure returns(uint128)
func (_Algebrapool *AlgebrapoolCaller) MaxLiquidityPerTick(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Algebrapool.contract.Call(opts, &out, "maxLiquidityPerTick")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxLiquidityPerTick is a free data retrieval call binding the contract method 0x70cf754a.
//
// Solidity: function maxLiquidityPerTick() pure returns(uint128)
func (_Algebrapool *AlgebrapoolSession) MaxLiquidityPerTick() (*big.Int, error) {
	return _Algebrapool.Contract.MaxLiquidityPerTick(&_Algebrapool.CallOpts)
}

// MaxLiquidityPerTick is a free data retrieval call binding the contract method 0x70cf754a.
//
// Solidity: function maxLiquidityPerTick() pure returns(uint128)
func (_Algebrapool *AlgebrapoolCallerSession) MaxLiquidityPerTick() (*big.Int, error) {
	return _Algebrapool.Contract.MaxLiquidityPerTick(&_Algebrapool.CallOpts)
}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 ) view returns(uint128 liquidity, uint32 lastLiquidityAddTimestamp, uint256 innerFeeGrowth0Token, uint256 innerFeeGrowth1Token, uint128 fees0, uint128 fees1)
func (_Algebrapool *AlgebrapoolCaller) Positions(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Liquidity                 *big.Int
	LastLiquidityAddTimestamp uint32
	InnerFeeGrowth0Token      *big.Int
	InnerFeeGrowth1Token      *big.Int
	Fees0                     *big.Int
	Fees1                     *big.Int
}, error) {
	var out []interface{}
	err := _Algebrapool.contract.Call(opts, &out, "positions", arg0)

	outstruct := new(struct {
		Liquidity                 *big.Int
		LastLiquidityAddTimestamp uint32
		InnerFeeGrowth0Token      *big.Int
		InnerFeeGrowth1Token      *big.Int
		Fees0                     *big.Int
		Fees1                     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Liquidity = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastLiquidityAddTimestamp = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.InnerFeeGrowth0Token = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.InnerFeeGrowth1Token = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Fees0 = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Fees1 = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 ) view returns(uint128 liquidity, uint32 lastLiquidityAddTimestamp, uint256 innerFeeGrowth0Token, uint256 innerFeeGrowth1Token, uint128 fees0, uint128 fees1)
func (_Algebrapool *AlgebrapoolSession) Positions(arg0 [32]byte) (struct {
	Liquidity                 *big.Int
	LastLiquidityAddTimestamp uint32
	InnerFeeGrowth0Token      *big.Int
	InnerFeeGrowth1Token      *big.Int
	Fees0                     *big.Int
	Fees1                     *big.Int
}, error) {
	return _Algebrapool.Contract.Positions(&_Algebrapool.CallOpts, arg0)
}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 ) view returns(uint128 liquidity, uint32 lastLiquidityAddTimestamp, uint256 innerFeeGrowth0Token, uint256 innerFeeGrowth1Token, uint128 fees0, uint128 fees1)
func (_Algebrapool *AlgebrapoolCallerSession) Positions(arg0 [32]byte) (struct {
	Liquidity                 *big.Int
	LastLiquidityAddTimestamp uint32
	InnerFeeGrowth0Token      *big.Int
	InnerFeeGrowth1Token      *big.Int
	Fees0                     *big.Int
	Fees1                     *big.Int
}, error) {
	return _Algebrapool.Contract.Positions(&_Algebrapool.CallOpts, arg0)
}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() pure returns(int24)
func (_Algebrapool *AlgebrapoolCaller) TickSpacing(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Algebrapool.contract.Call(opts, &out, "tickSpacing")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() pure returns(int24)
func (_Algebrapool *AlgebrapoolSession) TickSpacing() (*big.Int, error) {
	return _Algebrapool.Contract.TickSpacing(&_Algebrapool.CallOpts)
}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() pure returns(int24)
func (_Algebrapool *AlgebrapoolCallerSession) TickSpacing() (*big.Int, error) {
	return _Algebrapool.Contract.TickSpacing(&_Algebrapool.CallOpts)
}

// TickTable is a free data retrieval call binding the contract method 0xc677e3e0.
//
// Solidity: function tickTable(int16 ) view returns(uint256)
func (_Algebrapool *AlgebrapoolCaller) TickTable(opts *bind.CallOpts, arg0 int16) (*big.Int, error) {
	var out []interface{}
	err := _Algebrapool.contract.Call(opts, &out, "tickTable", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TickTable is a free data retrieval call binding the contract method 0xc677e3e0.
//
// Solidity: function tickTable(int16 ) view returns(uint256)
func (_Algebrapool *AlgebrapoolSession) TickTable(arg0 int16) (*big.Int, error) {
	return _Algebrapool.Contract.TickTable(&_Algebrapool.CallOpts, arg0)
}

// TickTable is a free data retrieval call binding the contract method 0xc677e3e0.
//
// Solidity: function tickTable(int16 ) view returns(uint256)
func (_Algebrapool *AlgebrapoolCallerSession) TickTable(arg0 int16) (*big.Int, error) {
	return _Algebrapool.Contract.TickTable(&_Algebrapool.CallOpts, arg0)
}

// Ticks is a free data retrieval call binding the contract method 0xf30dba93.
//
// Solidity: function ticks(int24 ) view returns(uint128 liquidityTotal, int128 liquidityDelta, uint256 outerFeeGrowth0Token, uint256 outerFeeGrowth1Token, int56 outerTickCumulative, uint160 outerSecondsPerLiquidity, uint32 outerSecondsSpent, bool initialized)
func (_Algebrapool *AlgebrapoolCaller) Ticks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LiquidityTotal           *big.Int
	LiquidityDelta           *big.Int
	OuterFeeGrowth0Token     *big.Int
	OuterFeeGrowth1Token     *big.Int
	OuterTickCumulative      *big.Int
	OuterSecondsPerLiquidity *big.Int
	OuterSecondsSpent        uint32
	Initialized              bool
}, error) {
	var out []interface{}
	err := _Algebrapool.contract.Call(opts, &out, "ticks", arg0)

	outstruct := new(struct {
		LiquidityTotal           *big.Int
		LiquidityDelta           *big.Int
		OuterFeeGrowth0Token     *big.Int
		OuterFeeGrowth1Token     *big.Int
		OuterTickCumulative      *big.Int
		OuterSecondsPerLiquidity *big.Int
		OuterSecondsSpent        uint32
		Initialized              bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LiquidityTotal = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LiquidityDelta = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.OuterFeeGrowth0Token = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.OuterFeeGrowth1Token = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.OuterTickCumulative = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.OuterSecondsPerLiquidity = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.OuterSecondsSpent = *abi.ConvertType(out[6], new(uint32)).(*uint32)
	outstruct.Initialized = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// Ticks is a free data retrieval call binding the contract method 0xf30dba93.
//
// Solidity: function ticks(int24 ) view returns(uint128 liquidityTotal, int128 liquidityDelta, uint256 outerFeeGrowth0Token, uint256 outerFeeGrowth1Token, int56 outerTickCumulative, uint160 outerSecondsPerLiquidity, uint32 outerSecondsSpent, bool initialized)
func (_Algebrapool *AlgebrapoolSession) Ticks(arg0 *big.Int) (struct {
	LiquidityTotal           *big.Int
	LiquidityDelta           *big.Int
	OuterFeeGrowth0Token     *big.Int
	OuterFeeGrowth1Token     *big.Int
	OuterTickCumulative      *big.Int
	OuterSecondsPerLiquidity *big.Int
	OuterSecondsSpent        uint32
	Initialized              bool
}, error) {
	return _Algebrapool.Contract.Ticks(&_Algebrapool.CallOpts, arg0)
}

// Ticks is a free data retrieval call binding the contract method 0xf30dba93.
//
// Solidity: function ticks(int24 ) view returns(uint128 liquidityTotal, int128 liquidityDelta, uint256 outerFeeGrowth0Token, uint256 outerFeeGrowth1Token, int56 outerTickCumulative, uint160 outerSecondsPerLiquidity, uint32 outerSecondsSpent, bool initialized)
func (_Algebrapool *AlgebrapoolCallerSession) Ticks(arg0 *big.Int) (struct {
	LiquidityTotal           *big.Int
	LiquidityDelta           *big.Int
	OuterFeeGrowth0Token     *big.Int
	OuterFeeGrowth1Token     *big.Int
	OuterTickCumulative      *big.Int
	OuterSecondsPerLiquidity *big.Int
	OuterSecondsSpent        uint32
	Initialized              bool
}, error) {
	return _Algebrapool.Contract.Ticks(&_Algebrapool.CallOpts, arg0)
}

// Timepoints is a free data retrieval call binding the contract method 0x74eceae6.
//
// Solidity: function timepoints(uint256 index) view returns(bool initialized, uint32 blockTimestamp, int56 tickCumulative, uint160 secondsPerLiquidityCumulative, uint88 volatilityCumulative, int24 averageTick, uint144 volumePerLiquidityCumulative)
func (_Algebrapool *AlgebrapoolCaller) Timepoints(opts *bind.CallOpts, index *big.Int) (struct {
	Initialized                   bool
	BlockTimestamp                uint32
	TickCumulative                *big.Int
	SecondsPerLiquidityCumulative *big.Int
	VolatilityCumulative          *big.Int
	AverageTick                   *big.Int
	VolumePerLiquidityCumulative  *big.Int
}, error) {
	var out []interface{}
	err := _Algebrapool.contract.Call(opts, &out, "timepoints", index)

	outstruct := new(struct {
		Initialized                   bool
		BlockTimestamp                uint32
		TickCumulative                *big.Int
		SecondsPerLiquidityCumulative *big.Int
		VolatilityCumulative          *big.Int
		AverageTick                   *big.Int
		VolumePerLiquidityCumulative  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Initialized = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.BlockTimestamp = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.TickCumulative = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.SecondsPerLiquidityCumulative = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VolatilityCumulative = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.AverageTick = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.VolumePerLiquidityCumulative = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Timepoints is a free data retrieval call binding the contract method 0x74eceae6.
//
// Solidity: function timepoints(uint256 index) view returns(bool initialized, uint32 blockTimestamp, int56 tickCumulative, uint160 secondsPerLiquidityCumulative, uint88 volatilityCumulative, int24 averageTick, uint144 volumePerLiquidityCumulative)
func (_Algebrapool *AlgebrapoolSession) Timepoints(index *big.Int) (struct {
	Initialized                   bool
	BlockTimestamp                uint32
	TickCumulative                *big.Int
	SecondsPerLiquidityCumulative *big.Int
	VolatilityCumulative          *big.Int
	AverageTick                   *big.Int
	VolumePerLiquidityCumulative  *big.Int
}, error) {
	return _Algebrapool.Contract.Timepoints(&_Algebrapool.CallOpts, index)
}

// Timepoints is a free data retrieval call binding the contract method 0x74eceae6.
//
// Solidity: function timepoints(uint256 index) view returns(bool initialized, uint32 blockTimestamp, int56 tickCumulative, uint160 secondsPerLiquidityCumulative, uint88 volatilityCumulative, int24 averageTick, uint144 volumePerLiquidityCumulative)
func (_Algebrapool *AlgebrapoolCallerSession) Timepoints(index *big.Int) (struct {
	Initialized                   bool
	BlockTimestamp                uint32
	TickCumulative                *big.Int
	SecondsPerLiquidityCumulative *big.Int
	VolatilityCumulative          *big.Int
	AverageTick                   *big.Int
	VolumePerLiquidityCumulative  *big.Int
}, error) {
	return _Algebrapool.Contract.Timepoints(&_Algebrapool.CallOpts, index)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Algebrapool *AlgebrapoolCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Algebrapool.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Algebrapool *AlgebrapoolSession) Token0() (common.Address, error) {
	return _Algebrapool.Contract.Token0(&_Algebrapool.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Algebrapool *AlgebrapoolCallerSession) Token0() (common.Address, error) {
	return _Algebrapool.Contract.Token0(&_Algebrapool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Algebrapool *AlgebrapoolCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Algebrapool.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Algebrapool *AlgebrapoolSession) Token1() (common.Address, error) {
	return _Algebrapool.Contract.Token1(&_Algebrapool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Algebrapool *AlgebrapoolCallerSession) Token1() (common.Address, error) {
	return _Algebrapool.Contract.Token1(&_Algebrapool.CallOpts)
}

// TotalFeeGrowth0Token is a free data retrieval call binding the contract method 0x6378ae44.
//
// Solidity: function totalFeeGrowth0Token() view returns(uint256)
func (_Algebrapool *AlgebrapoolCaller) TotalFeeGrowth0Token(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Algebrapool.contract.Call(opts, &out, "totalFeeGrowth0Token")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFeeGrowth0Token is a free data retrieval call binding the contract method 0x6378ae44.
//
// Solidity: function totalFeeGrowth0Token() view returns(uint256)
func (_Algebrapool *AlgebrapoolSession) TotalFeeGrowth0Token() (*big.Int, error) {
	return _Algebrapool.Contract.TotalFeeGrowth0Token(&_Algebrapool.CallOpts)
}

// TotalFeeGrowth0Token is a free data retrieval call binding the contract method 0x6378ae44.
//
// Solidity: function totalFeeGrowth0Token() view returns(uint256)
func (_Algebrapool *AlgebrapoolCallerSession) TotalFeeGrowth0Token() (*big.Int, error) {
	return _Algebrapool.Contract.TotalFeeGrowth0Token(&_Algebrapool.CallOpts)
}

// TotalFeeGrowth1Token is a free data retrieval call binding the contract method 0xecdecf42.
//
// Solidity: function totalFeeGrowth1Token() view returns(uint256)
func (_Algebrapool *AlgebrapoolCaller) TotalFeeGrowth1Token(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Algebrapool.contract.Call(opts, &out, "totalFeeGrowth1Token")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFeeGrowth1Token is a free data retrieval call binding the contract method 0xecdecf42.
//
// Solidity: function totalFeeGrowth1Token() view returns(uint256)
func (_Algebrapool *AlgebrapoolSession) TotalFeeGrowth1Token() (*big.Int, error) {
	return _Algebrapool.Contract.TotalFeeGrowth1Token(&_Algebrapool.CallOpts)
}

// TotalFeeGrowth1Token is a free data retrieval call binding the contract method 0xecdecf42.
//
// Solidity: function totalFeeGrowth1Token() view returns(uint256)
func (_Algebrapool *AlgebrapoolCallerSession) TotalFeeGrowth1Token() (*big.Int, error) {
	return _Algebrapool.Contract.TotalFeeGrowth1Token(&_Algebrapool.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0xa34123a7.
//
// Solidity: function burn(int24 bottomTick, int24 topTick, uint128 amount) returns(uint256 amount0, uint256 amount1)
func (_Algebrapool *AlgebrapoolTransactor) Burn(opts *bind.TransactOpts, bottomTick *big.Int, topTick *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Algebrapool.contract.Transact(opts, "burn", bottomTick, topTick, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xa34123a7.
//
// Solidity: function burn(int24 bottomTick, int24 topTick, uint128 amount) returns(uint256 amount0, uint256 amount1)
func (_Algebrapool *AlgebrapoolSession) Burn(bottomTick *big.Int, topTick *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Algebrapool.Contract.Burn(&_Algebrapool.TransactOpts, bottomTick, topTick, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xa34123a7.
//
// Solidity: function burn(int24 bottomTick, int24 topTick, uint128 amount) returns(uint256 amount0, uint256 amount1)
func (_Algebrapool *AlgebrapoolTransactorSession) Burn(bottomTick *big.Int, topTick *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Algebrapool.Contract.Burn(&_Algebrapool.TransactOpts, bottomTick, topTick, amount)
}

// Collect is a paid mutator transaction binding the contract method 0x4f1eb3d8.
//
// Solidity: function collect(address recipient, int24 bottomTick, int24 topTick, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_Algebrapool *AlgebrapoolTransactor) Collect(opts *bind.TransactOpts, recipient common.Address, bottomTick *big.Int, topTick *big.Int, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _Algebrapool.contract.Transact(opts, "collect", recipient, bottomTick, topTick, amount0Requested, amount1Requested)
}

// Collect is a paid mutator transaction binding the contract method 0x4f1eb3d8.
//
// Solidity: function collect(address recipient, int24 bottomTick, int24 topTick, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_Algebrapool *AlgebrapoolSession) Collect(recipient common.Address, bottomTick *big.Int, topTick *big.Int, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _Algebrapool.Contract.Collect(&_Algebrapool.TransactOpts, recipient, bottomTick, topTick, amount0Requested, amount1Requested)
}

// Collect is a paid mutator transaction binding the contract method 0x4f1eb3d8.
//
// Solidity: function collect(address recipient, int24 bottomTick, int24 topTick, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_Algebrapool *AlgebrapoolTransactorSession) Collect(recipient common.Address, bottomTick *big.Int, topTick *big.Int, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _Algebrapool.Contract.Collect(&_Algebrapool.TransactOpts, recipient, bottomTick, topTick, amount0Requested, amount1Requested)
}

// Flash is a paid mutator transaction binding the contract method 0x490e6cbc.
//
// Solidity: function flash(address recipient, uint256 amount0, uint256 amount1, bytes data) returns()
func (_Algebrapool *AlgebrapoolTransactor) Flash(opts *bind.TransactOpts, recipient common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _Algebrapool.contract.Transact(opts, "flash", recipient, amount0, amount1, data)
}

// Flash is a paid mutator transaction binding the contract method 0x490e6cbc.
//
// Solidity: function flash(address recipient, uint256 amount0, uint256 amount1, bytes data) returns()
func (_Algebrapool *AlgebrapoolSession) Flash(recipient common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _Algebrapool.Contract.Flash(&_Algebrapool.TransactOpts, recipient, amount0, amount1, data)
}

// Flash is a paid mutator transaction binding the contract method 0x490e6cbc.
//
// Solidity: function flash(address recipient, uint256 amount0, uint256 amount1, bytes data) returns()
func (_Algebrapool *AlgebrapoolTransactorSession) Flash(recipient common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _Algebrapool.Contract.Flash(&_Algebrapool.TransactOpts, recipient, amount0, amount1, data)
}

// Initialize is a paid mutator transaction binding the contract method 0xf637731d.
//
// Solidity: function initialize(uint160 initialPrice) returns()
func (_Algebrapool *AlgebrapoolTransactor) Initialize(opts *bind.TransactOpts, initialPrice *big.Int) (*types.Transaction, error) {
	return _Algebrapool.contract.Transact(opts, "initialize", initialPrice)
}

// Initialize is a paid mutator transaction binding the contract method 0xf637731d.
//
// Solidity: function initialize(uint160 initialPrice) returns()
func (_Algebrapool *AlgebrapoolSession) Initialize(initialPrice *big.Int) (*types.Transaction, error) {
	return _Algebrapool.Contract.Initialize(&_Algebrapool.TransactOpts, initialPrice)
}

// Initialize is a paid mutator transaction binding the contract method 0xf637731d.
//
// Solidity: function initialize(uint160 initialPrice) returns()
func (_Algebrapool *AlgebrapoolTransactorSession) Initialize(initialPrice *big.Int) (*types.Transaction, error) {
	return _Algebrapool.Contract.Initialize(&_Algebrapool.TransactOpts, initialPrice)
}

// Mint is a paid mutator transaction binding the contract method 0xaafe29c0.
//
// Solidity: function mint(address sender, address recipient, int24 bottomTick, int24 topTick, uint128 liquidityDesired, bytes data) returns(uint256 amount0, uint256 amount1, uint128 liquidityActual)
func (_Algebrapool *AlgebrapoolTransactor) Mint(opts *bind.TransactOpts, sender common.Address, recipient common.Address, bottomTick *big.Int, topTick *big.Int, liquidityDesired *big.Int, data []byte) (*types.Transaction, error) {
	return _Algebrapool.contract.Transact(opts, "mint", sender, recipient, bottomTick, topTick, liquidityDesired, data)
}

// Mint is a paid mutator transaction binding the contract method 0xaafe29c0.
//
// Solidity: function mint(address sender, address recipient, int24 bottomTick, int24 topTick, uint128 liquidityDesired, bytes data) returns(uint256 amount0, uint256 amount1, uint128 liquidityActual)
func (_Algebrapool *AlgebrapoolSession) Mint(sender common.Address, recipient common.Address, bottomTick *big.Int, topTick *big.Int, liquidityDesired *big.Int, data []byte) (*types.Transaction, error) {
	return _Algebrapool.Contract.Mint(&_Algebrapool.TransactOpts, sender, recipient, bottomTick, topTick, liquidityDesired, data)
}

// Mint is a paid mutator transaction binding the contract method 0xaafe29c0.
//
// Solidity: function mint(address sender, address recipient, int24 bottomTick, int24 topTick, uint128 liquidityDesired, bytes data) returns(uint256 amount0, uint256 amount1, uint128 liquidityActual)
func (_Algebrapool *AlgebrapoolTransactorSession) Mint(sender common.Address, recipient common.Address, bottomTick *big.Int, topTick *big.Int, liquidityDesired *big.Int, data []byte) (*types.Transaction, error) {
	return _Algebrapool.Contract.Mint(&_Algebrapool.TransactOpts, sender, recipient, bottomTick, topTick, liquidityDesired, data)
}

// SetCommunityFee is a paid mutator transaction binding the contract method 0xf0b9cf59.
//
// Solidity: function setCommunityFee(uint16 communityFee0, uint16 communityFee1) returns()
func (_Algebrapool *AlgebrapoolTransactor) SetCommunityFee(opts *bind.TransactOpts, communityFee0 uint16, communityFee1 uint16) (*types.Transaction, error) {
	return _Algebrapool.contract.Transact(opts, "setCommunityFee", communityFee0, communityFee1)
}

// SetCommunityFee is a paid mutator transaction binding the contract method 0xf0b9cf59.
//
// Solidity: function setCommunityFee(uint16 communityFee0, uint16 communityFee1) returns()
func (_Algebrapool *AlgebrapoolSession) SetCommunityFee(communityFee0 uint16, communityFee1 uint16) (*types.Transaction, error) {
	return _Algebrapool.Contract.SetCommunityFee(&_Algebrapool.TransactOpts, communityFee0, communityFee1)
}

// SetCommunityFee is a paid mutator transaction binding the contract method 0xf0b9cf59.
//
// Solidity: function setCommunityFee(uint16 communityFee0, uint16 communityFee1) returns()
func (_Algebrapool *AlgebrapoolTransactorSession) SetCommunityFee(communityFee0 uint16, communityFee1 uint16) (*types.Transaction, error) {
	return _Algebrapool.Contract.SetCommunityFee(&_Algebrapool.TransactOpts, communityFee0, communityFee1)
}

// SetIncentive is a paid mutator transaction binding the contract method 0x7c1fe0c8.
//
// Solidity: function setIncentive(address virtualPoolAddress) returns()
func (_Algebrapool *AlgebrapoolTransactor) SetIncentive(opts *bind.TransactOpts, virtualPoolAddress common.Address) (*types.Transaction, error) {
	return _Algebrapool.contract.Transact(opts, "setIncentive", virtualPoolAddress)
}

// SetIncentive is a paid mutator transaction binding the contract method 0x7c1fe0c8.
//
// Solidity: function setIncentive(address virtualPoolAddress) returns()
func (_Algebrapool *AlgebrapoolSession) SetIncentive(virtualPoolAddress common.Address) (*types.Transaction, error) {
	return _Algebrapool.Contract.SetIncentive(&_Algebrapool.TransactOpts, virtualPoolAddress)
}

// SetIncentive is a paid mutator transaction binding the contract method 0x7c1fe0c8.
//
// Solidity: function setIncentive(address virtualPoolAddress) returns()
func (_Algebrapool *AlgebrapoolTransactorSession) SetIncentive(virtualPoolAddress common.Address) (*types.Transaction, error) {
	return _Algebrapool.Contract.SetIncentive(&_Algebrapool.TransactOpts, virtualPoolAddress)
}

// SetLiquidityCooldown is a paid mutator transaction binding the contract method 0x289fe9b0.
//
// Solidity: function setLiquidityCooldown(uint32 newLiquidityCooldown) returns()
func (_Algebrapool *AlgebrapoolTransactor) SetLiquidityCooldown(opts *bind.TransactOpts, newLiquidityCooldown uint32) (*types.Transaction, error) {
	return _Algebrapool.contract.Transact(opts, "setLiquidityCooldown", newLiquidityCooldown)
}

// SetLiquidityCooldown is a paid mutator transaction binding the contract method 0x289fe9b0.
//
// Solidity: function setLiquidityCooldown(uint32 newLiquidityCooldown) returns()
func (_Algebrapool *AlgebrapoolSession) SetLiquidityCooldown(newLiquidityCooldown uint32) (*types.Transaction, error) {
	return _Algebrapool.Contract.SetLiquidityCooldown(&_Algebrapool.TransactOpts, newLiquidityCooldown)
}

// SetLiquidityCooldown is a paid mutator transaction binding the contract method 0x289fe9b0.
//
// Solidity: function setLiquidityCooldown(uint32 newLiquidityCooldown) returns()
func (_Algebrapool *AlgebrapoolTransactorSession) SetLiquidityCooldown(newLiquidityCooldown uint32) (*types.Transaction, error) {
	return _Algebrapool.Contract.SetLiquidityCooldown(&_Algebrapool.TransactOpts, newLiquidityCooldown)
}

// Swap is a paid mutator transaction binding the contract method 0x128acb08.
//
// Solidity: function swap(address recipient, bool zeroToOne, int256 amountRequired, uint160 limitSqrtPrice, bytes data) returns(int256 amount0, int256 amount1)
func (_Algebrapool *AlgebrapoolTransactor) Swap(opts *bind.TransactOpts, recipient common.Address, zeroToOne bool, amountRequired *big.Int, limitSqrtPrice *big.Int, data []byte) (*types.Transaction, error) {
	return _Algebrapool.contract.Transact(opts, "swap", recipient, zeroToOne, amountRequired, limitSqrtPrice, data)
}

// Swap is a paid mutator transaction binding the contract method 0x128acb08.
//
// Solidity: function swap(address recipient, bool zeroToOne, int256 amountRequired, uint160 limitSqrtPrice, bytes data) returns(int256 amount0, int256 amount1)
func (_Algebrapool *AlgebrapoolSession) Swap(recipient common.Address, zeroToOne bool, amountRequired *big.Int, limitSqrtPrice *big.Int, data []byte) (*types.Transaction, error) {
	return _Algebrapool.Contract.Swap(&_Algebrapool.TransactOpts, recipient, zeroToOne, amountRequired, limitSqrtPrice, data)
}

// Swap is a paid mutator transaction binding the contract method 0x128acb08.
//
// Solidity: function swap(address recipient, bool zeroToOne, int256 amountRequired, uint160 limitSqrtPrice, bytes data) returns(int256 amount0, int256 amount1)
func (_Algebrapool *AlgebrapoolTransactorSession) Swap(recipient common.Address, zeroToOne bool, amountRequired *big.Int, limitSqrtPrice *big.Int, data []byte) (*types.Transaction, error) {
	return _Algebrapool.Contract.Swap(&_Algebrapool.TransactOpts, recipient, zeroToOne, amountRequired, limitSqrtPrice, data)
}

// SwapSupportingFeeOnInputTokens is a paid mutator transaction binding the contract method 0x71334694.
//
// Solidity: function swapSupportingFeeOnInputTokens(address sender, address recipient, bool zeroToOne, int256 amountRequired, uint160 limitSqrtPrice, bytes data) returns(int256 amount0, int256 amount1)
func (_Algebrapool *AlgebrapoolTransactor) SwapSupportingFeeOnInputTokens(opts *bind.TransactOpts, sender common.Address, recipient common.Address, zeroToOne bool, amountRequired *big.Int, limitSqrtPrice *big.Int, data []byte) (*types.Transaction, error) {
	return _Algebrapool.contract.Transact(opts, "swapSupportingFeeOnInputTokens", sender, recipient, zeroToOne, amountRequired, limitSqrtPrice, data)
}

// SwapSupportingFeeOnInputTokens is a paid mutator transaction binding the contract method 0x71334694.
//
// Solidity: function swapSupportingFeeOnInputTokens(address sender, address recipient, bool zeroToOne, int256 amountRequired, uint160 limitSqrtPrice, bytes data) returns(int256 amount0, int256 amount1)
func (_Algebrapool *AlgebrapoolSession) SwapSupportingFeeOnInputTokens(sender common.Address, recipient common.Address, zeroToOne bool, amountRequired *big.Int, limitSqrtPrice *big.Int, data []byte) (*types.Transaction, error) {
	return _Algebrapool.Contract.SwapSupportingFeeOnInputTokens(&_Algebrapool.TransactOpts, sender, recipient, zeroToOne, amountRequired, limitSqrtPrice, data)
}

// SwapSupportingFeeOnInputTokens is a paid mutator transaction binding the contract method 0x71334694.
//
// Solidity: function swapSupportingFeeOnInputTokens(address sender, address recipient, bool zeroToOne, int256 amountRequired, uint160 limitSqrtPrice, bytes data) returns(int256 amount0, int256 amount1)
func (_Algebrapool *AlgebrapoolTransactorSession) SwapSupportingFeeOnInputTokens(sender common.Address, recipient common.Address, zeroToOne bool, amountRequired *big.Int, limitSqrtPrice *big.Int, data []byte) (*types.Transaction, error) {
	return _Algebrapool.Contract.SwapSupportingFeeOnInputTokens(&_Algebrapool.TransactOpts, sender, recipient, zeroToOne, amountRequired, limitSqrtPrice, data)
}

// AlgebrapoolBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Algebrapool contract.
type AlgebrapoolBurnIterator struct {
	Event *AlgebrapoolBurn // Event containing the contract specifics and raw log

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
func (it *AlgebrapoolBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlgebrapoolBurn)
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
		it.Event = new(AlgebrapoolBurn)
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
func (it *AlgebrapoolBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlgebrapoolBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlgebrapoolBurn represents a Burn event raised by the Algebrapool contract.
type AlgebrapoolBurn struct {
	Owner           common.Address
	BottomTick      *big.Int
	TopTick         *big.Int
	LiquidityAmount *big.Int
	Amount0         *big.Int
	Amount1         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c.
//
// Solidity: event Burn(address indexed owner, int24 indexed bottomTick, int24 indexed topTick, uint128 liquidityAmount, uint256 amount0, uint256 amount1)
func (_Algebrapool *AlgebrapoolFilterer) FilterBurn(opts *bind.FilterOpts, owner []common.Address, bottomTick []*big.Int, topTick []*big.Int) (*AlgebrapoolBurnIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var bottomTickRule []interface{}
	for _, bottomTickItem := range bottomTick {
		bottomTickRule = append(bottomTickRule, bottomTickItem)
	}
	var topTickRule []interface{}
	for _, topTickItem := range topTick {
		topTickRule = append(topTickRule, topTickItem)
	}

	logs, sub, err := _Algebrapool.contract.FilterLogs(opts, "Burn", ownerRule, bottomTickRule, topTickRule)
	if err != nil {
		return nil, err
	}
	return &AlgebrapoolBurnIterator{contract: _Algebrapool.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c.
//
// Solidity: event Burn(address indexed owner, int24 indexed bottomTick, int24 indexed topTick, uint128 liquidityAmount, uint256 amount0, uint256 amount1)
func (_Algebrapool *AlgebrapoolFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *AlgebrapoolBurn, owner []common.Address, bottomTick []*big.Int, topTick []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var bottomTickRule []interface{}
	for _, bottomTickItem := range bottomTick {
		bottomTickRule = append(bottomTickRule, bottomTickItem)
	}
	var topTickRule []interface{}
	for _, topTickItem := range topTick {
		topTickRule = append(topTickRule, topTickItem)
	}

	logs, sub, err := _Algebrapool.contract.WatchLogs(opts, "Burn", ownerRule, bottomTickRule, topTickRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlgebrapoolBurn)
				if err := _Algebrapool.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c.
//
// Solidity: event Burn(address indexed owner, int24 indexed bottomTick, int24 indexed topTick, uint128 liquidityAmount, uint256 amount0, uint256 amount1)
func (_Algebrapool *AlgebrapoolFilterer) ParseBurn(log types.Log) (*AlgebrapoolBurn, error) {
	event := new(AlgebrapoolBurn)
	if err := _Algebrapool.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlgebrapoolCollectIterator is returned from FilterCollect and is used to iterate over the raw logs and unpacked data for Collect events raised by the Algebrapool contract.
type AlgebrapoolCollectIterator struct {
	Event *AlgebrapoolCollect // Event containing the contract specifics and raw log

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
func (it *AlgebrapoolCollectIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlgebrapoolCollect)
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
		it.Event = new(AlgebrapoolCollect)
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
func (it *AlgebrapoolCollectIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlgebrapoolCollectIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlgebrapoolCollect represents a Collect event raised by the Algebrapool contract.
type AlgebrapoolCollect struct {
	Owner      common.Address
	Recipient  common.Address
	BottomTick *big.Int
	TopTick    *big.Int
	Amount0    *big.Int
	Amount1    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCollect is a free log retrieval operation binding the contract event 0x70935338e69775456a85ddef226c395fb668b63fa0115f5f20610b388e6ca9c0.
//
// Solidity: event Collect(address indexed owner, address recipient, int24 indexed bottomTick, int24 indexed topTick, uint128 amount0, uint128 amount1)
func (_Algebrapool *AlgebrapoolFilterer) FilterCollect(opts *bind.FilterOpts, owner []common.Address, bottomTick []*big.Int, topTick []*big.Int) (*AlgebrapoolCollectIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var bottomTickRule []interface{}
	for _, bottomTickItem := range bottomTick {
		bottomTickRule = append(bottomTickRule, bottomTickItem)
	}
	var topTickRule []interface{}
	for _, topTickItem := range topTick {
		topTickRule = append(topTickRule, topTickItem)
	}

	logs, sub, err := _Algebrapool.contract.FilterLogs(opts, "Collect", ownerRule, bottomTickRule, topTickRule)
	if err != nil {
		return nil, err
	}
	return &AlgebrapoolCollectIterator{contract: _Algebrapool.contract, event: "Collect", logs: logs, sub: sub}, nil
}

// WatchCollect is a free log subscription operation binding the contract event 0x70935338e69775456a85ddef226c395fb668b63fa0115f5f20610b388e6ca9c0.
//
// Solidity: event Collect(address indexed owner, address recipient, int24 indexed bottomTick, int24 indexed topTick, uint128 amount0, uint128 amount1)
func (_Algebrapool *AlgebrapoolFilterer) WatchCollect(opts *bind.WatchOpts, sink chan<- *AlgebrapoolCollect, owner []common.Address, bottomTick []*big.Int, topTick []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var bottomTickRule []interface{}
	for _, bottomTickItem := range bottomTick {
		bottomTickRule = append(bottomTickRule, bottomTickItem)
	}
	var topTickRule []interface{}
	for _, topTickItem := range topTick {
		topTickRule = append(topTickRule, topTickItem)
	}

	logs, sub, err := _Algebrapool.contract.WatchLogs(opts, "Collect", ownerRule, bottomTickRule, topTickRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlgebrapoolCollect)
				if err := _Algebrapool.contract.UnpackLog(event, "Collect", log); err != nil {
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

// ParseCollect is a log parse operation binding the contract event 0x70935338e69775456a85ddef226c395fb668b63fa0115f5f20610b388e6ca9c0.
//
// Solidity: event Collect(address indexed owner, address recipient, int24 indexed bottomTick, int24 indexed topTick, uint128 amount0, uint128 amount1)
func (_Algebrapool *AlgebrapoolFilterer) ParseCollect(log types.Log) (*AlgebrapoolCollect, error) {
	event := new(AlgebrapoolCollect)
	if err := _Algebrapool.contract.UnpackLog(event, "Collect", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlgebrapoolCommunityFeeIterator is returned from FilterCommunityFee and is used to iterate over the raw logs and unpacked data for CommunityFee events raised by the Algebrapool contract.
type AlgebrapoolCommunityFeeIterator struct {
	Event *AlgebrapoolCommunityFee // Event containing the contract specifics and raw log

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
func (it *AlgebrapoolCommunityFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlgebrapoolCommunityFee)
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
		it.Event = new(AlgebrapoolCommunityFee)
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
func (it *AlgebrapoolCommunityFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlgebrapoolCommunityFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlgebrapoolCommunityFee represents a CommunityFee event raised by the Algebrapool contract.
type AlgebrapoolCommunityFee struct {
	CommunityFee0New uint16
	CommunityFee1New uint16
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCommunityFee is a free log retrieval operation binding the contract event 0x370966829959865419a97fc8708e1d348a92142c2cfec7299e264677970174bc.
//
// Solidity: event CommunityFee(uint16 communityFee0New, uint16 communityFee1New)
func (_Algebrapool *AlgebrapoolFilterer) FilterCommunityFee(opts *bind.FilterOpts) (*AlgebrapoolCommunityFeeIterator, error) {

	logs, sub, err := _Algebrapool.contract.FilterLogs(opts, "CommunityFee")
	if err != nil {
		return nil, err
	}
	return &AlgebrapoolCommunityFeeIterator{contract: _Algebrapool.contract, event: "CommunityFee", logs: logs, sub: sub}, nil
}

// WatchCommunityFee is a free log subscription operation binding the contract event 0x370966829959865419a97fc8708e1d348a92142c2cfec7299e264677970174bc.
//
// Solidity: event CommunityFee(uint16 communityFee0New, uint16 communityFee1New)
func (_Algebrapool *AlgebrapoolFilterer) WatchCommunityFee(opts *bind.WatchOpts, sink chan<- *AlgebrapoolCommunityFee) (event.Subscription, error) {

	logs, sub, err := _Algebrapool.contract.WatchLogs(opts, "CommunityFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlgebrapoolCommunityFee)
				if err := _Algebrapool.contract.UnpackLog(event, "CommunityFee", log); err != nil {
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

// ParseCommunityFee is a log parse operation binding the contract event 0x370966829959865419a97fc8708e1d348a92142c2cfec7299e264677970174bc.
//
// Solidity: event CommunityFee(uint16 communityFee0New, uint16 communityFee1New)
func (_Algebrapool *AlgebrapoolFilterer) ParseCommunityFee(log types.Log) (*AlgebrapoolCommunityFee, error) {
	event := new(AlgebrapoolCommunityFee)
	if err := _Algebrapool.contract.UnpackLog(event, "CommunityFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlgebrapoolFeeIterator is returned from FilterFee and is used to iterate over the raw logs and unpacked data for Fee events raised by the Algebrapool contract.
type AlgebrapoolFeeIterator struct {
	Event *AlgebrapoolFee // Event containing the contract specifics and raw log

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
func (it *AlgebrapoolFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlgebrapoolFee)
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
		it.Event = new(AlgebrapoolFee)
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
func (it *AlgebrapoolFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlgebrapoolFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlgebrapoolFee represents a Fee event raised by the Algebrapool contract.
type AlgebrapoolFee struct {
	Fee uint16
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFee is a free log retrieval operation binding the contract event 0x598b9f043c813aa6be3426ca60d1c65d17256312890be5118dab55b0775ebe2a.
//
// Solidity: event Fee(uint16 fee)
func (_Algebrapool *AlgebrapoolFilterer) FilterFee(opts *bind.FilterOpts) (*AlgebrapoolFeeIterator, error) {

	logs, sub, err := _Algebrapool.contract.FilterLogs(opts, "Fee")
	if err != nil {
		return nil, err
	}
	return &AlgebrapoolFeeIterator{contract: _Algebrapool.contract, event: "Fee", logs: logs, sub: sub}, nil
}

// WatchFee is a free log subscription operation binding the contract event 0x598b9f043c813aa6be3426ca60d1c65d17256312890be5118dab55b0775ebe2a.
//
// Solidity: event Fee(uint16 fee)
func (_Algebrapool *AlgebrapoolFilterer) WatchFee(opts *bind.WatchOpts, sink chan<- *AlgebrapoolFee) (event.Subscription, error) {

	logs, sub, err := _Algebrapool.contract.WatchLogs(opts, "Fee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlgebrapoolFee)
				if err := _Algebrapool.contract.UnpackLog(event, "Fee", log); err != nil {
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

// ParseFee is a log parse operation binding the contract event 0x598b9f043c813aa6be3426ca60d1c65d17256312890be5118dab55b0775ebe2a.
//
// Solidity: event Fee(uint16 fee)
func (_Algebrapool *AlgebrapoolFilterer) ParseFee(log types.Log) (*AlgebrapoolFee, error) {
	event := new(AlgebrapoolFee)
	if err := _Algebrapool.contract.UnpackLog(event, "Fee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlgebrapoolFlashIterator is returned from FilterFlash and is used to iterate over the raw logs and unpacked data for Flash events raised by the Algebrapool contract.
type AlgebrapoolFlashIterator struct {
	Event *AlgebrapoolFlash // Event containing the contract specifics and raw log

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
func (it *AlgebrapoolFlashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlgebrapoolFlash)
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
		it.Event = new(AlgebrapoolFlash)
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
func (it *AlgebrapoolFlashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlgebrapoolFlashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlgebrapoolFlash represents a Flash event raised by the Algebrapool contract.
type AlgebrapoolFlash struct {
	Sender    common.Address
	Recipient common.Address
	Amount0   *big.Int
	Amount1   *big.Int
	Paid0     *big.Int
	Paid1     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFlash is a free log retrieval operation binding the contract event 0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633.
//
// Solidity: event Flash(address indexed sender, address indexed recipient, uint256 amount0, uint256 amount1, uint256 paid0, uint256 paid1)
func (_Algebrapool *AlgebrapoolFilterer) FilterFlash(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*AlgebrapoolFlashIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Algebrapool.contract.FilterLogs(opts, "Flash", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &AlgebrapoolFlashIterator{contract: _Algebrapool.contract, event: "Flash", logs: logs, sub: sub}, nil
}

// WatchFlash is a free log subscription operation binding the contract event 0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633.
//
// Solidity: event Flash(address indexed sender, address indexed recipient, uint256 amount0, uint256 amount1, uint256 paid0, uint256 paid1)
func (_Algebrapool *AlgebrapoolFilterer) WatchFlash(opts *bind.WatchOpts, sink chan<- *AlgebrapoolFlash, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Algebrapool.contract.WatchLogs(opts, "Flash", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlgebrapoolFlash)
				if err := _Algebrapool.contract.UnpackLog(event, "Flash", log); err != nil {
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

// ParseFlash is a log parse operation binding the contract event 0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633.
//
// Solidity: event Flash(address indexed sender, address indexed recipient, uint256 amount0, uint256 amount1, uint256 paid0, uint256 paid1)
func (_Algebrapool *AlgebrapoolFilterer) ParseFlash(log types.Log) (*AlgebrapoolFlash, error) {
	event := new(AlgebrapoolFlash)
	if err := _Algebrapool.contract.UnpackLog(event, "Flash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlgebrapoolIncentiveIterator is returned from FilterIncentive and is used to iterate over the raw logs and unpacked data for Incentive events raised by the Algebrapool contract.
type AlgebrapoolIncentiveIterator struct {
	Event *AlgebrapoolIncentive // Event containing the contract specifics and raw log

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
func (it *AlgebrapoolIncentiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlgebrapoolIncentive)
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
		it.Event = new(AlgebrapoolIncentive)
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
func (it *AlgebrapoolIncentiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlgebrapoolIncentiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlgebrapoolIncentive represents a Incentive event raised by the Algebrapool contract.
type AlgebrapoolIncentive struct {
	VirtualPoolAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterIncentive is a free log retrieval operation binding the contract event 0x915c5369e6580733735d1c2e30ca20dcaa395697a041033c9f35f80f53525e84.
//
// Solidity: event Incentive(address indexed virtualPoolAddress)
func (_Algebrapool *AlgebrapoolFilterer) FilterIncentive(opts *bind.FilterOpts, virtualPoolAddress []common.Address) (*AlgebrapoolIncentiveIterator, error) {

	var virtualPoolAddressRule []interface{}
	for _, virtualPoolAddressItem := range virtualPoolAddress {
		virtualPoolAddressRule = append(virtualPoolAddressRule, virtualPoolAddressItem)
	}

	logs, sub, err := _Algebrapool.contract.FilterLogs(opts, "Incentive", virtualPoolAddressRule)
	if err != nil {
		return nil, err
	}
	return &AlgebrapoolIncentiveIterator{contract: _Algebrapool.contract, event: "Incentive", logs: logs, sub: sub}, nil
}

// WatchIncentive is a free log subscription operation binding the contract event 0x915c5369e6580733735d1c2e30ca20dcaa395697a041033c9f35f80f53525e84.
//
// Solidity: event Incentive(address indexed virtualPoolAddress)
func (_Algebrapool *AlgebrapoolFilterer) WatchIncentive(opts *bind.WatchOpts, sink chan<- *AlgebrapoolIncentive, virtualPoolAddress []common.Address) (event.Subscription, error) {

	var virtualPoolAddressRule []interface{}
	for _, virtualPoolAddressItem := range virtualPoolAddress {
		virtualPoolAddressRule = append(virtualPoolAddressRule, virtualPoolAddressItem)
	}

	logs, sub, err := _Algebrapool.contract.WatchLogs(opts, "Incentive", virtualPoolAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlgebrapoolIncentive)
				if err := _Algebrapool.contract.UnpackLog(event, "Incentive", log); err != nil {
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

// ParseIncentive is a log parse operation binding the contract event 0x915c5369e6580733735d1c2e30ca20dcaa395697a041033c9f35f80f53525e84.
//
// Solidity: event Incentive(address indexed virtualPoolAddress)
func (_Algebrapool *AlgebrapoolFilterer) ParseIncentive(log types.Log) (*AlgebrapoolIncentive, error) {
	event := new(AlgebrapoolIncentive)
	if err := _Algebrapool.contract.UnpackLog(event, "Incentive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlgebrapoolInitializeIterator is returned from FilterInitialize and is used to iterate over the raw logs and unpacked data for Initialize events raised by the Algebrapool contract.
type AlgebrapoolInitializeIterator struct {
	Event *AlgebrapoolInitialize // Event containing the contract specifics and raw log

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
func (it *AlgebrapoolInitializeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlgebrapoolInitialize)
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
		it.Event = new(AlgebrapoolInitialize)
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
func (it *AlgebrapoolInitializeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlgebrapoolInitializeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlgebrapoolInitialize represents a Initialize event raised by the Algebrapool contract.
type AlgebrapoolInitialize struct {
	Price *big.Int
	Tick  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterInitialize is a free log retrieval operation binding the contract event 0x98636036cb66a9c19a37435efc1e90142190214e8abeb821bdba3f2990dd4c95.
//
// Solidity: event Initialize(uint160 price, int24 tick)
func (_Algebrapool *AlgebrapoolFilterer) FilterInitialize(opts *bind.FilterOpts) (*AlgebrapoolInitializeIterator, error) {

	logs, sub, err := _Algebrapool.contract.FilterLogs(opts, "Initialize")
	if err != nil {
		return nil, err
	}
	return &AlgebrapoolInitializeIterator{contract: _Algebrapool.contract, event: "Initialize", logs: logs, sub: sub}, nil
}

// WatchInitialize is a free log subscription operation binding the contract event 0x98636036cb66a9c19a37435efc1e90142190214e8abeb821bdba3f2990dd4c95.
//
// Solidity: event Initialize(uint160 price, int24 tick)
func (_Algebrapool *AlgebrapoolFilterer) WatchInitialize(opts *bind.WatchOpts, sink chan<- *AlgebrapoolInitialize) (event.Subscription, error) {

	logs, sub, err := _Algebrapool.contract.WatchLogs(opts, "Initialize")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlgebrapoolInitialize)
				if err := _Algebrapool.contract.UnpackLog(event, "Initialize", log); err != nil {
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

// ParseInitialize is a log parse operation binding the contract event 0x98636036cb66a9c19a37435efc1e90142190214e8abeb821bdba3f2990dd4c95.
//
// Solidity: event Initialize(uint160 price, int24 tick)
func (_Algebrapool *AlgebrapoolFilterer) ParseInitialize(log types.Log) (*AlgebrapoolInitialize, error) {
	event := new(AlgebrapoolInitialize)
	if err := _Algebrapool.contract.UnpackLog(event, "Initialize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlgebrapoolLiquidityCooldownIterator is returned from FilterLiquidityCooldown and is used to iterate over the raw logs and unpacked data for LiquidityCooldown events raised by the Algebrapool contract.
type AlgebrapoolLiquidityCooldownIterator struct {
	Event *AlgebrapoolLiquidityCooldown // Event containing the contract specifics and raw log

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
func (it *AlgebrapoolLiquidityCooldownIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlgebrapoolLiquidityCooldown)
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
		it.Event = new(AlgebrapoolLiquidityCooldown)
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
func (it *AlgebrapoolLiquidityCooldownIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlgebrapoolLiquidityCooldownIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlgebrapoolLiquidityCooldown represents a LiquidityCooldown event raised by the Algebrapool contract.
type AlgebrapoolLiquidityCooldown struct {
	LiquidityCooldown uint32
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLiquidityCooldown is a free log retrieval operation binding the contract event 0xb5e51602371b0e74f991b6e965cd7d32b4b14c7e6ede6d1298037650a0e1405f.
//
// Solidity: event LiquidityCooldown(uint32 liquidityCooldown)
func (_Algebrapool *AlgebrapoolFilterer) FilterLiquidityCooldown(opts *bind.FilterOpts) (*AlgebrapoolLiquidityCooldownIterator, error) {

	logs, sub, err := _Algebrapool.contract.FilterLogs(opts, "LiquidityCooldown")
	if err != nil {
		return nil, err
	}
	return &AlgebrapoolLiquidityCooldownIterator{contract: _Algebrapool.contract, event: "LiquidityCooldown", logs: logs, sub: sub}, nil
}

// WatchLiquidityCooldown is a free log subscription operation binding the contract event 0xb5e51602371b0e74f991b6e965cd7d32b4b14c7e6ede6d1298037650a0e1405f.
//
// Solidity: event LiquidityCooldown(uint32 liquidityCooldown)
func (_Algebrapool *AlgebrapoolFilterer) WatchLiquidityCooldown(opts *bind.WatchOpts, sink chan<- *AlgebrapoolLiquidityCooldown) (event.Subscription, error) {

	logs, sub, err := _Algebrapool.contract.WatchLogs(opts, "LiquidityCooldown")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlgebrapoolLiquidityCooldown)
				if err := _Algebrapool.contract.UnpackLog(event, "LiquidityCooldown", log); err != nil {
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

// ParseLiquidityCooldown is a log parse operation binding the contract event 0xb5e51602371b0e74f991b6e965cd7d32b4b14c7e6ede6d1298037650a0e1405f.
//
// Solidity: event LiquidityCooldown(uint32 liquidityCooldown)
func (_Algebrapool *AlgebrapoolFilterer) ParseLiquidityCooldown(log types.Log) (*AlgebrapoolLiquidityCooldown, error) {
	event := new(AlgebrapoolLiquidityCooldown)
	if err := _Algebrapool.contract.UnpackLog(event, "LiquidityCooldown", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlgebrapoolMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Algebrapool contract.
type AlgebrapoolMintIterator struct {
	Event *AlgebrapoolMint // Event containing the contract specifics and raw log

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
func (it *AlgebrapoolMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlgebrapoolMint)
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
		it.Event = new(AlgebrapoolMint)
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
func (it *AlgebrapoolMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlgebrapoolMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlgebrapoolMint represents a Mint event raised by the Algebrapool contract.
type AlgebrapoolMint struct {
	Sender          common.Address
	Owner           common.Address
	BottomTick      *big.Int
	TopTick         *big.Int
	LiquidityAmount *big.Int
	Amount0         *big.Int
	Amount1         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde.
//
// Solidity: event Mint(address sender, address indexed owner, int24 indexed bottomTick, int24 indexed topTick, uint128 liquidityAmount, uint256 amount0, uint256 amount1)
func (_Algebrapool *AlgebrapoolFilterer) FilterMint(opts *bind.FilterOpts, owner []common.Address, bottomTick []*big.Int, topTick []*big.Int) (*AlgebrapoolMintIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var bottomTickRule []interface{}
	for _, bottomTickItem := range bottomTick {
		bottomTickRule = append(bottomTickRule, bottomTickItem)
	}
	var topTickRule []interface{}
	for _, topTickItem := range topTick {
		topTickRule = append(topTickRule, topTickItem)
	}

	logs, sub, err := _Algebrapool.contract.FilterLogs(opts, "Mint", ownerRule, bottomTickRule, topTickRule)
	if err != nil {
		return nil, err
	}
	return &AlgebrapoolMintIterator{contract: _Algebrapool.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde.
//
// Solidity: event Mint(address sender, address indexed owner, int24 indexed bottomTick, int24 indexed topTick, uint128 liquidityAmount, uint256 amount0, uint256 amount1)
func (_Algebrapool *AlgebrapoolFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *AlgebrapoolMint, owner []common.Address, bottomTick []*big.Int, topTick []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var bottomTickRule []interface{}
	for _, bottomTickItem := range bottomTick {
		bottomTickRule = append(bottomTickRule, bottomTickItem)
	}
	var topTickRule []interface{}
	for _, topTickItem := range topTick {
		topTickRule = append(topTickRule, topTickItem)
	}

	logs, sub, err := _Algebrapool.contract.WatchLogs(opts, "Mint", ownerRule, bottomTickRule, topTickRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlgebrapoolMint)
				if err := _Algebrapool.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde.
//
// Solidity: event Mint(address sender, address indexed owner, int24 indexed bottomTick, int24 indexed topTick, uint128 liquidityAmount, uint256 amount0, uint256 amount1)
func (_Algebrapool *AlgebrapoolFilterer) ParseMint(log types.Log) (*AlgebrapoolMint, error) {
	event := new(AlgebrapoolMint)
	if err := _Algebrapool.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlgebrapoolSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Algebrapool contract.
type AlgebrapoolSwapIterator struct {
	Event *AlgebrapoolSwap // Event containing the contract specifics and raw log

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
func (it *AlgebrapoolSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlgebrapoolSwap)
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
		it.Event = new(AlgebrapoolSwap)
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
func (it *AlgebrapoolSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlgebrapoolSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlgebrapoolSwap represents a Swap event raised by the Algebrapool contract.
type AlgebrapoolSwap struct {
	Sender    common.Address
	Recipient common.Address
	Amount0   *big.Int
	Amount1   *big.Int
	Price     *big.Int
	Liquidity *big.Int
	Tick      *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, int256 amount0, int256 amount1, uint160 price, uint128 liquidity, int24 tick)
func (_Algebrapool *AlgebrapoolFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*AlgebrapoolSwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Algebrapool.contract.FilterLogs(opts, "Swap", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &AlgebrapoolSwapIterator{contract: _Algebrapool.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, int256 amount0, int256 amount1, uint160 price, uint128 liquidity, int24 tick)
func (_Algebrapool *AlgebrapoolFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *AlgebrapoolSwap, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Algebrapool.contract.WatchLogs(opts, "Swap", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlgebrapoolSwap)
				if err := _Algebrapool.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, int256 amount0, int256 amount1, uint160 price, uint128 liquidity, int24 tick)
func (_Algebrapool *AlgebrapoolFilterer) ParseSwap(log types.Log) (*AlgebrapoolSwap, error) {
	event := new(AlgebrapoolSwap)
	if err := _Algebrapool.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
