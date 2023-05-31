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

	tx, _, err := client.TransactionByHash(context.Background(), common.HexToHash("0x6c6856cdaabf7a6779f4649d74cf7b33681b21dc312ce60410a1d2748450dd66"))
	if err != nil {
		t.Fatal("TestInfo_GetArbTxProfit TransactionByHash failed!", "err", err)
	}

	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash("0x6c6856cdaabf7a6779f4649d74cf7b33681b21dc312ce60410a1d2748450dd66"))
	if err != nil {
		t.Fatal("TestInfo_GetArbTxProfit TransactionReceipt failed!", "err", err)
	}

	profit, slush, v3Hunting := info.GetArbTxProfit(tx, receipt.Logs, *tx.To())
	if profit == 0.0 {
		t.Fatal("TestInfo_GetArbTxProfit calculate profit failed!")
	}
	t.Logf("TestInfo_GetArbTxProfit, v3: %v, profit: %v, slush: %v", v3Hunting, profit, slush)
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
