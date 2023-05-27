package common

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

const (
	CLIENT_URL = "/home/bsc/.geth/geth.ipc"
	CHAIN_ID   = 56

	HopsMaxAllowed = 8

	MaxMatchedLoopsAllowed = 2

	ErrRonFiNo2     = "#2"
	ErrRonFiNo4     = "#4"
	ErrPANCAKEK     = "K"
	ErrNomiswapD    = "D"
	ErrTRANSFERFAIL = "TRANSFER_FAILED"
	ErrFORBIDDEN    = "FORBIDDEN"
	ErrOVERFLOW     = "OVERFLOW"

	LoopIdTopN = 6 // only keep top 6 of used loopId

	RedisMsgNewLoop         = "newLoop"
	RedisMsgNewKnowledge    = "newKnowledge"
	RedisMsgCancelLoop      = "cancelLoop"
	RedisMsgRestoreLoop     = "restoreLoop"
	RedisMsgCancelKnowledge = "cancelKnowledge"
	RedisMsgNewObsRouter    = "newObsRouter"
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
	//bscJsonRPCs = []string{
	//	"https://bsc-dataseed.binance.org/",
	//	"https://bsc-dataseed1.defibit.io/",core/vm/interpreter.go
	//	"https://bsc-dataseed1.ninicoin.io/",
	//}
	//
	//directRouteEndPoint = "https://api.nodereal.io/direct-route"

	ZeroAddress                 = common.HexToAddress("0x0000000000000000000000000000000000000000")
	OneAddress                  = common.HexToAddress("0x0000000000000000000000000000000000000001")
	RON_SWAP_F_ADDRESS          = common.HexToAddress("0x445b51E968bA40BcbebA6e74654ba1117b629925")
	RON_SWAP_X_ADDRESS          = common.HexToAddress("0x918F5Be29759552B2BeC5F9a1eb049c3060EB108")
	RON_SWAP_Z_ADDRESS          = common.HexToAddress("0xFA20Dfae69D86A0Bd64A391557bD2Fb60bEf8F1D")
	RON_V3_SWAP_ADDRESS         = common.HexToAddress("0x0C70A9dBC7e0704d344aa818FAE5cA6f0f73D534")
	PANCAKE_V3_TICKLENS_ADDRESS = common.HexToAddress("0x9a489505a00cE272eAa5e07Dba6491314CaE3796")
	PANCAKE_V3_FACOTRY_ADDRESS  = common.HexToAddress("0x0bfbcf9fa4f9c56b0f40a671ad40e0805a091865")

	TokenOSK = common.HexToAddress("0x04fA9Eb295266d9d4650EDCB879da204887Dc3Da") // A Special Care

	ZeroHash       = common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000")
	DefaultPoolFee = uint64(30)

	GasRough = [9]uint64{
		0, 100000, 150000, 210000, 280000, 350000, 420000, 500000, 600000,
	}

	WBNB = common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c")
	WETH = common.HexToAddress("0x2170Ed0880ac9A755fd29B2688956BD959F933F8")
	DAI  = common.HexToAddress("0x1AF3F329e8BE154074D8769D1FFa4eE058B1DBc3")
	USDC = common.HexToAddress("0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d")
	USDT = common.HexToAddress("0x55d398326f99059fF775485246999027B3197955")
	BUSD = common.HexToAddress("0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56")
	BTCB = common.HexToAddress("0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c")
	CAKE = common.HexToAddress("0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82")
	CHI  = common.HexToAddress("0x0000000000004946c0e9F43F4Dee607b0eF1fA1c")

	WrappedBNB = map[common.Address]struct{}{
		common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"): {},
		common.HexToAddress("0x0efb5FD2402A0967B92551d6AF54De148504A115"): {},
		common.HexToAddress("0x14233ba7106d1a3c2db6f9d6f5b216956a201072"): {},
		common.HexToAddress("0xa2ca18fc541b7b101c64e64bbc2834b05066248b"): {},
		common.HexToAddress("0x7f95b241e27fc6ac89390aa12ef40706f38254f6"): {},
		common.HexToAddress("0xfc425ea901671f4d6369b39c0fee4abed7ff9870"): {},
		common.HexToAddress("0xa5cce43a73d09ea3a19c9e1c2ccd51cf55db3699"): {},
	}

	TradableTokens = map[common.Address]string{
		WBNB: "WBNB",
		USDT: "USDT",
		USDC: "USDC",
		BUSD: "BUSD",
		WETH: "WETH",
		BTCB: "BTCB",
		CAKE: "CAKE",
		DAI:  "DAI",
	}
	BridgeTokens = map[common.Address]struct{}{
		WBNB: {},
		DAI:  {},
		USDC: {},
		USDT: {},
		BUSD: {},
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
