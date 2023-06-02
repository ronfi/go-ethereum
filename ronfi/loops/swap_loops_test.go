package loops

import (
	"context"
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	db "github.com/ethereum/go-ethereum/ronfi/db"
	"github.com/ethereum/go-ethereum/ronfi/defi"
)

func TestToLoopId(t *testing.T) {
	swapLoop := SwapLoop{
		TargetToken: common.HexToAddress("0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c"),
		Path: []common.Address{
			common.HexToAddress("0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c"),
			common.HexToAddress("0xea6df8a481e0fcbb5f713916c092eda16f93c731"),
			common.HexToAddress("0x8b4d4e4fb663725e19087d384b1c8c1a10cde9f7"),
			common.HexToAddress("0xd8fb853c205a0a49bc7fe0d5dc912fea9135bf7a"),
			common.HexToAddress("0x55d398326f99059ff775485246999027b3197955"),
			common.HexToAddress("0xbf188c8ca1f4b846fd274ee65d3f3bf41744f4ef"),
		},
		PoolFee:  []uint64{25, 525, 30},
		TokenFee: []uint64{0, 500, 0},
		Dirs:     []uint64{0, 500, 0},
		Index:    []uint64{8, 8, 8},
		LoopId:   common.HexToHash("0x81825ea35b5feb16da2045c3789814489ba39b78913c4b6d377d599b91dee4ab"),
	}

	loopId := swapLoop.toLoopId()
	if loopId != swapLoop.LoopId {
		t.Fatalf("expected: %v, bug got: %v", swapLoop.LoopId, loopId)
	}
}

func TestToKey(t *testing.T) {
	swapLoop := SwapLoop{
		TargetToken: common.HexToAddress("0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c"),
		Path: []common.Address{
			common.HexToAddress("0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c"),
			common.HexToAddress("0xea6df8a481e0fcbb5f713916c092eda16f93c731"),
			common.HexToAddress("0x8b4d4e4fb663725e19087d384b1c8c1a10cde9f7"),
			common.HexToAddress("0xd8fb853c205a0a49bc7fe0d5dc912fea9135bf7a"),
			common.HexToAddress("0x55d398326f99059ff775485246999027b3197955"),
			common.HexToAddress("0xbf188c8ca1f4b846fd274ee65d3f3bf41744f4ef"),
		},
		PoolFee:  []uint64{25, 525, 30},
		TokenFee: []uint64{0, 500, 0},
		Dirs:     []uint64{0, 500, 0},
		Index:    []uint64{8, 8, 8},
		LoopId:   common.HexToHash("0x81825ea35b5feb16da2045c3789814489ba39b78913c4b6d377d599b91dee4ab"),
		key:      "bb4cdb9c73bc095c-ea6df8a46f93c731-8b4d4e4f10cde9f7-d8fb853c9135bf7a-55d39832b3197955-bf188c8c1744f4ef",
	}

	key := swapLoop.toKey()
	if key != swapLoop.Key() {
		t.Fatalf("expected: %v, bug got: %v", swapLoop.Key(), key)
	}
}

func TestMatchLoops(t *testing.T) {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		t.Fatal("TestMatchLoops dial eth client failed!")
	}

	defer func() {
		if client != nil {
			client.Close()
		}
	}()

	conf := rcommon.MysqlConfig{
		DbHost: "176.9.120.196",
		DbPort: "3306",
		DbUser: "root",
		DbPass: "rkdb",
		DbData: "rkdb",
	}
	mysql := db.NewMysql(conf)
	if mysql == nil {
		t.Fatalf("TestMatchLoops NewMysql failed!")
	}

	di := defi.NewInfo(client, mysql)
	if di == nil {
		t.Fatalf("TestMatchLoops NewInfo failed!")
	}

	pairGasMap := mysql.LoadPairGas()
	t.Logf("LoadPairGas: %v", len(pairGasMap))

	feePatchMap := mysql.LoadFeePatch()
	t.Logf("LoadFeePatch: %v", len(feePatchMap))

	loopsIdMap := make(LIdMap)
	flashNokPairs := make(map[common.Address]uint64)

	// load loops
	loopsMap := LoadSwapLoops(
		mysql,
		di,
		loopsIdMap,
		feePatchMap,
		pairGasMap,
		flashNokPairs,
		false,
	)
	t.Logf("total %v loops keys, %v pairs, %v uni pairs", len(loopsMap.LoopsDB), len(loopsMap.AllPairsMap), len(loopsMap.UniPairsMap))

	archiveClient, err := ethclient.Dial("https://nd-814-711-835.p2pify.com/049f42c7290c310495b6940701e2ae14")
	if err != nil {
		t.Fatal("TestMatchLoops dial archive provider failed!")
	}

	defer func() {
		if client != nil {
			client.Close()
		}
	}()

	txHash := common.HexToHash("0x79a9b0cf0e3a2e0c1f2f8636ad4e651a46c0c64f659db7a450b4c7a9df17489b")
	tx, _, err := archiveClient.TransactionByHash(context.Background(), txHash)
	if err != nil {
		t.Fatalf("TestMatchLoops retrieve transaction failed!")
	}

	receipt, err := archiveClient.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		t.Fatalf("TestMatchLoops retrieve transaction receipt failed!")
	}

	swapPairsInfo := di.ExtractSwapPairInfo(loopsMap.AllPairsMap, nil, tx, rcommon.ZeroAddress, receipt.Logs, defi.RonFiExtractTypeHunting)
	if len(swapPairsInfo) == 0 {
		t.Fatalf("TestMatchLoops no swap info!")
	}

	pairsMatchedLoops := MatchLoops(swapPairsInfo, loopsMap, nil, 10)
	if len(pairsMatchedLoops) == 0 {
		t.Fatalf("TestMatchLoops no matched loops!")
	}

	//for i, swapInfos := range pairsMatchedLoops {
	//	for _, swapInfo := range swapInfos {
	//		t.Logf("pairId: %d, loopId: %v", i, swapInfo.toLoopId())
	//	}
	//}

	if pairsMatchedLoops[0][0].LoopId != common.HexToHash("0x4f3db4dad9961cf3233cba8fbd3e64b2369f6b73f96c1992909d327bb7e91cab") {
		t.Fatalf("expected: %v, but got: %v", "0x4f3db4dad9961cf3233cba8fbd3e64b2369f6b73f96c1992909d327bb7e91cab", pairsMatchedLoops[0][0].LoopId.String())
	}
}
