package defi

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
	"github.com/ethereum/go-ethereum/ronfi/db"
	"testing"
)

// https://nd-814-711-835.p2pify.com/049f42c7290c310495b6940701e2ae14
func TestInfo_GetPairInfo(t *testing.T) {
	type test struct {
		pairType    string
		pairAddress common.Address
		expected    PairInfo
	}

	tests := []test{
		{
			pairType:    "Uniswap V2",
			pairAddress: common.HexToAddress("0x63CFFc0A9E2648DeeEb9df98401B737BBd338325"),
			expected: PairInfo{
				Name:         "Uniswap V2",
				Index:        0x8,
				BothBriToken: false,
				KeyToken:     common.HexToAddress("0xd87069e48Bec111639448E80f8D5A34e27129F47"),
				Token0:       common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
				Token1:       common.HexToAddress("0xd87069e48Bec111639448E80f8D5A34e27129F47"),
				Factory:      common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"),
			},
		},
	}

	client, err := ethclient.Dial("https://nd-804-879-862.p2pify.com/ca0df7232f6a54347593373cfbf94df8")
	if err != nil {
		t.Fatal("TestInfo_GetPairInfo dial eth client failed!")
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
		DbData: "rkdb_eth",
	}
	dbInst := db.NewMysql(conf)
	if dbInst == nil {
		t.Fatalf("TestInfo_CheckIfObs NewMysql failed!")
	}
	info := NewInfo(client, dbInst)

	for _, tc := range tests {
		got := info.GetPairInfo(tc.pairAddress)
		if got == nil {
			t.Fatalf("TestInfo_GetPairInfo pairType: %v getPairInfo failed!", tc.pairType)
		} else {
			if !(*got).Equals(tc.expected) {
				t.Fatalf("pairType: %v, excepted:%v, got:%v", tc.pairType, tc.expected, got)
			}
		}
	}
}

func TestInfo_GetPairReserves(t *testing.T) {
	client, err := ethclient.Dial("https://nd-804-879-862.p2pify.com/ca0df7232f6a54347593373cfbf94df8")
	if err != nil {
		t.Fatal("TestInfo_GetPairReserves dial eth client failed!")
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
		DbData: "rkdb_eth",
	}
	dbInst := db.NewMysql(conf)
	if dbInst == nil {
		t.Fatalf("TestInfo_CheckIfObs NewMysql failed!")
	}
	info := NewInfo(client, dbInst)

	info.GetPairReserves(common.HexToAddress("0x63CFFc0A9E2648DeeEb9df98401B737BBd338325"))
}
