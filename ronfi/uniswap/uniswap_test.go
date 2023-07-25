package uniswap

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
	"github.com/ethereum/go-ethereum/ronfi/db"
	"github.com/ethereum/go-ethereum/ronfi/defi"
	common2 "github.com/ethereum/go-ethereum/ronfi/uniswap/common"
	v2 "github.com/ethereum/go-ethereum/ronfi/uniswap/v2"
	v3 "github.com/ethereum/go-ethereum/ronfi/uniswap/v3"
	"math/big"
	"testing"
)

func TestUniswap_Minimizer(t *testing.T) {
	testPolynomial := func(inst interface{}, x *big.Int) (*big.Int, *big.Int, []*big.Int, []*big.Int) {
		sec1 := new(big.Int).Mul(x, x)
		sec2 := new(big.Int).Mul(x, big.NewInt(20000))
		sec3 := big.NewInt(50000)
		res := new(big.Int).Add(new(big.Int).Sub(sec1, sec2), sec3)
		//fmt.Printf("x: %v, res: %v\n", x, res)
		return nil, res, nil, nil
	}

	a := big.NewInt(-1e6)
	b := big.NewInt(1e6)
	eps := big.NewInt(1)
	options := MinOptions{
		MaxIters: 10000,
		EPS:      eps,
	}
	if min, iters, err := minimizeScalar(nil, testPolynomial, []*big.Int{a, b}, &options); err == nil {
		t.Logf("iters: %v, x: %v, fun: %v\n", iters, min.x, min.fun)
	} else {
		t.Fatalf("error: %v\n", err)
	}
}

func TestUniswap_CycleSwap(t *testing.T) {
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
		DbData: "rkdb_eth",
	}
	mysqlInst := db.NewMysql(conf)
	if mysqlInst == nil {
		t.Fatalf("TestV3_Swap NewMysql failed!")
	}

	di := defi.NewInfo(client, mysqlInst)
	if di == nil {
		t.Fatalf("TestV3_Swap NewInfo failed!")
	}

	v2AmountIOs := make(map[common.Address]map[string]*big.Int)
	v3AmountIOs := make(map[common.Address]map[string]*v3.DetailOut)
	v2Pools := make(map[common.Address]*v2.Pool)
	v3Pools := make(map[common.Address]*v3.Pool)
	arb := FromAddress(
		di,
		nil,
		nil,
		v2AmountIOs,
		v3AmountIOs,
		v2Pools,
		v3Pools,
		common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"),
		nil,
		[]*Pool{
			{
				PoolAddr: common.HexToAddress("0x133B3D95bAD5405d14d53473671200e9342896BF"),
				PoolType: V3,
				PoolFee:  0,
				TickLens: common.HexToAddress("0x9a489505a00cE272eAa5e07Dba6491314CaE3796"),
			},
			{
				PoolAddr: common.HexToAddress("0xd710DaC10B7255fCc0eC929a0e742A18F25b6395"),
				PoolType: V3,
				PoolFee:  0,
				TickLens: common.HexToAddress("0x9a489505a00cE272eAa5e07Dba6491314CaE3796"),
			},
		},
	)

	v3States := make(map[common.Address]*v3.PoolState)
	arb.AutoUpdate(v3States)
	res := arb.CalculateArbitrage()
	if res != nil {
		fmt.Println("profitable: ", res.Profitable, "SwapAmount: ", res.SwapAmount, "BestProfit: ", res.BestProfit)
	} else {
		t.Fatalf("TestV3_Swap CalculateArbitrage failed!")
	}
}

