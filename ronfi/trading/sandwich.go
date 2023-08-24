package trading

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/txpool"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
	erc20token "github.com/ethereum/go-ethereum/ronfi/contracts/contract_erc20"
	ronswapv3fe "github.com/ethereum/go-ethereum/ronfi/contracts/contract_ronswapv3fe"
	v2 "github.com/ethereum/go-ethereum/ronfi/contracts/contract_v2pair"
	v3 "github.com/ethereum/go-ethereum/ronfi/contracts/contract_v3pool"
	"github.com/ethereum/go-ethereum/ronfi/defi"
	"github.com/ethereum/go-ethereum/ronfi/uniswap"
	uniswapV2 "github.com/ethereum/go-ethereum/ronfi/uniswap/v2"
	uniswapV3 "github.com/ethereum/go-ethereum/ronfi/uniswap/v3"
	"math/big"
	"strings"
)

var intGasPrice = new(big.Int).Mul(big.NewInt(75), big.NewInt(params.GWei))

type RonSandwich struct {
	client      *ethclient.Client
	di          *defi.Info
	ronSwapInst *ronswapv3fe.Ronswapv3fe
	privKey     *ecdsa.PrivateKey
	execAddress common.Address
	chain       *core.BlockChain
	chainConfig *params.ChainConfig
	block       *types.Block
	targetTx    *types.Transaction
	pool        *defi.SwapPairInfo
	appState    *state.StateDB
	lowerBound  *big.Int
	upperBound  *big.Int
	tolerance   *big.Int
}

func NewRonSandwich(
	client *ethclient.Client,
	di *defi.Info,
	ronSwapInst *ronswapv3fe.Ronswapv3fe,
	executorPrivKey *ecdsa.PrivateKey,
	executorAddress common.Address,
	chain *core.BlockChain,
	chainConfig *params.ChainConfig,
	block *types.Block,
	targetTx *types.Transaction,
	pairInfo *defi.SwapPairInfo,
	appState *state.StateDB) *RonSandwich {
	ronSandWich := &RonSandwich{
		client:      client,
		di:          di,
		ronSwapInst: ronSwapInst,
		privKey:     executorPrivKey,
		execAddress: executorAddress,
		chain:       chain,
		chainConfig: chainConfig,
		block:       block,
		targetTx:    targetTx,
		appState:    appState.Copy(),
	}

	ronSandWich.lowerBound = big.NewInt(0)
	balance := di.GetTokenBalance(txpool.RonFiSwapV3Address, rcommon.WETH)
	if balance == nil {
		log.Warn("RonFi Sandwich NewRonSandwich GetTokenBalance failed", "tx", targetTx.Hash(), "RonFiSwapV3Address", txpool.RonFiSwapV3Address)
		return nil
	} else {
		ronSandWich.upperBound = balance
	}
	ronSandWich.tolerance = new(big.Int).Exp(big.NewInt(10), big.NewInt(15), nil) //0.001 weth

	// only keep pair/pool whose tokenIn is WETH
	if pairInfo.TokenIn != rcommon.WETH {
		return nil
	}

	ronSandWich.pool = pairInfo
	return ronSandWich
}

type RonSandwichPrepRes struct {
	appState     *state.StateDB
	bLegAmountIn *big.Int
	tokenFee     *big.Int
	aLegGasUsed  uint64
	bLegGasUsed  uint64
}

