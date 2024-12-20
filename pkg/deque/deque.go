package deque

import "iter"

type Deque[T any] []T

func New[T any](values ...T) Deque[T] {
	return Deque[T](values)
}

func (d *Deque[T]) PushFront(v T) {
	*d = append([]T{v}, *d...)
}

func (d *Deque[T]) PushBack(v T) {
	*d = append(*d, v)
}

func (d *Deque[T]) PopFront() *T {
	if len(*d) == 0 {
		return nil
	}
	v := (*d)[0]
	*d = (*d)[1:]
	return &v
}

func (d *Deque[T]) PopBack() *T {
	if len(*d) == 0 {
		return nil
	}
	v := (*d)[len(*d)-1]
	*d = (*d)[:len(*d)-1]
	return &v
}

// Seq returns an iterator that yields elements from the deque until it is empty, popping elements as it goes.
func (d *Deque[T]) SeqFront() iter.Seq[T] {
	return func(yield func(T) bool) {
		for len(*d) > 0 {
			if !yield(*d.PopFront()) {
				return
			}
		}
	}
}

// SeqBack returns an iterator that yields elements from the deque until it is empty, popping elements from the back as it goes.
func (d *Deque[T]) SeqBack() iter.Seq[T] {
	return func(yield func(T) bool) {
		for len(*d) > 0 {
			if !yield(*d.PopBack()) {
				return
			}
		}
	}
}
