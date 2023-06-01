package mysql

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
	"testing"
)

func TestLoops(t *testing.T) {
	expectedLoopId := common.HexToHash("0x81825ea35b5feb16da2045c3789814489ba39b78913c4b6d377d599b91dee4ab")
	expectedPath := []string{
		"0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c",
		"0xea6df8a481e0fcbb5f713916c092eda16f93c731",
		"0x8b4d4e4fb663725e19087d384b1c8c1a10cde9f7",
		"0xd8fb853c205a0a49bc7fe0d5dc912fea9135bf7a",
		"0x55d398326f99059ff775485246999027b3197955",
		"0xbf188c8ca1f4b846fd274ee65d3f3bf41744f4ef",
	}

	conf := rcommon.MysqlConfig{
		DbHost: "176.9.120.196",
		DbPort: "3306",
		DbUser: "root",
		DbPass: "rkdb",
		DbData: "rkdb_eth",
	}
	mysql := NewMysql(conf)
	if mysql == nil {
		t.Fatalf("TestMatchLoops NewMysql failed!")
	}

	loops := mysql.LoadLoops()
	if loops == nil || len(loops) == 0 {
		t.Fatalf("nil LoadLoops")
	}

	found := false
	for _, loop := range loops {
		if loop.LoopId == expectedLoopId {
			if loop.Path[1] == common.HexToAddress(expectedPath[1]) {
				found = true
			}
		}
	}

	if !found {
		t.Fatalf("expected loop not found: %v", "0x81825ea35b5feb16da2045c3789814489ba39b78913c4b6d377d599b91dee4ab")
	}
}

func TestMysql_LoadPairGas(t *testing.T) {
	conf := rcommon.MysqlConfig{
		DbHost: "176.9.120.196",
		DbPort: "3306",
		DbUser: "root",
		DbPass: "rkdb",
		DbData: "rkdb_eth",
	}
	mysql := NewMysql(conf)
	if mysql == nil {
		t.Fatalf("TestMatchLoops NewMysql failed!")
	}

	pairGasMap := mysql.LoadPairGas()
	t.Logf("total pair gas: %v", len(pairGasMap))

	got := pairGasMap["0x0000f7b8776f4df08902db5dabee8dca289c9d77-1"]
	if got != 157653 {
		t.Fatalf("expected: %v, but got: %v", 157653, got)
	}
}

func TestMysql_LoadLoops(t *testing.T) {
	conf := rcommon.MysqlConfig{
		DbHost: "176.9.120.196",
		DbPort: "3306",
		DbUser: "root",
		DbPass: "rkdb",
		DbData: "rkdb_eth",
	}
	mysql := NewMysql(conf)
	if mysql == nil {
		t.Fatalf("TestMatchLoops NewMysql failed!")
	}

	loops := mysql.LoadLoops()
	cancelCount := 0
	for _, dbLoop := range loops {
		if dbLoop.Cancel {
			cancelCount++
		}
	}
	t.Logf("total loops: %v canceled loops: %v", len(loops), cancelCount)
	if len(loops) < 100000 {
		t.Fatalf("expected: >100000, but got: %v", len(loops))
	}
}

func TestMysql_LoadPairsInfo(t *testing.T) {
	conf := rcommon.MysqlConfig{
		DbHost: "176.9.120.196",
		DbPort: "3306",
		DbUser: "root",
		DbPass: "rkdb",
		DbData: "rkdb_eth",
	}
	mysql := NewMysql(conf)
	if mysql == nil {
		t.Fatalf("TestMysql_LoadPairsInfo NewMysql failed!")
	}

	infos := mysql.LoadPairsInfo()
	t.Logf("total length: %v, first element: %v", len(infos), infos[0])
	if len(infos) < 10000 {
		t.Fatalf("expected: >infos, but got: %v", len(infos))
	}
}