func (s *RonSandwich) prepare(pool *defi.SwapPairInfo, amountIn *big.Int) *RonSandwichPrepRes {
	var aLegAmountOut *big.Int
	appState := s.appState.Copy()
	// aLeg tx
	var (
		feeRate                  *big.Int
		tokenPairsAndFee         []*big.Int
		arbAIn, txFee            *big.Int
		aLegGasUsed, bLegGasUsed uint64
	)

	feeRate = big.NewInt(0)
	arbAIn = big.NewInt(0)
	txFee = big.NewInt(0)
	if res := s.executeABLegTx(pool, amountIn, appState, true, nil, tokenPairsAndFee, arbAIn, txFee); res == nil {
		log.Warn("RonFi Sandwich prepare aLegTx failed", "tx", s.targetTx.Hash(), "pair", pool.Address)
		return nil
	} else {
		aLegAmountOut = res.amountOut
		feeRate = res.feeRate
		aLegGasUsed = res.gasUsed
	}
	log.Info("RonFi Sandwich prepare aLegTx succeed!", "tx", s.targetTx.Hash(), "pair", pool.Address, "amountIn", amountIn, "amountOut", aLegAmountOut, "feeRate", feeRate)

	// target tx
	if applySuccess, reverted, _, err := applyTransaction(s.chain, s.chainConfig, s.block, s.targetTx, ronFiTxHash(s.targetTx.Hash()), appState); !applySuccess || reverted {
		log.Warn("RonFi Sandwich prepare applyTransaction targetTx failed", "tx", s.targetTx.Hash(), "pair", pool.Address, "err", err)
		return nil
	}

	// bLeg tx
	rPool := pool.Reverse()
	bLegAmountIn := aLegAmountOut
	if res := s.executeABLegTx(rPool, bLegAmountIn, appState, false, feeRate, tokenPairsAndFee, arbAIn, txFee); res == nil {
		log.Warn("RonFi Sandwich prepare bLegTx failed",
			"tx", s.targetTx.Hash(),
			"pair", rPool.Address,
			"tokenIn", rPool.TokenIn,
			"amountIn", bLegAmountIn)
	} else {
		log.Info("RonFi Sandwich prepare bLegTx succeed!", "tx", s.targetTx.Hash(), "pair", rPool.Address, "amountIn", bLegAmountIn, "amountOut", res.amountOut)
		if res.amountOut.Cmp(amountIn) <= 0 {
			log.Warn("RonFi Sandwich prepare not profitable", "tx", s.targetTx.Hash(), "pair", pool.Address, "amountIn", amountIn, "amountOut", res.amountOut)
		} else {
			bLegGasUsed = res.gasUsed
			return &RonSandwichPrepRes{
				appState:     appState,
				bLegAmountIn: bLegAmountIn,
				tokenFee:     feeRate,
				aLegGasUsed:  aLegGasUsed,
				bLegGasUsed:  bLegGasUsed,
			}
		}
	}

	return nil
}

func (s *RonSandwich) optimize(pool *defi.SwapPairInfo, amountIn *big.Int) bool {
	calculateF := func(appState *state.StateDB, pool *defi.SwapPairInfo, amountIn *big.Int) (*types.Transaction, bool, bool) {
		var (
			applySuccess, reverted bool
			//err                    string
			tokenPairsAndFee []*big.Int
		)

		arbAIn := big.NewInt(0)
		txFee := big.NewInt(0)
		payloads, _ := s.generatePayloads(pool, amountIn, nil, appState)
		if len(payloads) > 0 {
			// create tx and apply
			nonce := appState.GetNonce(s.execAddress)
			if frontRunTx := s.buildExecuteTx(payloads, true, tokenPairsAndFee, arbAIn, txFee, nonce, intGasPrice, GasLimitDefault); frontRunTx != nil {
				if applySuccess, reverted, _, _ = applyTransaction(s.chain, s.chainConfig, s.block, frontRunTx, ronFiTxHash(frontRunTx.Hash()), appState); applySuccess && !reverted {
					targetTxHash := ronFiTxHash(s.targetTx.Hash())
					if applySuccess, reverted, _, _ = applyTransaction(s.chain, s.chainConfig, s.block, s.targetTx, targetTxHash, appState); applySuccess && !reverted {
						//log.Info("RonFi Sandwich calculateF applyTransaction targetTx succeed!", "targetTx", s.targetTx.Hash().String(), "pair", pool.Address, "amountIn", amountIn)
						return frontRunTx, false, true
					} else {
						return nil, true, false
					}
				}
			}
		}

		return nil, false, false
	}

	return s.binarySearch(pool, calculateF, s.lowerBound, s.upperBound, s.tolerance, amountIn)
}

