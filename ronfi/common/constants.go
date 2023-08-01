package common

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

const (
	CLIENT_URL = "/home/bsc/.geth/geth.ipc"
	CHAIN_ID   = 1

	HopsMaxAllowed         = 8
	MaxMatchedLoopsAllowed = 2

	RedisMsgNewLoop         = "newLoop"
	RedisMsgNewKnowledge    = "newKnowledge"
	RedisMsgCancelLoop      = "cancelLoop"
	RedisMsgRestoreLoop     = "restoreLoop"
	RedisMsgCancelKnowledge = "cancelKnowledge"
	RedisMsgNewObsRouter    = "newObsRouter"
	RedisMsgNewDexRouter    = "newDexRouter"
	RedisMsgDelDexRouter    = "delDexRouter"
)

// internal
var (
	bigInt36      = new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e18))
	bigInt54      = new(big.Int).Mul(bigInt36, big.NewInt(1e18))
	decimalsToBig = []*big.Int{
		big.NewInt(1),
		big.NewInt(10),
		big.NewInt(100),
		big.NewInt(1000),
		big.NewInt(1e4),
		big.NewInt(1e5),
		big.NewInt(1e6),
		big.NewInt(1e7),
		big.NewInt(1e8),
		big.NewInt(1e9),
		big.NewInt(1e10),
		big.NewInt(1e11),
		big.NewInt(1e12),
		big.NewInt(1e13),
		big.NewInt(1e14),
		big.NewInt(1e15),
		big.NewInt(1e16),
		big.NewInt(1e17),
		big.NewInt(1e18),
		new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e1)), // 1e19
		new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e2)),
		new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e3)),
		new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e4)),
		new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e5)),
		new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e6)),
		new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e7)),
		new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e8)),
		new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e9)),
		new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e10)),
		new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e11)),
		new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e12)),
		new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e13)), // 1e31
		new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e14)),
		new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e15)),
		new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e16)),
		new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e17)),
		new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e18)),
		new(big.Int).Mul(bigInt36, big.NewInt(1e1)),
		new(big.Int).Mul(bigInt36, big.NewInt(1e2)),
		new(big.Int).Mul(bigInt36, big.NewInt(1e3)),
		new(big.Int).Mul(bigInt36, big.NewInt(1e4)),
		new(big.Int).Mul(bigInt36, big.NewInt(1e5)),
		new(big.Int).Mul(bigInt36, big.NewInt(1e6)),
		new(big.Int).Mul(bigInt36, big.NewInt(1e7)),
		new(big.Int).Mul(bigInt36, big.NewInt(1e8)),
		new(big.Int).Mul(bigInt36, big.NewInt(1e9)),
		new(big.Int).Mul(bigInt36, big.NewInt(1e11)),
		new(big.Int).Mul(bigInt36, big.NewInt(1e12)),
		new(big.Int).Mul(bigInt36, big.NewInt(1e13)),
		new(big.Int).Mul(bigInt36, big.NewInt(1e14)),
		new(big.Int).Mul(bigInt36, big.NewInt(1e15)),
		new(big.Int).Mul(bigInt36, big.NewInt(1e16)),
		new(big.Int).Mul(bigInt36, big.NewInt(1e17)),
		new(big.Int).Mul(bigInt36, big.NewInt(1e18)),
		new(big.Int).Mul(bigInt54, big.NewInt(1e1)),
		new(big.Int).Mul(bigInt54, big.NewInt(1e2)),
		new(big.Int).Mul(bigInt54, big.NewInt(1e3)),
		new(big.Int).Mul(bigInt54, big.NewInt(1e4)),
		new(big.Int).Mul(bigInt54, big.NewInt(1e5)),
		new(big.Int).Mul(bigInt54, big.NewInt(1e6)),
		new(big.Int).Mul(bigInt54, big.NewInt(1e7)),
		new(big.Int).Mul(bigInt54, big.NewInt(1e8)),
		new(big.Int).Mul(bigInt54, big.NewInt(1e9)),
		new(big.Int).Mul(bigInt54, big.NewInt(1e10)),
	}
)

// public
var (
	GasRough = [9]uint64{
		0, 100000, 150000, 210000, 280000, 350000, 420000, 500000, 600000,
	}

	ZeroAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")

	UniswapV3TicklensAddress = common.HexToAddress("0xbfd8137f7d1516D3ea5cA83523914859ec47F573")
	UniswapV3FacotryAddress  = common.HexToAddress("0x1F98431c8aD98523631AE4a59f267346ea31F984")

	WETH = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
	DAI  = common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F")
	USDC = common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
	USDT = common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7")
	BUSD = common.HexToAddress("0x4Fabb145d64652a948d72533023f6E7A623C7C53")
	USDP = common.HexToAddress("0x8E870D67F660D95d5be530380D0eC0bd388289E1")
	BTCB = common.HexToAddress("0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c")
	CHI  = common.HexToAddress("0x0000000000004946c0e9F43F4Dee607b0eF1fA1c")

	TradableTokens = map[common.Address]string{
		//USDT: "USDT",
		//USDC: "USDC",
		WETH: "WETH",
		//BTCB: "BTCB",
		//DAI:  "DAI",
	}

	OBSTradableTokens = map[common.Address]string{
		WETH: "WETH",
		USDT: "USDT",
		USDC: "USDC",
		USDP: "USDP",
	}

	BridgeTokens = map[common.Address]struct{}{
		DAI:  {},
		USDC: {},
		USDT: {},
		BUSD: {},
		WETH: {},
		BTCB: {},
	}
)
