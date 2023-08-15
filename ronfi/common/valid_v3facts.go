package common

import "github.com/ethereum/go-ethereum/common"

var (
	// ValidV3FactsMap is a list of valid v3 pools
	ValidV3FactsMap = map[common.Address]string{
		UniswapV3FacotryAddress:     "uniswap v3",
		SushiSwapV3FactoryAddress:   "sushiswap v3",
		PancakeSwapV3FactoryAddress: "pancake v3",
	}
)
