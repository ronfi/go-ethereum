package trading

import (
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/metachris/flashbotsrpc"
)

//const (
//	j               = "application/json"
//	flashbotURL     = "https://relay.flashbots.net"
//	stats           = "flashbots_getUserStats"
//	flashbotXHeader = "X-Flashbots-Signature"
//	p               = "POST"
//)

var (
	privateKey, _ = crypto.HexToECDSA(
		"22b48b8a5b7d55efafd943d2d25eb0d3cb756097ed3638316c8f364bb52f15bc",
	)

	//walletPrivateKey, _ = crypto.HexToECDSA(
	//	"3ef9395fcde9291cb1d3c9ab02d563597e8951cbdc422cd1e11776f93b8f224f",
	//)
)

//func flashbotHeader(signature []byte, privateKey *ecdsa.PrivateKey) string {
//	return crypto.PubkeyToAddress(privateKey.PublicKey).Hex() +
//		":" + hexutil.Encode(signature)
//}

func Flashbot(flashRpc *flashbotsrpc.BuilderBroadcastRPC, currentBlock uint64, txs []*types.Transaction) {
	var (
		dexTx []byte
		arbTx []byte
		err   error
	)
	dexTx, err = txs[0].MarshalBinary()
	if err != nil {
		log.Warn("Flashbot marshal dex tx failed", "dexTx", txs[0].Hash().String(), "err", err)
	}
	arbTx, err = txs[1].MarshalBinary()
	if err != nil {
		log.Warn("Flashbot marshal arb tx failed", "arbTx", txs[1].Hash().String(), "err", err)
	}

	expectedBlock := currentBlock + 1
	sendBundleArgs := flashbotsrpc.FlashbotsSendBundleRequest{
		Txs:         []string{fmt.Sprintf("0x%X", dexTx), fmt.Sprintf("0x%X", arbTx)},
		BlockNumber: fmt.Sprintf("0x%x", expectedBlock),
	}

	results := flashRpc.BroadcastBundle(privateKey, sendBundleArgs)
	for i, result := range results {
		if result.Err != nil {
			log.Warn("Falshbot broadcast failed", "tx", txs[i].Hash().String(), "err", result.Err)
		}
		log.Info("RonFi Flashbot", "response", result.BundleResponse.BundleHash)
	}
}
