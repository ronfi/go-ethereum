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
			proxy:        "Eip1167",
			proxyAddress: common.HexToAddress("0x716Cd63D7523D120a8Ab5bb97E149Aba5a726C76"),
			expected:     common.HexToAddress("0x850655492DC9e3a347209DD950d952306e19A9C1"),
		},
		{
			proxy:        "Eip1967Logic",
			proxyAddress: common.HexToAddress("0xe29F0B490F0d89CA7ACAc1C7BeD2E07eCAD65201"),
			expected:     common.HexToAddress("0x5713Ae21F4Bb696A877c90CCcAE310eFF4c4652A"),
		},
		{
			proxy:        "Eip1967Beacon",
			proxyAddress: common.HexToAddress("0xEEcc149386caD42784539Fc8449408E7956a3a8e"),
			expected:     common.HexToAddress("0x7F54C5d46fdd02d4db445D6F2dC44E642357fBE1"),
		},
		{
			proxy:        "OpenZeppelin",
			proxyAddress: common.HexToAddress("0x4b136E073D469f240983ACe06e3c5f776FCE45Ea"),
			expected:     common.HexToAddress("0x261701115BFC2e1A3Cf79bd197937f7b3016A427"),
		},
		{
			proxy:        "Eip897",
			proxyAddress: common.HexToAddress("0x99b0d6641A63Ce173E6EB063b3d3AED9A35Cf9bf"),
			expected:     common.HexToAddress("0x6f8A975758436A5Ec38d2f9d2336504430465517"),
		},
		{
			proxy:        "Eip1822",
			proxyAddress: common.HexToAddress("0xe87e15B9c7d989474Cb6d8c56b3DB4eFAD5b21E8"),
			expected:     common.HexToAddress("0xCDB432013B83559C7942925495ABa373e3Eb5815"),
		},
		{
			proxy:        "GnosisSafeProxy",
			proxyAddress: common.HexToAddress("0xccF45C824e26ff19597441B95586b7c7F536719d"),
			expected:     common.HexToAddress("0x2BB001433cf04c1f7d71E3c40FED66b2b563065E"),
		},
	}

	client, err := ethclient.Dial("http://localhost:8545")
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