func TestMysql_LoadPoolsInfo(t *testing.T) {
	conf := rcommon.MysqlConfig{
		DbHost: "176.9.120.196",
		DbPort: "3306",
		DbUser: "root",
		DbPass: "rkdb",
		DbData: "rkdb_eth",
	}
	mysql := NewMysql(conf)
	if mysql == nil {
		t.Fatalf("TestMysql_LoadPoolsInfo NewMysql failed!")
	}

	infos := mysql.LoadPoolsInfo()
	if len(infos) > 0 {
		t.Logf("first element: %v", infos[0])
	}
	t.Logf("total length: %v, first element: %v", len(infos), infos[0])
	if len(infos) < 10000 {
		t.Fatalf("expected: >infos, but got: %v", len(infos))
	}
}

func TestMysql_LoadTokensInfo(t *testing.T) {
	conf := rcommon.MysqlConfig{
		DbHost: "176.9.120.196",
		DbPort: "3306",
		DbUser: "root",
		DbPass: "rkdb",
		DbData: "rkdb_eth",
	}
	mysql := NewMysql(conf)
	if mysql == nil {
		t.Fatalf("TestMysql_LoadTokensInfo NewMysql failed!")
	}

	infos := mysql.LoadTokensInfo()
	t.Logf("first element: %v", infos[0])
	if len(infos) < 3000 {
		t.Fatalf("expected: >infos, but got: %v", len(infos))
	}
}

func TestMysql_UpdatePairGas(t *testing.T) {
	conf := rcommon.MysqlConfig{
		DbHost: "176.9.120.196",
		DbPort: "3306",
		DbUser: "root",
		DbPass: "rkdb",
		DbData: "rkdb_eth",
	}
	mysql := NewMysql(conf)
	if mysql == nil {
		t.Fatalf("TestMysql_UpdatePairGas NewMysql failed!")
	}

	pairDirGasMap := map[string]uint64{
		"0x00000c63ce9ab932fee29362845862242e7eb4a3-0": 126278,
		"0x00000c63ce9ab932fee29362845862242e7eb4a3-1": 126278,
	}

	lastId := mysql.UpdatePairGas(pairDirGasMap)
	if lastId == -1 {
		t.Fatalf("update pair gas failed")
	}

	t.Logf("lastId: %v", lastId)
}

func TestMysql_InsertLoop(t *testing.T) {
	dbLoop := &DBLoop{
		Key:    "bb4cdb9c73bc095c-3988f98b720d9f89-c10358f0534592ac-e30da8e8d9e344fd",
		LoopId: common.HexToHash("0xac7c6e7f26b17dc185a6e27c1bc133e0a6122c73279322d6f7898c00c4d5ddb9"),
		Path: []common.Address{
			common.HexToAddress("0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c"),
			common.HexToAddress("0x3988f98bf5cfdc41ef34a670cdbfdd24720d9f89"),
			common.HexToAddress("0xc10358f062663448a3489fc258139944534592ac"),
			common.HexToAddress("0xe30da8e84b561f72acc4b21f136c9025d9e344fd"),
		},
		PoolFee:   []uint64{25, 20},
		TokenFee:  []uint64{0, 0},
		Direction: []uint64{0, 1},
		Index:     []uint64{8, 8},
		Count:     1,
	}

	conf := rcommon.MysqlConfig{
		DbHost: "176.9.120.196",
		DbPort: "3306",
		DbUser: "root",
		DbPass: "rkdb",
		DbData: "rkdb_eth",
	}
	mysql := NewMysql(conf)
	if mysql == nil {
		t.Fatalf("TestMysql_UpdatePairGas NewMysql failed!")
	}

	mysql.InsertLoop(dbLoop.ToLoopRecord())
}

func TestMysql_InsertPairsInfo(t *testing.T) {
	pRecords := []*PairInfoRecord{
		&PairInfoRecord{
			"0xea6df8a481e0fcbb5f713916c092eda16f93c731",
			"Pancake LPs",
			8,
			false,
			"0x8b4d4e4fb663725e19087d384b1c8c1a10cde9f7",
			"0x8b4d4e4fb663725e19087d384b1c8c1a10cde9f7",
			"0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c",
			"0xca143ce32fe78f1f7019d7d551a6402fc5350c73",
		},
		&PairInfoRecord{
			"0xd8fb853c205a0a49bc7fe0d5dc912fea9135bf7a",
			"Pancake LPs",
			8,
			false,
			"0x8b4d4e4fb663725e19087d384b1c8c1a10cde9f7",
			"0x55d398326f99059ff775485246999027b3197955",
			"0x8b4d4e4fb663725e19087d384b1c8c1a10cde9f7",
			"0xca143ce32fe78f1f7019d7d551a6402fc5350c73",
		},
	}

	conf := rcommon.MysqlConfig{
		DbHost: "176.9.120.196",
		DbPort: "3306",
		DbUser: "root",
		DbPass: "rkdb",
		DbData: "rkdb_eth",
	}
	mysql := NewMysql(conf)
	if mysql == nil {
		t.Fatalf("TestMysql_InsertPairsInfo NewMysql failed!")
	}

	res := mysql.InsertPairsInfo(pRecords)
	if res == -1 {
		t.Fatalf("TestMysql_InsertPairsInfo failed!")
	}
}

