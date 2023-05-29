package stats

import (
	"encoding/binary"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/ronfi/defi"
	"github.com/ethereum/go-ethereum/ronfi/dexparser"
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
		isObs := s.di.CheckIfObsTx(nil, tx, receipt.Logs, *tx.To())
		if isObs {
			log.Info("RonFi arb obs", "arbTx", tx.Hash().String(), "methodId", methodID)
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
