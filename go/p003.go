package main
/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/

import (
	"fmt"
	"prime"
	"sort"
)

func main() {
	f := prime.Factors(600851475143)
	sort.Ints(f)
	if len(f) > 1 {
		fmt.Println(f[len(f) - 1])
	}
}

