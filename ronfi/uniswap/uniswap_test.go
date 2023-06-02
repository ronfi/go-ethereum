package uniswap

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
	"github.com/ethereum/go-ethereum/ronfi/defi"
	"github.com/ethereum/go-ethereum/ronfi/loops"
	"github.com/ethereum/go-ethereum/ronfi/mysql"
	"math/big"
	"testing"
)

func TestUniswap_Minimizer(t *testing.T) {
	testPolynomial := func(inst interface{}, x *big.Int) *big.Int {
		sec1 := new(big.Int).Mul(x, x)
		sec2 := new(big.Int).Mul(x, big.NewInt(20000))
		sec3 := big.NewInt(50000)
		res := new(big.Int).Add(new(big.Int).Sub(sec1, sec2), sec3)
		//fmt.Printf("x: %v, res: %v\n", x, res)
		return res
	}

	a := big.NewInt(-1e6)
	b := big.NewInt(1e6)
	eps := big.NewInt(1)
	options := MinOptions{
		MaxIters: 10000,
		EPS:      eps,
	}
	if min, err := minimizeScalar(nil, testPolynomial, []*big.Int{a, b}, &options); err == nil {
		t.Logf("x: %v, fun: %v\n", min.x, min.fun)
	} else {
		t.Fatalf("error: %v\n", err)
	}
}

func TestUniswap_CycleSwap(t *testing.T) {
	//client, err := ethclient.Dial("https://nd-804-879-862.p2pify.com/ca0df7232f6a54347593373cfbf94df8")
	client, err := ethclient.Dial("https://nd-814-711-835.p2pify.com/049f42c7290c310495b6940701e2ae14")
	if err != nil {
		t.Fatal("TestV3_Swap dial eth client failed!")
	}

	defer func() {
		if client != nil {
			client.Close()
		}
	}()

	if block, err := client.BlockNumber(context.Background()); err == nil {
		t.Logf("block: %v\n", block)
	}

	conf := rcommon.MysqlConfig{
		DbHost: "176.9.120.196",
		DbPort: "3306",
		DbUser: "root",
		DbPass: "rkdb",
		DbData: "rkdb",
	}
	mysqlInst := mysql.NewMysql(conf)
	if mysqlInst == nil {
		t.Fatalf("TestV3_Swap NewMysql failed!")
	}

	di := defi.NewInfo(client, mysqlInst)
	if di == nil {
		t.Fatalf("TestV3_Swap NewInfo failed!")
	}

	arb := FromAddress(
		di,
		nil,
		nil,
		common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"),
		[]*UniswapPool{
			{
				PoolAddr: common.HexToAddress("0x133B3D95bAD5405d14d53473671200e9342896BF"),
				PoolType: UniswapV3,
				PoolFee:  0,
				TickLens: common.HexToAddress("0x9a489505a00cE272eAa5e07Dba6491314CaE3796"),
			},
			{
				PoolAddr: common.HexToAddress("0xd710DaC10B7255fCc0eC929a0e742A18F25b6395"),
				PoolType: UniswapV3,
				PoolFee:  0,
				TickLens: common.HexToAddress("0x9a489505a00cE272eAa5e07Dba6491314CaE3796"),
			},
		},
	)

	arb.AutoUpdate(nil)
	res := arb.CalculateArbitrage()
	fmt.Println("profitable: ", res.Profitable, "SwapAmount: ", res.SwapAmount, "BestProfit: ", res.BestProfit)
}

