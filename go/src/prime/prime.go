package prime

import (
	"fmt"
)

func IsPrime(n int) bool {
	// IsPrime returns true if n has factors (other than itself and 1).
	for range Factor(n) {
		return true
	}
	return false
}

func PrintExclusions(n int, exclusions SieveQueue) {
	fmt.Printf("N: %d E: ", n)
	for _, e := range exclusions {
		fmt.Printf("(%d, %d) ", e.Next, e.Divisor)
	}
	fmt.Println()
}
