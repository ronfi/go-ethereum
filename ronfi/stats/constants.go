package stats

import "github.com/ethereum/go-ethereum/common"

const (
	MethodRonSwapX = 0x6326d217 // todo change to real method id

	TxGasUsedWarn = 95_000

	KncMinProfit         = 100.0
	KncObsDexMaxIndexGap = 5

	Ron  = "0"
	Obs1 = "1"
	Obs2 = "2"
	Obs3 = "3"
	Obs4 = "4"
	Obs5 = "5"
	Obs6 = "6"
	Obs7 = "7"
	Obs8 = "8"
	Obs9 = "9"
	Obsx = "x"
)

var (
	Obs1SwapAddr = common.HexToAddress("0xc11ce44147c9f6149fbe54adb0588523c38718d7")
	Obs2SwapAddr = common.HexToAddress("0xdB8D95f9a3180e00507c9745eEFc9287d553b481")
	Obs3SwapAddr = common.HexToAddress("0x2d2A7d56773ae7d5c7b9f1B57f7Be05039447B4D")
	Obs4SwapAddr = common.HexToAddress("0x0000E0Ca771e21bD00057F54A68C30D400000000")
	Obs5SwapAddr = common.HexToAddress("0x7efd91c8ca31e7452c5c4de00a8ca1c18910a35f")

	Obs5Method1 = uint32(0x007cca5a)
	Obs6Method  = uint32(0x0a000000)
)
