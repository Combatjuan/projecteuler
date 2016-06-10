package main
/*
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/

import (
	"fmt"
	"strconv"
)

func ToRunes(s string) []rune {
	u := make([]rune, 0)
	for _, r := range s {
		u = append(u, rune(r))
	}
	return u
}

func IsPalindrome(p string) bool {
	// A string is a palindrome if it reads the same forward and in reverse.
	runes := ToRunes(p)
	l := len(runes)
	for i := 0; i < l; i++ {
		if runes[i] != runes[l - 1 - i] {
			return false
		}
	}
	return true
}

func IsPalindromeInt(n int) bool {
	s := strconv.Itoa(n)
	return IsPalindrome(s)
}

func main() {
	var answer int
	for a := 100; a < 1000; a++ {
		for b := 100; b < a; b++ {
			product := a * b
			if product > answer {
				if IsPalindromeInt(product) {
					answer = product
				}
			}
		}
	}
	fmt.Println(answer)
}

