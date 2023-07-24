package v2

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
	"github.com/ethereum/go-ethereum/ronfi/defi"
	"github.com/ethereum/go-ethereum/ronfi/mysql"
	common2 "github.com/ethereum/go-ethereum/ronfi/uniswap/common"
	"testing"
)

func TestV2_CalculateTokensOutFromTokensIn(t *testing.T) {
	client, err := ethclient.Dial("https://nd-804-879-862.p2pify.com/ca0df7232f6a54347593373cfbf94df8")
	if err != nil {
		t.Fatal("TestV2_CalculateTokensOutFromTokensIn dial eth client failed!")
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
		t.Fatalf("TestV2_CalculateTokensOutFromTokensIn NewMysql failed!")
	}

	di := defi.NewInfo(client, mysqlInst)
	if di == nil {
		t.Fatalf("TestV2_CalculateTokensOutFromTokensIn NewInfo failed!")
	}

	pairAddr := common.HexToAddress("0xEFb47fcFCaD4F96c83D4ca676842fB03Ef20a477")
	v2Pool := NewV2Pool(di, pairAddr, 30, nil)
	v2Pool.UpdateReserves()
	amountOut := v2Pool.CalculateTokensOutFromTokensIn(
		common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
		common2.StrToBigInt("1000000000000000000"),
	)

	t.Logf("amountOut: %v", amountOut)
}

func TestV2_CalculateTokensInFromTokensOut(t *testing.T) {
	client, err := ethclient.Dial("https://nd-814-711-835.p2pify.com/049f42c7290c310495b6940701e2ae14")
	if err != nil {
		t.Fatal("TestV2_CalculateTokensInFromTokensOut dial eth client failed!")
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
		t.Fatalf("TestV2_CalculateTokensInFromTokensOut NewMysql failed!")
	}

	di := defi.NewInfo(client, mysqlInst)
	if di == nil {
		t.Fatalf("TestV2_CalculateTokensInFromTokensOut NewInfo failed!")
	}

	pairAddr := common.HexToAddress("0x16b9a82891338f9bA80E2D6970FddA79D1eb0daE")
	v2Pool := NewV2Pool(di, pairAddr, 30, nil)
	v2Pool.UpdateReserves()
	amountIn := v2Pool.CalculateTokensInFromTokensOut(
		common.HexToAddress("0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c"),
		common2.StrToBigInt("1000000000000000000"),
	)

	t.Logf("amountIn: %v", amountIn)
}

func TestV2_MaxTokensInFromTokensOut(t *testing.T) {
	client, err := ethclient.Dial("https://nd-814-711-835.p2pify.com/049f42c7290c310495b6940701e2ae14")
	if err != nil {
		t.Fatal("TestV2_MaxTokensInFromTokensOut dial eth client failed!")
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
		t.Fatalf("TestV2_MaxTokensInFromTokensOut NewMysql failed!")
	}

	di := defi.NewInfo(client, mysqlInst)
	if di == nil {
		t.Fatalf("TestV2_MaxTokensInFromTokensOut NewInfo failed!")
	}

	pairAddr := common.HexToAddress("0x16b9a82891338f9bA80E2D6970FddA79D1eb0daE")
	v2Pool := NewV2Pool(di, pairAddr, 30, nil)
	v2Pool.UpdateReserves()
	amountIn := v2Pool.MaxTokensInFromTokensOut(
		common.HexToAddress("0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c"),
	)

	t.Logf("amountIn: %v", amountIn)
}
