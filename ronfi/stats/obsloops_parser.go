package stats

import (
	"github.com/ethereum/go-ethereum/common"
)

type ObsParsedResult struct {
	AmountIn  string           `json:"amountIn"`
	Path      []common.Address `json:"path"`
	PoolFee   []uint64         `json:"poolFee"`
	TokenFee  []uint64         `json:"tokenFee"`
	Direction []uint64         `json:"direction"`
	HasV3     bool             `json:"hasV3"`
}

func (loop *ObsParsedResult) Equals(other *ObsParsedResult) bool {
	if len(loop.Path) != len(other.Path) {
		return false
	} else {
		for i := 0; i < len(loop.Path); i++ {
			if loop.Path[i] != other.Path[i] {
				return false
			}
		}
	}

	if len(loop.PoolFee) != len(other.PoolFee) {
		return false
	} else {
		for i := 0; i < len(loop.PoolFee); i++ {
			if loop.PoolFee[i] != other.PoolFee[i] {
				return false
			}
		}
	}

	if len(loop.TokenFee) != len(other.TokenFee) {
		return false
	} else {
		for i := 0; i < len(loop.TokenFee); i++ {
			if loop.TokenFee[i] != other.TokenFee[i] {
				return false
			}
		}
	}

	if len(loop.Direction) != len(other.Direction) {
		return false
	} else {
		for i := 0; i < len(loop.Direction); i++ {
			if loop.Direction[i] != other.Direction[i] {
				return false
			}
		}
	}

	return true
}

