package prime

import (
//	"math/big"
)

func Factors(n int) []int {
	factors := make([]int, 0)
	returnChannel := Factor(n)
	for {
		n, ok := <- returnChannel
		if !ok {
			return factors
		} else {
			factors = append(factors, n)
		}
	}
	return factors
}

func Factor(n int) chan int {
	ch := make(chan int, 3)
	if n < 2 {
		close(ch)
		return ch
	}
	go func() {
		var div int = 2
		for {
			if n % div == 0 {
				n /= div
				ch <- div
				div = 2
			} else {
				if div * div > n {
					if n > 1 {
						ch <- n
					}
					break
				}
				div++
			}
		}
		close(ch)
	}()
	return ch
}

//!func BigFactor(n *big.Int) chan big.Int {
//!	ch := make(chan big.Int, 32)
//!	if n.Cmp(big.NewInt(2)) < 0 {
//!		close(ch)
//!		return ch
//!	}
//!	go func() {
//!		var rem big.Int
//!		d := big.NewInt(2)
//!		zero := big.NewInt(0)
//!		one := big.NewInt(1)
//!		for {
//!			rem.Mod(n, d)
//!			if rem.Cmp(zero) == 0 {
//!				n.Div(n, d)
//!
//!				value := big.NewInt(0)
//!				value.Set(d)
//!				ch <- *value
//!			} else {
//!				d.Add(d, one)
//!				var temp big.Int
//!				temp.Mul(d, d)
//!				if n.Cmp(&temp) < 0 {
//!					break
//!				}
//!			}
//!		}
//!		close(ch)
//!	}()
//!	return ch
//!}

//!func BigFactors(n *big.Int) []big.Int {
//!	factors := make([]big.Int, 0)
//!	for f := range BigFactor(n) {
//!		factors = append(factors, f)
//!	}
//!	return factors
//!}

