package trading

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/mclock"
	"github.com/ethereum/go-ethereum/core/txpool"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
	ronswapv3fe "github.com/ethereum/go-ethereum/ronfi/contracts/contract_ronswapv3fe"
	"github.com/ethereum/go-ethereum/ronfi/uniswap"
	"github.com/metachris/flashbotsrpc"
	"math/big"
	"runtime"
	"sync/atomic"
)

var (
	ronV3Swap *ronswapv3fe.Ronswapv3fe
	flashRpc  *flashbotsrpc.BuilderBroadcastRPC
)

func (w *Worker) InitRonFiSwap() bool {
	var err error

	if ronV3Swap, err = ronswapv3fe.NewRonswapv3fe(txpool.RonFiSwapV3Address, w.client); err != nil {
		log.Error("RonFi InitRonFiSwap New transactor instance Failed on Ronswapv3fe contract", "err", err)
		return false
	}

	// create flash rpc
	if flashRpc = flashbotsrpc.NewBuilderBroadcastRPC(falshbotURLs); err != nil {
		log.Error("RonFi InitRonFiSwap create flashbot rpc failed!", "err", err)
		return false
	}

	return true
}

func (w *Worker) DexSwapHunting(executorPrivKey *ecdsa.PrivateKey, executorAddress *common.Address, tx *types.Transaction, profitMin *big.Int, gasLimit uint64, bestProfit *uniswap.CycleWithProfit, dryRun bool, handlerStartTime mclock.AbsTime) (bool, *types.Transaction) {
	options, err := bind.NewKeyedTransactorWithChainID(executorPrivKey, big.NewInt(rcommon.CHAIN_ID))
	if err != nil {
		log.Error("RonFi swap transaction, get options failed!", "reason", err)
		return false, nil
	}

	w.nonceLock.Lock()
	options.Nonce = new(big.Int).SetUint64(w.nonceState.GetNonce(*executorAddress))
	w.nonceLock.Unlock()

	options.Value = big.NewInt(0)
	options.GasLimit = gasLimit
	//if tx.GasPrice().Cmp(GasPriceMaxAllowed) >= 0 {
	//	options.GasPrice = GasPriceMinAllowedBusy
	//	w.ReportSkipReason(tx, SkipReasonGasPriceHigh, "")
	//} else {
	//	options.GasPrice = tx.GasPrice()
	//}

	//var baseFee *big.Int
	//if feeHis, err := w.client.FeeHistory(context.Background(), 1, new(big.Int).SetUint64(w.currentBlockNum), nil); err != nil {
	//	log.Warn("RonFi swap transaction, FeeHistory failed, err=%s", err)
	//} else {
	//	baseFee = feeHis.BaseFee[0]
	//}

	options.GasPrice = w.gasPrice
	options.NoSend = true //only return signedTx

	path := make([]common.Address, 0, len(bestProfit.Cycle.PoolAddresses)*2)
	dirs := make([]uint8, 0, len(bestProfit.Cycle.PoolAddresses))
	poolFee := make([]uint64, 0, len(bestProfit.Cycle.PoolAddresses))
	tokenFee := make([]uint64, 0, len(bestProfit.Cycle.PoolAddresses))
	poolType := make([]uint8, 0, len(bestProfit.Cycle.PoolAddresses))
	if len(bestProfit.Cycle.PoolAddresses) != len(bestProfit.Cycle.SwapVectors) {
		log.Error("RonFi swap transaction, swapVectors and poolAddress length mismatch")
		return false, nil
	}

	for i := 0; i < len(bestProfit.Cycle.PoolAddresses); i++ {
		pool := bestProfit.Cycle.PoolAddresses[i]
		swapVector := bestProfit.Cycle.SwapVectors[i]
		token := swapVector.TokenOut
		dir := uint8(0)
		if !swapVector.ZeroForOne {
			dir = 1
		}
		dirs = append(dirs, dir)
		path = append(path, pool)
		path = append(path, token)
		poolFee = append(poolFee, swapVector.PoolFee)
		tokenFee = append(tokenFee, swapVector.TokenFee)
		poolType = append(poolType, swapVector.PoolType)
	}

	tokenPairsAndFee := make([]*big.Int, 2*len(bestProfit.Cycle.PoolAddresses))
	for i := 0; i < len(bestProfit.Cycle.PoolAddresses); i++ {
		tmp := new(big.Int).SetBytes(path[2*i].Bytes())
		tmp = tmp.Add(tmp, new(big.Int).Lsh(new(big.Int).SetUint64(poolFee[i]), 160))
		tmp = tmp.Add(tmp, new(big.Int).Lsh(new(big.Int).SetUint64(uint64(dirs[i])), 176))
		tmp = tmp.Add(tmp, new(big.Int).Lsh(new(big.Int).SetUint64(uint64(poolType[i])), 192))
		tokenPairsAndFee[2*i] = tmp

		tmp = new(big.Int).SetBytes(path[2*i+1].Bytes())
		tmp = tmp.Add(tmp, new(big.Int).Lsh(new(big.Int).SetUint64(tokenFee[i]), 160))
		tokenPairsAndFee[2*i+1] = tmp
	}
	amountIn := bestProfit.Profit.SwapAmount

	var arbTx *types.Transaction
	if arb, e := ronV3Swap.RonSwapV3(options, tokenPairsAndFee, amountIn); e != nil {
		log.Error("RonFi swap transaction, create swap transaction failed!", "reason", e)
		return false, nil
	} else {
		arbTx = arb
	}

	arbTx.RonTxType = types.RonTxIsArbTx
	arbTx.RonPeerName = "self"

	log.Info("RonFi broadcast transaction",
		"dexTx", tx.Hash(),
		"arbTx", arbTx.Hash(),
		"gasPrice", arbTx.GasPrice(),
		"gasLimit", gasLimit,
		"nonce", arbTx.Nonce(),
		"runningGoRoutine", runtime.NumGoroutine(),
		"runningApp", atomic.LoadInt64(&w.runningApp),
		"elapsed", mclock.Since(handlerStartTime).String())

	if !dryRun {
		txs := make([]*types.Transaction, 0, 2)
		txs = append(txs, tx)
		txs = append(txs, arbTx)
		Flashbot(flashRpc, w.currentBlock, w.currentBlockNum, txs)
		return true, arbTx

	} else {
		log.Info("RonFi ", "dexTx", tx.Hash().String(), "arbTx", arbTx.Hash().String(), "dryRun", dryRun)
	}

	return false, nil
}
