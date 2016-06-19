package main

/*
Implements Project Euler Problem 10:
https://projecteuler.net/problem=10
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/

import (
	"fmt"
	"prime"
)

func main() {
	fmt.Println("Project Euler Problem 10")
	sum := 0
	for p := range prime.Sieve() {
		if p > 2000000 {
			break
		} else {
			sum += p
		}
	}
	fmt.Println(sum)
}