func TestUniswap_V3Loops(t *testing.T) {
	client, err := ethclient.Dial("https://nd-804-879-862.p2pify.com/ca0df7232f6a54347593373cfbf94df8")
	if err != nil {
		t.Fatal("TestV3_Swap dial eth client failed!")
	}

	defer func() {
		if client != nil {
			client.Close()
		}
	}()

	if block, err := client.BlockNumber(context.Background()); err == nil {
		t.Logf("block: %v\n", block)
	}

	conf := rcommon.MysqlConfig{
		DbHost: "176.9.120.196",
		DbPort: "3306",
		DbUser: "root",
		DbPass: "rkdb",
		DbData: "rkdb",
	}
	mysqlInst := mysql.NewMysql(conf)
	if mysqlInst == nil {
		t.Fatalf("TestV3_Swap NewMysql failed!")
	}

	di := defi.NewInfo(client, mysqlInst)
	if di == nil {
		t.Fatalf("TestV3_Swap NewInfo failed!")
	}

	oskLimitedPairsMap := mysqlInst.LoadOskLimitedPairs()
	pairGasMap := mysqlInst.LoadPairGas()
	feePatchMap := mysqlInst.LoadFeePatch()
	flashNokPairs := make(map[common.Address]uint64)
	loopsIdMap := make(loops.LIdMap)
	loopsMap := loops.LoadSwapLoops(
		mysqlInst,
		di,
		loopsIdMap,
		feePatchMap,
		pairGasMap,
		oskLimitedPairsMap,
		flashNokPairs,
		false)

	poolsInfo := di.GetAllPoolInfo()
	v3Loops := NewV3Loops(di, loopsMap.UniPairsMap, poolsInfo, loopsMap.AllTokensMap, pairGasMap)
	arbs := v3Loops.FindLoops(&Edge{
		Source: common.HexToAddress("0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82"),
		Target: common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"),
		Tag: &TaggedEdge{
			Pair:     common.HexToAddress("0xAfB2Da14056725E3BA3a30dD846B6BBbd7886c56"),
			Dir:      0,
			PoolType: UniswapV3,
		},
	})

	t.Logf("Total cylces: %v\n", len(arbs))
	for _, arb := range arbs {
		t.Logf("arb: %v\n", arb.String())
	}
}

func TestUniswap_LocalFork(t *testing.T) {
	//client, err := ethclient.Dial("https://nd-814-711-835.p2pify.com/049f42c7290c310495b6940701e2ae14")
	client, err := ethclient.Dial("http://127.0.0.1:9545")
	if err != nil {
		t.Fatal("TestUniswap_LocalFork dial eth client failed!")
	}

	defer func() {
		if client != nil {
			client.Close()
		}
	}()

	if block, err := client.BlockNumber(context.Background()); err == nil {
		fmt.Printf("block: %v\n", block)
	}

	conf := rcommon.MysqlConfig{
		DbHost: "176.9.120.196",
		DbPort: "3306",
		DbUser: "root",
		DbPass: "rkdb",
		DbData: "rkdb",
	}
	mysqlInst := mysql.NewMysql(conf)
	if mysqlInst == nil {
		t.Fatalf("TestV3_Swap NewMysql failed!")
	}

	di := defi.NewInfo(client, mysqlInst)
	if di == nil {
		t.Fatalf("TestV3_Swap NewInfo failed!")
	}

	arb := FromAddress(
		di,
		nil,
		nil,
		common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"),
		[]*UniswapPool{
			{
				PoolAddr: common.HexToAddress("0x133B3D95bAD5405d14d53473671200e9342896BF"),
				PoolType: UniswapV3,
				PoolFee:  0,
				TickLens: common.HexToAddress("0x9a489505a00cE272eAa5e07Dba6491314CaE3796"),
			},
			{
				PoolAddr: common.HexToAddress("0xd710DaC10B7255fCc0eC929a0e742A18F25b6395"),
				PoolType: UniswapV3,
				PoolFee:  0,
				TickLens: common.HexToAddress("0x9a489505a00cE272eAa5e07Dba6491314CaE3796"),
			},
		},
	)

	arb.AutoUpdate(nil)
	res := arb.CalculateArbitrage()
	fmt.Printf("profitable: %v, SwapAmount: %v, BestProfit: %v\n", res.Profitable, res.SwapAmount, res.BestProfit)
}
