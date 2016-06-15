package main
/*
Implements Project Euler Problem 6:
https://projecteuler.net/problem=6
The sum of the squares of the first ten natural numbers is,
1^2 + 2^2 + ... + 10^2 = 385

The square of the sum of the first ten natural numbers is,
(1 + 2 + ... + 10)^2 = 552 = 3025

Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/

import (
	"fmt"
)

func squareOfSums(first, last int) int {
	var total int
	for i := first; i <= last; i++ {
		total += i
	}
	return total * total
}

func sumOfSquares(first, last int) int {
	var total int
	for i := first; i <= last; i++ {
		total += i * i
	}
	return total
}

func main() {
	fmt.Println(squareOfSums(1, 100) - sumOfSquares(1, 100))
}
