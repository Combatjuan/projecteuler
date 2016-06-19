package main

/*
Implements Project Euler Problem 14:
https://projecteuler.net/problem=14
The following iterative sequence is defined for the set of positive integers:

n = n/2 (n is even)
n = 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:
13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1

It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.
*/

import (
	"fmt"
	"seq"
)

//!func Collatz(n int) chan int {
//!	ch := make(chan int, 0)
//!	go func() {
//!		defer close(ch)
//!		for {
//!			ch <- n
//!			if n == 1 {
//!				return
//!			} else if n%2 == 0 {
//!				n = n / 2
//!			} else {
//!				n = n*3 + 1
//!			}
//!		}
//!	}()
//!	return ch
//!}

func main() {
	fmt.Println("Project Euler Problem 14")
	max := 0
	winner := 0

	// A smarter thing to do would be to memoize value in Collatz.
	// If we encounter a number, we'll know how many more terms that one takes.
	// But this completes in under a minute.
	for n := 1; n < 1000000; n++ {
		l := make([]int, 0)
		for n := range seq.Collatz(n) {
			l = append(l, n)
		}
		if len(l) > max {
			max = len(l)
			winner = n
		}
	}
	fmt.Println(max)
	fmt.Println(winner)
}
