package common

import "github.com/ethereum/go-ethereum/common"

var (
	// ValidV3FactsMap is a list of valid v3 pools
	ValidV3FactsMap = map[common.Address]string{
		common.HexToAddress("0x1F98431c8aD98523631AE4a59f267346ea31F984"): "uniswap v3",
		common.HexToAddress("0xbACEB8eC6b9355Dfc0269C18bac9d6E2Bdc29C4F"): "sushiswap v3",
		common.HexToAddress("0x0BFbCF9fa4f9C56B0F40a671Ad40E0805A091865"): "pancake v3",
	}
)
