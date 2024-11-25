package models

import "math/big"

type FibonacciResult struct {
	Value *big.Int
}

func CalculateFibonacci(n int) *big.Int {
	if n <= 1 {
		return big.NewInt(int64(n))
	}

	a, b := big.NewInt(0), big.NewInt(1)
	for i := 2; i <= n; i++ {
		a.Add(a, b)
		a, b = b, a
	}
	return b
}
