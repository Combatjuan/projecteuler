package main

/*
Implements Project Euler Problem 15:
https://projecteuler.net/problem=15
Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down, there are exactly 6 routes to the bottom right corner.



How many such routes are there through a 20×20 grid?
*/

import (
	"fmt"
	"math/big"
	"peutil"
)

func main() {
	fmt.Println("Project Euler Problem 15")
	// Turns out this is just a multiset permutation problem.
	// If the grid is NxN then you must move down (N-1) times and right (N-1) times in some order.
	// So the answer is (2*19)!/(19!*19!)
	var answer big.Int
	var numerator, denominator, partdenom big.Int
	numerator = peutil.BigFactorial(2 * 19)
	partdenom = peutil.BigFactorial(19)
	denominator.Mul(&partdenom, &partdenom)
	answer.Div(&numerator, &denominator)
	fmt.Println(answer.String())
}

/*
3: 6
X-X-X
| | |
X-X-X
| | |
X-X-X

DDRR
DRDR
DRRD
RDDR
RDRD
RRDD
*/
