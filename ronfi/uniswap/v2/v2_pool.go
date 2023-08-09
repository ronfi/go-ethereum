package v2

import (
	"encoding/binary"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/ronfi/defi"
	v3 "github.com/ethereum/go-ethereum/ronfi/uniswap/v3"
	"math/big"
)

const FeeBase = 10000

type PoolState struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
	Stamp    uint32
}

type Pool struct {
	di       *defi.Info
	statedb  *state.StateDB
	Fee      uint64
	Address  common.Address
	PairInfo *defi.PairInfo
	Name     string
	State    *PoolState
}

func NewV2Pool(di *defi.Info, address common.Address, fee uint64, statedb *state.StateDB) *Pool {
	if di == nil {
		log.Warn("RonFi NewV2Pool di is nil")
		return nil
	}

	pairInfo := di.GetPairInfo(address)
	if pairInfo == nil {
		log.Warn("RonFi NewV2Pool pairInfo is nil", "address", address.Hex())
		return nil
	}
	v2Pool := &Pool{
		di:       di,
		statedb:  statedb,
		Address:  address,
		Fee:      fee,
		PairInfo: pairInfo,
	}

	if token0 := di.GetTokenInfo(pairInfo.Token0); token0 != nil {
		token0Symbol := di.GetTokenInfo(pairInfo.Token0).Symbol
		if token1 := di.GetTokenInfo(pairInfo.Token1); token1 != nil {
			token1Symbol := di.GetTokenInfo(pairInfo.Token1).Symbol
			v2Pool.Name = fmt.Sprintf("%s (%s-%s V2, %d)", address, token0Symbol, token1Symbol, fee)
		} else {
			log.Warn("RonFi NewV2Pool token1 is nil", "address", pairInfo.Token1)
			return nil
		}
	} else {
		log.Warn("RonFi NewV2Pool token0 is nil", "address", pairInfo.Token0)
		return nil
	}

	return v2Pool
}

func (p *Pool) UpdateReserves(v2State *PoolState) bool {
	if v2State != nil {
		p.State = v2State
		return true
	} else {
		if p.PairInfo != nil {
			if p.PairInfo.Index != 0xFF && p.statedb != nil {
				res := p.statedb.GetState(p.Address, common.BigToHash(new(big.Int).SetUint64(p.PairInfo.Index)))
				reserve0 := new(big.Int).SetBytes(res[18:32])
				reserve1 := new(big.Int).SetBytes(res[4:18])
				stamp := binary.BigEndian.Uint32(res[0:4])

				p.State = &PoolState{
					Reserve0: reserve0,
					Reserve1: reserve1,
					Stamp:    stamp,
				}

				return true
			} else {
				if res := p.di.GetPairReserves(p.Address); res != nil {
					p.State = &PoolState{
						Reserve0: res.Reserve0,
						Reserve1: res.Reserve1,
						Stamp:    res.Timestamp,
					}

					return true
				}
			}
		}
	}

	return false
}

func (p *Pool) CalculateTokensOutFromTokensIn(tokenIn common.Address, amountIn *big.Int) *big.Int {
	var (
		reserveIn  *big.Int
		reserveOut *big.Int
	)

	if amountIn == nil || amountIn.Cmp(v3.ZERO) <= 0 {
		log.Warn("RonFi CalculateTokensOutFromTokensIn cannot be negative!", "pool", p.Address, "tokenIn", tokenIn)
		return nil
	}

	if tokenIn == p.PairInfo.Token0 {
		reserveIn = p.State.Reserve0
		reserveOut = p.State.Reserve1
	} else if tokenIn == p.PairInfo.Token1 {
		reserveIn = p.State.Reserve1
		reserveOut = p.State.Reserve0
	} else {
		log.Warn("RonFi CalculateTokensOutFromTokensIn Could not identify tokenIn!")
		return nil
	}

	amountInWithFee := new(big.Int).Mul(amountIn, new(big.Int).Sub(big.NewInt(FeeBase), big.NewInt(int64(p.Fee))))
	numerator := new(big.Int).Mul(amountInWithFee, reserveOut)
	denominator := new(big.Int).Add(new(big.Int).Mul(reserveIn, big.NewInt(FeeBase)), amountInWithFee)
	return new(big.Int).Div(numerator, denominator)
}

func (p *Pool) CalculateTokensInFromTokensOut(tokenOut common.Address, amountOut *big.Int) *big.Int {
	var (
		reserveIn  *big.Int
		reserveOut *big.Int
	)

	if amountOut == nil || amountOut.Cmp(v3.ZERO) <= 0 {
		log.Warn("RonFi CalculateTokensInFromTokensOut cannot be negative!", "pool", p.Address, "tokenOut", tokenOut)
		return nil
	}

	if tokenOut == p.PairInfo.Token0 {
		reserveIn = p.State.Reserve1
		reserveOut = p.State.Reserve0
	} else if tokenOut == p.PairInfo.Token1 {
		reserveIn = p.State.Reserve0
		reserveOut = p.State.Reserve1
	} else {
		log.Warn("RonFi CalculateTokensInFromTokensOut Could not identify tokenOut!", "pool", p.Address, "tokenOut", tokenOut)
		return nil
	}

	if amountOut.Cmp(reserveOut) >= 0 {
		//log.Warn("RonFi CalculateTokensInFromTokensOut amountOut is bigger than reserveOut!", "pool", p.Address, "tokenOut", tokenOut)
		return nil
	}

	numerator := new(big.Int).Mul(new(big.Int).Mul(reserveIn, amountOut), big.NewInt(FeeBase))
	denominator := new(big.Int).Mul(new(big.Int).Sub(reserveOut, amountOut), new(big.Int).Sub(big.NewInt(FeeBase), big.NewInt(int64(p.Fee))))
	return new(big.Int).Add(new(big.Int).Div(numerator, denominator), big.NewInt(1))
}

func (p *Pool) MaxTokensInFromTokensOut(tokenOut common.Address) (amountIn *big.Int) {
	var (
		reserveIn  *big.Int
		reserveOut *big.Int
	)

	if tokenOut == p.PairInfo.Token0 {
		reserveIn = p.State.Reserve1
		reserveOut = p.State.Reserve0
	} else if tokenOut == p.PairInfo.Token1 {
		reserveIn = p.State.Reserve0
		reserveOut = p.State.Reserve1
	} else {
		log.Warn("RonFi CalculateTokensInFromTokensOut Could not identify tokenOut!", "pool", p.Address, "tokenOut", tokenOut)
		return nil
	}

	numerator := new(big.Int).Mul(new(big.Int).Mul(reserveIn, new(big.Int).Sub(reserveOut, big.NewInt(1))), big.NewInt(FeeBase))
	denominator := new(big.Int).Mul(new(big.Int).Sub(reserveOut, big.NewInt(1)), new(big.Int).Sub(big.NewInt(FeeBase), big.NewInt(int64(p.Fee))))
	return new(big.Int).Add(new(big.Int).Div(numerator, denominator), big.NewInt(1))
}
