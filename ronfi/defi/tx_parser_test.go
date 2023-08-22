package defi

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
	"github.com/ethereum/go-ethereum/ronfi/db"
	"math/big"
	"testing"
)

func TestInfo_GetArbTxProfit(t *testing.T) {
	client, err := ethclient.Dial("https://nd-804-879-862.p2pify.com/ca0df7232f6a54347593373cfbf94df8")
	if err != nil {
		t.Fatal("TestInfo_GetArbTxProfit dial eth client failed!")
	}

	InitRonFiOracle(client)

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
	dbInst := db.NewMysql(conf)
	if dbInst == nil {
		t.Fatalf("TestInfo_CheckIfObs NewMysql failed!")
	}
	signer := types.NewEIP155Signer(big.NewInt(1))
	info := NewInfo(client, dbInst, signer)

	tx, _, err := client.TransactionByHash(context.Background(), common.HexToHash("0x0a288cca5cb3e9349722a940d39b3b1ba3ee1902d3b77ffc30353614013e9a06"))
	if err != nil {
		t.Fatal("TestInfo_GetArbTxProfit TransactionByHash failed!", "err", err)
	}

	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash("0x0a288cca5cb3e9349722a940d39b3b1ba3ee1902d3b77ffc30353614013e9a06"))
	if err != nil {
		t.Fatal("TestInfo_GetArbTxProfit TransactionReceipt failed!", "err", err)
	}

	profit, v3Hunting, isArbTx := info.GetArbTxProfit(tx, receipt.Logs, *tx.To())
	if profit == 0.0 {
		t.Fatalf("TestInfo_GetArbTxProfit calculate profit failed, v3: %v, profit: %v, isArbTx: %v", v3Hunting, profit, isArbTx)
	}
	t.Logf("TestInfo_GetArbTxProfit, v3: %v, profit: %v, isArbTx: %v", v3Hunting, profit, isArbTx)
}

func TestInfo_CheckIfObs(t *testing.T) {
	client, err := ethclient.Dial("https://nd-804-879-862.p2pify.com/ca0df7232f6a54347593373cfbf94df8")
	if err != nil {
		t.Fatal("TestInfo_CheckIfObs dial eth client failed!")
	}

	InitRonFiOracle(client)

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
	dbInst := db.NewMysql(conf)
	if dbInst == nil {
		t.Fatalf("TestInfo_CheckIfObs NewMysql failed!")
	}

	config := &params.ChainConfig{
		ChainID: big.NewInt(1),
	}
	signer := types.MakeSigner(config, big.NewInt(17034870), 1681266455)
	info := NewInfo(client, dbInst, signer)

	tx, _, err := client.TransactionByHash(context.Background(), common.HexToHash("0x7058aee966509fd467f8c35e390235d9180cca8dce5f92cc3f4f8180c9519161"))
	if err != nil {
		t.Fatal("TestInfo_CheckIfObs TransactionByHash failed!", "err", err)
	}

	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash("0x7058aee966509fd467f8c35e390235d9180cca8dce5f92cc3f4f8180c9519161"))
	if err != nil {
		t.Fatal("TestInfo_CheckIfObs TransactionReceipt failed!", "err", err)
	}

	isDex, isObs := info.CheckIfObsTx(tx, receipt.Logs, *tx.To())
	t.Logf("isDex: %v, isObs: %v", isDex, isObs)
}
