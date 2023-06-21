package stats

import (
	"github.com/ethereum/go-ethereum/common"
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
	"github.com/ethereum/go-ethereum/ronfi/db"
	"github.com/ethereum/go-ethereum/ronfi/defi"
	"math/big"
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

func parseObs5Input(di *defi.Info, input []byte) (*ObsParsedResult, *db.DBLoop, bool) {
	size := len(input)
	minSize := 4 + 32*5 + 32*5 // for 2 hops
	if size < minSize {
		return nil, nil, false
	}

	hops := int(size-(4+32*5)-32) / 32 / 2

	addrs := make([]common.Address, hops)
	path := make([]common.Address, hops*2)
	poolFee := make([]uint64, hops)
	tokenFee := make([]uint64, hops)
	direction := make([]uint64, hops)
	dbDirs := make([]uint64, hops)
	index := make([]uint64, hops)
	amountIn := new(big.Int).SetBytes(input[size-32:]) // 32 bytes
	pos := 4 + 32*5
	for i := 0; i < hops; i++ {
		pair := common.BytesToAddress(input[pos+12 : pos+32])
		addrs[i] = pair
		pos += 32
	}

	for i := 0; i < hops; i++ {
		dir := uint64(input[pos+31])
		if dir == 0 {
			dir = 1
		} else {
			dir = 0
		}
		direction[i] = dir
		pos += 32
	}

	toDbLoop := true
	for i := 0; i < hops; i++ {
		pair := addrs[i]

		var (
			token0     common.Address
			token1     common.Address
			v2ResIndex uint64
			v3Pair     bool
		)

		pairInfo := di.GetPairInfo(pair)
		if pairInfo == nil {
			// check if v3 pool
			poolInfo := di.GetPoolInfo(pair)
			if poolInfo == nil {
				toDbLoop = false
				break
			} else {
				token0 = poolInfo.Token0
				token1 = poolInfo.Token1
				v2ResIndex = 0xFF
				v3Pair = true
			}
		} else {
			token0 = pairInfo.Token0
			token1 = pairInfo.Token1
			v2ResIndex = pairInfo.Index
			poolFee[i] = 30
		}

		if toDbLoop {
			dir := direction[i]
			if dir == 0 {
				path[i*2] = token0
			} else {
				path[i*2] = token1
			}
			path[i*2+1] = pair
			if v3Pair {
				dir += 4
			}
			dbDirs[i] = dir
			index[i] = v2ResIndex
		}
	}

	var (
		dbLoop  *db.DBLoop
		obsLoop *ObsParsedResult
	)

	obsLoop = &ObsParsedResult{
		AmountIn:  amountIn.String(),
		Path:      addrs,
		PoolFee:   poolFee,
		TokenFee:  tokenFee,
		Direction: direction,
	}

	if toDbLoop {
		dbLoop := &db.DBLoop{
			Key:       rcommon.ToKey(path),
			LoopId:    rcommon.ToLoopId(path, poolFee, tokenFee),
			Path:      path,
			PoolFee:   poolFee,
			TokenFee:  tokenFee,
			Direction: dbDirs,
			Index:     index,
			Count:     1,
		}

		return obsLoop, dbLoop, true
	} else {
		return obsLoop, dbLoop, false
	}
}

func parseObs6Input(di *defi.Info, input []byte) (*ObsParsedResult, *db.DBLoop, bool) {
	size := len(input)
	minSize := 77 // for 2 hops
	if size < minSize {
		return nil, nil, false
	}

	hops := 2

	addrs := make([]common.Address, hops)
	path := make([]common.Address, hops*2)
	poolFee := make([]uint64, hops)
	tokenFee := make([]uint64, hops)
	direction := make([]uint64, hops)
	dbDirs := make([]uint64, hops)
	index := make([]uint64, hops)
	amountIn := new(big.Int).SetBytes(input[4:13]) // 32 bytes
	pos := 4 + 9
	for i := 0; i < hops; i++ {
		pair := common.BytesToAddress(input[pos : pos+20])
		addrs[i] = pair
		pos += 20
	}

	toDbLoop := true
	var tokenIn, tokenOut common.Address
	tokenIn = rcommon.WETH
	for i := 0; i < hops; i++ {
		pair := addrs[i]

		var (
			token0     common.Address
			token1     common.Address
			v2ResIndex uint64
			v3Pair     bool
		)

		pairInfo := di.GetPairInfo(pair)
		if pairInfo == nil {
			// check if v3 pool
			poolInfo := di.GetPoolInfo(pair)
			if poolInfo == nil {
				toDbLoop = false
				break
			} else {
				token0 = poolInfo.Token0
				token1 = poolInfo.Token1
				v2ResIndex = 0xFF
				v3Pair = true
			}
		} else {
			token0 = pairInfo.Token0
			token1 = pairInfo.Token1
			v2ResIndex = pairInfo.Index
			poolFee[i] = 30
		}

		dir := uint64(0)
		if tokenIn == token0 {
			dir = 0
			tokenOut = token1
		} else {
			dir = 1
			tokenOut = token0
		}

		if toDbLoop {
			path[i*2] = tokenIn
			path[i*2+1] = pair
			direction[i] = dir
			if v3Pair {
				dir += 4
			}
			dbDirs[i] = dir
			index[i] = v2ResIndex
		}

		tokenIn = tokenOut
		tokenOut = rcommon.ZeroAddress
	}

	var (
		dbLoop  *db.DBLoop
		obsLoop *ObsParsedResult
	)

	obsLoop = &ObsParsedResult{
		AmountIn:  amountIn.String(),
		Path:      addrs,
		PoolFee:   poolFee,
		TokenFee:  tokenFee,
		Direction: direction,
	}

	if toDbLoop {
		dbLoop := &db.DBLoop{
			Key:       rcommon.ToKey(path),
			LoopId:    rcommon.ToLoopId(path, poolFee, tokenFee),
			Path:      path,
			PoolFee:   poolFee,
			TokenFee:  tokenFee,
			Direction: dbDirs,
			Index:     index,
			Count:     1,
		}

		return obsLoop, dbLoop, true
	} else {
		return obsLoop, dbLoop, false
	}
}
