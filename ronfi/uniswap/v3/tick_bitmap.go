package v3

import (
	"github.com/ethereum/go-ethereum/log"
	"math/big"
)

type TickBitMap map[int16]*big.Int

func (tb TickBitMap) min() int {
	var min int
	for k := range tb {
		if min == 0 || k < int16(min) {
			min = int(k)
		}
	}

	return min
}

func (tb TickBitMap) max() int {
	var max int
	for k := range tb {
		if k > int16(max) {
			max = int(k)
		}
	}

	return max
}

func flipTick(tickBitMap TickBitMap, tick, tickSpacing int) {
	if tick%tickSpacing != 0 {
		log.Warn("RonFi flipTick Tick is not correctly spaced!")
		return
	}

	wordPos, bitPos := position(tick / tickSpacing)

	if val, ok := tickBitMap[wordPos]; !ok {
		log.Warn("RonFi flipTick Called flipTick on missing word")
		return
	} else {
		mask := new(big.Int).Lsh(big.NewInt(1), uint(bitPos))
		val.Xor(val, mask)
		tickBitMap[wordPos] = val
	}
}

func nextInitializedTickWithinOneWord(tickBitMap TickBitMap, tick, tickSpacing int, lte bool) (int16, int, bool, bool) {
	var (
		nextTick          int
		initializedStatus bool
		masked            *big.Int
	)

	compressed := tick / tickSpacing
	if tick < 0 && tick%tickSpacing != 0 {
		compressed--
	}

	if lte {
		wordPos, bitPos := position(compressed)
		mask := new(big.Int).Add(
			new(big.Int).Sub(
				new(big.Int).Lsh(big.NewInt(1), uint(bitPos)),
				big.NewInt(1),
			),
			new(big.Int).Lsh(big.NewInt(1), uint(bitPos)))

		bitmapWord, exist := tickBitMap[wordPos]
		if !exist {
			//log.Warn("nextInitializedTickWithinOneWord Called on missing word",
			//	"wordPos", wordPos, "bitPos", bitPos, "compressed", compressed, "tick", tick, "tickSpacing", tickSpacing, "lte", lte)
			return wordPos, nextTick, initializedStatus, false
		}

		masked = new(big.Int).And(bitmapWord, mask)
		if masked.BitLen() > 0 {
			initializedStatus = true
		}

		if initializedStatus {
			nextTick = (compressed - (int(bitPos) - mostSignificantBit(masked))) * tickSpacing
		} else {
			nextTick = (compressed - int(bitPos)) * tickSpacing
		}
		//log.Info("nextInitializedTickWithinOneWord", "tick", tick, "nextTick", nextTick, "compressed", compressed, "bitmap", fmt.Sprintf("%x", bitmapWord))
	} else {
		wordPos, bitPos := position(compressed + 1)
		mask := new(big.Int).Not(new(big.Int).Sub(
			new(big.Int).Lsh(big.NewInt(1), uint(bitPos)),
			big.NewInt(1),
		))

		bitmapWord, exist := tickBitMap[wordPos]
		if !exist {
			//log.Warn("nextInitializedTickWithinOneWord Called on missing word",
			//	"wordPos", wordPos, "bitPos", bitPos, "compressed", compressed, "tick", tick, "tickSpacing", tickSpacing, "lte", lte)
			return wordPos, nextTick, initializedStatus, false
		}

		masked = new(big.Int).And(bitmapWord, mask)
		if masked.BitLen() > 0 {
			initializedStatus = true
		}

		if initializedStatus {
			nextTick = (compressed + 1 + (leastSignificantBit(masked) - int(bitPos))) * tickSpacing
		} else {
			nextTick = (compressed + 1 + int(0xFF-bitPos)) * tickSpacing
		}
		//log.Info("nextInitializedTickWithinOneWord", "tick", tick, "nextTick", nextTick, "compressed", compressed, "bitmap", fmt.Sprintf("%x", bitmapWord))
	}

	return 0, nextTick, initializedStatus, true
}

func position(tick int) (int16, uint8) {
	wordPos := int16(tick >> 8)
	bitPos := uint8(tick%256) & 0xff

	return wordPos, bitPos
}
