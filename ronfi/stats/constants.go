package stats

import "github.com/ethereum/go-ethereum/common"

const (
	MethodRonSwapX  = 0x6326d217
	MethodRonSwapD  = 0x34ec0942
	MethodRonSwapZD = 0x6a610127 // Contract: RonChiSwapZ, Method: ronSwapD, Without CHI
	MethodRonSwapZC = 0xfe5856bb // Contract: RonChiSwapZ, Method: ronSwapAuto, With CHI
	MethodRonExecPL = 0x69bf3560

	TxGasUsedWarn = 95_000

	KncMinProfit         = 100.0
	KncObsDexMaxIndexGap = 5

	Ron  = "0"
	Obs1 = "1"
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
	//ObsASwapAddr1 = common.HexToAddress("0x0000000009B6FEcc7a941111d72f072f13E8aa5E") // 一个实施三明治攻击型套利的合约
	//ObsASwapAddr2 = common.HexToAddress("0x27a0b602e95f7712b4f4ec5e56e47e165a520594") // 一个实施三明治攻击型套利的合约
	//ObsASwapAddr3 = common.HexToAddress("0xabf1ebfec1f4c4d68716453975099db1c43820e3") // 一个实施三明治攻击型套利的合约
	//ObsASwapAddr4 = common.HexToAddress("0xf922ef0c6bddf275bd9ecdca0d1f3dffbce243f4") // 一个实施三明治攻击型套利的合约
	//ObsASwapAddr5 = common.HexToAddress("0x9c6b06f282ae033c25c33b5a720c6a10557b58e4") // 一个实施三明治攻击型套利的合约
	//ObsASwapAddr6 = common.HexToAddress("0xcf2314ff5bc54ba6e3688cda497eed9515b2f1a2") // 一个实施三明治攻击型套利的合约
	//Obs3SwapAddr  = common.HexToAddress("0xa8cC4B94624B31465e777f06ee3591fd5AFe5821")
	//Obs6Method    = uint32(0x9c47d752) // "0xba6d76f1fa25727bc54756fb9a5b79bb05fda737", "0xbbc09C2bBD752d8f489c7A6167F36a44028a69BC", "0x1c84183e211a7e69ee455e6c7d88cebb2af685cd"
	//Obs7SwapAddr  = common.HexToAddress("0x0000000000008AfdAcc486225455281F614843e7")
	//Obs1Method1   = uint32(0x1de9c881)
	//Obs1Method2   = uint32(0xecfa311d)
	//Obs3MethodWD  = uint32(0x8e19f1bf)
	//Obs3MethodWdA = uint32(0x736fe565)
	//Obs4Method1   = uint32(0x444412c5) // "0xfbaebcb3c22ba9acb3a64eb0be82c9a46c56deea" WBNB, "0x2ec5a20f3a23d83f9a1ad1eb57da607f8e4f9e55" USDT
	//Obs4Method2   = uint32(0x4ef2ba16) // "0x3f32d3eaabdcb3a9fe8101e58cf40621b7d43e15" WBNB, "0xbcccf200acaf3bd0d653c05e7787cec8fbabbd58" USDT
	//Obs4Method3   = uint32(0xd3678529) // "0x5e8ce43f99aeb594d29c0a92b3bff04990deb903" USDT, "0x270eb2621eb213a0600a03151af8f9be520e56d4" WBNB
	//Obs4Method4   = uint32(0x2e6b295d) // "0x911c6f41f62ae50277dcd2254753bc832aab904b" WBNB, "0xba5276f63492b351c7227a4f285593cefa250ad3" USDT
	//Obs5Method1   = uint32(0xd7eeee04) // "0xe13ce3d0388032dca4b6cd4059a0bf6ed9dcc005", "0xc582840b5360929DEFb2C4a1578baFD95a5f1c11"
	//Obs8Method    = uint32(0x258b1b69) // "0xb67054f21eef394285e19f76d32e1dba917ce2d7"
	//Obs5Method2   = uint32(0x2049db2c) // "0x60975e212ead6d8cea65e2863d278599bad0a7fa", "0xA88C8d63AC966ba2C384aA906981eC1daa116Ab3", "0x1c3bF265Db354d36A85c98249Af3Bcf53ED5E0f5", "0x4C51dA43594312f74d567AB0034C8CD8F22d4E13"
	//Obs9Method    = uint32(0xc0cac0ca)

	Obs1SwapAddr = common.HexToAddress("0xc11ce44147c9f6149fbe54adb0588523c38718d7")
)
