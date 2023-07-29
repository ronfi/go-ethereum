package trading

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"testing"
)

func TestFlashbot(t *testing.T) {
	//client, err := ethclient.Dial("https://nd-804-879-862.p2pify.com/ca0df7232f6a54347593373cfbf94df8")
	client, err := ethclient.Dial("https://goerli.infura.io/v3/6b589162f066453fa7769e842067a1a4")
	if err != nil {
		t.Fatal("TestInfo_GetPairInfo dial eth client failed!")
	}

	defer func() {
		client.Close()
	}()

	block, e := client.BlockNumber(context.Background())
	if e != nil {
		t.Fatal("TestInfo_GetPairInfo get current block failed!", "err", e)
		return
	}
	t.Logf("current block: %d", block)

	//Flashbot(big.NewInt(int64(block)))
}
