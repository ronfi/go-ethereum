package stats

import (
	"encoding/binary"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/mclock"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
	"github.com/ethereum/go-ethereum/ronfi/defi"
	"github.com/ethereum/go-ethereum/ronfi/dexparser"
	"math/big"
	"sort"
	"strings"
)

type TargetDexInfo struct {
	Index   int
	Success bool
}

type DexPairInfo struct {
	address common.Address
	token0  defi.TokenInfo
	token1  defi.TokenInfo
}

const maxLenCheckSum = 256 // Circular Buffer for Input Data Checksum
var (
	PrevBlockTargetDexTxs = map[uint64]TargetDexInfo{}
	PrevBlockTxs          int
	Gap0ToStats           = map[string]int{}

	MyInputCheckSums = [maxLenCheckSum]uint64{}
	MyIndexCheckSums = 0
	InputCheckSums   = [maxLenCheckSum]uint64{}
	IndexCheckSums   = 0
)

func (s *Stats) report(header *types.Header) {
	bc := s.chain
	if bc == nil {
		log.Error("RonFi Stats report: chain issue, bc == nil")
		return
	}

	blockNumber := header.Number.Uint64()
	block := bc.GetBlockByNumber(blockNumber)
	if block == nil {
		log.Warn("RonFi Stats report: block == nil", "blockNumber", blockNumber)
		return
	}

	receipts := bc.GetReceiptsByHash(block.Hash())
	signer := types.MakeSigner(s.chain.Config(), header.Number, header.Time)
	var methodID uint32
	blockTxs := block.Transactions()
	TargetDexTxs := make(map[uint64]TargetDexInfo, len(blockTxs))
	defer func() {
		PrevBlockTargetDexTxs = TargetDexTxs
		PrevBlockTxs = len(blockTxs)
	}()

	// pre-processing first, to collect the dex txs short hash
	for _, tx := range blockTxs {
		to := tx.To()
		data := tx.Data()
		if to == nil || len(data) == 0 {
			continue
		}
		if len(data) >= 4 {
			methodID = binary.BigEndian.Uint32(tx.Data()[:4])
		} else {
			methodID = 0
		}

		if _, ok := dexparser.ParsableDexMethods[methodID]; ok {
			txLookup := bc.GetTransactionLookup(tx.Hash())
			if txLookup == nil || txLookup.BlockHash != block.Hash() {
				// maybe chain reorg
				continue
			}
		}
	}

	// pre-processing, find out all target dex txs index. If a target missing, it could have been overwritten and not on chain anymore!
	for _, tx := range blockTxs {
		dexTxId := tx.Hash().Uint64()
		if info, exist := TargetDexTxs[dexTxId]; exist {
			txLookup := bc.GetTransactionLookup(tx.Hash())
			if txLookup == nil || txLookup.BlockHash != block.Hash() {
				// maybe chain reorg
				continue
			}
			info.Index = int(txLookup.Index)
			if txLookup.Index < uint64(len(receipts)) {
				if receipt := receipts[txLookup.Index]; receipt.Status == 1 {
					info.Success = true
				}
			}

			TargetDexTxs[dexTxId] = info
		}
	}

	for _, tx := range blockTxs {
		to := tx.To()
		if to == nil {
			continue
		}

		data := tx.Data()
		if len(data) >= 4 {
			methodID = binary.BigEndian.Uint32(tx.Data()[:4])
		} else {
			methodID = 0
		}

		txLookup := bc.GetTransactionLookup(tx.Hash())
		if txLookup == nil || txLookup.BlockHash != block.Hash() {
			// maybe chain reorg
			continue
		}
		if txLookup.Index >= uint64(len(receipts)) {
			log.Error("RonFi report", "arbTx", tx.Hash().String(), "BlockIndex", txLookup.BlockIndex, "Index", txLookup.Index, "receipts", len(receipts))
			continue
		}
		receipt := receipts[txLookup.Index]

		isDex, isObs := s.di.CheckIfObsTx(nil, tx, receipt.Logs, *tx.To())
		if isDex {
			if receipt.Status == 1 {
				var token common.Address
				amountF := 0.0
				found := false
				for _, receiptLog := range receipt.Logs {
					if len(receiptLog.Topics) > 0 && receiptLog.Topics[0] == state.TokenTransferEvent {
						token = receiptLog.Address
						if tInfo := s.di.GetTokenInfo(token); tInfo != nil {
							if len(receiptLog.Data) != 32 {
								continue
							}
							amount := new(big.Int).SetBytes(receiptLog.Data[:32])
							amountF = rcommon.ToFloat(amount, tInfo.Decimals)
							found = true
							break //todo: check the wrong logic here! why break? missed stats for other transferred tokens in this tx.
						}
					} else {
						continue
					}
				}

				if found {
					swapPairsInfo := s.di.ExtractSwapPairInfo(nil, nil, tx, *tx.To(), receipt.Logs, defi.RonFiExtractTypeStats)
					if len(swapPairsInfo) > 0 {
						// stats dex volume
						v, exist := s.dexTokensVol[token]
						if !exist {
							s.dexTokensVol[token] = v
						} else {
							s.dexTokensVol[token] = v + amountF
						}

						amountInUSD := defi.GetAmountInUSD(amountF, token) //todo: check what happen if this token is not in my limited list?
						for _, pair := range swapPairsInfo {
							old, ok := s.dexPairsVol[pair.Address]
							if !ok {
								s.dexPairsVol[pair.Address] = amountInUSD
							} else {
								s.dexPairsVol[pair.Address] = old + amountInUSD
							}
						}
					}
				}
			}
		} else if isObs {
			number := block.NumberU64()
			from, _ := types.Sender(signer, tx)

			var obsId ObsId
			switch *to {
			case Obs1SwapAddr:
				obsId = Obs1
			case Obs2SwapAddr:
				obsId = Obs2
			case Obs3SwapAddr:
				obsId = Obs3
			default:
				obsId = Obsx
			}
			s.obsReport(obsId, number, tx, from, tx.Hash().String(), methodID, data, receipt)
		}
	}
}

