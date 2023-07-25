package trading

//
//import (
//	"crypto/ecdsa"
//	"github.com/ethereum/go-ethereum/accounts/abi/bind"
//	"github.com/ethereum/go-ethereum/common"
//	"github.com/ethereum/go-ethereum/common/mclock"
//	"github.com/ethereum/go-ethereum/core/types"
//	"github.com/ethereum/go-ethereum/log"
//	ronswapv3 "github.com/ethereum/go-ethereum/ronfi/contracts/contract_ronswapv3"
//	"github.com/ethereum/go-ethereum/ronfi/loops"
//	"github.com/ethereum/go-ethereum/ronfi/uniswap"
//	"github.com/ethereum/go-ethereum/rpc"
//	"math/big"
//	"os"
//	"runtime"
//	"strings"
//	"sync/atomic"
//
//	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
//	ronswapz "github.com/ethereum/go-ethereum/ronfi/contracts/contract_ronswapz"
//)
//
//var (
//	ronV3Swap *ronswapv3.Ronswapv3
//	ronSwapz  *ronswapz.Ronswapz
//	hostname  string
//	hostId    uint64
//)
//
//func (w *Worker) InitRonFiSwap() bool {
//	var err error
//
//	if ronSwapz, err = ronswapz.NewRonswapz(rcommon.RON_SWAP_X_ADDRESS, w.client); err != nil {
//		log.Error("RonFi InitRonFiSwap New transactor instance Failed on RonSwapZ contract", "err", err)
//		return false
//	}
//
//	if w.v3Hunting {
//		if ronV3Swap, err = ronswapv3.NewRonswapv3(rcommon.RON_V3_SWAP_ADDRESS, w.client); err != nil {
//			log.Error("RonFi InitRonFiSwap New transactor instance Failed on RonV3Swap contract", "err", err)
//			return false
//		}
//	}
//
//	if name, er := os.Hostname(); er == nil {
//		hostname = name
//		log.Info("RonFi InitRonFiSwap Finished", "hostname", hostname)
//		switch hostname {
//		case "aba":
//			hostId = 10
//		case "abb":
//			hostId = 1
//		case "d18a":
//			hostId = 2
//		case "d18b":
//			hostId = 3
//		case "c12c":
//			hostId = 4
//		case "c12h":
//			hostId = 5
//		case "c12i":
//			hostId = 6
//		case "c12j":
//			hostId = 7
//		case "c12k":
//			hostId = 8
//		case "c24a":
//			hostId = 16
//		case "c24b":
//			hostId = 17
//		case "m5f":
//			hostId = 18
//		case "m5m":
//			hostId = 19
//		case "m5n":
//			hostId = 20
//		case "m5p":
//			hostId = 21
//		case "m5q":
//			hostId = 22
//		case "m5s":
//			hostId = 23
//		case "m5t":
//			hostId = 24
//		case "m5u":
//			hostId = 25
//		case "m5v":
//			hostId = 26
//		case "m5x":
//			hostId = 27
//		case "m5y":
//			hostId = 28
//		case "hza":
//			hostId = 30
//		case "hzb":
//			hostId = 31
//		case "hzc":
//			hostId = 32
//		case "hzd":
//			hostId = 33
//		case "m5a":
//			hostId = 34
//		case "m5c":
//			hostId = 35
//		case "m5e":
//			hostId = 36
//		case "m0z":
//			hostId = 37
//		case "m4z":
//			hostId = 38
//		case "m5z":
//			hostId = 39
//		default:
//			hostId = 100
//		}
//	} else {
//		hostId = 100
//	}
//
//	return true
//}
//
//func (w *Worker) DexSwapHunting(executorPrivKey *ecdsa.PrivateKey, executorAddress *common.Address, firstChance bool, tx *types.Transaction, kid int, loop *loops.SwapLoop, profitMin *big.Int, gasLimit uint64, handlerStartTime mclock.AbsTime, dryRun bool, pairId int, copyId int, profitLevel byte) (bool, *types.Transaction) {
//	options, err := bind.NewKeyedTransactorWithChainID(executorPrivKey, big.NewInt(rcommon.CHAIN_ID))
//	if err != nil {
//		log.Error("RonFi swap transaction, get options failed!", "reason", err)
//		return false, nil
//	}
//
//	w.nonceLock.Lock()
//	options.Nonce = new(big.Int).SetUint64(w.nonceState.GetNonce(*executorAddress))
//	w.nonceLock.Unlock()
//
//	options.Value = big.NewInt(0)
//	options.GasLimit = gasLimit
//	if tx.GasPrice().Cmp(GasPriceMaxAllowed) >= 0 {
//		options.GasPrice = GasPriceMinAllowedBusy
//		w.ReportSkipReason(tx, SkipReasonGasPriceHigh, "")
//	} else {
//		options.GasPrice = tx.GasPrice()
//	}
//	options.NoSend = true //only return signedTx
//
//	hops := len(loop.Dirs)
//	tokenPairsAndFee := make([]*big.Int, 2*hops)
//	{
//		loopPath := loop.Path
//		for i := 0; i < hops; i++ {
//			tmp := new(big.Int).SetBytes(loopPath[2*i+1].Bytes())
//			tmp = tmp.Add(tmp, new(big.Int).Lsh(new(big.Int).SetUint64(loop.PoolFee[i]), 160))
//			tmp = tmp.Add(tmp, new(big.Int).Lsh(new(big.Int).SetUint64(loop.Dirs[i]), 176)) // 1 byte for dir
//			if w.isP2pHunting && i == 0 && kid > 0 {
//				tmp = tmp.Add(tmp, new(big.Int).Lsh(new(big.Int).SetUint64(uint64(kid)), 184)) // 4bytes for history knowledge id
//			}
//			tokenPairsAndFee[2*i] = tmp
//
//			if i == hops-1 {
//				tmp = new(big.Int).SetBytes(loop.TargetToken.Bytes())
//			} else {
//				tmp = new(big.Int).SetBytes(loopPath[2*i+2].Bytes())
//			}
//			tmp = tmp.Add(tmp, new(big.Int).Lsh(new(big.Int).SetUint64(loop.TokenFee[i]), 160))
//			tokenPairsAndFee[2*i+1] = tmp
//		}
//	}
//
//	shortHash := tx.Hash().Uint64()
//	short := new(big.Int).Lsh(new(big.Int).SetUint64(shortHash), 192)                   // 192 bits = 24 bytes, this is to get the first 8-bytes of the tx hash
//	short.Add(short, new(big.Int).Lsh(new(big.Int).SetUint64(hostId), 184))             // 1 byte reserved for hostId, 184 bits = 23 bytes
//	short.Add(short, new(big.Int).Lsh(new(big.Int).SetUint64(loop.FlashOk), 176))       // 1 byte reserved for flashOk, 176 bits = 22 bytes
//	short.Add(short, new(big.Int).Lsh(new(big.Int).SetBytes(executorAddress[:7]), 120)) // 8 byte reserved for wallet, 120 bits = 15 bytes
//	short.Add(short, new(big.Int).Lsh(new(big.Int).SetBytes([]byte{profitLevel}), 112)) // 1 byte reserved for profitLevel, 112 bits = 14 bytes
//	profitMin = new(big.Int).Add(profitMin, short)                                      // {8-bytes-tx-hash, 1-byte-hostId, 1-byte-flashOk, 7-bytes-wallet, 1-byte-profitLevel, 14-bytes-profit-min}
//
//	var arbTx *types.Transaction
//	if rpc.UseChiToken {
//		if arb, e := ronSwapz.RonSwapAuto(options, tokenPairsAndFee, profitMin); e != nil {
//			log.Error("RonFi NewTransaction fail", "dexTx", tx.Hash().String(), "elapsed", mclock.Since(handlerStartTime), "err", e)
//			return false, nil
//		} else {
//			arbTx = arb
//		}
//	} else {
//		if arb, e := ronSwapz.RonSwapD(options, tokenPairsAndFee, profitMin); e != nil {
//			log.Error("RonFi NewTransaction fail", "dexTx", tx.Hash().String(), "elapsed", mclock.Since(handlerStartTime), "err", e)
//			return false, nil
//		} else {
//			arbTx = arb
//		}
//	}
//
//	arbTx.RonTxType = types.RonTxIsArbTx
//	arbTx.RonPeerName = "self"
//
//	if firstChance {
//		elapsed := mclock.Since(handlerStartTime)
//		log.Info("RonFi broadcast transaction",
//			"dexTx", tx.Hash(),
//			"pair", pairId,
//			"arbTx", arbTx.Hash(),
//			"gasPrice", arbTx.GasPrice(),
//			"hops", len(loop.Fees),
//			"loopId", loop.LoopId,
//			"gasLimit", gasLimit,
//			"nonce", arbTx.Nonce(),
//			"elapsed", elapsed)
//	}
//
//	if !dryRun {
//		// broadcast it immediately
//		w.eth.RonBroadcastTransaction(arbTx)
//
//		{
//			elapsed := mclock.Since(handlerStartTime)
//			log.Info("RonFi", "dexTx", tx.Hash().String(), "arbTx", arbTx.Hash(), "elapsed", elapsed)
//
//			if !w.isP2pHunting && w.checkArbTx && copyId == 0 {
//				if w.CheckArbTx(executorPrivKey, loop, tx, arbTx) {
//					return true, arbTx
//				} else {
//					return false, arbTx
//				}
//			} else {
//				return true, arbTx
//			}
//		}
//	} else {
//		//Test: broadcast dexTx in dryRun mode.
//		w.eth.RonBroadcastTransaction(tx)
//		log.Info("RonFi ", "dexTx", tx.Hash().String(), "arbTx", arbTx.Hash().String(), "handler", mclock.Since(handlerStartTime), "dryRun", dryRun)
//		if !w.isP2pHunting && w.checkArbTx {
//			if w.CheckArbTx(executorPrivKey, loop, tx, arbTx) {
//				return true, arbTx
//			} else {
//				return false, arbTx
//			}
//		} else {
//			return true, arbTx
//		}
//	}
//}
//
//func (w *Worker) CheckArbTx(executorPrivKey *ecdsa.PrivateKey, loop *loops.SwapLoop, tx *types.Transaction, arbTx *types.Transaction) bool {
//	//notes: since we only care about the 'Pancake: K', 'Transfer_Fail' and 'Forbidden' error for Cancel, it does not mather whether we use new updated appState here or the original appState!
//	appState := w.stateDbProduceOneCopy() // CheckArbTx is not a time critical task! so we use slow stateDB copy here.
//	if appState == nil {
//		return true
//	}
//
//	applySuccess, reverted, e := w.applyTransaction(arbTx, arbTx.Hash(), appState)
//	if !applySuccess || reverted {
//		hasTheseSuffix := strings.HasSuffix(e, rcommon.ErrRonFiNo4) ||
//			strings.HasSuffix(e, rcommon.ErrPANCAKEK) ||
//			strings.HasSuffix(e, rcommon.ErrNomiswapD) ||
//			strings.HasSuffix(e, rcommon.ErrTRANSFERFAIL) ||
//			strings.HasSuffix(e, rcommon.ErrFORBIDDEN) ||
//			strings.HasSuffix(e, rcommon.ErrOVERFLOW)
//		if reverted && hasTheseSuffix {
//			loop.Cancel = true // once a cancel happen, this loop will be cancelled also, never use it anymore!
//		}
//
//		if !w.dryRun {
//			if reverted && hasTheseSuffix {
//				// cancel arbTx
//				cancelTx := w.CancelTransaction(executorPrivKey, arbTx)
//				if cancelTx != nil {
//					log.Warn("RonFi arb tx cancel", "dexTx", tx.Hash(), "arbTx", arbTx.Hash(), "loopId", loop.LoopId, "cancelTx", cancelTx.Hash().String(), "err", e)
//				} else {
//					log.Warn("RonFi arb tx callStatic fail, and its cancel tx fail on creation", "dexTx", tx.Hash(), "arbTx", arbTx.Hash(), "err", e)
//				}
//			} else {
//				log.Warn("RonFi arb tx cancel ignored", "dexTx", tx.Hash(), "arbTx", arbTx.Hash(), "loopId", loop.LoopId, "err", e)
//				return true
//			}
//		} else {
//			if reverted && hasTheseSuffix {
//				log.Warn("RonFi arb tx cancel", "dexTx", tx.Hash(), "arbTx", arbTx.Hash(), "loopId", loop.LoopId, "err", e)
//			} else {
//				log.Warn("RonFi arb tx cancel ignored", "dexTx", tx.Hash(), "arbTx", arbTx.Hash(), "loopId", loop.LoopId, "err", e)
//				return true
//			}
//		}
//		return false
//	} else {
//		return true
//	}
//}
//
//func (w *Worker) CancelTransaction(executorPrivKey *ecdsa.PrivateKey, tx *types.Transaction) *types.Transaction {
//	var data []byte
//	gasPrice := tx.GasPrice()
//	updatedGasPrice := new(big.Int).Add(gasPrice, new(big.Int).Div(gasPrice, big.NewInt(10)))
//	updatedGasPrice = new(big.Int).Add(updatedGasPrice, new(big.Int).SetUint64(10))
//	cancelTx := types.NewTransaction(tx.Nonce(), *tx.To(), big.NewInt(0), uint64(25000), updatedGasPrice, data)
//	signedTx, err := types.SignTx(cancelTx, types.NewEIP155Signer(w.chainConfig.ChainID), executorPrivKey)
//	if err != nil {
//		log.Warn("RonFi arb cancelTransaction SignTx failed!", "tx", tx.Hash().String())
//	} else {
//		atomic.AddUint64(&w.totalArbTxsCanceled, 1)
//		signedTx.RonTxType = types.RonTxIsArbTx
//		w.eth.RonBroadcastTransaction(signedTx)
//
//		return signedTx
//	}
//
//	return nil
//}
//
//func (w *Worker) DexV3SwapHunting(executorPrivKey *ecdsa.PrivateKey, executorAddress *common.Address, tx *types.Transaction, profitMin *big.Int, gasLimit uint64, bestProfit *uniswap.CycleWithProfit, dryRun bool, handlerStartTime mclock.AbsTime) (bool, *types.Transaction) {
//	options, err := bind.NewKeyedTransactorWithChainID(executorPrivKey, big.NewInt(rcommon.CHAIN_ID))
//	if err != nil {
//		log.Error("RonFi swap transaction, get options failed!", "reason", err)
//		return false, nil
//	}
//
//	w.nonceLock.Lock()
//	options.Nonce = new(big.Int).SetUint64(w.nonceState.GetNonce(*executorAddress))
//	w.nonceLock.Unlock()
//
//	options.Value = big.NewInt(0)
//	options.GasLimit = gasLimit
//	if tx.GasPrice().Cmp(GasPriceMaxAllowed) >= 0 {
//		options.GasPrice = GasPriceMinAllowedBusy
//		w.ReportSkipReason(tx, SkipReasonGasPriceHigh, "")
//	} else {
//		options.GasPrice = tx.GasPrice()
//	}
//	options.NoSend = true //only return signedTx
//
//	path := make([]common.Address, 0, len(bestProfit.Cycle.PoolAddresses)*2)
//	dirs := make([]uint8, 0, len(bestProfit.Cycle.PoolAddresses))
//	poolFee := make([]uint64, 0, len(bestProfit.Cycle.PoolAddresses))
//	tokenFee := make([]uint64, 0, len(bestProfit.Cycle.PoolAddresses))
//	poolType := make([]uint8, 0, len(bestProfit.Cycle.PoolAddresses))
//	if len(bestProfit.Cycle.PoolAddresses) != len(bestProfit.Cycle.SwapVectors) {
//		log.Error("RonFi swap transaction, swapVectors and poolAddress length mismatch")
//		return false, nil
//	}
//
//	for i := 0; i < len(bestProfit.Cycle.PoolAddresses); i++ {
//		pool := bestProfit.Cycle.PoolAddresses[i]
//		swapVector := bestProfit.Cycle.SwapVectors[i]
//		token := swapVector.TokenOut
//		dir := uint8(0)
//		if !swapVector.ZeroForOne {
//			dir = 1
//		}
//		dirs = append(dirs, dir)
//		path = append(path, pool)
//		path = append(path, token)
//		poolFee = append(poolFee, swapVector.PoolFee)
//		tokenFee = append(tokenFee, swapVector.TokenFee)
//		poolType = append(poolType, swapVector.PoolType)
//	}
//
//	tokenPairsAndFee := make([]*big.Int, 2*len(bestProfit.Cycle.PoolAddresses))
//	for i := 0; i < len(bestProfit.Cycle.PoolAddresses); i++ {
//		tmp := new(big.Int).SetBytes(path[2*i].Bytes())
//		tmp = tmp.Add(tmp, new(big.Int).Lsh(new(big.Int).SetUint64(poolFee[i]), 160))
//		tmp = tmp.Add(tmp, new(big.Int).Lsh(new(big.Int).SetUint64(uint64(dirs[i])), 176))
//		tmp = tmp.Add(tmp, new(big.Int).Lsh(new(big.Int).SetUint64(uint64(poolType[i])), 192))
//		tokenPairsAndFee[2*i] = tmp
//
//		tmp = new(big.Int).SetBytes(path[2*i+1].Bytes())
//		tmp = tmp.Add(tmp, new(big.Int).Lsh(new(big.Int).SetUint64(tokenFee[i]), 160))
//		tokenPairsAndFee[2*i+1] = tmp
//	}
//	amountIn := bestProfit.Profit.SwapAmount
//
//	var arbTx *types.Transaction
//	if arb, e := ronV3Swap.RonSwapV3(options, tokenPairsAndFee, amountIn); e != nil {
//		log.Error("RonFi swap transaction, create swap transaction failed!", "reason", e)
//		return false, nil
//	} else {
//		arbTx = arb
//	}
//
//	arbTx.RonTxType = types.RonTxIsArbTx
//	arbTx.RonPeerName = "self"
//
//	log.Info("RonFi broadcast transaction",
//		"dexTx", tx.Hash(),
//		"arbTx", arbTx.Hash(),
//		"gasPrice", arbTx.GasPrice(),
//		"gasLimit", gasLimit,
//		"nonce", arbTx.Nonce(),
//		"runningGoRoutine", runtime.NumGoroutine(),
//		"runningApp", atomic.LoadInt64(&w.runningApp),
//		"elapsed", mclock.Since(handlerStartTime).String())
//
//	if !dryRun {
//		// broadcast it immediately
//		w.eth.RonBroadcastTransaction(arbTx)
//		return true, arbTx
//
//	} else {
//		log.Info("RonFi ", "dexTx", tx.Hash().String(), "arbTx", arbTx.Hash().String(), "dryRun", dryRun)
//	}
//
//	return false, nil
//}