func (s *RonSandwich) binarySearch(pool *defi.SwapPairInfo, calculateF func(appState *state.StateDB, pool *defi.SwapPairInfo, amountIn *big.Int) (*types.Transaction, bool, bool), left, right *big.Int, tolerance *big.Int, amountIn *big.Int) bool {
	gap := new(big.Int).Sub(right, left)
	if gap.Cmp(tolerance) > 0 {
		mid := new(big.Int).Add(left, new(big.Int).Div(gap, big.NewInt(2)))
		_, terminate, ok := calculateF(s.appState.Copy(), pool, mid)
		if terminate {
			// apply target tx failed, then exit the iteration
			return false
		}
		if ok {
			// Number go up
			amountIn.Set(mid)
			return s.binarySearch(pool, calculateF, mid, right, tolerance, amountIn)
		} else {
			// Number go down
			return s.binarySearch(pool, calculateF, left, mid, tolerance, amountIn)
		}
	}

	if amountIn.Cmp(big.NewInt(0)) > 0 {
		return true
	}

	return false
}

func (s *RonSandwich) genTransTxPayload(token common.Address, to common.Address, amount *big.Int) *ronswapv3fe.RonSwapV3FEPayload {
	var payload *ronswapv3fe.RonSwapV3FEPayload

	if amount == nil || amount.Cmp(big.NewInt(0)) <= 0 {
		return nil
	}

	// Define ABI
	tokenAbi, err := abi.JSON(strings.NewReader(erc20token.TokenABI))
	if err != nil {
		log.Warn("RonFi optimize creating ABI failed!", "err", err)
		return nil
	}
	// Encode arguments
	encodedParams, err := tokenAbi.Pack("transfer", to, amount)
	if err != nil {
		log.Warn("Ronfi optimize encoding arguments failed!", "err", err)
		return nil
	}
	value := big.NewInt(0)
	payload = &ronswapv3fe.RonSwapV3FEPayload{
		Target: token,
		Data:   encodedParams,
		Value:  value,
	}

	return payload
}

