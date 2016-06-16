package prime

import (
	"container/heap"
)

type sieveRecord struct {
	Next int
	Divisor int
	index int
}

type SieveQueue []*sieveRecord

func (sq SieveQueue) Len() int { return len(sq) }

func (sq SieveQueue) Less(i, j int) bool {
	return sq[i].Next < sq[j].Next;
}

func (sq SieveQueue) Swap(i, j int) {
	sq[i], sq[j] = sq[j], sq[i]
	sq[i].index = i
	sq[j].index = j
}

func (sq *SieveQueue) Push(x interface{}) {
	n := len(*sq)
	item := x.(*sieveRecord)
	item.index = n
	*sq = append(*sq, item)
}

func (sq *SieveQueue) Pop() interface{} {
	old := *sq
	n := len(old)
	item := old[n-1]
	item.index = -1 // For safety?
	*sq = old[0:n-1]
	return item
}

func (sq *SieveQueue) update(item *sieveRecord, divisor int, next int) {
	item.Divisor = divisor
	item.Next = next
	heap.Fix(sq, item.index)
}

