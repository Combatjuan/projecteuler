package seq

import (
	"math/big"
)

func Fibonacci() chan *big.Int {
	ch := make(chan *big.Int, 0)
	go func() {
		nminus1 := big.NewInt(0)
		n := big.NewInt(1)
		for {
			nminus1.Add(n, nminus1)
			nminus1, n = n, nminus1
			ch <- n
		}
		close(ch)
	}()
	return ch
}

func Triangle() chan int {
	ch := make(chan int, 0)
	go func() {
		sum := 0
		for n := 1; true; n++ {
			sum += n
			ch <- sum
		}
		close(ch)
	}()
	return ch
}

func Collatz(n int) chan int {
	ch := make(chan int, 0)
	go func() {
		defer close(ch)
		for {
			ch <- n
			if n == 1 {
				return
			} else if n%2 == 0 {
				n = n / 2
			} else {
				n = n*3 + 1
			}
		}
	}()
	return ch
}
