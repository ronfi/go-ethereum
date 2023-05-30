package defi

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
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

	tx, _, err := client.TransactionByHash(context.Background(), common.HexToHash("0xc1f0619383bfce756ffb9498d7078ae031d4a9f74e1e7d453d6e142c24ccee2d"))
	if err != nil {
		t.Fatal("TestInfo_GetArbTxProfit TransactionByHash failed!", "err", err)
	}

	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash("0xc1f0619383bfce756ffb9498d7078ae031d4a9f74e1e7d453d6e142c24ccee2d"))
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

	tx, _, err := client.TransactionByHash(context.Background(), common.HexToHash("0xdf0d24e771e8931f9914fd08c116f61d4c8a8e040e9b1ad6f0a15a897c395f04"))
	if err != nil {
		t.Fatal("TestInfo_GetArbTxProfit TransactionByHash failed!", "err", err)
	}

	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash("0xdf0d24e771e8931f9914fd08c116f61d4c8a8e040e9b1ad6f0a15a897c395f04"))
	if err != nil {
		t.Fatal("TestInfo_GetArbTxProfit TransactionReceipt failed!", "err", err)
	}

	isObs := false
	obsAddr := *tx.To()
	swapPairsInfo := info.ExtractSwapPairInfo(nil, nil, tx, *tx.To(), receipt.Logs, RonFiExtractTypeStats)
	if len(swapPairsInfo) > 1 {
		head := swapPairsInfo[0]
		tail := swapPairsInfo[len(swapPairsInfo)-1]
		if head.TokenIn != rcommon.ZeroAddress && tail.TokenOut != rcommon.ZeroAddress &&
			head.TokenIn == tail.TokenOut && tail.AmountOut != nil &&
			head.AmountIn != nil && tail.AmountOut.Cmp(head.AmountIn) >= 0 {
			if head.Address != tail.Address || len(swapPairsInfo) > 2 {
				isObs = true
			}
		}
		if !isObs {
			return
		}

		// loops linkage check
		for i := 1; i < len(swapPairsInfo); i++ {
			tail = swapPairsInfo[i]
			if head.TokenOut == rcommon.ZeroAddress ||
				tail.TokenIn == rcommon.ZeroAddress ||
				head.TokenOut != tail.TokenIn ||
				head.To == rcommon.ZeroAddress ||
				(head.To != tail.Address && head.To != obsAddr) {
				isObs = false
				break
			}
			head = tail
		}
	}

	t.Logf("isObs: %v", isObs)
}
