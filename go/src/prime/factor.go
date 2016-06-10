package prime

import (
	"math/big"
)

func Factor(n int) chan int {
	ch := make(chan int, 32)
	go func() {
		var div int = 2
		for {
			if n % div == 0 {
				n /= div
				ch <- div
			} else {
				div++
				if div * div > n {
					break
				}
			}
		}
		close(ch)
	}()
	return ch
}

func BigFactor(n *big.Int) chan big.Int {
	ch := make(chan big.Int, 32)
	go func() {
		var rem big.Int
		d := big.NewInt(2)
		zero := big.NewInt(0)
		one := big.NewInt(1)
		for {
			rem.Mod(n, d)
			if rem.Cmp(zero) == 0 {
				n.Div(n, d)

				value := big.NewInt(0)
				value.Set(d)
				ch <- *value
			} else {
				d.Add(d, one)
				var temp big.Int
				temp.Mul(d, d)
				if n.Cmp(&temp) < 0 {
					break
				}
			}
		}
		close(ch)
	}()
	return ch
}

func Factors(n int) []int {
	factors := make([]int, 0)
	for f := range Factor(n) {
		factors = append(factors, f)
	}
	return factors
}

func BigFactors(n *big.Int) []big.Int {
	factors := make([]big.Int, 0)
	for f := range BigFactor(n) {
		factors = append(factors, f)
	}
	return factors
}

