// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package DppAdvanced

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

// PMMPricingPMMState is an auto generated low-level Go binding around an user-defined struct.
type PMMPricingPMMState struct {
	I  *big.Int
	K  *big.Int
	B  *big.Int
	Q  *big.Int
	B0 *big.Int
	Q0 *big.Int
	R  uint8
}

// DppAdvancedABI is the input ABI used to generate the binding from.
const DppAdvancedABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"assetTo\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"}],\"name\":\"DODOFlashLoan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"DODOSwap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLpFeeRate\",\"type\":\"uint256\"}],\"name\":\"LpFeeRateChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferPrepared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumPMMPricing.RState\",\"name\":\"newRState\",\"type\":\"uint8\"}],\"name\":\"RChange\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_BASE_PRICE_CUMULATIVE_LAST_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_BASE_RESERVE_\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"\",\"type\":\"uint112\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_BASE_TARGET_\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"\",\"type\":\"uint112\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_BASE_TOKEN_\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_BLOCK_TIMESTAMP_LAST_\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_IS_OPEN_TWAP_\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_I_\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_K_\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_LP_FEE_RATE_\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_MAINTAINER_\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_MT_FEE_RATE_MODEL_\",\"outputs\":[{\"internalType\":\"contractIFeeRateModel\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_NEW_OWNER_\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_OWNER_\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_QUOTE_RESERVE_\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"\",\"type\":\"uint112\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_QUOTE_TARGET_\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"\",\"type\":\"uint112\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_QUOTE_TOKEN_\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_RState_\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quoteAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"assetTo\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"flashLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBaseInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"input\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"midPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPMMState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"K\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"B\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Q\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"B0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Q0\",\"type\":\"uint256\"},{\"internalType\":\"enumPMMPricing.RState\",\"name\":\"R\",\"type\":\"uint8\"}],\"internalType\":\"structPMMPricing.PMMState\",\"name\":\"state\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPMMStateForCall\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"K\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"B\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Q\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"B0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Q0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"R\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getQuoteInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"input\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lpFeeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mtFeeRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVaultReserve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"baseReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quoteReserve\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maintainer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"baseTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lpFeeRate\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mtFeeRateModel\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"k\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOpenTWAP\",\"type\":\"bool\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"initOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"payBaseAmount\",\"type\":\"uint256\"}],\"name\":\"querySellBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"receiveQuoteAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mtFee\",\"type\":\"uint256\"},{\"internalType\":\"enumPMMPricing.RState\",\"name\":\"newRState\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"newBaseTarget\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"payQuoteAmount\",\"type\":\"uint256\"}],\"name\":\"querySellQuote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"receiveBaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mtFee\",\"type\":\"uint256\"},{\"internalType\":\"enumPMMPricing.RState\",\"name\":\"newRState\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"newQuoteTarget\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ratioSync\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetTo\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newLpFeeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newI\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newK\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseOutAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quoteOutAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBaseReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minQuoteReserve\",\"type\":\"uint256\"}],\"name\":\"reset\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"retrieve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"sellBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"receiveQuoteAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"sellQuote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"receiveBaseAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newLpFeeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newI\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newK\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBaseReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minQuoteReserve\",\"type\":\"uint256\"}],\"name\":\"tuneParameters\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newI\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBaseReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minQuoteReserve\",\"type\":\"uint256\"}],\"name\":\"tunePrice\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// DppAdvanced is an auto generated Go binding around an Ethereum contract.
type DppAdvanced struct {
	DppAdvancedCaller     // Read-only binding to the contract
	DppAdvancedTransactor // Write-only binding to the contract
	DppAdvancedFilterer   // Log filterer for contract events
}

