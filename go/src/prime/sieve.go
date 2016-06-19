package prime

import (
	"container/heap"
)

type SieveRecord struct {
	Next    int
	Divisor int
}

type SieveQueue []SieveRecord

func (sq SieveQueue) Len() int { return len(sq) }

func (sq SieveQueue) Less(i, j int) bool {
	return sq[i].Next < sq[j].Next
}

func (sq SieveQueue) Swap(i, j int) {
	sq[i], sq[j] = sq[j], sq[i]
}

func (sq *SieveQueue) Push(item interface{}) {
	*sq = append(*sq, item.(SieveRecord))
}

func (sq *SieveQueue) Pop() interface{} {
	old := *sq
	n := len(old)
	item := old[n-1]
	*sq = old[0 : n-1]
	return item
}

//func Sieve(start, end int) chan int {
func Sieve() chan int {
	ch := make(chan int, 0)
	sq := &SieveQueue{}
	heap.Init(sq)

	go func() {
		var next SieveRecord
		exclusions := &SieveQueue{}
		heap.Init(exclusions)
		next = SieveRecord{3, 0}
		n := 2

		// We count up numbers
		for n = 2; true; n++ {
			if n < next.Next {
				// We found a non-excluded number.  It's prime.
				ch <- n
				heap.Push(exclusions, SieveRecord{n + n, n})
				if next.Divisor == 0 {
					next = heap.Pop(exclusions).(SieveRecord)
				}
			} else {
				// Cross off numbers and move us forward.
				for n == next.Next {
					heap.Push(exclusions, SieveRecord{n + next.Divisor, next.Divisor})
					next = heap.Pop(exclusions).(SieveRecord)
				}
			}
		}
	}()

	return ch
}

/*func LinearSieve() chan int {
	ch := make(chan int, 0)

	go func() {
		primes := make([]SieveRecord, 0)

		n := 2
		for {
			var newPrime SieveRecord
			for _, prime := range primes {
				if n == prime.Next {
					primes
				}
			}
			if newPrime.Divisor != 0 {
				primes = append(primes, newPrime
			}
			n++
		}
	}()
	return ch
}
*/
