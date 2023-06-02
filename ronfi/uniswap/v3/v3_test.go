package v3

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
	"github.com/ethereum/go-ethereum/ronfi/defi"
	"github.com/ethereum/go-ethereum/ronfi/mysql"
	common2 "github.com/ethereum/go-ethereum/ronfi/uniswap/common"
	"math/big"
	"testing"
)

func TestV3_getSqrtRatioAtTick(t *testing.T) {
	type test struct {
		name string
		tick int
		want *big.Int
	}

	tests := []test{
		{name: "-12345", tick: -12345, want: common2.StrToBigInt("42739035517269358503607398648")},
		{name: "0", tick: 0, want: common2.StrToBigInt("79228162514264337593543950336")},
		{name: "12345", tick: 12345, want: common2.StrToBigInt("146870458338965608271414022015")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSqrtRatioAtTick(tt.tick); got.Cmp(tt.want) != 0 {
				t.Errorf("getSqrtRatioAtTick() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV3_getTickAtSqrtRatio(t *testing.T) {
	type test struct {
		name string
		r    *big.Int
		want int
	}

	tests := []test{
		{name: "79228162514264337593543950336", r: common2.StrToBigInt("79228162514264337593543950336"), want: 0},
		{name: "146870458338965608271414022015", r: common2.StrToBigInt("146870458338965608271414022015"), want: 12345},
		{name: "42739035517269358503607398648", r: common2.StrToBigInt("42739035517269358503607398648"), want: -12345},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, ok := getTickAtSqrtRatio(tt.r); !ok || got != tt.want {
				t.Errorf("getTickAtSqrtRatio() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV3_mostSignificantBit(t *testing.T) {
	type test struct {
		name string
		v    *big.Int
		want int
	}

	tests := []test{
		{name: "0", v: big.NewInt(0), want: -1},
		{name: "1", v: common2.HexToBigInt("1"), want: 0},
		{name: "1000000000", v: common2.HexToBigInt("1000000000"), want: 36},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostSignificantBit(tt.v); got != tt.want {
				t.Errorf("mostSignificantBit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV3_leastSignificantBit(t *testing.T) {
	type test struct {
		name string
		v    *big.Int
		want int
	}

	tests := []test{
		{name: "0", v: big.NewInt(0), want: -1},
		{name: "1", v: common2.HexToBigInt("1"), want: 0},
		{name: "1000000000", v: common2.HexToBigInt("1000000010"), want: 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leastSignificantBit(tt.v); got != tt.want {
				t.Errorf("leastSignificantBit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV3_mulDiv(t *testing.T) {
	type args struct {
		a           *big.Int
		b           *big.Int
		denominator *big.Int
	}
	type test struct {
		name string
		arg  args
		want *big.Int
	}

	tests := []test{
		{
			name: "1",
			arg: args{
				a:           common2.HexToBigInt("1000"),
				b:           common2.HexToBigInt("2000"),
				denominator: common2.HexToBigInt("200"),
			},
			want: common2.HexToBigInt("10000"),
		},
		{
			name: "2",
			arg: args{
				a:           common2.HexToBigInt("100000000000000000"),
				b:           common2.HexToBigInt("20000000000000000000000000"),
				denominator: common2.HexToBigInt("200"),
			},
			want: common2.StrToBigInt("1461501637330902918203684832716283019655932542976"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mulDiv(tt.arg.a, tt.arg.b, tt.arg.denominator); got.Cmp(tt.want) != 0 {
				t.Errorf("mulDiv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV3_Swap(t *testing.T) {
	client, err := ethclient.Dial("https://nd-804-879-862.p2pify.com/049f42c7290c310495b6940701e2ae14")
	if err != nil {
		t.Fatal("TestV3_Swap dial eth client failed!")
	}

	defer func() {
		if client != nil {
			client.Close()
		}
	}()

	if block, err := client.BlockNumber(context.Background()); err == nil {
		t.Logf("block: %v\n", block)
	}

	conf := rcommon.MysqlConfig{
		DbHost: "176.9.120.196",
		DbPort: "3306",
		DbUser: "root",
		DbPass: "rkdb",
		DbData: "rkdb",
	}
	mysqlInst := mysql.NewMysql(conf)
	if mysqlInst == nil {
		t.Fatalf("TestV3_Swap NewMysql failed!")
	}

	di := defi.NewInfo(client, mysqlInst)
	if di == nil {
		t.Fatalf("TestV3_Swap NewInfo failed!")
	}

	poolAddr := common.HexToAddress("0xb4E9DeA6105089f15685508B8EF2e7f7F5A1B16D")
	TickLensAddr := common.HexToAddress("0x9a489505a00cE272eAa5e07Dba6491314CaE3796")

	v3Pool := NewV3Pool(di, poolAddr, TickLensAddr, nil)
	v3Pool.UpdatePoolState(nil)
	amountOut := v3Pool.CalculateTokensOutFromTokensIn(
		common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"),
		common2.StrToBigInt("100000000000000000"),
	)

	t.Logf("amountOut: %v", amountOut)
}
