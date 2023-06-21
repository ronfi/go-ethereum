package stats

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
	db "github.com/ethereum/go-ethereum/ronfi/db"
	"github.com/ethereum/go-ethereum/ronfi/defi"
	"testing"
)

func TestParserObs5Input(t *testing.T) {
	client, err := ethclient.Dial("https://nd-804-879-862.p2pify.com/ca0df7232f6a54347593373cfbf94df8")
	if err != nil {
		t.Fatal("TestParserObs5Input dial eth client failed!")
	}

	defer func() {
		if client != nil {
			client.Close()
		}
	}()

	testLoop := &db.DBLoop{
		Path: []common.Address{
			common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
			common.HexToAddress("0x40643d019bcb381800f0c2d593a4a6472ccd7493"),
			common.HexToAddress("0x2eCBa91da63C29EA80Fbe7b52632CA2d1F8e5Be0"),
			common.HexToAddress("0x253c06b22d4b67b675ea7d2fd48422e62c394d7e"),
		},
		PoolFee:   []uint64{0, 0},
		TokenFee:  []uint64{0, 0},
		Direction: []uint64{1, 4},
		Index:     []uint64{8, 255},
		Count:     1,
	}

	testObsLoop := &ObsParsedResult{
		Path: []common.Address{
			common.HexToAddress("0x40643d019bcb381800f0c2d593a4a6472ccd7493"),
			common.HexToAddress("0x253c06b22d4b67b675ea7d2fd48422e62c394d7e"),
		},
		PoolFee:   []uint64{0, 0},
		TokenFee:  []uint64{0, 0},
		Direction: []uint64{1, 0},
	}

	tx, _, err := client.TransactionByHash(context.Background(), common.HexToHash("0x12b095a340e92a78323c57328077d124e4b474e57a2ac8de00ae7faae67c228d"))
	if err != nil {
		t.Fatalf("RPC Client TransactionByHash Failed!")
	}

	input := tx.Data()

	conf := rcommon.MysqlConfig{
		DbHost: "176.9.120.196",
		DbPort: "3306",
		DbUser: "root",
		DbPass: "rkdb",
		DbData: "rkdb_eth",
	}
	dbInst := db.NewMysql(conf)
	di := defi.NewInfo(client, dbInst)
	if obsLoop, dbLoop, ok := parseObs5Input(di, input); ok {
		if !dbLoop.Equals(testLoop) {
			t.Fatalf("TestParserObs5Input: expected: %v, but got: %v", testLoop, dbLoop)
		}

		if !obsLoop.Equals(testObsLoop) {
			t.Fatalf("TestParserObs5Input failed on obsLoop: expected: %v, but got: %v", testObsLoop, obsLoop)
		}
	} else {
		t.Fatalf("TestParserObs5Input failed on parseObs5")
	}
}

func TestParserObs6Input(t *testing.T) {
	client, err := ethclient.Dial("https://nd-804-879-862.p2pify.com/ca0df7232f6a54347593373cfbf94df8")
	if err != nil {
		t.Fatal("TestParserObs6Input dial eth client failed!")
	}

	defer func() {
		if client != nil {
			client.Close()
		}
	}()

	testLoop := &db.DBLoop{
		Path: []common.Address{
			common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
			common.HexToAddress("0x6ecbfbee5e615051e865970da9361c18628b1aa9"),
			common.HexToAddress("0xbf6693cDc1ba57D2E9a2D0c4bC1FB27C927bb9Ba"),
			common.HexToAddress("0xc04df2479913774d811368316f5aa521151be1a9"),
		},
		PoolFee:   []uint64{30, 0},
		TokenFee:  []uint64{0, 0},
		Direction: []uint64{1, 4},
		Index:     []uint64{8, 255},
		Count:     1,
	}

	testObsLoop := &ObsParsedResult{
		Path: []common.Address{
			common.HexToAddress("0x6ecbfbee5e615051e865970da9361c18628b1aa9"),
			common.HexToAddress("0xc04df2479913774d811368316f5aa521151be1a9"),
		},
		PoolFee:   []uint64{30, 0},
		TokenFee:  []uint64{0, 0},
		Direction: []uint64{1, 0},
	}

	tx, _, err := client.TransactionByHash(context.Background(), common.HexToHash("0xe4132de87fee3a2ad3d3b1575b208d51831d06a14d8070dbaa0078934400212d"))
	if err != nil {
		t.Fatalf("RPC Client TransactionByHash Failed!")
	}

	input := tx.Data()

	conf := rcommon.MysqlConfig{
		DbHost: "176.9.120.196",
		DbPort: "3306",
		DbUser: "root",
		DbPass: "rkdb",
		DbData: "rkdb_eth",
	}
	dbInst := db.NewMysql(conf)
	di := defi.NewInfo(client, dbInst)
	if obsLoop, dbLoop, ok := parseObs6Input(di, input); ok {
		if !dbLoop.Equals(testLoop) {
			t.Fatalf("TestParserObs6Input: expected: %v, but got: %v", testLoop, dbLoop)
		}

		if !obsLoop.Equals(testObsLoop) {
			t.Fatalf("TestParserObs6Input failed on obsLoop: expected: %v, but got: %v", testObsLoop, obsLoop)
		}
	} else {
		t.Fatalf("TestParserObs6Input failed on parseObs6")
	}
}