func (s *Stats) dexReport(block uint64, txHash string, receipt *types.Receipt) {
	switch len(receipt.Logs) {
	case 0:
		s.dexTotalFail++
		log.Info("RonFi arb dex", "dexTx", txHash, "b", block, "index", receipt.TransactionIndex, "fail", s.dexTotalFail)
	default:
		log.Info("RonFi arb dex", "dexTx", txHash, "b", block, "index", receipt.TransactionIndex)
	}
}

func (s *Stats) obsReport(
	id ObsId,
	block uint64,
	tx *types.Transaction,
	from common.Address,
	txHash string,
	methodID uint32,
	data []byte,
	receipt *types.Receipt) {
	if len(data) < 4 {
		log.Warn(fmt.Sprintf("RonFi arb obs%s action unknown", id), "tx", tx.Hash().String(), "b", block)
		return
	}

	s.obsStats.updateTotalArbTxs(id)
	simpleSum := s.parseObs(id, block, tx, from, txHash, methodID, data, receipt)
	txFeeInBnb := new(big.Int).Mul(tx.GasPrice(), new(big.Int).SetUint64(receipt.GasUsed))
	txFee := rcommon.EthBigInt2Float64(txFeeInBnb)
	price := defi.GetTradingTokenPrice(rcommon.USDC)
	txFeeInUSD := price * txFee
	s.obsStats.updateTotalFee(id, txFeeInUSD)

	// check duplicated input[]
	isCopyHunting := false
	for _, oldSum := range InputCheckSums {
		if oldSum == simpleSum {
			isCopyHunting = true
			break
		}
	}
	if !isCopyHunting {
		InputCheckSums[IndexCheckSums%maxLenCheckSum] = simpleSum
		IndexCheckSums++
	} else {
		s.obsStats.updateTotalCopied(id)
	}

	suffix := ""
	abbrFrom := prefix6HexString(from.String())

	status := false
	switch id {
	default:
		if len(receipt.Logs) > 2 {
			status = true
		}
	}

	if !status {
		s.obsStats.updateTotalFail(id)
		if !isCopyHunting {
			suffix = "         fail"
		} else {
			suffix = "    copy fail"
		}
		log.Info(fmt.Sprintf("RonFi arb obs%s", id)+suffix,
			"arbTx", txHash,
			"b", block,
			"index", receipt.TransactionIndex,
			"e", abbrFrom,
			"f", rcommon.Float2Str(txFeeInUSD, 3),
			"#", s.obsStats.getTotalFail(id))
	} else {
		s.obsStats.updateTotalSuccess(id)
		if !isCopyHunting {
			suffix = "         success"
		} else {
			suffix = "    copy success"
		}
		allRatio := s.obsStats.ratio(id)
		profit, v3Hunting := s.di.GetArbTxProfit(tx, receipt.Logs, *tx.To())
		s.obsStats.updateGrossProfit(id, profit)
		if v3Hunting {
			s.obsStats.updateGrossV3Profit(id, profit)
		}

		log.Info(fmt.Sprintf("RonFi arb obs%s", id)+suffix,
			"arbTx", txHash,
			"v3", v3Hunting,
			"b", block,
			"index", receipt.TransactionIndex,
			"e", abbrFrom,
			"f", rcommon.Float2Str(txFeeInUSD, 3),
			"t", s.obsStats.getTotalSuccess(id),
			"ok(%)", rcommon.Float2Str(allRatio, 1),
			"p", rcommon.Float2Str(profit, 3))
	}
}

