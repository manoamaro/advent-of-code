package set

import "fmt"

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

func Empty[T comparable]() Set[T] {
	return Set[T]{data: make(map[T]bool)}
}

func FromSlice[T comparable](slice []T) Set[T] {
	s := make(map[T]bool)
	for _, v := range slice {
		s[v] = true
	}
	return Set[T]{data: s}
}

func (s *Set[T]) Add(v T) {
	(*s).data[v] = true
}

func (s *Set[T]) AddAll(v []T) {
	for _, value := range v {
		s.data[value] = true
	}
}

func (s *Set[T]) First() *T {
	for v := range s.data {
		return &v
	}
	return nil
}

func (s Set[T]) Slice() []T {
	slice := make([]T, 0, len(s.data))
	for v := range s.data {
		slice = append(slice, v)
	}
	return slice
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

func (s Set[T]) Contains(v T) bool {
	_, ok := s.data[v]
	return ok
}

func (s Set[T]) Equals(other Set[T]) bool {
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

func (s Set[T]) Len() int {
	return len(s.data)
}

func (s Set[T]) String() string {
	return fmt.Sprintf("%v", s.Slice())
}
