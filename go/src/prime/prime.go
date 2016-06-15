package prime

func IsPrime(n int) bool {
	// IsPrime returns true if n has factors (other than itself and 1).
	for range Factor(n) {
		return true
	}
	return false
}

//func Sieve(start, end int) chan int {
func Sieve() chan int {
	ch := make(chan int, 0)
	sieve := SieveQueue{sieveRecord{2, 4}}

	go func() {
		var next sieveRecord
		n := 2
		ch <- 2
		for {
			for {
				next = sieve.Pop()
				if n < next.next {
					sieve.Push(sieveRecord{n, n + n, 0})
					ch <-n
					n++
				} else {
					break
				}
			}
			sieve.Push(sieveRecord{next.divisor, next.divisor + n.next, 0})
		}
	}()

	return ch
}

