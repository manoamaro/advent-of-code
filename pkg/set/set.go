package set

import (
	"fmt"
	"iter"
	"math"
)

type Set[T comparable] struct {
	data map[T]bool
}

func New[T comparable](items ...T) Set[T] {
	s := make(map[T]bool)
	for _, v := range items {
		s[v] = true
	}
	return Set[T]{data: s}
}

func (s *Set[T]) Add(v T) {
	(*s).data[v] = true
}

func (s *Set[T]) Copy() Set[T] {
	newSet := New[T]()
	for v := range s.data {
		newSet.Add(v)
	}
	return newSet
}

func (s *Set[T]) First() *T {
	for v := range s.data {
		return &v
	}
	return nil
}

func (s *Set[T]) Slice() []T {
	slice := make([]T, 0, len(s.data))
	for v := range s.data {
		slice = append(slice, v)
	}
	return slice
}

func (s *Set[T]) Seq() iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range s.data {
			if !yield(v) {
				return
			}
		}
	}
}

func (s *Set[T]) Remove(v T) {
	delete(s.data, v)
}

func (s *Set[T]) RemoveAll(v []T) {
	for _, value := range v {
		delete(s.data, value)
	}
}

func (s *Set[T]) Clear() {
	s.data = make(map[T]bool)
}

func (s *Set[T]) Contains(v T) bool {
	_, ok := s.data[v]
	return ok
}

func (s *Set[T]) Union(other Set[T]) Set[T] {
	result := New[T]()
	for v := range s.data {
		result.Add(v)
	}
	for v := range other.data {
		result.Add(v)
	}
	return result
}

func (s *Set[T]) Intersection(other Set[T]) Set[T] {
	result := New[T]()
	for v := range s.data {
		if other.Contains(v) {
			result.Add(v)
		}
	}
	return result
}

func (s *Set[T]) Difference(other Set[T]) Set[T] {
	result := New[T]()
	for v := range s.data {
		if !other.Contains(v) {
			result.Add(v)
		}
	}
	return result
}

func (s *Set[T]) SymmetricDifference(other Set[T]) Set[T] {
	result := New[T]()
	for v := range s.data {
		if !other.Contains(v) {
			result.Add(v)
		}
	}
	for v := range other.data {
		if !s.Contains(v) {
			result.Add(v)
		}
	}
	return result
}

func (s *Set[T]) CountFunc(f func(T) bool) int {
	count := 0
	for v := range s.data {
		if f(v) {
			count++
		}
	}
	return count
}

func (s *Set[T]) MaxFunc(f func(T) int) T {
	max := math.MinInt
	var maxV T
	for v := range s.data {
		if f(v) > max {
			max = f(v)
			maxV = v
		}
	}
	return maxV
}

func (s *Set[T]) Equals(other Set[T]) bool {
	if len(s.data) != len(other.data) {
		return false
	}
	for k := range s.data {
		if !other.Contains(k) {
			return false
		}
	}
	return true
}

func (s *Set[T]) Len() int {
	return len(s.data)
}

func (s *Set[T]) String() string {
	return fmt.Sprintf("%v", s.Slice())
}
