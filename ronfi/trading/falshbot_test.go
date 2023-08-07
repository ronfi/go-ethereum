package trading

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/metachris/flashbotsrpc"
	"log"
	"math/big"
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
	// Replace with your own sender and recipient addresses
	senderAddress := common.HexToAddress("0xe9Abf280082df456a728Eadfc1C60eab79Bb1D9f")
	recipientAddress := common.HexToAddress("0x152f6653A84d52856327D30BabE61823dcF0c65a")

	privateKey, _ = crypto.HexToECDSA(
		"3ef9395fcde9291cb1d3c9ab02d563597e8951cbdc422cd1e11776f93b8f224f",
	)

	// Get the current nonce for the sender address
	nonce, err := client.PendingNonceAt(context.Background(), senderAddress)
	if err != nil {
		log.Fatalf("Failed to get nonce: %v", err)
	}

	// Set the gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())

	// Set the gas limit
	gasLimit := uint64(21000)

	// Set the amount to send (in wei)
	amount := big.NewInt(10000000000000000) // 0.1 ETH

	// Create the transaction
	//tx := types.NewTransaction(nonce, recipientAddress, amount, gasLimit, gasPrice, nil)

	// Sign the transaction
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		t.Fatalf("Failed to get chain ID: %v", err)
	}

	// Create the transaction
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		To:       &recipientAddress,
		Value:    amount,
		Data:     nil,
	})

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		t.Fatalf("Failed to sign transaction: %v", err)
	}
	t.Logf("signedTx ChainId: %v", signedTx.ChainId())

	// Convert signed transaction to raw bytes
	rawTxBytes, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		log.Fatalf("Failed to encode transaction to bytes: %v", err)
	}

	// Convert raw bytes to hexadecimal string
	rawTxHex := fmt.Sprintf("0x%x", rawTxBytes)

	rpc := flashbotsrpc.New("https://relay-goerli.flashbots.net")
	opts := flashbotsrpc.FlashbotsSendBundleRequest{
		Txs:         []string{rawTxHex},
		BlockNumber: fmt.Sprintf("0x%x", block+1),
	}

	result, err := rpc.FlashbotsSendBundle(privateKey, opts)
	if err != nil {
		if errors.Is(err, flashbotsrpc.ErrRelayErrorResponse) { // user/tx error, rather than JSON or network error
			fmt.Println(err.Error())
		} else {
			fmt.Printf("error: %+v\n", err)
		}
		return
	}

	// Print result
	fmt.Printf("%+v\n", result)

	//Flashbot(big.NewInt(int64(block)))
}
