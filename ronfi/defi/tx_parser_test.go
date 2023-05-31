package defi

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
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

	info := NewInfo(client)

	tx, _, err := client.TransactionByHash(context.Background(), common.HexToHash("0xe769a0cee33813b6126d0e76b5fc8b57a25d6e73bedf36b3035346b49120a450"))
	if err != nil {
		t.Fatal("TestInfo_GetArbTxProfit TransactionByHash failed!", "err", err)
	}

	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash("0xe769a0cee33813b6126d0e76b5fc8b57a25d6e73bedf36b3035346b49120a450"))
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
		t.Fatal("TestInfo_GetArbTxProfit dial eth client failed!")
	}

	InitRonFiOracle(client)

	defer func() {
		if client != nil {
			client.Close()
		}
	}()

	info := NewInfo(client)

	tx, _, err := client.TransactionByHash(context.Background(), common.HexToHash("0xa7e4481a58f1390b9256dd52af7e22cecf13ac4124212960fb2e4735d45a6d89"))
	if err != nil {
		t.Fatal("TestInfo_GetArbTxProfit TransactionByHash failed!", "err", err)
	}

	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash("0xa7e4481a58f1390b9256dd52af7e22cecf13ac4124212960fb2e4735d45a6d89"))
	if err != nil {
		t.Fatal("TestInfo_GetArbTxProfit TransactionReceipt failed!", "err", err)
	}

	isDex, isObs := info.CheckIfObsTx(nil, tx, receipt.Logs, *tx.To())
	t.Logf("isDex: %v, isObs: %v", isDex, isObs)
}