func (s *Stats) parseObs(
	id ObsId,
	block uint64,
	tx *types.Transaction,
	from common.Address,
	txHash string,
	methodID uint32,
	data []byte,
	receipt *types.Receipt) (simpleSum uint64) {
	switch id {
	default:
		if len(data) > 4 {
			simpleSum = simpleCheckSum(data[4:]) + block + binary.BigEndian.Uint64(from[12:])
		} else {
			simpleSum = block + binary.BigEndian.Uint64(from[12:])
		}
	}

	return
}

func (s *Stats) getPairNames(pair common.Address) (string, string, string) {
	pairName := ""
	token0Sym := ""
	token1Sym := ""
	if pInfo := s.di.GetPairInfo(pair); pInfo != nil {
		pairName = pInfo.Name
		if t0Info := s.di.GetTokenInfo(pInfo.Token0); t0Info != nil {
			token0Sym = t0Info.Symbol
		}
		if t1Info := s.di.GetTokenInfo(pInfo.Token1); t1Info != nil {
			token1Sym = t1Info.Symbol
		}
	} else if pInfo := s.di.GetPoolInfo(pair); pInfo != nil {
		pairName = pInfo.Name
		if t0Info := s.di.GetTokenInfo(pInfo.Token0); t0Info != nil {
			token0Sym = t0Info.Symbol
		}
		if t1Info := s.di.GetTokenInfo(pInfo.Token1); t1Info != nil {
			token1Sym = t1Info.Symbol
		}
	}

	return pairName, token0Sym, token1Sym
}

func (s *Stats) pairStatsReport() {
	// obs4
	for _, id := range PairsObsIds {
		switch id {
		case Obs1, Obs3, Obs4, Obs5, Obs7:
			{
				ss := s.obsPairStats.topN(id, 20)
				for i, pf := range ss {
					pairName, token0Sym, token1Sym := s.getPairNames(pf.Key)
					countInRonSwap := s.obsPairStats.freqInObs(id, pf.Key)
					log.Info(fmt.Sprintf("RonFi arb obs%s pair stats", id),
						"rank",
						i+1,
						"pair", strings.ToLower(pf.Key.String()),
						"dex", pairName,
						"t0", token0Sym,
						"t1", token1Sym,
						"count", pf.Value,
						"ronfi", countInRonSwap)

				}
			}
			log.Info("RonFi arb") // empty line to separate display, for easy of life
		case Ron:
			{
				ss := s.obsPairStats.topN(id, 20)
				for i, pf := range ss {
					pairName, token0Sym, token1Sym := s.getPairNames(pf.Key)
					countInObs4Swap := s.obsPairStats.freqInObs(Obs4, pf.Key)
					countInObs1Swap := s.obsPairStats.freqInObs(Obs1, pf.Key)
					countInObs3Swap := s.obsPairStats.freqInObs(Obs3, pf.Key)
					countInObs5Swap := s.obsPairStats.freqInObs(Obs5, pf.Key)
					countInObs7Swap := s.obsPairStats.freqInObs(Obs7, pf.Key)

					log.Info("RonFi arb ron pair stats",
						"rank", i+1,
						"pair", strings.ToLower(pf.Key.String()),
						"dex", pairName,
						"t0", token0Sym,
						"t1", token1Sym,
						"count", pf.Value,
						"obs4", countInObs4Swap,
						"obs1", countInObs1Swap,
						"obs3", countInObs3Swap,
						"obs5", countInObs5Swap,
						"obs7", countInObs7Swap)
				}
			}
		}
	}
}

