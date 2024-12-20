package queue

import (
	"container/heap"
	"iter"
)

type PQItem[T any] struct {
	Value    T
	Priority int
	index    int
}

type PriorityQueue[T any] []*PQItem[T]

func NewPriorityQueue[T any]() *PriorityQueue[T] {
	res := make(PriorityQueue[T], 0)
	heap.Init(&res)
	return &res
}

func (pq PriorityQueue[T]) Len() int { return len(pq) }

func (pq PriorityQueue[T]) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue[T]) Push(x any) {
	n := len(*pq)
	item := x.(*PQItem[T])
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue[T]) PushValue(value T, priority int) {
	heap.Push(pq, &PQItem[T]{Value: value, Priority: priority})
}

func (pq *PriorityQueue[T]) PopValue() T {
	return heap.Pop(pq).(*PQItem[T]).Value
}

func (pq *PriorityQueue[T]) Seq() iter.Seq[T] {
	return func(yield func(T) bool) {
		for len(*pq) > 0 {
			if !yield(pq.PopValue()) {
				break
			}
		}
	}
}

func (pq *PriorityQueue[T]) SeqPriority() iter.Seq2[T, int] {
	return func(yield func(T, int) bool) {
		for len(*pq) > 0 {
			item := heap.Pop(pq).(*PQItem[T])
			if !yield(item.Value, item.Priority) {
				break
			}
		}
	}
}
