package queue

import "iter"

type Queue[T any] []T

func New[T any](values ...T) Queue[T] {
	q := make(Queue[T], len(values))
	copy(q, values)
	return q
}

func (q *Queue[T]) Push(v T) {
	*q = append(*q, v)
}

func (q *Queue[T]) Pop() *T {
	if len(*q) == 0 {
		return nil
	}
	v := (*q)[0]
	*q = (*q)[1:]
	return &v
}

// Seq returns an iterator that yields elements from the queue until it is empty, popping elements as it goes.
func (q *Queue[T]) Seq() iter.Seq[T] {
	return func(yield func(T) bool) {
		for len(*q) > 0 {
			if !yield(*q.Pop()) {
				return
			}
		}
	}
}