func (s *Stats) dexVolumeReport() {
	// sort pairs by volume in descending
	type kv struct {
		Key   common.Address
		Value float64
	}
	var ss []kv
	for k, v := range s.dexPairsVol {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	log.Info("RonFi arb top20 dex pairs:")
	for index, v := range ss {
		if index < 20 {
			pair := v.Key
			vol := v.Value
			pairName := ""
			token0Sym := ""
			token1Sym := ""
			if pairInfo, ok := s.topDexPairsInfo[pair]; !ok {
				pairName, token0Sym, token1Sym = s.getPairNames(pair)
			} else {
				token0Sym = pairInfo.token0.Symbol
				token1Sym = pairInfo.token1.Symbol
			}
			info := fmt.Sprintf("%s (%-15s) -- (%-8s , %-8s)", strings.ToLower(pair.String()), pairName, token0Sym, token1Sym)
			log.Info("RonFi arb top dex pairs:", "top", index+1, "pair", info, "volume($M)", vol/1_000_000)
		} else {
			break
		}
	}
}

func (s *Stats) obsProfitReport(id ObsId) {
	log.Info("RonFi arb")

	shortFee := s.obsStats.shortFee(id)
	log.Info(fmt.Sprintf("RonFi arb obs%s bot fee: (short)", id), "$", rcommon.Float2Str(shortFee, 2))
	s.obsStats.updateReportedTotalFee(id, s.obsStats.getTotalFee(id))

	log.Info(fmt.Sprintf("RonFi arb obs%s bot fee: (total)", id), "$", rcommon.Float2Str(s.obsStats.getTotalFee(id), 2))
	copyRate := s.obsStats.copyRate(id)
	log.Info(fmt.Sprintf("RonFi arb obs%s txs copied", id), "copyRate(%)", copyRate, "totalCopied", s.obsStats.getTotalCopied(id))

	obsTotalArbTxs := s.obsStats.getTotalArbTxs(id)
	txSent := obsTotalArbTxs - s.obsStats.getReportedTotalArbTxs(id)
	s.obsStats.updateReportedTotalArbTxs(id, obsTotalArbTxs)

	gross := s.obsStats.getGrossProfit(id) - s.obsStats.getPrevGrossProfit(id)
	grossV3 := s.obsStats.getPrevGrossV3Profit(id) - s.obsStats.getPrevGrossV3Profit(id)
	net := gross - shortFee
	log.Info(fmt.Sprintf("RonFi arb obs%s bot profit: (short)", id),
		"gross($)", rcommon.Float2Str(gross, 2),
		"net($)", rcommon.Float2Str(net, 2),
		"grossV3($)", rcommon.Float2Str(grossV3, 2),
		"txs", txSent,
	)
	s.obsStats.updatePrevGrossProfit(id, s.obsStats.getGrossProfit(id))
	s.obsStats.updatePrevGrossV3Profit(id, s.obsStats.getGrossV3Profit(id))

	huntingRate := float64(obsTotalArbTxs) / float64(s.txCount) * 1000
	gross = s.obsStats.getGrossProfit(id)
	grossV3 = s.obsStats.getGrossV3Profit(id)
	net = gross - s.obsStats.getTotalFee(id)
	log.Info(fmt.Sprintf("RonFi arb obs%s bot profit: (total)", id),
		"gross($)", rcommon.Float2Str(gross, 2),
		"net($)", rcommon.Float2Str(net, 2),
		"grossV3($)", rcommon.Float2Str(grossV3, 2),
		"huntingRate(‰)", huntingRate,
		"txs", obsTotalArbTxs,
		"runtime", mclock.Since(s.startTime),
	)
}
