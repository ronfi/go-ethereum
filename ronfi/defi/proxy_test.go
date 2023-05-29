package defi

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"testing"
)

func TestProxy(t *testing.T) {
	type test struct {
		proxy        string
		proxyAddress common.Address
		expected     common.Address
	}

	tests := []test{
		{
			proxy:        "Eip1167 Minimal",
			proxyAddress: common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
			expected:     common.HexToAddress("0xa2327a938Febf5FEC13baCFb16Ae10EcBc4cbDCF"),
		},
		{
			proxy:        "Eip1967Logic",
			proxyAddress: common.HexToAddress("0x6b4eE7352406707003bC6f6b96595FD35925af48"),
			expected:     common.HexToAddress("0xB0187178095e751E817dDC2e8EBe3fb1F7E46842"),
		},
		{
			proxy:        "Eip1967Beacon",
			proxyAddress: common.HexToAddress("0xcEb49cB8E765E930B7290Ac42E8d26a99EbAa4Ce"),
			expected:     common.HexToAddress("0x4A5aD53ed70357961e58Faf304CCfd06180D2C30"),
		},
		{
			proxy:        "OpenZeppelin",
			proxyAddress: common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
			expected:     common.HexToAddress("0xa2327a938Febf5FEC13baCFb16Ae10EcBc4cbDCF"),
		},
		{
			proxy:        "Eip897",
			proxyAddress: common.HexToAddress("0xA8E03D0E75FA9a0bBa097266FC8e14284118e0EF"),
			expected:     common.HexToAddress("0x164Ac16e514f4e47d201aFF991740E166d409b8A"),
		},
		{
			proxy:        "Eip1822",
			proxyAddress: common.HexToAddress("0x261b45D85cCFeAbb11F022eBa346ee8D1cd488c0"),
			expected:     common.HexToAddress("0xf4dD399f6584B4FFd925C86782025A4282429D7c"),
		},
		{
			proxy:        "GnosisSafeProxy",
			proxyAddress: common.HexToAddress("0x526254C12360D68D8F90c0cD9cFE5a32002e845A"),
			expected:     common.HexToAddress("0xd9Db270c1B5E3Bd161E8c8503c55cEABeE709552"),
		},
	}

	client, err := ethclient.Dial("https://nd-804-879-862.p2pify.com/ca0df7232f6a54347593373cfbf94df8")
	if err != nil {
		t.Fatal("TestProxy dial eth client failed!")
	}

	defer func() {
		if client != nil {
			client.Close()
		}
	}()

	p := NewProxy(client)

	for _, tc := range tests {
		if got, ok := p.detectProxyTarget(tc.proxyAddress); ok {
			if got != tc.expected {
				t.Errorf("proxy: %v, excepted:%v, got:%v", tc.proxy, tc.expected, got)
			}
		} else {
			t.Errorf("proxy: %v, detect real address failed!", tc.proxy)
		}
	}
}