func (s *RonSandwich) genSwapTxPayload(appState *state.StateDB, pool *defi.SwapPairInfo, amountIn *big.Int) (payload ronswapv3fe.RonSwapV3FEPayload, amountOut *big.Int, ok bool) {
	if !pool.V3 {
		v2PoolInst := uniswapV2.NewV2Pool(s.di, pool.Address, 30, appState)
		if v2PoolInst == nil {
			log.Warn("RonFi Sandwich genSwapTxPayload NewV2Pool failed", "pair", pool.Address)
			return
		}

		if !v2PoolInst.UpdateReserves(nil) {
			//log.Warn("RonFi Sandwich genSwapTxPayload UpdateReserves failed", "pair", pool.Address)
			return
		}
		amountOut = v2PoolInst.CalculateTokensOutFromTokensIn(pool.TokenIn, amountIn)
		if amountOut == nil || amountOut.Cmp(big.NewInt(0)) <= 0 {
			//log.Warn("RonFi Sandwich genSwapTxPayload CalculateTokensOutFromTokensIn failed", "pair", pool.Address)
			return
		}

		amount0Out := big.NewInt(0)
		amount1Out := big.NewInt(0)
		if pool.Dir == 0 {
			amount1Out = amountOut
		} else {
			amount0Out = amountOut
		}

		if amount0Out.Cmp(big.NewInt(0)) <= 0 && amount1Out.Cmp(big.NewInt(0)) <= 0 {
			log.Warn("RonFi Sandwich genSwapTxPayload amount0Out and amount1Out both <= 0", "pair", pool.Address)
			return
		}

		// Define ABI
		v2PairAbi, err := abi.JSON(strings.NewReader(v2.V2pairABI))
		if err != nil {
			log.Warn("RonFi genSwapTxPayload v2 creating ABI failed!", "err", err)
			return
		}

		var encodedParams []byte
		// Encode arguments
		encodedParams, err = v2PairAbi.Pack("swap", amount0Out, amount1Out, txpool.RonFiSwapV3Address, []byte{})
		if err != nil {
			log.Warn("RonFi genSwapTxPayload v2 encoding arguments failed!", "err", err)
			return
		}

		//log.Info("RonFi genSwapTxPayload", "pair", pool.Address, "amountIn", amountIn, "amountOut", amountOut, "amount0Out", amount0Out, "amount1Out", amount1Out)
		value := big.NewInt(0)
		payload = ronswapv3fe.RonSwapV3FEPayload{
			Target: pool.Address,
			Data:   encodedParams,
			Value:  value,
		}
	} else {
		poolInfo := s.di.GetPoolInfo(pool.Address)
		if poolInfo == nil {
			log.Warn("RonFi Sandwich genSwapTxPayload GetPoolInfo failed", "pool", pool.Address)
			return
		} else {
			tickLens := rcommon.GetV3TickLens(poolInfo.Factory)
			if tickLens == rcommon.ZeroAddress {
				log.Warn("RonFi Sandwich genSwapTxPayload GetV3TickLens failed", "pool", pool.Address)
				return
			}

			v3PoolInst := uniswapV3.NewV3Pool(s.di, s.targetTx, pool.Address, tickLens, appState)
			if v3PoolInst == nil {
				log.Warn("RonFi Sandwich genSwapTxPayload NewV3Pool failed", "pool", pool.Address)
				return
			}

			if !v3PoolInst.UpdatePoolState(nil) {
				//log.Warn("RonFi Sandwich genSwapTxPayload UpdatePoolState failed", "pool", pool.Address)
				return
			}
			amountOut, _ = v3PoolInst.CalculateTokensOutFromTokensIn(pool.TokenIn, amountIn)
			if amountOut == nil || amountOut.Cmp(big.NewInt(0)) <= 0 {
				//log.Warn("RonFi Sandwich genSwapTxPayload CalculateTokensOutFromTokensIn failed", "pair", pool.Address)
				return
			}

			// Define ABI
			v3PoolAbi, err := abi.JSON(strings.NewReader(v3.V3poolABI))
			if err != nil {
				log.Warn("RonFi Sandwich genSwapTxPayload v3 creating ABI failed!", "err", err)
				return
			}

			var encodedParams []byte
			encodedParams, err = v3PoolAbi.Pack("swap", txpool.RonFiSwapV3Address, pool.Dir == 0, amountIn, v3PoolInst.State.SqrtPriceX96, []byte{})
			if err != nil {
				log.Warn("RonFi Sandwich genSwapTxPayload v3  encoding arguments failed!", "err", err)
				return
			}
			value := big.NewInt(0)

			payload = ronswapv3fe.RonSwapV3FEPayload{
				Target: pool.Address,
				Data:   encodedParams,
				Value:  value,
			}
		}
	}

	ok = true
	return
}

func (s *RonSandwich) generatePayloads(pool *defi.SwapPairInfo, amountIn *big.Int, tokenFee *big.Int, appState *state.StateDB) ([]ronswapv3fe.RonSwapV3FEPayload, *big.Int) {
	payloads := make([]ronswapv3fe.RonSwapV3FEPayload, 0, 10)
	if !pool.V3 {
		if payload := s.genTransTxPayload(pool.TokenIn, pool.Address, amountIn); payload != nil {
			payloads = append(payloads, *payload)
		} else {
			return payloads, nil
		}
	}

	realAmountIn := amountIn
	if tokenFee != nil && tokenFee.Cmp(big.NewInt(0)) > 0 {
		realAmountIn = new(big.Int).Div(new(big.Int).Mul(amountIn, tokenFee), big.NewInt(10000))
	}
	if payload, amountOut, ok := s.genSwapTxPayload(appState, pool, realAmountIn); ok && amountOut != nil && amountOut.Cmp(big.NewInt(0)) > 0 {
		payloads = append(payloads, payload)
		return payloads, amountOut
	}

	return payloads, nil
}

