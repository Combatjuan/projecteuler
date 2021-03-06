package main

/*
Implements Project Euler Problem 7:
https://projecteuler.net/problem=7
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?
*/

import (
	"fmt"
	"prime"
)

func main() {
	fmt.Println("Project Euler Problem 7")
	count := 1
	for p := range prime.Sieve() {
		if count == 10001 {
			fmt.Println(p)
			break
		}
		count++
	}
}
