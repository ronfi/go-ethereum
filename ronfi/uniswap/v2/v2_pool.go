package v2

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/ronfi/defi"
	v3 "github.com/ethereum/go-ethereum/ronfi/uniswap/v3"
	"math/big"
)

const (
	V2_FEE_BASE = 10000
)

type PoolState struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
}

type V2Pool struct {
	di       *defi.Info
	statedb  *state.StateDB
	Fee      uint64
	Address  common.Address
	PairInfo *defi.PairInfo
	Name     string
	State    *PoolState
}

func NewV2Pool(di *defi.Info, address common.Address, fee uint64, statedb *state.StateDB) *V2Pool {
	if di == nil {
		log.Warn("RonFi NewV2Pool di is nil")
		return nil
	}

	pairInfo := di.GetPairInfo(address)
	if pairInfo == nil {
		log.Warn("RonFi NewV2Pool pairInfo is nil", "address", address.Hex())
		return nil
	}
	v2Pool := &V2Pool{
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

func (p *V2Pool) UpdateReserves() {
	if p.PairInfo != nil {
		if p.PairInfo.Index != 0xFF && p.statedb != nil {
			res := p.statedb.GetState(p.Address, common.BigToHash(new(big.Int).SetUint64(p.PairInfo.Index)))
			reserve0 := new(big.Int).SetBytes(res[18:32])
			reserve1 := new(big.Int).SetBytes(res[4:18])
			p.State = &PoolState{
				Reserve0: reserve0,
				Reserve1: reserve1,
			}
		} else {
			if res := p.di.GetPairReserves(p.Address); res != nil {
				p.State = &PoolState{
					Reserve0: res.Reserve0,
					Reserve1: res.Reserve1,
				}
			}
		}
	}

	//fmt.Println("RonFi V2Pool UpdateReserves", p.Name, p.State.Reserve0, p.State.Reserve1)
}

func (p *V2Pool) CalculateTokensOutFromTokensIn(tokenIn common.Address, amountIn *big.Int) *big.Int {
	var (
		reserveIn  *big.Int
		reserveOut *big.Int
	)

	if amountIn.Cmp(v3.ZERO) < 0 {
		log.Warn("RonFi calculateTokensOutFromTokensIn cannot be negative!")
		return nil
	}

	if tokenIn == p.PairInfo.Token0 {
		reserveIn = p.State.Reserve0
		reserveOut = p.State.Reserve1
	} else if tokenIn == p.PairInfo.Token1 {
		reserveIn = p.State.Reserve1
		reserveOut = p.State.Reserve0
	} else {
		log.Warn("RonFi calculateTokensOutFromTokensIn Could not identify tokenIn!")
		return nil
	}

	amountInWithFee := new(big.Int).Mul(amountIn, new(big.Int).Sub(big.NewInt(V2_FEE_BASE), big.NewInt(int64(p.Fee))))
	numerator := new(big.Int).Mul(amountInWithFee, reserveOut)
	denominator := new(big.Int).Add(new(big.Int).Mul(reserveIn, big.NewInt(V2_FEE_BASE)), amountInWithFee)
	return new(big.Int).Div(numerator, denominator)
}