func TestUniswap_CalculateMaxInAmount(t *testing.T) {
	client, err := ethclient.Dial("https://nd-804-879-862.p2pify.com/ca0df7232f6a54347593373cfbf94df8")
	if err != nil {
		t.Fatal("TestUniswap_CalculateMaxInAmount dial eth client failed!")
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
		DbData: "rkdb_eth",
	}
	mysqlInst := db.NewMysql(conf)
	if mysqlInst == nil {
		t.Fatalf("TestUniswap_CalculateMaxInAmount NewMysql failed!")
	}

	di := defi.NewInfo(client, mysqlInst)
	if di == nil {
		t.Fatalf("TestUniswap_CalculateMaxInAmount NewInfo failed!")
	}

	v2AmountIOs := make(map[common.Address]map[string]*big.Int)
	v3AmountIOs := make(map[common.Address]map[string]*v3.DetailOut)
	v2Pools := make(map[common.Address]*v2.Pool)
	v3Pools := make(map[common.Address]*v3.Pool)
	arb := FromAddress(
		di,
		nil,
		nil,
		v2AmountIOs,
		v3AmountIOs,
		v2Pools,
		v3Pools,
		common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"),
		nil,
		[]*Pool{
			{
				PoolAddr: common.HexToAddress("0x133B3D95bAD5405d14d53473671200e9342896BF"),
				PoolType: V3,
				PoolFee:  0,
				TickLens: common.HexToAddress("0x9a489505a00cE272eAa5e07Dba6491314CaE3796"),
			},
			{
				PoolAddr: common.HexToAddress("0xd710DaC10B7255fCc0eC929a0e742A18F25b6395"),
				PoolType: V3,
				PoolFee:  0,
				TickLens: common.HexToAddress("0x9a489505a00cE272eAa5e07Dba6491314CaE3796"),
			},
		},
	)

	v3States := make(map[common.Address]*v3.PoolState)
	arb.AutoUpdate(v3States)
	maxIn := arb.CalculateMaxInAmount(common2.StrToBigInt("1000000000000000000"), 0)
	t.Logf("maxIn: %v\n", maxIn)
}

func TestUniswap_NewV3Loops(t *testing.T) {
	client, err := ethclient.Dial("https://nd-804-879-862.p2pify.com/ca0df7232f6a54347593373cfbf94df8")
	if err != nil {
		t.Fatal("TestUniswap_NewV3Loops dial eth client failed!")
	}

	defer func() {
		if client != nil {
			client.Close()
		}
	}()

	conf := rcommon.MysqlConfig{
		DbHost: "176.9.120.196",
		DbPort: "3306",
		DbUser: "root",
		DbPass: "rkdb",
		DbData: "rkdb_eth",
	}
	mysqlInst := db.NewMysql(conf)
	if mysqlInst == nil {
		t.Fatalf("TestUniswap_NewV3Loops NewMysql failed!")
	}

	di := defi.NewInfo(client, mysqlInst)
	if di == nil {
		t.Fatalf("TestUniswap_NewV3Loops NewInfo failed!")
	}

	pairsInfo := make(defi.PairInfoMap)
	for address, info := range di.GetAllPairInfo() {
		if _, ok := rcommon.ValidV2FactsMap[info.Factory]; ok {
			pairsInfo[address] = info
		}
	}
	t.Logf("pairsInfo: %v\n", len(pairsInfo))

	poolsInfo := make(map[common.Address]*defi.PoolInfo)
	for address, info := range di.GetAllPoolInfo() {
		if _, ok := rcommon.ValidV3FactsMap[info.Factory]; ok {
			poolsInfo[address] = info
		}
	}
	t.Logf("poolsInfo: %v\n", len(poolsInfo))

	tokensInfo := di.GetAllTokenInfo()
	t.Logf("tokensInfo: %v\n", len(tokensInfo))

	pairGasMap := mysqlInst.LoadPairGas()
	t.Logf("pairGasMap: %v\n", len(pairGasMap))

	v3Loops := NewV3Loops(di, pairsInfo, poolsInfo, tokensInfo, pairGasMap)
	if v3Loops == nil {
		t.Fatalf("TestUniswap_NewV3Loops NewV3Loops failed!")
	}

	taggedEdge := &TaggedEdge{
		Pair:     common.HexToAddress("0xD9F3a5A0e7499149538cDEE8B804a5Ac4f4944Fc"),
		Dir:      1,
		PoolType: V2,
		GasNeed:  500000,
	}
	edge := &Edge{
		Source: common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
		Target: common.HexToAddress("0x8185Bc4757572Da2a610f887561c32298f1A5748"),
		Tag:    taggedEdge,
	}

	arbs := v3Loops.FindLoops(edge)
	if arbs == nil {
		t.Fatalf("TestUniswap_NewV3Loops FindLoops failed!")
	} else {
		t.Logf("total %v arbs found!\n", len(arbs))
		for _, arb := range arbs {
			t.Logf("arb: %v\n", arb.String())
		}
	}
}
