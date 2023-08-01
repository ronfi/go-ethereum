package defi

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
	"github.com/ethereum/go-ethereum/ronfi/db"
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
	info := NewInfo(client, dbInst)

	tx, _, err := client.TransactionByHash(context.Background(), common.HexToHash("0x906b030e9a18e702ce211d00b343ac1e0999eb2e4e5dbb5c84e6f0256646cb41"))
	if err != nil {
		t.Fatal("TestInfo_GetArbTxProfit TransactionByHash failed!", "err", err)
	}

	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash("0x906b030e9a18e702ce211d00b343ac1e0999eb2e4e5dbb5c84e6f0256646cb41"))
	if err != nil {
		t.Fatal("TestInfo_GetArbTxProfit TransactionReceipt failed!", "err", err)
	}

	profit, v3Hunting := info.GetArbTxProfit(tx, receipt.Logs, *tx.To())
	if profit == 0.0 {
		t.Fatal("TestInfo_GetArbTxProfit calculate profit failed!")
	}
	t.Logf("TestInfo_GetArbTxProfit, v3: %v, profit: %v", v3Hunting, profit)
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
	info := NewInfo(client, dbInst)

	tx, _, err := client.TransactionByHash(context.Background(), common.HexToHash("0x5d098158f7e6cbedff76ae96923f7e1e9456a23f9012f3a5ba208b3fba90bcf2"))
	if err != nil {
		t.Fatal("TestInfo_CheckIfObs TransactionByHash failed!", "err", err)
	}

	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash("0x5d098158f7e6cbedff76ae96923f7e1e9456a23f9012f3a5ba208b3fba90bcf2"))
	if err != nil {
		t.Fatal("TestInfo_CheckIfObs TransactionReceipt failed!", "err", err)
	}

	isDex, isObs := info.CheckIfObsTx(tx, receipt.Logs, *tx.To())
	t.Logf("isDex: %v, isObs: %v", isDex, isObs)
}
