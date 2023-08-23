package defi

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
	"github.com/ethereum/go-ethereum/ronfi/db"
	"math/big"
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
				CanFlashLoan: true,
				KeyToken:     common.HexToAddress("0xd87069e48Bec111639448E80f8D5A34e27129F47"),
				Token0:       common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
				Token1:       common.HexToAddress("0xd87069e48Bec111639448E80f8D5A34e27129F47"),
				Factory:      common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"),
			},
		},
		{
			pairType:    "PolkaBridge AMM",
			pairAddress: common.HexToAddress("0xdda0A346D267a48EC74CC68979584d85501fe5D5"),
			expected: PairInfo{
				Name:         "PolkaBridge AMM",
				Index:        0x9,
				BothBriToken: true,
				CanFlashLoan: false,
				KeyToken:     common.HexToAddress("0x0000000000000000000000000000000000000000"),
				Token0:       common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
				Token1:       common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"),
				Factory:      common.HexToAddress("0xf5d8E98ef1f66f1cAb949A2555837FE755fc2D68"),
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
		t.Fatalf("TestInfo_GetPairInfo NewMysql failed!")
	}

	config := &params.ChainConfig{
		ChainID: big.NewInt(1),
	}
	signer := types.MakeSigner(config, big.NewInt(17034870), 1681266455)
	info := NewInfo(client, dbInst, signer)

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
	config := &params.ChainConfig{
		ChainID: big.NewInt(1),
	}
	signer := types.MakeSigner(config, big.NewInt(17034870), 1681266455)
	info := NewInfo(client, dbInst, signer)

	resv := info.GetPairReserves(common.HexToAddress("0xB4e16d0168e52d35CaCD2c6185b44281Ec28C9Dc"))
	if resv == nil {
		t.Fatalf("TestInfo_GetPairReserves getPairReserves failed!")
	} else {
		t.Logf("TestInfo_GetPairReserves reserve0: %v, reserve1: %v, timeStamp: %v", resv.Reserve0, resv.Reserve1, resv.Timestamp)
	}
}

func TestInfo_CheckIfSandwichAttack(t *testing.T) {
	client, err := ethclient.Dial("https://nd-804-879-862.p2pify.com/ca0df7232f6a54347593373cfbf94df8")
	if err != nil {
		t.Fatal("TestInfo_CheckIfSandwichAttack dial eth client failed!")
	}

	defer func() {
		if client != nil {
			client.Close()
		}
	}()

	InitRonFiOracle(client)

	conf := rcommon.MysqlConfig{
		DbHost: "176.9.120.196",
		DbPort: "3306",
		DbUser: "root",
		DbPass: "rkdb",
		DbData: "rkdb_eth",
	}
	dbInst := db.NewMysql(conf)
	if dbInst == nil {
		t.Fatalf("TestInfo_CheckIfSandwichAttack NewMysql failed!")
	}

	config := &params.ChainConfig{
		ChainID: big.NewInt(1),
	}
	signer := types.MakeSigner(config, big.NewInt(17034870), 1681266455)
	info := NewInfo(client, dbInst, signer)

	// case 1: normal swap in aleg/bleg
	//a: 0x88d834c48c5116250e84c1f2257a108d827fd0106fc57d1e139a5eae53833bf1
	//t: 0x9eca86325196b80097401d5417ecfee2b6acd4aee04d895993f61587c91e6d7c
	//b: 0x9b64b58b5b16bafc16b5d8029fd63432136c73fd6d5526052209ee304d417d84

	// case 2: loop in aleg/bleg
	// a: 0xc36a5c1ae3c869e6c5dda5b0f9cf949c8c2bfacf73f3a1fcf3df3fec77fd54ba
	// t: 0x34549b17a9e3e4f9ffb0021f45d3ddfa0f5e6d8499fce954778a84e5615a1b11
	// b: 0xe659554ef79d1b13aa99e5af830437c45b3f2d89f134b2cb6c55eb843ab6f912
	aLegTxHash := common.HexToHash("0x59f0dd5939a2682bd0664610c9693d5fd0ea33bc06f5adbc76d7081dfe9c7bde")
	if aLegTx, _, err := client.TransactionByHash(context.Background(), aLegTxHash); err == nil {
		if aLegReceipt, err := client.TransactionReceipt(context.Background(), aLegTxHash); err == nil {
			aLeg := &TxAndReceipt{
				Tx:      aLegTx,
				Receipt: aLegReceipt,
			}

			targetTxhash := common.HexToHash("0xc03dfc0d55b64cd554947c2a42ae0b60076eeb9677bf6e5d83e19853e5772f85")
			if targetTx, _, err := client.TransactionByHash(context.Background(), targetTxhash); err == nil {
				if targetReceipt, err := client.TransactionReceipt(context.Background(), targetTxhash); err == nil {
					target := &TxAndReceipt{
						Tx:      targetTx,
						Receipt: targetReceipt,
					}

					bLegTxhash := common.HexToHash("0xf2162a71f51b002631cd53a9e58d8e071a4f6a6ae97184d457e98547b8ad0928")
					if bLegTx, _, err := client.TransactionByHash(context.Background(), bLegTxhash); err == nil {
						if bLegReceipt, err := client.TransactionReceipt(context.Background(), bLegTxhash); err == nil {
							bLeg := &TxAndReceipt{
								Tx:      bLegTx,
								Receipt: bLegReceipt,
							}

							if attacker, profit, netProfit, ok := info.CheckIfSandwichAttack(aLeg, target, bLeg); ok {
								t.Logf("TestInfo_CheckIfSandwichAttack attacker: %v, profit: %v, netProfit: %v", attacker, profit, netProfit)
							} else {
								t.Fatal("TestInfo_CheckIfSandwichAttack CheckIfSandwichAttack failed!")
							}
						}
					}
				}
			}
		}
	}
}
