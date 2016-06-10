package prime

func IsPrime(n int) bool {
	// IsPrime returns true if n has factors (other than itself and 1).
	for range Factor(n) {
		return true
	}
	return false
}

