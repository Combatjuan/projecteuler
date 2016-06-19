package main

/*
Implements Project Euler Problem 9:
https://projecteuler.net/problem=9
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
 a2 + b2 = c2

For example, 32 + 42 = 9 + 16 = 25 = 52.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.Find the product abc.
*/

import (
	"fmt"
)

func main() {
	fmt.Println("Project Euler Problem 9")
	for a := 1; a < 1000; a++ {
		for b := a; b < 1000; b++ {
			c := 1000 - a - b
			if a*a+b*b == c*c {
				fmt.Println(a * b * c)
				return
			} else if a*a == b*b+c*c {
				fmt.Println(a * b * c)
				return
			} else if a*a+c*c == b*b {
				fmt.Println(a * b * c)
				return
			}
		}
	}
}
