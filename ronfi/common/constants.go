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

	RedisMsgNewLoop      = "newLoop"
	RedisMsgCancelLoop   = "cancelLoop"
	RedisMsgRestoreLoop  = "restoreLoop"
	RedisMsgNewObsRouter = "newObsRouter"
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
	BTCB = common.HexToAddress("0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c")
	CHI  = common.HexToAddress("0x0000000000004946c0e9F43F4Dee607b0eF1fA1c")

	TradableTokens = map[common.Address]string{
		USDT: "USDT",
		USDC: "USDC",
		WETH: "WETH",
		BTCB: "BTCB",
		DAI:  "DAI",
	}

	BridgeTokens = map[common.Address]struct{}{
		DAI:  {},
		USDC: {},
		USDT: {},
		WETH: {},
		BTCB: {},
	}

	DexMethodsTypical = map[uint64]string{
		0x38ed1739: "swapExactTokensForTokens", // amountIn fixed
		0x5c11d795: "swapExactTokensForTokensSupportingFeeOnTransferTokens",
		0x7ff36ab5: "swapExactETHForTokens", // amountIn fixed,	ETH -> TOK
		0xb6f9de95: "swapExactETHForTokensSupportingFeeOnTransferTokens",
		0x9cf68911: "swapExactBNBForTokens", // amountIn fixed
		0x18cbafe5: "swapExactTokensForETH", // amountIn fixed,	TOK -> ETH
		0x791ac947: "swapExactTokensForETHSupportingFeeOnTransferTokens",
		0x5d616c5b: "swapExactTokensForBNB", // amountIn fixed,	TOK -> BNB
		0xd46d2f83: "swapExactTokensForBNBSupportingFeeOnTransferTokens",
		0x8803dbee: "swapTokensForExactTokens", // amountOut fixed
		0x4a25d94a: "swapTokensForExactETH",    // amountOut fixed,	TOK -> ETH
		0xfb3bdb41: "swapETHForExactTokens",    // amountOut fixed,	ETH -> TOK
		0x8332a963: "swapBNBForExactTokens",    // amountOut fixed
		0xd67b571e: "swapTokensForExactBNB",    // amountOut fixed,   TOK -> BNB
		0xe8e33700: "addLiquidity",             //(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline)
		0x57528cca: "addLiquidity",             //(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, uint256 startingSwapTime)
		0xf305d719: "addLiquidityETH",
		0x6bb6a6f6: "addLiquidityETH", //(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, uint256 startingSwapTime)
		0xeaaed442: "addLiquidityBNB", //(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountBNBMin, address to, uint256 deadline)
		0xbaa2abde: "removeLiquidity",
		0x2195995c: "removeLiquidityWithPermit",
		0x02751cec: "removeLiquidityETH",
		0xaf2979eb: "removeLiquidityETHSupportingFeeOnTransferTokens",
		0xded9382a: "removeLiquidityETHWithPermit",
		0x5b0d5984: "removeLiquidityETHWithPermitSupportingFeeOnTransferTokens",
		0xe0588488: "removeLiquidityBNB",
		0x34a0772d: "removeLiquidityBNBWithPermit",
		0x685a0a34: "removeLiquidityBNBWithPermitSupportingFeeOnTransferTokens", //(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountBNBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s)
		0x5f575529: "metamask swap",                                             // Function: swap(bytes data)  MetaMask also use it
		0x7c025200: "1inch",
		0x2e95b6c8: "1inch",
		0xcba9bc66: "Agg1Router",
		0x415565b0: "ZeroEx",
		0xc43c9ef6: "ZeroEx",
		0x441a3e70: "Withdraw", //PlanetFinance
		0xe2bbb158: "Deposit",  //PlanetFinance
		0xd72ef771: "work",     // AdminUpgradeabilityProxy
		0xd29a0025: "Kill",     // AdminUpgradeabilityProxy
	}
)
