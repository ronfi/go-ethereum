package trading

import (
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/metachris/flashbotsrpc"
)

var (
	privateKey, _ = crypto.HexToECDSA(
		"22b48b8a5b7d55efafd943d2d25eb0d3cb756097ed3638316c8f364bb52f15bc",
	)
)

func Flashbot(flashRpc *flashbotsrpc.BuilderBroadcastRPC, block *types.Block, currentBlock uint64, txs []*types.Transaction) {
	dexTxBytes, err := txs[0].MarshalBinary()
	if err != nil {
		log.Warn("RonFi Flashbot encode dexTx failed", "err", err)
	}
	dexTxHex := fmt.Sprintf("0x%x", dexTxBytes)

	arbTxBytes, err := rlp.EncodeToBytes(txs[1])
	if err != nil {
		log.Warn("RonFi Flashbot encode arbTx failed", "err", err)
	}
	arbTxHex := fmt.Sprintf("0x%x", arbTxBytes)

	expectedBlock := currentBlock + 1

	simulateFlashbotBundle(dexTxHex, arbTxHex, expectedBlock)

	sendBundleArgs := flashbotsrpc.FlashbotsSendBundleRequest{
		Txs:         []string{dexTxHex, arbTxHex},
		BlockNumber: fmt.Sprintf("0x%x", expectedBlock),
	}

	results := flashRpc.BroadcastBundle(privateKey, sendBundleArgs)
	for _, result := range results {
		if result.Err != nil {
			log.Warn("RonFi Falshbot broadcast failed", "tx", txs[1].Hash().String(), "err", result.Err)
		} else {
			log.Info("RonFi Flashbot", "response", result.BundleResponse.BundleHash)
		}
	}
}

func simulateFlashbotBundle(dexTxHex, arbTxHex string, expectedBlock uint64) {
	// simulate
	flashBotRpc := flashbotsrpc.NewFlashbotsRPC("https://relay.flashbots.net")

	callBundleArgs := flashbotsrpc.FlashbotsCallBundleParam{
		Txs:              []string{dexTxHex, arbTxHex},
		BlockNumber:      fmt.Sprintf("0x%x", expectedBlock),
		StateBlockNumber: "latest",
	}
	result, err := flashBotRpc.FlashbotsCallBundle(privateKey, callBundleArgs)
	if err != nil {
		log.Warn("RonFi Falshbot call failed", "err", err)
	} else {
		log.Info("RonFi Flashbot", "response", result)
	}
}

func FlashbotSandWich(flashRpc *flashbotsrpc.BuilderBroadcastRPC, block *types.Block, currentBlock uint64, txs []*types.Transaction) {
	if aLegTxData, err := txs[0].MarshalBinary(); err != nil {
		log.Warn("RonFi Flashbot encode aLegTx failed", "err", err)
		return
	} else {
		aLegTxHex := fmt.Sprintf("0x%x", aLegTxData)
		if targetTx, err := txs[1].MarshalBinary(); err != nil {
			log.Warn("RonFi Flashbot encode targetTx failed", "err", err)
			return
		} else {
			targetTxHex := fmt.Sprintf("0x%x", targetTx)

			if bLegTxData, err := txs[2].MarshalBinary(); err != nil {
				log.Warn("RonFi Flashbot encode bLegTx failed", "err", err)
				return
			} else {
				bLegTxHex := fmt.Sprintf("0x%x", bLegTxData)

				expectedBlock := currentBlock + 1

				simulateFlashbotSandWichBundle(aLegTxHex, targetTxHex, bLegTxHex, expectedBlock)

				sendBundleArgs := flashbotsrpc.FlashbotsSendBundleRequest{
					Txs:         []string{aLegTxHex, targetTxHex, bLegTxHex},
					BlockNumber: fmt.Sprintf("0x%x", expectedBlock),
				}

				results := flashRpc.BroadcastBundle(privateKey, sendBundleArgs)
				for _, result := range results {
					if result.Err != nil {
						log.Warn("RonFi FlashbotSandWich broadcast failed", "tx", txs[1].Hash().String(), "err", result.Err)
					} else {
						log.Info("RonFi FlashbotSandWich", "response", result.BundleResponse.BundleHash)
					}
				}
			}
		}
	}
}

func simulateFlashbotSandWichBundle(aLegTxHex, targetTxHex, bLegTxHex string, expectedBlock uint64) {
	// simulate
	flashBotRpc := flashbotsrpc.NewFlashbotsRPC("https://relay.flashbots.net")

	callBundleArgs := flashbotsrpc.FlashbotsCallBundleParam{
		Txs:              []string{aLegTxHex, targetTxHex, bLegTxHex},
		BlockNumber:      fmt.Sprintf("0x%x", expectedBlock),
		StateBlockNumber: "latest",
	}
	result, err := flashBotRpc.FlashbotsCallBundle(privateKey, callBundleArgs)
	if err != nil {
		log.Warn("RonFi simulateFlashbotSandWichBundle call failed", "err", err)
	} else {
		log.Info("RonFi simulateFlashbotSandWichBundle", "response", result)
	}
}
