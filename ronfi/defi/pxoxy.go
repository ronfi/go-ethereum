package defi

import (
	"bytes"
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
	"sync"
)

var (
	Eip1167BytecodePrefix = []byte{0x36, 0x3d, 0x3d, 0x37, 0x3d, 0x3d, 0x3d, 0x36, 0x3d}
	Eip1167BytecodeSuffix = []byte{0x57, 0xfd, 0x5b, 0xf3}

	ProxyBySlots = []common.Hash{
		// Eip1967LogicSlot obtained as bytes32(uint256(keccak256('eip1967.proxy.implementation')) - 1)
		common.HexToHash("0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc"),
		// Eip1967BeaconSlot obtained as bytes32(uint256(keccak256('eip1967.proxy.beacon')) - 1)
		common.HexToHash("0xa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d50"),
		//OpenZeppelinSlot obtained as keccak256("org.zeppelinos.proxy.implementation")
		common.HexToHash("0x7050c9e0f4ca769c69bd3a8ef740bc37934f8e2c036e5a723fd8ee048ed3f8c3"),
		// Eip1822LogicSlot obtained as keccak256("PROXIABLE")
		common.HexToHash("0xc5f16f0fcc639fa48a6947836d9850f504798523bf8c9a3a87d5876cf622bcf7"),
	}

	Eip1167BeaconMethods = [][]byte{
		// bytes4(keccak256("implementation()")) padded to 32 bytes
		[]byte{0x5c, 0x60, 0xda, 0x1b},
		// bytes4(keccak256("childImplementation()")) padded to 32 bytes
		// some implementations use this over the standard method name so that the beacon contract is not detected as an EIP-897 proxy itself
		[]byte{0xda, 0x52, 0x57, 0x16},
	}

	ProxyByCalls = [][]byte{
		// Eip897Interface bytes4(keccak256("implementation()")) padded to 32 bytes
		[]byte{0x5c, 0x60, 0xda, 0x1b},
		// GnosisSafeProxyInterface bytes4(keccak256("masterCopy()")) padded to 32 bytes
		[]byte{0xa6, 0x19, 0x48, 0x6e},
	}
)

type Proxy struct {
	client *ethclient.Client

	proxyMap map[common.Address]common.Address
	lock     sync.RWMutex
}

func eip1167Detect(code []byte) (target common.Address, ok bool) {
	if len(code) < len(Eip1167BytecodePrefix) || len(code) < len(Eip1167BytecodeSuffix) {
		return
	}

	prefix := code[0:len(Eip1167BytecodePrefix)]
	suffix := code[len(code)-len(Eip1167BytecodeSuffix):]
	if !bytes.Equal(prefix, Eip1167BytecodePrefix) || !bytes.Equal(suffix, Eip1167BytecodeSuffix) {
		return
	}

	//// detect length of address (20 bytes non-optimized, 0 < N < 20 bytes for vanity addresses)
	pushNHex := code[len(Eip1167BytecodePrefix)]
	//// push1 ... push20 use opcodes 0x60 ... 0x73
	addressLength := int(pushNHex) - 0x5f

	if addressLength < 1 || addressLength > 20 {
		return
	}

	target = common.BytesToAddress(code[len(Eip1167BytecodePrefix)+1 : len(Eip1167BytecodePrefix)+1+addressLength])
	ok = true

	return
}

func getAddressByCall(client *ethclient.Client, msg ethereum.CallMsg) (common.Address, bool) {
	if val, err := client.CallContract(context.Background(), msg, nil); err == nil {
		target := common.BytesToAddress(val)
		if target != rcommon.ZeroAddress {
			return target, true
		}
	}

	return common.Address{}, false
}

func NewProxy(client *ethclient.Client) *Proxy {
	return &Proxy{
		client:   client,
		proxyMap: make(map[common.Address]common.Address),
		lock:     sync.RWMutex{},
	}
}

func (p Proxy) detectProxyTarget(proxyAddress common.Address) (target common.Address, ok bool) {
	p.lock.RLock()
	if target, ok = p.proxyMap[proxyAddress]; ok {
		p.lock.RUnlock()
		return
	}
	p.lock.RUnlock()

	defer func() {
		if target != rcommon.ZeroAddress {
			p.lock.Lock()
			if _, exist := p.proxyMap[proxyAddress]; !exist {
				p.proxyMap[proxyAddress] = target
			}
			p.lock.Unlock()
		}
	}()

	// eip1167
	if code, err := p.client.CodeAt(context.Background(), proxyAddress, nil); err == nil {
		if target, ok = eip1167Detect(code); ok {
			return
		}
	}

	// Proxy By Slots
	for i, hash := range ProxyBySlots {
		if val, err := p.client.StorageAt(context.Background(), proxyAddress, hash, nil); err == nil {
			target = common.BytesToAddress(val)
			if target != rcommon.ZeroAddress {
				if i != 1 {
					ok = true
					return
				} else {
					// eip1967 beacon
					for _, method := range Eip1167BeaconMethods {
						msg := ethereum.CallMsg{To: &target, Data: method}
						if target, ok = getAddressByCall(p.client, msg); ok && target != rcommon.ZeroAddress {
							return
						}
					}
				}
			}
		}
	}

	// Proxy By Calls
	for _, method := range ProxyByCalls {
		msg := ethereum.CallMsg{To: &proxyAddress, Data: method}
		if target, ok = getAddressByCall(p.client, msg); ok && target != rcommon.ZeroAddress {
			return
		}
	}

	return
}