//func parseObs1Input(di *defi.Info, tx *types.Transaction, methodId uint32, input []byte) ([]*ObsParsedResult, *mysql.DBLoop, bool) {
//	size := len(input)
//	if methodId == Obs1Method1 {
//		minSize := 4 + 32*7
//		if size < minSize {
//			return nil, nil, false
//		}
//		hops := (size - 32*5 - 4) / 32
//		addrs := make([]common.Address, hops)
//		path := make([]common.Address, hops*2)
//		fees := new(big.Int).SetBytes(input[100:132])
//		poolFee := make([]uint64, hops)
//		tokenFee := make([]uint64, hops) // just filling zero here
//		direction := make([]uint64, hops)
//		index := make([]uint64, hops)
//		amountIn := new(big.Int).SetBytes(input[4:36])
//		if hops > 5 || amountIn.BitLen() == 0 {
//			// must be a mistake
//			return nil, nil, false
//		}
//		pos := 164
//		{
//			remaining := new(big.Int)
//			if remaining.Rsh(fees, uint(1+hops*18)).BitLen() != 0 {
//				//log.Info("RonFi arb obs1 data parse", "tx", tx, "remaining", remaining)
//				fees = new(big.Int).Rsh(fees, 2) // remaining 2 bit unused
//			} else {
//				fees = new(big.Int).Rsh(fees, 1) // remaining 1 bit unused
//			}
//		}
//
//		toDbLoop := true
//		for i := 0; i < hops; i++ {
//			pair := common.BytesToAddress(input[pos+12 : pos+32])
//			pairInfo := di.GetPairInfo(pair)
//			if pairInfo == nil {
//				toDbLoop = false
//			}
//
//			poolFee[i] = new(big.Int).Mod(fees, new(big.Int).SetUint64(65536)).Uint64()
//			fees = new(big.Int).Rsh(fees, 16)               // 16 bit fee
//			dir := 3 - uint64(fees.Bit(0)+(fees.Bit(1)<<1)) // 2 -> 1, 3 -> 0. Even higher bit is unknown usage, we reserve it for the future.
//			direction[i] = dir
//
//			if toDbLoop {
//				index[i] = pairInfo.Index
//				tokenIn := pairInfo.Token0
//				if dir&0x1 == 1 {
//					tokenIn = pairInfo.Token1
//				}
//				path[2*i] = tokenIn
//				path[2*i+1] = pair
//			}
//			addrs[i] = pair
//
//			fees = new(big.Int).Rsh(fees, 2) // 2 bit direction
//			pos += 32
//		}
//		if fees.BitLen() != 0 { // fees field still has something?
//			log.Warn("RonFi arb obs1 data parse fail", "tx", tx.Hash().String(), "feesRemaining", fees)
//			return nil, nil, false
//		}
//
//		var (
//			dbLoop  *mysql.DBLoop
//			obsLoop *ObsParsedResult
//		)
//
//		obsLoop = &ObsParsedResult{
//			AmountIn:  amountIn.String(),
//			Path:      addrs,
//			PoolFee:   poolFee,
//			TokenFee:  tokenFee,
//			Direction: direction,
//		}
//
//		if toDbLoop {
//			dbLoop = &mysql.DBLoop{
//				Kid:       -1,
//				Key:       rcommon.ToKey(path),
//				LoopId:    rcommon.ToLoopId(path, poolFee, tokenFee),
//				Path:      path,
//				PoolFee:   poolFee,
//				TokenFee:  tokenFee,
//				Direction: direction,
//				Index:     index,
//				Count:     1,
//			}
//
//			return []*ObsParsedResult{obsLoop}, dbLoop, true
//		} else {
//			return []*ObsParsedResult{obsLoop}, dbLoop, false
//		}
//	} else if methodId == Obs1Method2 {
//		minSize := 4 + 32*22
//		if size < minSize {
//			return nil, nil, false
//		}
//
//		amountInPos := 4 + 5*32 // 164
//		txsCnt := int(new(big.Int).SetBytes(input[amountInPos : amountInPos+32]).Int64())
//		if txsCnt > 100 || txsCnt < 1 {
//			return nil, nil, false
//		}
//		result := make([]*ObsParsedResult, txsCnt)
//		hopsPos := amountInPos + (txsCnt+1)*32*2
//		hopsCnt := int(new(big.Int).SetBytes(input[hopsPos : hopsPos+32]).Int64())
//		if hopsCnt > 1000 || hopsCnt < 2 {
//			return nil, nil, false
//		}
//
//		hopsPos += 32                             // Now points to 1st pair
//		poolFeePos := hopsPos + (hopsCnt+1)*32    // Now points to pool fees of 1st transaction
//		swapsCntPos := poolFeePos + (txsCnt+1)*32 // Now points to 1st transaction swaps number (2 swaps, 3 swaps for example)
//
//		for i := 0; i < txsCnt; i++ {
//			obsLoop := new(ObsParsedResult)
//			result[i] = obsLoop
//
//			amountInPos += 32
//			amountIn := new(big.Int).SetBytes(input[amountInPos : amountInPos+32])
//
//			swapsCnt := int(new(big.Int).SetBytes(input[swapsCntPos : swapsCntPos+32]).Int64())
//			swapsCntPos += 32
//			if swapsCnt > 8 || swapsCnt < 2 {
//				return nil, nil, false
//			}
//
//			if swapsCnt > 5 || amountIn.BitLen() == 0 {
//				// must be a mistake
//				return nil, nil, false
//			}
//
//			result[i].AmountIn = amountIn.String()
//			result[i].PoolFee = make([]uint64, swapsCnt)
//			result[i].TokenFee = make([]uint64, swapsCnt) // just filling zero here
//			result[i].Direction = make([]uint64, swapsCnt)
//			fees := new(big.Int).SetBytes(input[poolFeePos : poolFeePos+32])
//			poolFeePos += 32 // Now points to pool fees of next transaction
//
//			result[i].Path = make([]common.Address, swapsCnt)
//			fees = new(big.Int).Rsh(fees, 1) // remaining 1 bit unused
//			for j := 0; j < swapsCnt; j++ {
//				result[i].Path[j] = common.BytesToAddress(input[hopsPos+12 : hopsPos+32])
//				result[i].PoolFee[j] = uint64(new(big.Int).Mod(fees, new(big.Int).SetUint64(65536)).Uint64())
//				fees = new(big.Int).Rsh(fees, 16) // 16 bit fee
//				dir := fees.Bit(0) + (fees.Bit(1) << 1)
//				fees = new(big.Int).Rsh(fees, 2)         // 2 bit direction
//				result[i].Direction[j] = uint64(3 - dir) // 2 -> 1, 3 -> 0. Even higher bit is unknown usage, we reserve it for the future.
//
//				hopsPos += 32
//			}
//		}
//		return result, nil, false
//	}
//
//	return nil, nil, false
//}
//
//func parseObs3Input(di *defi.Info, input []byte, from *common.Address) (*ObsParsedResult, *mysql.DBLoop, bool) {
//	size := len(input)
//	minSize := 58 + 24*2 // for 2 hops
//	if size < minSize {
//		return nil, nil, false
//	}
//
//	hops := (size - 58) / 24
//	addrs := make([]common.Address, hops)
//	path := make([]common.Address, hops*2)
//	poolFee := make([]uint64, hops)
//	tokenFee := make([]uint64, hops)
//	direction := make([]uint64, hops)
//	index := make([]uint64, hops)
//	amountIn := new(big.Int).SetBytes(input[5:21]) // 16 bytes
//	token0 := new(big.Int).SetBytes(input[21:41])
//	path[0] = common.BigToAddress(token0)
//	fromInt := new(big.Int).SetBytes(from.Bytes())
//	pos := 42
//	secret := new(big.Int).Xor(amountIn, fromInt)
//	secret = new(big.Int).Xor(secret, token0)
//	secret = new(big.Int).Lsh(secret, 1)
//
//	toDbLoop := true
//	for i := 0; i < hops; i++ {
//		addrInt := new(big.Int).SetBytes(input[pos+4 : pos+24])
//		addrInt = new(big.Int).Xor(addrInt, secret)
//		pair := common.BigToAddress(addrInt)
//		pairInfo := di.GetPairInfo(pair)
//		if pairInfo == nil {
//			toDbLoop = false
//		}
//
//		if toDbLoop {
//			tokenIn := path[2*i]
//			path[2*i+1] = pair
//			tokenOut := tokenIn
//			if tokenIn == pairInfo.Token0 {
//				direction[i] = 0
//				tokenOut = pairInfo.Token1
//			} else if tokenIn == pairInfo.Token1 {
//				direction[i] = 1
//				tokenOut = pairInfo.Token0
//			} else {
//				toDbLoop = false
//			}
//			if i != 0 {
//				path[2*i] = tokenOut
//			}
//
//			index[i] = pairInfo.Index
//		}
//		addrs[i] = pair
//
//		poolFee[i] = uint64(input[pos])
//		direction[i] = uint64(input[pos+1])
//		tokenFee[i] = uint64(binary.BigEndian.Uint16(input[pos+2 : pos+4]))
//		pos += 24
//	}
//
//	var (
//		dbLoop  *mysql.DBLoop
//		obsLoop *ObsParsedResult
//	)
//
//	obsLoop = &ObsParsedResult{
//		AmountIn:  amountIn.String(),
//		Path:      addrs,
//		PoolFee:   poolFee,
//		TokenFee:  tokenFee,
//		Direction: direction,
//	}
//
//	if toDbLoop {
//		dbLoop := &mysql.DBLoop{
//			Kid:       -1,
//			Key:       rcommon.ToKey(path),
//			LoopId:    rcommon.ToLoopId(path, poolFee, tokenFee),
//			Path:      path,
//			PoolFee:   poolFee,
//			TokenFee:  tokenFee,
//			Direction: direction,
//			Index:     index,
//			Count:     1,
//		}
//
//		return obsLoop, dbLoop, true
//	} else {
//		return obsLoop, dbLoop, false
//	}
//}
//
//func parseObs4Input(di *defi.Info, input []byte) (*ObsParsedResult, *mysql.DBLoop, bool) {
//	size := len(input)
//	minSize := 4 + 32*3 + 32*4 // for 2 hops
//	if size < minSize {
//		return nil, nil, false
//	}
//	pathLen := (size - 100) / 32
//	hops := pathLen / 2
//	addrs := make([]common.Address, pathLen)
//	path := make([]common.Address, pathLen)
//	poolFee := make([]uint64, hops)
//	tokenFee := make([]uint64, hops)
//	direction := make([]uint64, hops)
//	index := make([]uint64, hops)
//	amountIn := new(big.Int).SetBytes(input[4:36]) // 32 bytes
//	token0 := new(big.Int).SetBytes(input[size-20 : size])
//	path[0] = common.BigToAddress(token0)
//	pos := 100
//
//	toDbLoop := true
//	for i := 0; i < pathLen; i++ {
//		addr := common.BytesToAddress(input[pos+12 : pos+32])
//		addrs[i] = addr
//
//		if i+1 < pathLen { // last one (token) has been put as addr[0], i.e. token0
//			path[i+1] = addr
//		}
//		if i%2 == 0 {
//			pairInfo := di.GetPairInfo(addr)
//			if pairInfo == nil {
//				toDbLoop = false
//			}
//
//			if toDbLoop {
//				tokenIn := path[i]
//				if tokenIn == pairInfo.Token0 {
//					direction[i/2] = 0
//				} else if tokenIn == pairInfo.Token1 {
//					direction[i/2] = 1
//				} else {
//					toDbLoop = false
//				}
//
//				index[i/2] = pairInfo.Index
//			}
//
//			poolFee[i/2] = uint64(10000 - int(binary.BigEndian.Uint16(input[pos+10:pos+12])))
//		} else {
//			tokenFee[i/2] = uint64(10000 - int(binary.BigEndian.Uint16(input[pos+10:pos+12])))
//		}
//		pos += 32
//	}
//
//	var (
//		dbLoop  *mysql.DBLoop
//		obsLoop *ObsParsedResult
//	)
//
//	obsLoop = &ObsParsedResult{
//		AmountIn:  amountIn.String(),
//		Path:      addrs,
//		PoolFee:   poolFee,
//		TokenFee:  tokenFee,
//		Direction: direction,
//	}
//
//	if toDbLoop {
//		dbLoop := &mysql.DBLoop{
//			Kid:       -1,
//			Key:       rcommon.ToKey(path),
//			LoopId:    rcommon.ToLoopId(path, poolFee, tokenFee),
//			Path:      path,
//			PoolFee:   poolFee,
//			TokenFee:  tokenFee,
//			Direction: direction,
//			Index:     index,
//			Count:     1,
//		}
//
//		return obsLoop, dbLoop, true
//	} else {
//		return obsLoop, dbLoop, false
//	}
//}
//
//func parseObs5Input(di *defi.Info, input []byte) (*ObsParsedResult, *mysql.DBLoop, bool) {
//	size := len(input)
//	minSize := 4 + 32*7 + 32*5*2 // for 2 hops
//	if size < minSize {
//		return nil, nil, false
//	}
//	hops := int(input[4+32*7-1]) // array size field
//	if hops > (size-228)/32/5 {
//		hops = (size - 228) / 32 / 5
//	}
//
//	hasV3 := false
//	addrs := make([]common.Address, hops)
//	path := make([]common.Address, hops*2)
//	poolFee := make([]uint64, hops)
//	tokenFee := make([]uint64, hops)
//	direction := make([]uint64, hops)
//	index := make([]uint64, hops)
//	amountIn := new(big.Int).SetBytes(input[4:36]) // 32 bytes
//	pos := 4 + 32*7
//	tokenFee[0] = 0
//	for i := 0; i < hops; i++ {
//		pair := common.BytesToAddress(input[pos+12 : pos+32])
//		addrs[i] = pair
//
//		path[2*i+1] = pair
//		pos += 32
//		tokenIn := common.BytesToAddress(input[pos+12 : pos+32])
//		pos += 32
//		dir := uint64(input[pos+31])
//		if dir == 0x38 {
//			hasV3 = true
//		}
//		direction[i] = dir
//		pos += 32
//		poolFee[i] = uint64(10000 - int(binary.BigEndian.Uint16(input[pos+30:pos+32])))
//		pos += 32
//		tf := uint64(10000 - int(binary.BigEndian.Uint16(input[pos+30:pos+32])))
//		if i >= hops-1 {
//			path[0] = tokenIn // last token must be the token0
//		} else {
//			path[2*i+2] = tokenIn
//			tokenFee[i/2+1] = tf
//		}
//		pos += 32
//	}
//
//	toDbLoop := true
//	for i := 0; i < hops*2; i += 2 {
//		tokenIn := path[i]
//		pair := path[i+1]
//
//		var (
//			token0     common.Address
//			token1     common.Address
//			v2ResIndex uint64
//			v3Pair     bool
//		)
//
//		if direction[i/2] != 0x38 {
//			pairInfo := di.GetPairInfo(pair)
//			if pairInfo == nil {
//				toDbLoop = false
//				break
//			}
//			token0 = pairInfo.Token0
//			token1 = pairInfo.Token1
//			v2ResIndex = pairInfo.Index
//		} else {
//			poolInfo := di.GetPoolInfo(pair)
//			if poolInfo == nil {
//				toDbLoop = false
//				break
//			}
//			token0 = poolInfo.Token0
//			token1 = poolInfo.Token1
//			v2ResIndex = 0xFF
//			v3Pair = true
//		}
//
//		if toDbLoop {
//			dir := uint64(0)
//			if v3Pair {
//				dir += 4
//			}
//			if tokenIn == token0 {
//				direction[i/2] = dir
//			} else if tokenIn == token1 {
//				dir += 1
//				direction[i/2] = dir
//			} else {
//				toDbLoop = false
//			}
//			index[i/2] = v2ResIndex
//		}
//	}
//
//	var (
//		dbLoop  *mysql.DBLoop
//		obsLoop *ObsParsedResult
//	)
//
//	obsLoop = &ObsParsedResult{
//		AmountIn:  amountIn.String(),
//		Path:      addrs,
//		PoolFee:   poolFee,
//		TokenFee:  tokenFee,
//		Direction: direction,
//		HasV3:     hasV3,
//	}
//
//	if toDbLoop {
//		dbLoop := &mysql.DBLoop{
//			Kid:       -1,
//			Key:       rcommon.ToKey(path),
//			LoopId:    rcommon.ToLoopId(path, poolFee, tokenFee),
//			Path:      path,
//			PoolFee:   poolFee,
//			TokenFee:  tokenFee,
//			Direction: direction,
//			Index:     index,
//			Count:     1,
//			HasV3:     hasV3,
//		}
//
//		return obsLoop, dbLoop, true
//	} else {
//		return obsLoop, dbLoop, false
//	}
//}
//
//func parseObs6Input(di *defi.Info, input []byte) (*ObsParsedResult, *mysql.DBLoop, bool) {
//	size := len(input)
//	minSize := 4 + 32*7*2 + 32*3 // for 2 hops
//	if size < minSize {
//		return nil, nil, false
//	}
//
//	methodID := binary.BigEndian.Uint32(input[:4])
//	if methodID == Obs6Method {
//		hops := int(input[4+32*3-1])
//		pathLen := hops * 2
//		addrs := make([]common.Address, hops)
//		path := make([]common.Address, pathLen)
//		poolFee := make([]uint64, hops)
//		tokenFee := make([]uint64, hops)
//		direction := make([]uint64, hops) // actually, here is an unknown field, not the direction, we just want to reuse this field to keep this data.
//		index := make([]uint64, hops)
//		//amountIn := new(big.Int).SetBytes(input[4:16]) // 12 bytes
//
//		toDbLoop := true
//		pos := 4 + 32*3
//		for i := 0; i < hops; i++ {
//			pos += 32
//			pair := common.BytesToAddress(input[pos+12 : pos+32])
//			addrs[i] = pair
//			pos += 32 * 3
//			tokenIn := common.BytesToAddress(input[pos+12 : pos+32])
//
//			pairInfo := di.GetPairInfo(pair)
//			if pairInfo == nil {
//				toDbLoop = false
//			}
//
//			dir := uint64(0)
//			if toDbLoop {
//				if tokenIn == pairInfo.Token0 {
//					dir = 0
//				} else if tokenIn == pairInfo.Token1 {
//					dir = 1
//				} else {
//					toDbLoop = false
//				}
//				index[i] = pairInfo.Index
//				path[2*i] = tokenIn
//				path[2*i+1] = pair
//				direction[i] = dir
//			}
//
//			pos += 32
//			poolFee[i] = 10000 - uint64(math.Ceil(float64(binary.BigEndian.Uint32(input[pos+28:pos+32]))/10.0))
//			pos += 32
//			tokenFee[i] = 10000 - uint64(math.Ceil(float64(binary.BigEndian.Uint32(input[pos+28:pos+32]))/10.0))
//			pos += 32
//		}
//
//		var (
//			dbLoop  *mysql.DBLoop
//			obsLoop *ObsParsedResult
//		)
//
//		obsLoop = &ObsParsedResult{
//			AmountIn:  new(big.Int).String(),
//			Path:      addrs,
//			PoolFee:   poolFee,
//			TokenFee:  tokenFee,
//			Direction: direction,
//		}
//
//		if toDbLoop {
//			dbLoop := &mysql.DBLoop{
//				Kid:       -1,
//				Key:       rcommon.ToKey(path),
//				LoopId:    rcommon.ToLoopId(path, poolFee, tokenFee),
//				Path:      path,
//				PoolFee:   poolFee,
//				TokenFee:  tokenFee,
//				Direction: direction,
//				Index:     index,
//				Count:     1,
//			}
//
//			return obsLoop, dbLoop, true
//		} else {
//			return obsLoop, dbLoop, false
//		}
//	}
//
//	return nil, nil, false
//}
//
//func parseObs7Input(di *defi.Info, input []byte) (*ObsParsedResult, *mysql.DBLoop, bool) {
//	size := len(input)
//	minSize := 4 + 32*4 // for 2 hops
//	if size < minSize {
//		return nil, nil, false
//	}
//
//	pathLen := (size - 4) / 32
//	hops := pathLen / 2
//	addrs := make([]common.Address, pathLen)
//	path := make([]common.Address, pathLen)
//	poolFee := make([]uint64, hops)
//	tokenFee := make([]uint64, hops)
//	direction := make([]uint64, hops) // actually, here is an unknown field, not the direction, we just want to reuse this field to keep this data.
//	index := make([]uint64, hops)
//	amountIn := new(big.Int).SetBytes(input[4:16]) // 12 bytes
//	token0 := new(big.Int).SetBytes(input[16:36])
//	path[0] = common.BigToAddress(token0)
//	pos := 36
//
//	// pairs
//	for i := 0; i < hops; i++ {
//		path[2*i+1] = common.BytesToAddress(input[pos+12 : pos+32])
//		poolFee[i] = uint64(10000 - int(binary.BigEndian.Uint16(input[pos+10:pos+12])))
//		if i > 0 {
//			direction[i] = uint64(binary.BigEndian.Uint16(input[pos+8 : pos+10])) // I don't know whether it only has 1 byte there, normally I just saw 2 or 0. But for safety, I keep more data if I can.
//		}
//		pos += 32
//	}
//
//	// tokens
//	for i := 0; i < hops-1; i++ {
//		path[2*i+2] = common.BytesToAddress(input[pos+12 : pos+32])
//		tokenFee[i+1] = uint64(binary.BigEndian.Uint16(input[pos+10 : pos+12])) // So far, I never saw any non-zero value here, but let's just keep it here to avoid any surprise.
//		pos += 32
//	}
//
//	toDbLoop := true
//	for i := 0; i < pathLen; i += 2 {
//		tokenIn := path[i]
//		pair := path[i+1]
//		addrs[i] = path[i]
//		addrs[i+1] = path[i+1]
//
//		pairInfo := di.GetPairInfo(pair)
//		if pairInfo == nil {
//			toDbLoop = false
//		}
//
//		dir := uint64(0)
//
//		if toDbLoop {
//			if tokenIn == pairInfo.Token0 {
//				dir = 0
//			} else if tokenIn == pairInfo.Token1 {
//				dir = 1
//			} else {
//				toDbLoop = false
//			}
//			index[i/2] = pairInfo.Index
//		}
//
//		dir |= (direction[i/2] & 0x1) << 1
//		direction[i/2] = dir
//	}
//
//	var (
//		dbLoop  *mysql.DBLoop
//		obsLoop *ObsParsedResult
//	)
//
//	obsLoop = &ObsParsedResult{
//		AmountIn:  amountIn.String(),
//		Path:      addrs,
//		PoolFee:   poolFee,
//		TokenFee:  tokenFee,
//		Direction: direction,
//	}
//
//	if toDbLoop {
//		dbLoop := &mysql.DBLoop{
//			Kid:       -1,
//			Key:       rcommon.ToKey(path),
//			LoopId:    rcommon.ToLoopId(path, poolFee, tokenFee),
//			Path:      path,
//			PoolFee:   poolFee,
//			TokenFee:  tokenFee,
//			Direction: direction,
//			Index:     index,
//			Count:     1,
//		}
//
//		return obsLoop, dbLoop, true
//	} else {
//		return obsLoop, dbLoop, false
//	}
//}
//
//func parseRonFiInput(di *defi.Info, methodId uint32, input []byte) (*ObsParsedResult, *mysql.DBLoop, bool) {
//	size := len(input)
//	switch methodId {
//	case MethodRonSwapX, MethodRonSwapD:
//		minSize := 196 + 32*2 // for 2 hops
//		if size < minSize {
//			return nil, nil, false
//		}
//		hops := (size - 196) / 32
//		path := make([]common.Address, hops*2)
//		poolFee := make([]uint64, hops)
//		tokenFee := make([]uint64, hops)
//		direction := make([]uint64, hops)
//		index := make([]uint64, hops)
//		amountIn := new(big.Int).SetBytes(input[68:100])
//		token0 := new(big.Int).SetBytes(input[164:196])
//		path[0] = common.BigToAddress(token0)
//		rawDirection := binary.BigEndian.Uint16(input[66:68])
//		pos := 196
//		feeIndex := 64
//
//		toDbLoop := true
//		for i := 0; i < hops; i++ {
//			addrInt := new(big.Int).SetBytes(input[pos : pos+32])
//			pair := common.BigToAddress(addrInt)
//			pairInfo := di.GetPairInfo(pair)
//			if pairInfo == nil {
//				toDbLoop = false
//			}
//
//			tokenIn := path[i]
//			tokenOut := tokenIn
//
//			if toDbLoop {
//				if tokenIn == pairInfo.Token0 {
//					tokenOut = pairInfo.Token1
//				} else if tokenIn == pairInfo.Token1 {
//					tokenOut = pairInfo.Token0
//				} else {
//					toDbLoop = false
//				}
//				path[2*i+1] = pair
//				if i != hops-1 {
//					path[2*i+2] = tokenOut
//				}
//				index[i] = pairInfo.Index
//			}
//
//			poolFee[i] = uint64(binary.BigEndian.Uint16(input[feeIndex : feeIndex+2]))
//			direction[i] = uint64((rawDirection >> (2 * i)) & 3)
//			pos += 32
//			feeIndex -= 2
//		}
//
//		var (
//			dbLoop  *mysql.DBLoop
//			obsLoop *ObsParsedResult
//		)
//
//		obsLoop = &ObsParsedResult{
//			AmountIn:  amountIn.String(),
//			Path:      path,
//			PoolFee:   poolFee,
//			TokenFee:  tokenFee,
//			Direction: direction,
//		}
//
//		if toDbLoop {
//			dbLoop := &mysql.DBLoop{
//				Kid:       -1,
//				Key:       rcommon.ToKey(path),
//				LoopId:    rcommon.ToLoopId(path, poolFee, tokenFee),
//				Path:      path,
//				PoolFee:   poolFee,
//				TokenFee:  tokenFee,
//				Direction: direction,
//				Index:     index,
//				Count:     1,
//			}
//
//			return obsLoop, dbLoop, true
//		} else {
//			return obsLoop, dbLoop, false
//		}
//
//	case MethodRonSwapZD, MethodRonSwapZC:
//		minSize := 4 + 32*3 + 32*4 // for 2 hops
//		if size < minSize {
//			return nil, nil, false
//		}
//
//		pathLen := (size - 100) / 32
//		hops := pathLen / 2
//		path := make([]common.Address, pathLen)
//		poolFee := make([]uint64, hops)
//		tokenFee := make([]uint64, hops)
//		direction := make([]uint64, hops)
//		index := make([]uint64, hops)
//		kid := -1
//		amountIn := new(big.Int).SetBytes(input[54:68]) // 14 bytes = 112 bits
//		token0 := new(big.Int).SetBytes(input[size-20 : size])
//		path[0] = common.BigToAddress(token0)
//		pos := 100
//
//		toDbLoop := true
//		for i := 0; i < pathLen; i++ {
//			addr := common.BytesToAddress(input[pos+12 : pos+32])
//			if i+1 < pathLen { // last one (token) has been put as addr[0], i.e. token0
//				path[i+1] = addr
//			}
//			if i%2 == 0 {
//				pairInfo := di.GetPairInfo(addr)
//				if pairInfo == nil {
//					toDbLoop = false
//				}
//
//				if toDbLoop {
//					index[i/2] = pairInfo.Index
//				}
//
//				if i == 0 {
//					kid = int(binary.BigEndian.Uint32(input[pos+5 : pos+9]))
//				}
//				poolFee[i/2] = uint64(binary.BigEndian.Uint16(input[pos+10 : pos+12]))
//				direction[i/2] = uint64(input[pos+9])
//			} else {
//				tokenFee[i/2] = uint64(binary.BigEndian.Uint16(input[pos+10 : pos+12]))
//			}
//			pos += 32
//		}
//
//		var (
//			dbLoop  *mysql.DBLoop
//			obsLoop *ObsParsedResult
//		)
//
//		obsLoop = &ObsParsedResult{
//			AmountIn:  amountIn.String(),
//			Path:      path,
//			PoolFee:   poolFee,
//			TokenFee:  tokenFee,
//			Direction: direction,
//		}
//
//		if toDbLoop {
//			dbLoop := &mysql.DBLoop{
//				Kid:       kid,
//				Key:       rcommon.ToKey(path),
//				LoopId:    rcommon.ToLoopId(path, poolFee, tokenFee),
//				Path:      path,
//				PoolFee:   poolFee,
//				TokenFee:  tokenFee,
//				Direction: direction,
//				Index:     index,
//				Count:     1,
//			}
//
//			return obsLoop, dbLoop, true
//		} else {
//			return obsLoop, dbLoop, false
//		}
//	case MethodRonExecPL:
//		return nil, nil, true
//	default:
//		return nil, nil, false
//	}
//}
