package trading

import (
	"github.com/ethereum/go-ethereum/params"
	"math/big"
	"time"
)

const (
	DexTxInQueueSkipMax   = 128 //enlarge skip queue max
	DexTxInQueueFilterMin = DexTxInQueueSkipMax / 16

	maxWalletNonceKnownWallet = 128 // for each shard, the total is this multiply 256 (MapsetShardCount), i.e. maximum 32K SenderNonce reusing can be detected
	maxStateDbsCopies         = 32
	ChanReadTimeout           = 1 * time.Second
)

var (
	GasPriceMaxAllowed     = big.NewInt(0).Mul(big.NewInt(75), big.NewInt(params.GWei))           // 50 Gwei will be an amazing price
	GasPriceMinAllowedIdle = big.NewInt(0).Mul(big.NewInt(1), big.NewInt(params.GWei))            // 1 Gwei is the minimum required price
	GasPriceMinAllowedBusy = big.NewInt(0).Mul(big.NewInt(5_010_000_000), big.NewInt(params.Wei)) // 5.01 Gwei is the minimum required price in case the whole blockchain is busy
	GasPriceMaxAllowedCopy = big.NewInt(0).Mul(big.NewInt(50), big.NewInt(params.GWei))           // 10 Gwei is the maximum allowed price for arb tx copy

	GasLimitDefault = uint64(500000)

	falshbotURLs = []string{
		"https://relay.flashbots.net",
		"https://rpc.titanbuilder.xyz",
		"https://builder0x69.io",
		"https://rpc.beaverbuild.org",
		"https://rsync-builder.xyz",
		"https://api.blocknative.com/v1/auction",
		// "https://mev.api.blxrbdn.com", # Authentication required
		"https://eth-builder.com",
		"https://builder.gmbit.co/rpc",
		"https://buildai.net",
		"https://rpc.payload.de",
		"https://rpc.lightspeedbuilder.info",
		"https://rpc.nfactorial.xyz",
	}
)
