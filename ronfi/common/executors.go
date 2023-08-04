package common

import "github.com/ethereum/go-ethereum/common"

// AllTradingExecutors 5*160+1*800 = total 1600 wallets, for 6 hunting servers
var AllTradingExecutors = []common.Address{ // Total 560 Executors, Split into 5 Groups which has 80 for each, Plus 4 Groups which has 40 for each
}

var AllV3TradingExecutors = []common.Address{
	common.HexToAddress("0xEE86fA28Fc3C206BCbb48fD8C3c75e9326609779"),
}