// DppAdvancedCaller is an auto generated read-only Go binding around an Ethereum contract.
type DppAdvancedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DppAdvancedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DppAdvancedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DppAdvancedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DppAdvancedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DppAdvancedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DppAdvancedSession struct {
	Contract     *DppAdvanced      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DppAdvancedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DppAdvancedCallerSession struct {
	Contract *DppAdvancedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DppAdvancedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DppAdvancedTransactorSession struct {
	Contract     *DppAdvancedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DppAdvancedRaw is an auto generated low-level Go binding around an Ethereum contract.
type DppAdvancedRaw struct {
	Contract *DppAdvanced // Generic contract binding to access the raw methods on
}

// DppAdvancedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DppAdvancedCallerRaw struct {
	Contract *DppAdvancedCaller // Generic read-only contract binding to access the raw methods on
}

// DppAdvancedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DppAdvancedTransactorRaw struct {
	Contract *DppAdvancedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDppAdvanced creates a new instance of DppAdvanced, bound to a specific deployed contract.
func NewDppAdvanced(address common.Address, backend bind.ContractBackend) (*DppAdvanced, error) {
	contract, err := bindDppAdvanced(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DppAdvanced{DppAdvancedCaller: DppAdvancedCaller{contract: contract}, DppAdvancedTransactor: DppAdvancedTransactor{contract: contract}, DppAdvancedFilterer: DppAdvancedFilterer{contract: contract}}, nil
}

// NewDppAdvancedCaller creates a new read-only instance of DppAdvanced, bound to a specific deployed contract.
func NewDppAdvancedCaller(address common.Address, caller bind.ContractCaller) (*DppAdvancedCaller, error) {
	contract, err := bindDppAdvanced(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DppAdvancedCaller{contract: contract}, nil
}

// NewDppAdvancedTransactor creates a new write-only instance of DppAdvanced, bound to a specific deployed contract.
func NewDppAdvancedTransactor(address common.Address, transactor bind.ContractTransactor) (*DppAdvancedTransactor, error) {
	contract, err := bindDppAdvanced(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DppAdvancedTransactor{contract: contract}, nil
}

// NewDppAdvancedFilterer creates a new log filterer instance of DppAdvanced, bound to a specific deployed contract.
func NewDppAdvancedFilterer(address common.Address, filterer bind.ContractFilterer) (*DppAdvancedFilterer, error) {
	contract, err := bindDppAdvanced(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DppAdvancedFilterer{contract: contract}, nil
}

// bindDppAdvanced binds a generic wrapper to an already deployed contract.
func bindDppAdvanced(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DppAdvancedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DppAdvanced *DppAdvancedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DppAdvanced.Contract.DppAdvancedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DppAdvanced *DppAdvancedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DppAdvanced.Contract.DppAdvancedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DppAdvanced *DppAdvancedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DppAdvanced.Contract.DppAdvancedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DppAdvanced *DppAdvancedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DppAdvanced.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DppAdvanced *DppAdvancedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DppAdvanced.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DppAdvanced *DppAdvancedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DppAdvanced.Contract.contract.Transact(opts, method, params...)
}

// BASEPRICECUMULATIVELAST is a free data retrieval call binding the contract method 0xfe24cb7f.
//
// Solidity: function _BASE_PRICE_CUMULATIVE_LAST_() view returns(uint256)
func (_DppAdvanced *DppAdvancedCaller) BASEPRICECUMULATIVELAST(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DppAdvanced.contract.Call(opts, &out, "_BASE_PRICE_CUMULATIVE_LAST_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASEPRICECUMULATIVELAST is a free data retrieval call binding the contract method 0xfe24cb7f.
//
// Solidity: function _BASE_PRICE_CUMULATIVE_LAST_() view returns(uint256)
func (_DppAdvanced *DppAdvancedSession) BASEPRICECUMULATIVELAST() (*big.Int, error) {
	return _DppAdvanced.Contract.BASEPRICECUMULATIVELAST(&_DppAdvanced.CallOpts)
}

// BASEPRICECUMULATIVELAST is a free data retrieval call binding the contract method 0xfe24cb7f.
//
// Solidity: function _BASE_PRICE_CUMULATIVE_LAST_() view returns(uint256)
func (_DppAdvanced *DppAdvancedCallerSession) BASEPRICECUMULATIVELAST() (*big.Int, error) {
	return _DppAdvanced.Contract.BASEPRICECUMULATIVELAST(&_DppAdvanced.CallOpts)
}

// BASERESERVE is a free data retrieval call binding the contract method 0x7d721504.
//
// Solidity: function _BASE_RESERVE_() view returns(uint112)
func (_DppAdvanced *DppAdvancedCaller) BASERESERVE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DppAdvanced.contract.Call(opts, &out, "_BASE_RESERVE_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASERESERVE is a free data retrieval call binding the contract method 0x7d721504.
//
// Solidity: function _BASE_RESERVE_() view returns(uint112)
func (_DppAdvanced *DppAdvancedSession) BASERESERVE() (*big.Int, error) {
	return _DppAdvanced.Contract.BASERESERVE(&_DppAdvanced.CallOpts)
}

// BASERESERVE is a free data retrieval call binding the contract method 0x7d721504.
//
// Solidity: function _BASE_RESERVE_() view returns(uint112)
func (_DppAdvanced *DppAdvancedCallerSession) BASERESERVE() (*big.Int, error) {
	return _DppAdvanced.Contract.BASERESERVE(&_DppAdvanced.CallOpts)
}

// BASETARGET is a free data retrieval call binding the contract method 0xe539ef49.
//
// Solidity: function _BASE_TARGET_() view returns(uint112)
func (_DppAdvanced *DppAdvancedCaller) BASETARGET(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DppAdvanced.contract.Call(opts, &out, "_BASE_TARGET_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASETARGET is a free data retrieval call binding the contract method 0xe539ef49.
//
// Solidity: function _BASE_TARGET_() view returns(uint112)
func (_DppAdvanced *DppAdvancedSession) BASETARGET() (*big.Int, error) {
	return _DppAdvanced.Contract.BASETARGET(&_DppAdvanced.CallOpts)
}

// BASETARGET is a free data retrieval call binding the contract method 0xe539ef49.
//
// Solidity: function _BASE_TARGET_() view returns(uint112)
func (_DppAdvanced *DppAdvancedCallerSession) BASETARGET() (*big.Int, error) {
	return _DppAdvanced.Contract.BASETARGET(&_DppAdvanced.CallOpts)
}

// BASETOKEN is a free data retrieval call binding the contract method 0x4a248d2a.
//
// Solidity: function _BASE_TOKEN_() view returns(address)
func (_DppAdvanced *DppAdvancedCaller) BASETOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DppAdvanced.contract.Call(opts, &out, "_BASE_TOKEN_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BASETOKEN is a free data retrieval call binding the contract method 0x4a248d2a.
//
// Solidity: function _BASE_TOKEN_() view returns(address)
func (_DppAdvanced *DppAdvancedSession) BASETOKEN() (common.Address, error) {
	return _DppAdvanced.Contract.BASETOKEN(&_DppAdvanced.CallOpts)
}

// BASETOKEN is a free data retrieval call binding the contract method 0x4a248d2a.
//
// Solidity: function _BASE_TOKEN_() view returns(address)
func (_DppAdvanced *DppAdvancedCallerSession) BASETOKEN() (common.Address, error) {
	return _DppAdvanced.Contract.BASETOKEN(&_DppAdvanced.CallOpts)
}

// BLOCKTIMESTAMPLAST is a free data retrieval call binding the contract method 0x880a4d87.
//
// Solidity: function _BLOCK_TIMESTAMP_LAST_() view returns(uint32)
func (_DppAdvanced *DppAdvancedCaller) BLOCKTIMESTAMPLAST(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DppAdvanced.contract.Call(opts, &out, "_BLOCK_TIMESTAMP_LAST_")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// BLOCKTIMESTAMPLAST is a free data retrieval call binding the contract method 0x880a4d87.
//
// Solidity: function _BLOCK_TIMESTAMP_LAST_() view returns(uint32)
func (_DppAdvanced *DppAdvancedSession) BLOCKTIMESTAMPLAST() (uint32, error) {
	return _DppAdvanced.Contract.BLOCKTIMESTAMPLAST(&_DppAdvanced.CallOpts)
}

// BLOCKTIMESTAMPLAST is a free data retrieval call binding the contract method 0x880a4d87.
//
// Solidity: function _BLOCK_TIMESTAMP_LAST_() view returns(uint32)
func (_DppAdvanced *DppAdvancedCallerSession) BLOCKTIMESTAMPLAST() (uint32, error) {
	return _DppAdvanced.Contract.BLOCKTIMESTAMPLAST(&_DppAdvanced.CallOpts)
}

// ISOPENTWAP is a free data retrieval call binding the contract method 0x2df6cb48.
//
// Solidity: function _IS_OPEN_TWAP_() view returns(bool)
func (_DppAdvanced *DppAdvancedCaller) ISOPENTWAP(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DppAdvanced.contract.Call(opts, &out, "_IS_OPEN_TWAP_")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ISOPENTWAP is a free data retrieval call binding the contract method 0x2df6cb48.
//
// Solidity: function _IS_OPEN_TWAP_() view returns(bool)
func (_DppAdvanced *DppAdvancedSession) ISOPENTWAP() (bool, error) {
	return _DppAdvanced.Contract.ISOPENTWAP(&_DppAdvanced.CallOpts)
}

// ISOPENTWAP is a free data retrieval call binding the contract method 0x2df6cb48.
//
// Solidity: function _IS_OPEN_TWAP_() view returns(bool)
func (_DppAdvanced *DppAdvancedCallerSession) ISOPENTWAP() (bool, error) {
	return _DppAdvanced.Contract.ISOPENTWAP(&_DppAdvanced.CallOpts)
}

// I is a free data retrieval call binding the contract method 0xf811d692.
//
// Solidity: function _I_() view returns(uint128)
func (_DppAdvanced *DppAdvancedCaller) I(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DppAdvanced.contract.Call(opts, &out, "_I_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// I is a free data retrieval call binding the contract method 0xf811d692.
//
// Solidity: function _I_() view returns(uint128)
func (_DppAdvanced *DppAdvancedSession) I() (*big.Int, error) {
	return _DppAdvanced.Contract.I(&_DppAdvanced.CallOpts)
}

// I is a free data retrieval call binding the contract method 0xf811d692.
//
// Solidity: function _I_() view returns(uint128)
func (_DppAdvanced *DppAdvancedCallerSession) I() (*big.Int, error) {
	return _DppAdvanced.Contract.I(&_DppAdvanced.CallOpts)
}

// K is a free data retrieval call binding the contract method 0xec2fd46d.
//
// Solidity: function _K_() view returns(uint64)
func (_DppAdvanced *DppAdvancedCaller) K(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _DppAdvanced.contract.Call(opts, &out, "_K_")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// K is a free data retrieval call binding the contract method 0xec2fd46d.
//
// Solidity: function _K_() view returns(uint64)
func (_DppAdvanced *DppAdvancedSession) K() (uint64, error) {
	return _DppAdvanced.Contract.K(&_DppAdvanced.CallOpts)
}

// K is a free data retrieval call binding the contract method 0xec2fd46d.
//
// Solidity: function _K_() view returns(uint64)
func (_DppAdvanced *DppAdvancedCallerSession) K() (uint64, error) {
	return _DppAdvanced.Contract.K(&_DppAdvanced.CallOpts)
}

// LPFEERATE is a free data retrieval call binding the contract method 0xab44a7a3.
//
// Solidity: function _LP_FEE_RATE_() view returns(uint64)
func (_DppAdvanced *DppAdvancedCaller) LPFEERATE(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _DppAdvanced.contract.Call(opts, &out, "_LP_FEE_RATE_")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LPFEERATE is a free data retrieval call binding the contract method 0xab44a7a3.
//
// Solidity: function _LP_FEE_RATE_() view returns(uint64)
func (_DppAdvanced *DppAdvancedSession) LPFEERATE() (uint64, error) {
	return _DppAdvanced.Contract.LPFEERATE(&_DppAdvanced.CallOpts)
}

// LPFEERATE is a free data retrieval call binding the contract method 0xab44a7a3.
//
// Solidity: function _LP_FEE_RATE_() view returns(uint64)
func (_DppAdvanced *DppAdvancedCallerSession) LPFEERATE() (uint64, error) {
	return _DppAdvanced.Contract.LPFEERATE(&_DppAdvanced.CallOpts)
}

// MAINTAINER is a free data retrieval call binding the contract method 0x4322ec83.
//
// Solidity: function _MAINTAINER_() view returns(address)
func (_DppAdvanced *DppAdvancedCaller) MAINTAINER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DppAdvanced.contract.Call(opts, &out, "_MAINTAINER_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MAINTAINER is a free data retrieval call binding the contract method 0x4322ec83.
//
// Solidity: function _MAINTAINER_() view returns(address)
func (_DppAdvanced *DppAdvancedSession) MAINTAINER() (common.Address, error) {
	return _DppAdvanced.Contract.MAINTAINER(&_DppAdvanced.CallOpts)
}

// MAINTAINER is a free data retrieval call binding the contract method 0x4322ec83.
//
// Solidity: function _MAINTAINER_() view returns(address)
func (_DppAdvanced *DppAdvancedCallerSession) MAINTAINER() (common.Address, error) {
	return _DppAdvanced.Contract.MAINTAINER(&_DppAdvanced.CallOpts)
}

// MTFEERATEMODEL is a free data retrieval call binding the contract method 0xf6b06e70.
//
// Solidity: function _MT_FEE_RATE_MODEL_() view returns(address)
func (_DppAdvanced *DppAdvancedCaller) MTFEERATEMODEL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DppAdvanced.contract.Call(opts, &out, "_MT_FEE_RATE_MODEL_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MTFEERATEMODEL is a free data retrieval call binding the contract method 0xf6b06e70.
//
// Solidity: function _MT_FEE_RATE_MODEL_() view returns(address)
func (_DppAdvanced *DppAdvancedSession) MTFEERATEMODEL() (common.Address, error) {
	return _DppAdvanced.Contract.MTFEERATEMODEL(&_DppAdvanced.CallOpts)
}

// MTFEERATEMODEL is a free data retrieval call binding the contract method 0xf6b06e70.
//
// Solidity: function _MT_FEE_RATE_MODEL_() view returns(address)
func (_DppAdvanced *DppAdvancedCallerSession) MTFEERATEMODEL() (common.Address, error) {
	return _DppAdvanced.Contract.MTFEERATEMODEL(&_DppAdvanced.CallOpts)
}

// NEWOWNER is a free data retrieval call binding the contract method 0x8456db15.
//
// Solidity: function _NEW_OWNER_() view returns(address)
func (_DppAdvanced *DppAdvancedCaller) NEWOWNER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DppAdvanced.contract.Call(opts, &out, "_NEW_OWNER_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NEWOWNER is a free data retrieval call binding the contract method 0x8456db15.
//
// Solidity: function _NEW_OWNER_() view returns(address)
func (_DppAdvanced *DppAdvancedSession) NEWOWNER() (common.Address, error) {
	return _DppAdvanced.Contract.NEWOWNER(&_DppAdvanced.CallOpts)
}

// NEWOWNER is a free data retrieval call binding the contract method 0x8456db15.
//
// Solidity: function _NEW_OWNER_() view returns(address)
func (_DppAdvanced *DppAdvancedCallerSession) NEWOWNER() (common.Address, error) {
	return _DppAdvanced.Contract.NEWOWNER(&_DppAdvanced.CallOpts)
}

// OWNER is a free data retrieval call binding the contract method 0x16048bc4.
//
// Solidity: function _OWNER_() view returns(address)
func (_DppAdvanced *DppAdvancedCaller) OWNER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DppAdvanced.contract.Call(opts, &out, "_OWNER_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OWNER is a free data retrieval call binding the contract method 0x16048bc4.
//
// Solidity: function _OWNER_() view returns(address)
func (_DppAdvanced *DppAdvancedSession) OWNER() (common.Address, error) {
	return _DppAdvanced.Contract.OWNER(&_DppAdvanced.CallOpts)
}

// OWNER is a free data retrieval call binding the contract method 0x16048bc4.
//
// Solidity: function _OWNER_() view returns(address)
func (_DppAdvanced *DppAdvancedCallerSession) OWNER() (common.Address, error) {
	return _DppAdvanced.Contract.OWNER(&_DppAdvanced.CallOpts)
}

// QUOTERESERVE is a free data retrieval call binding the contract method 0xbbf5ce78.
//
// Solidity: function _QUOTE_RESERVE_() view returns(uint112)
func (_DppAdvanced *DppAdvancedCaller) QUOTERESERVE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DppAdvanced.contract.Call(opts, &out, "_QUOTE_RESERVE_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QUOTERESERVE is a free data retrieval call binding the contract method 0xbbf5ce78.
//
// Solidity: function _QUOTE_RESERVE_() view returns(uint112)
func (_DppAdvanced *DppAdvancedSession) QUOTERESERVE() (*big.Int, error) {
	return _DppAdvanced.Contract.QUOTERESERVE(&_DppAdvanced.CallOpts)
}

// QUOTERESERVE is a free data retrieval call binding the contract method 0xbbf5ce78.
//
// Solidity: function _QUOTE_RESERVE_() view returns(uint112)
func (_DppAdvanced *DppAdvancedCallerSession) QUOTERESERVE() (*big.Int, error) {
	return _DppAdvanced.Contract.QUOTERESERVE(&_DppAdvanced.CallOpts)
}

// QUOTETARGET is a free data retrieval call binding the contract method 0x77f58657.
//
// Solidity: function _QUOTE_TARGET_() view returns(uint112)
func (_DppAdvanced *DppAdvancedCaller) QUOTETARGET(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DppAdvanced.contract.Call(opts, &out, "_QUOTE_TARGET_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QUOTETARGET is a free data retrieval call binding the contract method 0x77f58657.
//
// Solidity: function _QUOTE_TARGET_() view returns(uint112)
func (_DppAdvanced *DppAdvancedSession) QUOTETARGET() (*big.Int, error) {
	return _DppAdvanced.Contract.QUOTETARGET(&_DppAdvanced.CallOpts)
}

// QUOTETARGET is a free data retrieval call binding the contract method 0x77f58657.
//
// Solidity: function _QUOTE_TARGET_() view returns(uint112)
func (_DppAdvanced *DppAdvancedCallerSession) QUOTETARGET() (*big.Int, error) {
	return _DppAdvanced.Contract.QUOTETARGET(&_DppAdvanced.CallOpts)
}

// QUOTETOKEN is a free data retrieval call binding the contract method 0xd4b97046.
//
// Solidity: function _QUOTE_TOKEN_() view returns(address)
func (_DppAdvanced *DppAdvancedCaller) QUOTETOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DppAdvanced.contract.Call(opts, &out, "_QUOTE_TOKEN_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QUOTETOKEN is a free data retrieval call binding the contract method 0xd4b97046.
//
// Solidity: function _QUOTE_TOKEN_() view returns(address)
func (_DppAdvanced *DppAdvancedSession) QUOTETOKEN() (common.Address, error) {
	return _DppAdvanced.Contract.QUOTETOKEN(&_DppAdvanced.CallOpts)
}

// QUOTETOKEN is a free data retrieval call binding the contract method 0xd4b97046.
//
// Solidity: function _QUOTE_TOKEN_() view returns(address)
func (_DppAdvanced *DppAdvancedCallerSession) QUOTETOKEN() (common.Address, error) {
	return _DppAdvanced.Contract.QUOTETOKEN(&_DppAdvanced.CallOpts)
}

// RState is a free data retrieval call binding the contract method 0xbf357dae.
//
// Solidity: function _RState_() view returns(uint32)
func (_DppAdvanced *DppAdvancedCaller) RState(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DppAdvanced.contract.Call(opts, &out, "_RState_")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RState is a free data retrieval call binding the contract method 0xbf357dae.
//
// Solidity: function _RState_() view returns(uint32)
func (_DppAdvanced *DppAdvancedSession) RState() (uint32, error) {
	return _DppAdvanced.Contract.RState(&_DppAdvanced.CallOpts)
}

// RState is a free data retrieval call binding the contract method 0xbf357dae.
//
// Solidity: function _RState_() view returns(uint32)
func (_DppAdvanced *DppAdvancedCallerSession) RState() (uint32, error) {
	return _DppAdvanced.Contract.RState(&_DppAdvanced.CallOpts)
}

// GetBaseInput is a free data retrieval call binding the contract method 0x65f6fcbb.
//
// Solidity: function getBaseInput() view returns(uint256 input)
func (_DppAdvanced *DppAdvancedCaller) GetBaseInput(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DppAdvanced.contract.Call(opts, &out, "getBaseInput")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBaseInput is a free data retrieval call binding the contract method 0x65f6fcbb.
//
// Solidity: function getBaseInput() view returns(uint256 input)
func (_DppAdvanced *DppAdvancedSession) GetBaseInput() (*big.Int, error) {
	return _DppAdvanced.Contract.GetBaseInput(&_DppAdvanced.CallOpts)
}

// GetBaseInput is a free data retrieval call binding the contract method 0x65f6fcbb.
//
// Solidity: function getBaseInput() view returns(uint256 input)
func (_DppAdvanced *DppAdvancedCallerSession) GetBaseInput() (*big.Int, error) {
	return _DppAdvanced.Contract.GetBaseInput(&_DppAdvanced.CallOpts)
}

// GetMidPrice is a free data retrieval call binding the contract method 0xee27c689.
//
// Solidity: function getMidPrice() view returns(uint256 midPrice)
func (_DppAdvanced *DppAdvancedCaller) GetMidPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DppAdvanced.contract.Call(opts, &out, "getMidPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMidPrice is a free data retrieval call binding the contract method 0xee27c689.
//
// Solidity: function getMidPrice() view returns(uint256 midPrice)
func (_DppAdvanced *DppAdvancedSession) GetMidPrice() (*big.Int, error) {
	return _DppAdvanced.Contract.GetMidPrice(&_DppAdvanced.CallOpts)
}

// GetMidPrice is a free data retrieval call binding the contract method 0xee27c689.
//
// Solidity: function getMidPrice() view returns(uint256 midPrice)
func (_DppAdvanced *DppAdvancedCallerSession) GetMidPrice() (*big.Int, error) {
	return _DppAdvanced.Contract.GetMidPrice(&_DppAdvanced.CallOpts)
}

// GetPMMState is a free data retrieval call binding the contract method 0xa382d1b9.
//
// Solidity: function getPMMState() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint8) state)
func (_DppAdvanced *DppAdvancedCaller) GetPMMState(opts *bind.CallOpts) (PMMPricingPMMState, error) {
	var out []interface{}
	err := _DppAdvanced.contract.Call(opts, &out, "getPMMState")

	if err != nil {
		return *new(PMMPricingPMMState), err
	}

	out0 := *abi.ConvertType(out[0], new(PMMPricingPMMState)).(*PMMPricingPMMState)

	return out0, err

}

// GetPMMState is a free data retrieval call binding the contract method 0xa382d1b9.
//
// Solidity: function getPMMState() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint8) state)
func (_DppAdvanced *DppAdvancedSession) GetPMMState() (PMMPricingPMMState, error) {
	return _DppAdvanced.Contract.GetPMMState(&_DppAdvanced.CallOpts)
}

// GetPMMState is a free data retrieval call binding the contract method 0xa382d1b9.
//
// Solidity: function getPMMState() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint8) state)
func (_DppAdvanced *DppAdvancedCallerSession) GetPMMState() (PMMPricingPMMState, error) {
	return _DppAdvanced.Contract.GetPMMState(&_DppAdvanced.CallOpts)
}

// GetPMMStateForCall is a free data retrieval call binding the contract method 0xfd1ed7e9.
//
// Solidity: function getPMMStateForCall() view returns(uint256 i, uint256 K, uint256 B, uint256 Q, uint256 B0, uint256 Q0, uint256 R)
func (_DppAdvanced *DppAdvancedCaller) GetPMMStateForCall(opts *bind.CallOpts) (struct {
	I  *big.Int
	K  *big.Int
	B  *big.Int
	Q  *big.Int
	B0 *big.Int
	Q0 *big.Int
	R  *big.Int
}, error) {
	var out []interface{}
	err := _DppAdvanced.contract.Call(opts, &out, "getPMMStateForCall")

	outstruct := new(struct {
		I  *big.Int
		K  *big.Int
		B  *big.Int
		Q  *big.Int
		B0 *big.Int
		Q0 *big.Int
		R  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.I = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.K = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.B = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Q = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.B0 = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Q0 = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.R = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPMMStateForCall is a free data retrieval call binding the contract method 0xfd1ed7e9.
//
// Solidity: function getPMMStateForCall() view returns(uint256 i, uint256 K, uint256 B, uint256 Q, uint256 B0, uint256 Q0, uint256 R)
func (_DppAdvanced *DppAdvancedSession) GetPMMStateForCall() (struct {
	I  *big.Int
	K  *big.Int
	B  *big.Int
	Q  *big.Int
	B0 *big.Int
	Q0 *big.Int
	R  *big.Int
}, error) {
	return _DppAdvanced.Contract.GetPMMStateForCall(&_DppAdvanced.CallOpts)
}

// GetPMMStateForCall is a free data retrieval call binding the contract method 0xfd1ed7e9.
//
// Solidity: function getPMMStateForCall() view returns(uint256 i, uint256 K, uint256 B, uint256 Q, uint256 B0, uint256 Q0, uint256 R)
func (_DppAdvanced *DppAdvancedCallerSession) GetPMMStateForCall() (struct {
	I  *big.Int
	K  *big.Int
	B  *big.Int
	Q  *big.Int
	B0 *big.Int
	Q0 *big.Int
	R  *big.Int
}, error) {
	return _DppAdvanced.Contract.GetPMMStateForCall(&_DppAdvanced.CallOpts)
}

// GetQuoteInput is a free data retrieval call binding the contract method 0x71f9100c.
//
// Solidity: function getQuoteInput() view returns(uint256 input)
func (_DppAdvanced *DppAdvancedCaller) GetQuoteInput(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DppAdvanced.contract.Call(opts, &out, "getQuoteInput")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuoteInput is a free data retrieval call binding the contract method 0x71f9100c.
//
// Solidity: function getQuoteInput() view returns(uint256 input)
func (_DppAdvanced *DppAdvancedSession) GetQuoteInput() (*big.Int, error) {
	return _DppAdvanced.Contract.GetQuoteInput(&_DppAdvanced.CallOpts)
}

// GetQuoteInput is a free data retrieval call binding the contract method 0x71f9100c.
//
// Solidity: function getQuoteInput() view returns(uint256 input)
func (_DppAdvanced *DppAdvancedCallerSession) GetQuoteInput() (*big.Int, error) {
	return _DppAdvanced.Contract.GetQuoteInput(&_DppAdvanced.CallOpts)
}

// GetUserFeeRate is a free data retrieval call binding the contract method 0x44096609.
//
// Solidity: function getUserFeeRate(address user) view returns(uint256 lpFeeRate, uint256 mtFeeRate)
func (_DppAdvanced *DppAdvancedCaller) GetUserFeeRate(opts *bind.CallOpts, user common.Address) (struct {
	LpFeeRate *big.Int
	MtFeeRate *big.Int
}, error) {
	var out []interface{}
	err := _DppAdvanced.contract.Call(opts, &out, "getUserFeeRate", user)

	outstruct := new(struct {
		LpFeeRate *big.Int
		MtFeeRate *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LpFeeRate = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MtFeeRate = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetUserFeeRate is a free data retrieval call binding the contract method 0x44096609.
//
// Solidity: function getUserFeeRate(address user) view returns(uint256 lpFeeRate, uint256 mtFeeRate)
func (_DppAdvanced *DppAdvancedSession) GetUserFeeRate(user common.Address) (struct {
	LpFeeRate *big.Int
	MtFeeRate *big.Int
}, error) {
	return _DppAdvanced.Contract.GetUserFeeRate(&_DppAdvanced.CallOpts, user)
}

// GetUserFeeRate is a free data retrieval call binding the contract method 0x44096609.
//
// Solidity: function getUserFeeRate(address user) view returns(uint256 lpFeeRate, uint256 mtFeeRate)
func (_DppAdvanced *DppAdvancedCallerSession) GetUserFeeRate(user common.Address) (struct {
	LpFeeRate *big.Int
	MtFeeRate *big.Int
}, error) {
	return _DppAdvanced.Contract.GetUserFeeRate(&_DppAdvanced.CallOpts, user)
}

// GetVaultReserve is a free data retrieval call binding the contract method 0x36223ce9.
//
// Solidity: function getVaultReserve() view returns(uint256 baseReserve, uint256 quoteReserve)
func (_DppAdvanced *DppAdvancedCaller) GetVaultReserve(opts *bind.CallOpts) (struct {
	BaseReserve  *big.Int
	QuoteReserve *big.Int
}, error) {
	var out []interface{}
	err := _DppAdvanced.contract.Call(opts, &out, "getVaultReserve")

	outstruct := new(struct {
		BaseReserve  *big.Int
		QuoteReserve *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BaseReserve = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.QuoteReserve = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetVaultReserve is a free data retrieval call binding the contract method 0x36223ce9.
//
// Solidity: function getVaultReserve() view returns(uint256 baseReserve, uint256 quoteReserve)
func (_DppAdvanced *DppAdvancedSession) GetVaultReserve() (struct {
	BaseReserve  *big.Int
	QuoteReserve *big.Int
}, error) {
	return _DppAdvanced.Contract.GetVaultReserve(&_DppAdvanced.CallOpts)
}

// GetVaultReserve is a free data retrieval call binding the contract method 0x36223ce9.
//
// Solidity: function getVaultReserve() view returns(uint256 baseReserve, uint256 quoteReserve)
func (_DppAdvanced *DppAdvancedCallerSession) GetVaultReserve() (struct {
	BaseReserve  *big.Int
	QuoteReserve *big.Int
}, error) {
	return _DppAdvanced.Contract.GetVaultReserve(&_DppAdvanced.CallOpts)
}

// QuerySellBase is a free data retrieval call binding the contract method 0x79a04876.
//
// Solidity: function querySellBase(address trader, uint256 payBaseAmount) view returns(uint256 receiveQuoteAmount, uint256 mtFee, uint8 newRState, uint256 newBaseTarget)
func (_DppAdvanced *DppAdvancedCaller) QuerySellBase(opts *bind.CallOpts, trader common.Address, payBaseAmount *big.Int) (struct {
	ReceiveQuoteAmount *big.Int
	MtFee              *big.Int
	NewRState          uint8
	NewBaseTarget      *big.Int
}, error) {
	var out []interface{}
	err := _DppAdvanced.contract.Call(opts, &out, "querySellBase", trader, payBaseAmount)

	outstruct := new(struct {
		ReceiveQuoteAmount *big.Int
		MtFee              *big.Int
		NewRState          uint8
		NewBaseTarget      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReceiveQuoteAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MtFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.NewRState = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.NewBaseTarget = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// QuerySellBase is a free data retrieval call binding the contract method 0x79a04876.
//
// Solidity: function querySellBase(address trader, uint256 payBaseAmount) view returns(uint256 receiveQuoteAmount, uint256 mtFee, uint8 newRState, uint256 newBaseTarget)
func (_DppAdvanced *DppAdvancedSession) QuerySellBase(trader common.Address, payBaseAmount *big.Int) (struct {
	ReceiveQuoteAmount *big.Int
	MtFee              *big.Int
	NewRState          uint8
	NewBaseTarget      *big.Int
}, error) {
	return _DppAdvanced.Contract.QuerySellBase(&_DppAdvanced.CallOpts, trader, payBaseAmount)
}

// QuerySellBase is a free data retrieval call binding the contract method 0x79a04876.
//
// Solidity: function querySellBase(address trader, uint256 payBaseAmount) view returns(uint256 receiveQuoteAmount, uint256 mtFee, uint8 newRState, uint256 newBaseTarget)
func (_DppAdvanced *DppAdvancedCallerSession) QuerySellBase(trader common.Address, payBaseAmount *big.Int) (struct {
	ReceiveQuoteAmount *big.Int
	MtFee              *big.Int
	NewRState          uint8
	NewBaseTarget      *big.Int
}, error) {
	return _DppAdvanced.Contract.QuerySellBase(&_DppAdvanced.CallOpts, trader, payBaseAmount)
}

// QuerySellQuote is a free data retrieval call binding the contract method 0x66410a21.
//
// Solidity: function querySellQuote(address trader, uint256 payQuoteAmount) view returns(uint256 receiveBaseAmount, uint256 mtFee, uint8 newRState, uint256 newQuoteTarget)
func (_DppAdvanced *DppAdvancedCaller) QuerySellQuote(opts *bind.CallOpts, trader common.Address, payQuoteAmount *big.Int) (struct {
	ReceiveBaseAmount *big.Int
	MtFee             *big.Int
	NewRState         uint8
	NewQuoteTarget    *big.Int
}, error) {
	var out []interface{}
	err := _DppAdvanced.contract.Call(opts, &out, "querySellQuote", trader, payQuoteAmount)

	outstruct := new(struct {
		ReceiveBaseAmount *big.Int
		MtFee             *big.Int
		NewRState         uint8
		NewQuoteTarget    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReceiveBaseAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MtFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.NewRState = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.NewQuoteTarget = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// QuerySellQuote is a free data retrieval call binding the contract method 0x66410a21.
//
// Solidity: function querySellQuote(address trader, uint256 payQuoteAmount) view returns(uint256 receiveBaseAmount, uint256 mtFee, uint8 newRState, uint256 newQuoteTarget)
func (_DppAdvanced *DppAdvancedSession) QuerySellQuote(trader common.Address, payQuoteAmount *big.Int) (struct {
	ReceiveBaseAmount *big.Int
	MtFee             *big.Int
	NewRState         uint8
	NewQuoteTarget    *big.Int
}, error) {
	return _DppAdvanced.Contract.QuerySellQuote(&_DppAdvanced.CallOpts, trader, payQuoteAmount)
}

// QuerySellQuote is a free data retrieval call binding the contract method 0x66410a21.
//
// Solidity: function querySellQuote(address trader, uint256 payQuoteAmount) view returns(uint256 receiveBaseAmount, uint256 mtFee, uint8 newRState, uint256 newQuoteTarget)
func (_DppAdvanced *DppAdvancedCallerSession) QuerySellQuote(trader common.Address, payQuoteAmount *big.Int) (struct {
	ReceiveBaseAmount *big.Int
	MtFee             *big.Int
	NewRState         uint8
	NewQuoteTarget    *big.Int
}, error) {
	return _DppAdvanced.Contract.QuerySellQuote(&_DppAdvanced.CallOpts, trader, payQuoteAmount)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_DppAdvanced *DppAdvancedCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DppAdvanced.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_DppAdvanced *DppAdvancedSession) Version() (string, error) {
	return _DppAdvanced.Contract.Version(&_DppAdvanced.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_DppAdvanced *DppAdvancedCallerSession) Version() (string, error) {
	return _DppAdvanced.Contract.Version(&_DppAdvanced.CallOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_DppAdvanced *DppAdvancedTransactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DppAdvanced.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_DppAdvanced *DppAdvancedSession) ClaimOwnership() (*types.Transaction, error) {
	return _DppAdvanced.Contract.ClaimOwnership(&_DppAdvanced.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_DppAdvanced *DppAdvancedTransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _DppAdvanced.Contract.ClaimOwnership(&_DppAdvanced.TransactOpts)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xd0a494e4.
//
// Solidity: function flashLoan(uint256 baseAmount, uint256 quoteAmount, address assetTo, bytes data) returns()
func (_DppAdvanced *DppAdvancedTransactor) FlashLoan(opts *bind.TransactOpts, baseAmount *big.Int, quoteAmount *big.Int, assetTo common.Address, data []byte) (*types.Transaction, error) {
	return _DppAdvanced.contract.Transact(opts, "flashLoan", baseAmount, quoteAmount, assetTo, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xd0a494e4.
//
// Solidity: function flashLoan(uint256 baseAmount, uint256 quoteAmount, address assetTo, bytes data) returns()
func (_DppAdvanced *DppAdvancedSession) FlashLoan(baseAmount *big.Int, quoteAmount *big.Int, assetTo common.Address, data []byte) (*types.Transaction, error) {
	return _DppAdvanced.Contract.FlashLoan(&_DppAdvanced.TransactOpts, baseAmount, quoteAmount, assetTo, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xd0a494e4.
//
// Solidity: function flashLoan(uint256 baseAmount, uint256 quoteAmount, address assetTo, bytes data) returns()
func (_DppAdvanced *DppAdvancedTransactorSession) FlashLoan(baseAmount *big.Int, quoteAmount *big.Int, assetTo common.Address, data []byte) (*types.Transaction, error) {
	return _DppAdvanced.Contract.FlashLoan(&_DppAdvanced.TransactOpts, baseAmount, quoteAmount, assetTo, data)
}

// Init is a paid mutator transaction binding the contract method 0x01a3c30b.
//
// Solidity: function init(address owner, address maintainer, address baseTokenAddress, address quoteTokenAddress, uint256 lpFeeRate, address mtFeeRateModel, uint256 k, uint256 i, bool isOpenTWAP) returns()
func (_DppAdvanced *DppAdvancedTransactor) Init(opts *bind.TransactOpts, owner common.Address, maintainer common.Address, baseTokenAddress common.Address, quoteTokenAddress common.Address, lpFeeRate *big.Int, mtFeeRateModel common.Address, k *big.Int, i *big.Int, isOpenTWAP bool) (*types.Transaction, error) {
	return _DppAdvanced.contract.Transact(opts, "init", owner, maintainer, baseTokenAddress, quoteTokenAddress, lpFeeRate, mtFeeRateModel, k, i, isOpenTWAP)
}

// Init is a paid mutator transaction binding the contract method 0x01a3c30b.
//
// Solidity: function init(address owner, address maintainer, address baseTokenAddress, address quoteTokenAddress, uint256 lpFeeRate, address mtFeeRateModel, uint256 k, uint256 i, bool isOpenTWAP) returns()
func (_DppAdvanced *DppAdvancedSession) Init(owner common.Address, maintainer common.Address, baseTokenAddress common.Address, quoteTokenAddress common.Address, lpFeeRate *big.Int, mtFeeRateModel common.Address, k *big.Int, i *big.Int, isOpenTWAP bool) (*types.Transaction, error) {
	return _DppAdvanced.Contract.Init(&_DppAdvanced.TransactOpts, owner, maintainer, baseTokenAddress, quoteTokenAddress, lpFeeRate, mtFeeRateModel, k, i, isOpenTWAP)
}

// Init is a paid mutator transaction binding the contract method 0x01a3c30b.
//
// Solidity: function init(address owner, address maintainer, address baseTokenAddress, address quoteTokenAddress, uint256 lpFeeRate, address mtFeeRateModel, uint256 k, uint256 i, bool isOpenTWAP) returns()
func (_DppAdvanced *DppAdvancedTransactorSession) Init(owner common.Address, maintainer common.Address, baseTokenAddress common.Address, quoteTokenAddress common.Address, lpFeeRate *big.Int, mtFeeRateModel common.Address, k *big.Int, i *big.Int, isOpenTWAP bool) (*types.Transaction, error) {
	return _DppAdvanced.Contract.Init(&_DppAdvanced.TransactOpts, owner, maintainer, baseTokenAddress, quoteTokenAddress, lpFeeRate, mtFeeRateModel, k, i, isOpenTWAP)
}

// InitOwner is a paid mutator transaction binding the contract method 0x0d009297.
//
// Solidity: function initOwner(address newOwner) returns()
func (_DppAdvanced *DppAdvancedTransactor) InitOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DppAdvanced.contract.Transact(opts, "initOwner", newOwner)
}

// InitOwner is a paid mutator transaction binding the contract method 0x0d009297.
//
// Solidity: function initOwner(address newOwner) returns()
func (_DppAdvanced *DppAdvancedSession) InitOwner(newOwner common.Address) (*types.Transaction, error) {
	return _DppAdvanced.Contract.InitOwner(&_DppAdvanced.TransactOpts, newOwner)
}

// InitOwner is a paid mutator transaction binding the contract method 0x0d009297.
//
// Solidity: function initOwner(address newOwner) returns()
func (_DppAdvanced *DppAdvancedTransactorSession) InitOwner(newOwner common.Address) (*types.Transaction, error) {
	return _DppAdvanced.Contract.InitOwner(&_DppAdvanced.TransactOpts, newOwner)
}

// RatioSync is a paid mutator transaction binding the contract method 0xc57a5d03.
//
// Solidity: function ratioSync() returns()
func (_DppAdvanced *DppAdvancedTransactor) RatioSync(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DppAdvanced.contract.Transact(opts, "ratioSync")
}

// RatioSync is a paid mutator transaction binding the contract method 0xc57a5d03.
//
// Solidity: function ratioSync() returns()
func (_DppAdvanced *DppAdvancedSession) RatioSync() (*types.Transaction, error) {
	return _DppAdvanced.Contract.RatioSync(&_DppAdvanced.TransactOpts)
}

// RatioSync is a paid mutator transaction binding the contract method 0xc57a5d03.
//
// Solidity: function ratioSync() returns()
func (_DppAdvanced *DppAdvancedTransactorSession) RatioSync() (*types.Transaction, error) {
	return _DppAdvanced.Contract.RatioSync(&_DppAdvanced.TransactOpts)
}

// Reset is a paid mutator transaction binding the contract method 0x8ff3928c.
//
// Solidity: function reset(address assetTo, uint256 newLpFeeRate, uint256 newI, uint256 newK, uint256 baseOutAmount, uint256 quoteOutAmount, uint256 minBaseReserve, uint256 minQuoteReserve) returns(bool)
func (_DppAdvanced *DppAdvancedTransactor) Reset(opts *bind.TransactOpts, assetTo common.Address, newLpFeeRate *big.Int, newI *big.Int, newK *big.Int, baseOutAmount *big.Int, quoteOutAmount *big.Int, minBaseReserve *big.Int, minQuoteReserve *big.Int) (*types.Transaction, error) {
	return _DppAdvanced.contract.Transact(opts, "reset", assetTo, newLpFeeRate, newI, newK, baseOutAmount, quoteOutAmount, minBaseReserve, minQuoteReserve)
}

// Reset is a paid mutator transaction binding the contract method 0x8ff3928c.
//
// Solidity: function reset(address assetTo, uint256 newLpFeeRate, uint256 newI, uint256 newK, uint256 baseOutAmount, uint256 quoteOutAmount, uint256 minBaseReserve, uint256 minQuoteReserve) returns(bool)
func (_DppAdvanced *DppAdvancedSession) Reset(assetTo common.Address, newLpFeeRate *big.Int, newI *big.Int, newK *big.Int, baseOutAmount *big.Int, quoteOutAmount *big.Int, minBaseReserve *big.Int, minQuoteReserve *big.Int) (*types.Transaction, error) {
	return _DppAdvanced.Contract.Reset(&_DppAdvanced.TransactOpts, assetTo, newLpFeeRate, newI, newK, baseOutAmount, quoteOutAmount, minBaseReserve, minQuoteReserve)
}

// Reset is a paid mutator transaction binding the contract method 0x8ff3928c.
//
// Solidity: function reset(address assetTo, uint256 newLpFeeRate, uint256 newI, uint256 newK, uint256 baseOutAmount, uint256 quoteOutAmount, uint256 minBaseReserve, uint256 minQuoteReserve) returns(bool)
func (_DppAdvanced *DppAdvancedTransactorSession) Reset(assetTo common.Address, newLpFeeRate *big.Int, newI *big.Int, newK *big.Int, baseOutAmount *big.Int, quoteOutAmount *big.Int, minBaseReserve *big.Int, minQuoteReserve *big.Int) (*types.Transaction, error) {
	return _DppAdvanced.Contract.Reset(&_DppAdvanced.TransactOpts, assetTo, newLpFeeRate, newI, newK, baseOutAmount, quoteOutAmount, minBaseReserve, minQuoteReserve)
}

// Retrieve is a paid mutator transaction binding the contract method 0x28c4e24c.
//
// Solidity: function retrieve(address to, address token, uint256 amount) returns()
func (_DppAdvanced *DppAdvancedTransactor) Retrieve(opts *bind.TransactOpts, to common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DppAdvanced.contract.Transact(opts, "retrieve", to, token, amount)
}

// Retrieve is a paid mutator transaction binding the contract method 0x28c4e24c.
//
// Solidity: function retrieve(address to, address token, uint256 amount) returns()
func (_DppAdvanced *DppAdvancedSession) Retrieve(to common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DppAdvanced.Contract.Retrieve(&_DppAdvanced.TransactOpts, to, token, amount)
}

// Retrieve is a paid mutator transaction binding the contract method 0x28c4e24c.
//
// Solidity: function retrieve(address to, address token, uint256 amount) returns()
func (_DppAdvanced *DppAdvancedTransactorSession) Retrieve(to common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DppAdvanced.Contract.Retrieve(&_DppAdvanced.TransactOpts, to, token, amount)
}

// SellBase is a paid mutator transaction binding the contract method 0xbd6015b4.
//
// Solidity: function sellBase(address to) returns(uint256 receiveQuoteAmount)
func (_DppAdvanced *DppAdvancedTransactor) SellBase(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _DppAdvanced.contract.Transact(opts, "sellBase", to)
}

// SellBase is a paid mutator transaction binding the contract method 0xbd6015b4.
//
// Solidity: function sellBase(address to) returns(uint256 receiveQuoteAmount)
func (_DppAdvanced *DppAdvancedSession) SellBase(to common.Address) (*types.Transaction, error) {
	return _DppAdvanced.Contract.SellBase(&_DppAdvanced.TransactOpts, to)
}

// SellBase is a paid mutator transaction binding the contract method 0xbd6015b4.
//
// Solidity: function sellBase(address to) returns(uint256 receiveQuoteAmount)
func (_DppAdvanced *DppAdvancedTransactorSession) SellBase(to common.Address) (*types.Transaction, error) {
	return _DppAdvanced.Contract.SellBase(&_DppAdvanced.TransactOpts, to)
}

// SellQuote is a paid mutator transaction binding the contract method 0xdd93f59a.
//
// Solidity: function sellQuote(address to) returns(uint256 receiveBaseAmount)
func (_DppAdvanced *DppAdvancedTransactor) SellQuote(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _DppAdvanced.contract.Transact(opts, "sellQuote", to)
}

// SellQuote is a paid mutator transaction binding the contract method 0xdd93f59a.
//
// Solidity: function sellQuote(address to) returns(uint256 receiveBaseAmount)
func (_DppAdvanced *DppAdvancedSession) SellQuote(to common.Address) (*types.Transaction, error) {
	return _DppAdvanced.Contract.SellQuote(&_DppAdvanced.TransactOpts, to)
}

// SellQuote is a paid mutator transaction binding the contract method 0xdd93f59a.
//
// Solidity: function sellQuote(address to) returns(uint256 receiveBaseAmount)
func (_DppAdvanced *DppAdvancedTransactorSession) SellQuote(to common.Address) (*types.Transaction, error) {
	return _DppAdvanced.Contract.SellQuote(&_DppAdvanced.TransactOpts, to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DppAdvanced *DppAdvancedTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DppAdvanced.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DppAdvanced *DppAdvancedSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DppAdvanced.Contract.TransferOwnership(&_DppAdvanced.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DppAdvanced *DppAdvancedTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DppAdvanced.Contract.TransferOwnership(&_DppAdvanced.TransactOpts, newOwner)
}

// TuneParameters is a paid mutator transaction binding the contract method 0x3b20884a.
//
// Solidity: function tuneParameters(uint256 newLpFeeRate, uint256 newI, uint256 newK, uint256 minBaseReserve, uint256 minQuoteReserve) returns(bool)
func (_DppAdvanced *DppAdvancedTransactor) TuneParameters(opts *bind.TransactOpts, newLpFeeRate *big.Int, newI *big.Int, newK *big.Int, minBaseReserve *big.Int, minQuoteReserve *big.Int) (*types.Transaction, error) {
	return _DppAdvanced.contract.Transact(opts, "tuneParameters", newLpFeeRate, newI, newK, minBaseReserve, minQuoteReserve)
}

// TuneParameters is a paid mutator transaction binding the contract method 0x3b20884a.
//
// Solidity: function tuneParameters(uint256 newLpFeeRate, uint256 newI, uint256 newK, uint256 minBaseReserve, uint256 minQuoteReserve) returns(bool)
func (_DppAdvanced *DppAdvancedSession) TuneParameters(newLpFeeRate *big.Int, newI *big.Int, newK *big.Int, minBaseReserve *big.Int, minQuoteReserve *big.Int) (*types.Transaction, error) {
	return _DppAdvanced.Contract.TuneParameters(&_DppAdvanced.TransactOpts, newLpFeeRate, newI, newK, minBaseReserve, minQuoteReserve)
}

// TuneParameters is a paid mutator transaction binding the contract method 0x3b20884a.
//
// Solidity: function tuneParameters(uint256 newLpFeeRate, uint256 newI, uint256 newK, uint256 minBaseReserve, uint256 minQuoteReserve) returns(bool)
func (_DppAdvanced *DppAdvancedTransactorSession) TuneParameters(newLpFeeRate *big.Int, newI *big.Int, newK *big.Int, minBaseReserve *big.Int, minQuoteReserve *big.Int) (*types.Transaction, error) {
	return _DppAdvanced.Contract.TuneParameters(&_DppAdvanced.TransactOpts, newLpFeeRate, newI, newK, minBaseReserve, minQuoteReserve)
}

// TunePrice is a paid mutator transaction binding the contract method 0x10d76460.
//
// Solidity: function tunePrice(uint256 newI, uint256 minBaseReserve, uint256 minQuoteReserve) returns(bool)
func (_DppAdvanced *DppAdvancedTransactor) TunePrice(opts *bind.TransactOpts, newI *big.Int, minBaseReserve *big.Int, minQuoteReserve *big.Int) (*types.Transaction, error) {
	return _DppAdvanced.contract.Transact(opts, "tunePrice", newI, minBaseReserve, minQuoteReserve)
}

// TunePrice is a paid mutator transaction binding the contract method 0x10d76460.
//
// Solidity: function tunePrice(uint256 newI, uint256 minBaseReserve, uint256 minQuoteReserve) returns(bool)
func (_DppAdvanced *DppAdvancedSession) TunePrice(newI *big.Int, minBaseReserve *big.Int, minQuoteReserve *big.Int) (*types.Transaction, error) {
	return _DppAdvanced.Contract.TunePrice(&_DppAdvanced.TransactOpts, newI, minBaseReserve, minQuoteReserve)
}

// TunePrice is a paid mutator transaction binding the contract method 0x10d76460.
//
// Solidity: function tunePrice(uint256 newI, uint256 minBaseReserve, uint256 minQuoteReserve) returns(bool)
func (_DppAdvanced *DppAdvancedTransactorSession) TunePrice(newI *big.Int, minBaseReserve *big.Int, minQuoteReserve *big.Int) (*types.Transaction, error) {
	return _DppAdvanced.Contract.TunePrice(&_DppAdvanced.TransactOpts, newI, minBaseReserve, minQuoteReserve)
}

// DppAdvancedDODOFlashLoanIterator is returned from FilterDODOFlashLoan and is used to iterate over the raw logs and unpacked data for DODOFlashLoan events raised by the DppAdvanced contract.
type DppAdvancedDODOFlashLoanIterator struct {
	Event *DppAdvancedDODOFlashLoan // Event containing the contract specifics and raw log

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
func (it *DppAdvancedDODOFlashLoanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DppAdvancedDODOFlashLoan)
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
		it.Event = new(DppAdvancedDODOFlashLoan)
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
func (it *DppAdvancedDODOFlashLoanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DppAdvancedDODOFlashLoanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DppAdvancedDODOFlashLoan represents a DODOFlashLoan event raised by the DppAdvanced contract.
type DppAdvancedDODOFlashLoan struct {
	Borrower    common.Address
	AssetTo     common.Address
	BaseAmount  *big.Int
	QuoteAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDODOFlashLoan is a free log retrieval operation binding the contract event 0x0b82e93068db15abd9fbb2682c65462ea8a0a10582dce93a5664818e296f54eb.
//
// Solidity: event DODOFlashLoan(address borrower, address assetTo, uint256 baseAmount, uint256 quoteAmount)
func (_DppAdvanced *DppAdvancedFilterer) FilterDODOFlashLoan(opts *bind.FilterOpts) (*DppAdvancedDODOFlashLoanIterator, error) {

	logs, sub, err := _DppAdvanced.contract.FilterLogs(opts, "DODOFlashLoan")
	if err != nil {
		return nil, err
	}
	return &DppAdvancedDODOFlashLoanIterator{contract: _DppAdvanced.contract, event: "DODOFlashLoan", logs: logs, sub: sub}, nil
}

// WatchDODOFlashLoan is a free log subscription operation binding the contract event 0x0b82e93068db15abd9fbb2682c65462ea8a0a10582dce93a5664818e296f54eb.
//
// Solidity: event DODOFlashLoan(address borrower, address assetTo, uint256 baseAmount, uint256 quoteAmount)
func (_DppAdvanced *DppAdvancedFilterer) WatchDODOFlashLoan(opts *bind.WatchOpts, sink chan<- *DppAdvancedDODOFlashLoan) (event.Subscription, error) {

	logs, sub, err := _DppAdvanced.contract.WatchLogs(opts, "DODOFlashLoan")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DppAdvancedDODOFlashLoan)
				if err := _DppAdvanced.contract.UnpackLog(event, "DODOFlashLoan", log); err != nil {
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

// ParseDODOFlashLoan is a log parse operation binding the contract event 0x0b82e93068db15abd9fbb2682c65462ea8a0a10582dce93a5664818e296f54eb.
//
// Solidity: event DODOFlashLoan(address borrower, address assetTo, uint256 baseAmount, uint256 quoteAmount)
func (_DppAdvanced *DppAdvancedFilterer) ParseDODOFlashLoan(log types.Log) (*DppAdvancedDODOFlashLoan, error) {
	event := new(DppAdvancedDODOFlashLoan)
	if err := _DppAdvanced.contract.UnpackLog(event, "DODOFlashLoan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DppAdvancedDODOSwapIterator is returned from FilterDODOSwap and is used to iterate over the raw logs and unpacked data for DODOSwap events raised by the DppAdvanced contract.
type DppAdvancedDODOSwapIterator struct {
	Event *DppAdvancedDODOSwap // Event containing the contract specifics and raw log

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
func (it *DppAdvancedDODOSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DppAdvancedDODOSwap)
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
		it.Event = new(DppAdvancedDODOSwap)
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
func (it *DppAdvancedDODOSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DppAdvancedDODOSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DppAdvancedDODOSwap represents a DODOSwap event raised by the DppAdvanced contract.
type DppAdvancedDODOSwap struct {
	FromToken  common.Address
	ToToken    common.Address
	FromAmount *big.Int
	ToAmount   *big.Int
	Trader     common.Address
	Receiver   common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDODOSwap is a free log retrieval operation binding the contract event 0xc2c0245e056d5fb095f04cd6373bc770802ebd1e6c918eb78fdef843cdb37b0f.
//
// Solidity: event DODOSwap(address fromToken, address toToken, uint256 fromAmount, uint256 toAmount, address trader, address receiver)
func (_DppAdvanced *DppAdvancedFilterer) FilterDODOSwap(opts *bind.FilterOpts) (*DppAdvancedDODOSwapIterator, error) {

	logs, sub, err := _DppAdvanced.contract.FilterLogs(opts, "DODOSwap")
	if err != nil {
		return nil, err
	}
	return &DppAdvancedDODOSwapIterator{contract: _DppAdvanced.contract, event: "DODOSwap", logs: logs, sub: sub}, nil
}

// WatchDODOSwap is a free log subscription operation binding the contract event 0xc2c0245e056d5fb095f04cd6373bc770802ebd1e6c918eb78fdef843cdb37b0f.
//
// Solidity: event DODOSwap(address fromToken, address toToken, uint256 fromAmount, uint256 toAmount, address trader, address receiver)
func (_DppAdvanced *DppAdvancedFilterer) WatchDODOSwap(opts *bind.WatchOpts, sink chan<- *DppAdvancedDODOSwap) (event.Subscription, error) {

	logs, sub, err := _DppAdvanced.contract.WatchLogs(opts, "DODOSwap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DppAdvancedDODOSwap)
				if err := _DppAdvanced.contract.UnpackLog(event, "DODOSwap", log); err != nil {
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

// ParseDODOSwap is a log parse operation binding the contract event 0xc2c0245e056d5fb095f04cd6373bc770802ebd1e6c918eb78fdef843cdb37b0f.
//
// Solidity: event DODOSwap(address fromToken, address toToken, uint256 fromAmount, uint256 toAmount, address trader, address receiver)
func (_DppAdvanced *DppAdvancedFilterer) ParseDODOSwap(log types.Log) (*DppAdvancedDODOSwap, error) {
	event := new(DppAdvancedDODOSwap)
	if err := _DppAdvanced.contract.UnpackLog(event, "DODOSwap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DppAdvancedLpFeeRateChangeIterator is returned from FilterLpFeeRateChange and is used to iterate over the raw logs and unpacked data for LpFeeRateChange events raised by the DppAdvanced contract.
type DppAdvancedLpFeeRateChangeIterator struct {
	Event *DppAdvancedLpFeeRateChange // Event containing the contract specifics and raw log

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
func (it *DppAdvancedLpFeeRateChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DppAdvancedLpFeeRateChange)
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
		it.Event = new(DppAdvancedLpFeeRateChange)
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
func (it *DppAdvancedLpFeeRateChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DppAdvancedLpFeeRateChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DppAdvancedLpFeeRateChange represents a LpFeeRateChange event raised by the DppAdvanced contract.
type DppAdvancedLpFeeRateChange struct {
	NewLpFeeRate *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLpFeeRateChange is a free log retrieval operation binding the contract event 0x9950d5a2f2c7264863d40100bf993f0cdbc4711806caba6284d07e80fd500879.
//
// Solidity: event LpFeeRateChange(uint256 newLpFeeRate)
func (_DppAdvanced *DppAdvancedFilterer) FilterLpFeeRateChange(opts *bind.FilterOpts) (*DppAdvancedLpFeeRateChangeIterator, error) {

	logs, sub, err := _DppAdvanced.contract.FilterLogs(opts, "LpFeeRateChange")
	if err != nil {
		return nil, err
	}
	return &DppAdvancedLpFeeRateChangeIterator{contract: _DppAdvanced.contract, event: "LpFeeRateChange", logs: logs, sub: sub}, nil
}

// WatchLpFeeRateChange is a free log subscription operation binding the contract event 0x9950d5a2f2c7264863d40100bf993f0cdbc4711806caba6284d07e80fd500879.
//
// Solidity: event LpFeeRateChange(uint256 newLpFeeRate)
func (_DppAdvanced *DppAdvancedFilterer) WatchLpFeeRateChange(opts *bind.WatchOpts, sink chan<- *DppAdvancedLpFeeRateChange) (event.Subscription, error) {

	logs, sub, err := _DppAdvanced.contract.WatchLogs(opts, "LpFeeRateChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DppAdvancedLpFeeRateChange)
				if err := _DppAdvanced.contract.UnpackLog(event, "LpFeeRateChange", log); err != nil {
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

// ParseLpFeeRateChange is a log parse operation binding the contract event 0x9950d5a2f2c7264863d40100bf993f0cdbc4711806caba6284d07e80fd500879.
//
// Solidity: event LpFeeRateChange(uint256 newLpFeeRate)
func (_DppAdvanced *DppAdvancedFilterer) ParseLpFeeRateChange(log types.Log) (*DppAdvancedLpFeeRateChange, error) {
	event := new(DppAdvancedLpFeeRateChange)
	if err := _DppAdvanced.contract.UnpackLog(event, "LpFeeRateChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DppAdvancedOwnershipTransferPreparedIterator is returned from FilterOwnershipTransferPrepared and is used to iterate over the raw logs and unpacked data for OwnershipTransferPrepared events raised by the DppAdvanced contract.
type DppAdvancedOwnershipTransferPreparedIterator struct {
	Event *DppAdvancedOwnershipTransferPrepared // Event containing the contract specifics and raw log

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
func (it *DppAdvancedOwnershipTransferPreparedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DppAdvancedOwnershipTransferPrepared)
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
		it.Event = new(DppAdvancedOwnershipTransferPrepared)
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
func (it *DppAdvancedOwnershipTransferPreparedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DppAdvancedOwnershipTransferPreparedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DppAdvancedOwnershipTransferPrepared represents a OwnershipTransferPrepared event raised by the DppAdvanced contract.
type DppAdvancedOwnershipTransferPrepared struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferPrepared is a free log retrieval operation binding the contract event 0xdcf55418cee3220104fef63f979ff3c4097ad240c0c43dcb33ce837748983e62.
//
// Solidity: event OwnershipTransferPrepared(address indexed previousOwner, address indexed newOwner)
func (_DppAdvanced *DppAdvancedFilterer) FilterOwnershipTransferPrepared(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DppAdvancedOwnershipTransferPreparedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DppAdvanced.contract.FilterLogs(opts, "OwnershipTransferPrepared", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DppAdvancedOwnershipTransferPreparedIterator{contract: _DppAdvanced.contract, event: "OwnershipTransferPrepared", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferPrepared is a free log subscription operation binding the contract event 0xdcf55418cee3220104fef63f979ff3c4097ad240c0c43dcb33ce837748983e62.
//
// Solidity: event OwnershipTransferPrepared(address indexed previousOwner, address indexed newOwner)
func (_DppAdvanced *DppAdvancedFilterer) WatchOwnershipTransferPrepared(opts *bind.WatchOpts, sink chan<- *DppAdvancedOwnershipTransferPrepared, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DppAdvanced.contract.WatchLogs(opts, "OwnershipTransferPrepared", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DppAdvancedOwnershipTransferPrepared)
				if err := _DppAdvanced.contract.UnpackLog(event, "OwnershipTransferPrepared", log); err != nil {
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

// ParseOwnershipTransferPrepared is a log parse operation binding the contract event 0xdcf55418cee3220104fef63f979ff3c4097ad240c0c43dcb33ce837748983e62.
//
// Solidity: event OwnershipTransferPrepared(address indexed previousOwner, address indexed newOwner)
func (_DppAdvanced *DppAdvancedFilterer) ParseOwnershipTransferPrepared(log types.Log) (*DppAdvancedOwnershipTransferPrepared, error) {
	event := new(DppAdvancedOwnershipTransferPrepared)
	if err := _DppAdvanced.contract.UnpackLog(event, "OwnershipTransferPrepared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DppAdvancedOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DppAdvanced contract.
type DppAdvancedOwnershipTransferredIterator struct {
	Event *DppAdvancedOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DppAdvancedOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DppAdvancedOwnershipTransferred)
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
		it.Event = new(DppAdvancedOwnershipTransferred)
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
func (it *DppAdvancedOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DppAdvancedOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DppAdvancedOwnershipTransferred represents a OwnershipTransferred event raised by the DppAdvanced contract.
type DppAdvancedOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DppAdvanced *DppAdvancedFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DppAdvancedOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DppAdvanced.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DppAdvancedOwnershipTransferredIterator{contract: _DppAdvanced.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DppAdvanced *DppAdvancedFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DppAdvancedOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DppAdvanced.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DppAdvancedOwnershipTransferred)
				if err := _DppAdvanced.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DppAdvanced *DppAdvancedFilterer) ParseOwnershipTransferred(log types.Log) (*DppAdvancedOwnershipTransferred, error) {
	event := new(DppAdvancedOwnershipTransferred)
	if err := _DppAdvanced.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DppAdvancedRChangeIterator is returned from FilterRChange and is used to iterate over the raw logs and unpacked data for RChange events raised by the DppAdvanced contract.
type DppAdvancedRChangeIterator struct {
	Event *DppAdvancedRChange // Event containing the contract specifics and raw log

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
func (it *DppAdvancedRChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DppAdvancedRChange)
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
		it.Event = new(DppAdvancedRChange)
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
func (it *DppAdvancedRChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DppAdvancedRChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DppAdvancedRChange represents a RChange event raised by the DppAdvanced contract.
type DppAdvancedRChange struct {
	NewRState uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRChange is a free log retrieval operation binding the contract event 0xdf176ad18be4f9f32efaa32f06e9d1175476504739a745f1399a6d3fa4b75917.
//
// Solidity: event RChange(uint8 newRState)
func (_DppAdvanced *DppAdvancedFilterer) FilterRChange(opts *bind.FilterOpts) (*DppAdvancedRChangeIterator, error) {

	logs, sub, err := _DppAdvanced.contract.FilterLogs(opts, "RChange")
	if err != nil {
		return nil, err
	}
	return &DppAdvancedRChangeIterator{contract: _DppAdvanced.contract, event: "RChange", logs: logs, sub: sub}, nil
}

// WatchRChange is a free log subscription operation binding the contract event 0xdf176ad18be4f9f32efaa32f06e9d1175476504739a745f1399a6d3fa4b75917.
//
// Solidity: event RChange(uint8 newRState)
func (_DppAdvanced *DppAdvancedFilterer) WatchRChange(opts *bind.WatchOpts, sink chan<- *DppAdvancedRChange) (event.Subscription, error) {

	logs, sub, err := _DppAdvanced.contract.WatchLogs(opts, "RChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DppAdvancedRChange)
				if err := _DppAdvanced.contract.UnpackLog(event, "RChange", log); err != nil {
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

// ParseRChange is a log parse operation binding the contract event 0xdf176ad18be4f9f32efaa32f06e9d1175476504739a745f1399a6d3fa4b75917.
//
// Solidity: event RChange(uint8 newRState)
func (_DppAdvanced *DppAdvancedFilterer) ParseRChange(log types.Log) (*DppAdvancedRChange, error) {
	event := new(DppAdvancedRChange)
	if err := _DppAdvanced.contract.UnpackLog(event, "RChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