func (s *RonSandwich) buildExecuteTx(payloads []ronswapv3fe.RonSwapV3FEPayload, isAleg bool, tokenPairsAndFee []*big.Int, amountIn, txFee *big.Int, nonce uint64, gasPrice *big.Int, gas uint64) *types.Transaction {
	if options, err := bind.NewKeyedTransactorWithChainID(s.privKey, big.NewInt(rcommon.CHAIN_ID)); err != nil {
		return nil
	} else {
		options.Nonce = new(big.Int).SetUint64(nonce)
		options.Value = big.NewInt(0)
		options.GasLimit = gas
		options.GasPrice = gasPrice
		options.NoSend = true //only return signedTx

		if tx, err := s.ronSwapInst.RonSandwich(options, isAleg, payloads, tokenPairsAndFee, amountIn, txFee); err == nil {
			return tx
		}
	}

	return nil
}

type SandwichRes struct {
	tx        *types.Transaction
	feeRate   *big.Int
	amountOut *big.Int
	gasUsed   uint64
}

func (s *RonSandwich) executeABLegTx(pool *defi.SwapPairInfo, amountIn *big.Int, appState *state.StateDB, isAleg bool, tokenFee *big.Int, tokenPairsAndFee []*big.Int, arbAIn, txFee *big.Int) *SandwichRes {
	var (
		applySuccess, reverted bool
		gasUsed                uint64
		err                    string
		swapAmountOut, feeRate *big.Int
	)

	feeRate = big.NewInt(0)
	payloads, amountOut := s.generatePayloads(pool, amountIn, tokenFee, appState)
	if len(payloads) == 0 {
		log.Warn("RonFi executeABLegTx payloads empty", "tx", s.targetTx.Hash(), "pair", pool.Address, "isAleg", isAleg)
		return nil
	}

	if tx := s.buildExecuteTx(payloads, isAleg, tokenPairsAndFee, arbAIn, txFee, appState.GetNonce(s.execAddress), intGasPrice, GasLimitDefault); tx != nil {
		applySuccess, reverted, gasUsed, err = applyTransaction(s.chain, s.chainConfig, s.block, tx, ronFiTxHash(tx.Hash()), appState)
		if !applySuccess || reverted {
			log.Warn("RonFi executeABLegTx applyTransaction failed",
				"tx", s.targetTx.Hash(),
				"pair", pool.Address,
				"tokenIn", pool.TokenIn,
				"amountIn", amountIn,
				"amountOut", amountOut,
				"gasUsed", gasUsed,
				"isAleg", isAleg,
				"applySuccess", applySuccess,
				"reverted", reverted,
				"err", err)
			return nil
		}

		vlogs := appState.GetLogs(ronFiTxHash(tx.Hash()), s.block.NumberU64(), common.Hash{})
		if isAleg {
			swapPairsInfo := s.di.ExtractSwapPairInfo(tx, *tx.To(), vlogs, defi.RonFiExtractTypeStats)
			if len(swapPairsInfo) == 0 {
				return nil
			}

			info := swapPairsInfo[0]
			swapAmountOut = info.AmountOut
		}

		if bAmountOut := s.di.ExtractTransferAmount(vlogs, pool.TokenOut, pool.Address, txpool.RonFiSwapV3Address); bAmountOut != nil && bAmountOut.Cmp(big.NewInt(0)) > 0 {
			if isAleg && bAmountOut.Cmp(swapAmountOut) != 0 {
				feeRate = new(big.Int).Div(new(big.Int).Mul(big.NewInt(10000), new(big.Int).Sub(swapAmountOut, bAmountOut)), swapAmountOut)
				rem := new(big.Int).Rem(feeRate, big.NewInt(10))
				if rem.Cmp(big.NewInt(0)) != 0 {
					feeRate = new(big.Int).Add(feeRate, big.NewInt(1))
				}

				log.Warn("RonFi executeABLegTx bAmountOut != swapAmountOut",
					"tx", s.targetTx.Hash(),
					"pair", pool.Address,
					"tokenIn", pool.TokenIn,
					"amountIn", amountIn,
					"isAleg", isAleg,
					"feeRate", feeRate,
					"swapAmountOut", swapAmountOut,
					"bAmountOut", bAmountOut)
			}
			return &SandwichRes{
				tx:        tx,
				feeRate:   feeRate,
				amountOut: bAmountOut,
				gasUsed:   gasUsed,
			}
		}
	}

	return nil
}

