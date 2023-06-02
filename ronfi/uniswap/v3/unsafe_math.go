package v3

import "math/big"

func divRoundingUp(x *big.Int, y *big.Int) *big.Int {
	solution := new(big.Int).Div(x, y)
	zero := big.NewInt(0)
	if new(big.Int).Mod(x, y).Cmp(zero) > 0 {
		solution = new(big.Int).Add(solution, big.NewInt(1))
	}

	return solution
}
