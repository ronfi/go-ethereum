package stats

import (
	"encoding/binary"
	"github.com/ethereum/go-ethereum/common"
)

func prefix6HexString(longHexStr string) string {
	result := make([]byte, 0, 6)
	result = append(result, []byte(longHexStr[2:8])...)

	return string(result)
}

func simpleCheckSum(input []byte) uint64 {
	sum := uint64(0)
	pos := 0
	for i := 0; i < len(input)/4; i++ {
		unit := binary.BigEndian.Uint32(input[pos : pos+4])
		sum += uint64(unit)
		pos += 4
	}
	return sum
}

func pathSimpleCheckSum(paths []common.Address) uint64 {
	sum := uint64(0)
	for _, path := range paths {
		sum += simpleCheckSum(path.Bytes())
	}
	return sum
}
