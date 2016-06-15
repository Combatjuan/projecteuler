package main
/*
Implements Project Euler Problem 5:
https://projecteuler.net/problem=5

2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

import (
	"fmt"
)

func main() {
	var i int = 20
	fmt.Println("Project Euler Problem 5")
	// A better way to do this would be to factor all the numbers and multiple the distinct factors
	// their maximum count.  Must faster.  But this brute force method takes less than a second
	// for the desired input.  So... meh.
	for {
		notFound := false
		for j := 20; j > 1; j-- {
			if i % j != 0 {
				notFound = true
				break;
			}
		}
		if notFound == false {
			fmt.Println(i)
			return
		}
		i++
	}
}
