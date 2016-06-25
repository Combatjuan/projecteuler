package peutil

import (
	"math/big"
)

func Factorial(n int) int {
	product := 1
	for i := 2; i <= n; i++ {
		product *= i
	}
	return product
}

func BigFactorial(n int) big.Int {
	product := big.NewInt(1)
	var i int64
	var max int64 = int64(n)
	for i = 2; i <= max; i++ {
		multiplier := big.NewInt(i)
		product.Mul(product, multiplier)
	}
	return *product
}