func (s *RonSandwich) buildArbTx(appState *state.StateDB, cycle *uniswap.LPCycle, amountIn *big.Int) (*types.Transaction, bool) {
	options, err := bind.NewKeyedTransactorWithChainID(s.privKey, big.NewInt(rcommon.CHAIN_ID))
	if err != nil {
		log.Error("RonFi buildArbTx, get options failed!", "reason", err)
		return nil, false
	}

	options.Nonce = new(big.Int).SetUint64(appState.GetNonce(s.execAddress))

	options.Value = big.NewInt(0)
	options.GasPrice = intGasPrice
	options.GasLimit = GasLimitOneHop * 10
	options.NoSend = true //only return signedTx

	path := make([]common.Address, 0, len(cycle.PoolAddresses)*2)
	dirs := make([]uint8, 0, len(cycle.PoolAddresses))
	poolFee := make([]uint64, 0, len(cycle.PoolAddresses))
	tokenFee := make([]uint64, 0, len(cycle.PoolAddresses))
	poolType := make([]uint8, 0, len(cycle.PoolAddresses))
	if len(cycle.PoolAddresses) != len(cycle.SwapVectors) {
		log.Error("RonFi buildArbTx, swapVectors and poolAddress length mismatch")
		return nil, false
	}

	for i := 0; i < len(cycle.PoolAddresses); i++ {
		pool := cycle.PoolAddresses[i]
		swapVector := cycle.SwapVectors[i]
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

	tokenPairsAndFee := make([]*big.Int, 2*len(cycle.PoolAddresses))
	for i := 0; i < len(cycle.PoolAddresses); i++ {
		tmp := new(big.Int).SetBytes(path[2*i].Bytes())
		tmp = tmp.Add(tmp, new(big.Int).Lsh(new(big.Int).SetUint64(poolFee[i]), 160))
		tmp = tmp.Add(tmp, new(big.Int).Lsh(new(big.Int).SetUint64(uint64(dirs[i])), 176))
		tmp = tmp.Add(tmp, new(big.Int).Lsh(new(big.Int).SetUint64(uint64(poolType[i])), 192))
		tokenPairsAndFee[2*i] = tmp

		tmp = new(big.Int).SetBytes(path[2*i+1].Bytes())
		tmp = tmp.Add(tmp, new(big.Int).Lsh(new(big.Int).SetUint64(tokenFee[i]), 160))
		tokenPairsAndFee[2*i+1] = tmp
	}

	var arbTx *types.Transaction
	if arb, e := ronV3Swap.RonSwapV3(options, tokenPairsAndFee, amountIn); e != nil {
		log.Error("RonFi buildArbTx, create swap transaction failed!", "reason", e)
		return nil, false
	} else {
		arbTx = arb
	}

	arbTx.RonTxType = types.RonTxIsArbTx
	arbTx.RonPeerName = "self"

	return arbTx, true
}