func TestMysql_InsertPoolsInfo(t *testing.T) {
	pRecords := []*PoolInfoRecord{
		&PoolInfoRecord{
			"0x36696169c63e42cd08ce11f5deebbcebae652050",
			"WBNB-USDT V3, 500",
			"0x55d398326f99059ff775485246999027b3197955",
			"0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c",
			500,
			10,
		},
	}

	conf := rcommon.MysqlConfig{
		DbHost: "176.9.120.196",
		DbPort: "3306",
		DbUser: "root",
		DbPass: "rkdb",
		DbData: "rkdb_eth",
	}
	mysql := NewMysql(conf)
	if mysql == nil {
		t.Fatalf("TestMysql_InsertPoolsInfo NewMysql failed!")
	}

	res := mysql.InsertPoolsInfo(pRecords)
	if res == -1 {
		t.Fatalf("TestMysql_InsertPoolsInfo failed!")
	}
}

func TestMysql_InsertTokensInfo(t *testing.T) {
	tRecords := []*TokenInfoRecord{
		&TokenInfoRecord{
			"0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c",
			"WBNB",
			18,
		},
		&TokenInfoRecord{
			"0x8b4d4e4fb663725e19087d384b1c8c1a10cde9f7",
			"ZYC",
			18,
		},
	}

	conf := rcommon.MysqlConfig{
		DbHost: "176.9.120.196",
		DbPort: "3306",
		DbUser: "root",
		DbPass: "rkdb",
		DbData: "rkdb_eth",
	}
	mysql := NewMysql(conf)
	if mysql == nil {
		t.Fatalf("TestMysql_UpdatePairGas NewMysql failed!")
	}

	res := mysql.InsertTokensInfo(tRecords)
	if res == -1 {
		t.Fatalf("TestMysql_InsertTokensInfo failed!")
	}
}

func TestMysql_LoopRecordToDBLoop(t *testing.T) {
	var record LoopRecord
	jsonData := "{\"loopsId\":\"0x907e2367a24c6bbee78c96da7d6b9e23ed5e32d14c1dea4ce57292a73b1eed1f\",\"path\":\"[\\\"0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c\\\", \\\"0x16b9a82891338f9bA80E2D6970FddA79D1eb0daE\\\", \\\"0x55d398326f99059fF775485246999027B3197955\\\", \\\"0x5F75fc9A9082c770Db0899885A55f676eac6784a\\\", \\\"0xEad44e4f49b697ab98b15a4859c74E2846A9bFFD\\\", \\\"0xdD66694eCdf1dD86751e8003a21Bc5c41F2742D9\\\", \\\"0x3Ef3f6Cf9f867A33C4109208a6bd043A9Ef0E091\\\", \\\"0x78A3E54dEA17C04B396AB81F88Db185719a9A9Be\\\", \\\"0x55d398326f99059fF775485246999027B3197955\\\", \\\"0x16b9a82891338f9bA80E2D6970FddA79D1eb0daE\\\"]\",\"poolFee\":\"[25, 25, 25, 25, 25]\",\"tokenFee\":\"[0, 0, 0, 0, 0]\",\"direction\":\"[1, 0, 1, 0, 0]\",\"index\":\"[8, 8, 8, 8, 8]\",\"counts\":1, \"hasV3\":0}"
	json.Unmarshal([]byte(jsonData), &record)

	record.ToDBLoop()
	t.Logf("done!")
}
