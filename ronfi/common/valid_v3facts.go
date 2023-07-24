package common

import "github.com/ethereum/go-ethereum/common"

var (
	// ValidV3FactsMap is a list of valid v3 pools
	ValidV3FactsMap = map[common.Address]string{
		common.HexToAddress("0x1F98431c8aD98523631AE4a59f267346ea31F984"): "uniswap v3",
	}
)
