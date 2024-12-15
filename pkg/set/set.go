package set

type Set[T comparable] map[T]struct{}

func New[T comparable](items ...T) Set[T] {
	s := make(Set[T])
	for _, v := range items {
		s[v] = struct{}{}
	}
	return s
}

func FromSlice[T comparable](slice []T) Set[T] {
	s := make(Set[T])
	for _, v := range slice {
		s[v] = struct{}{}
	}
	return s
}

func (s *Set[T]) Add(v T) {
	(*s)[v] = struct{}{}
}

func (s *Set[T]) AddAll(v []T) {
	for _, value := range v {
		(*s)[value] = struct{}{}
	}
}

func (s *Set[T]) First() *T {
	for v := range *s {
		return &v
	}
	return nil
}

func (s *Set[T]) Slice() []T {
	slice := make([]T, 0, len(*s))
	for v := range *s {
		slice = append(slice, v)
	}
	return slice
}

func (s *Set[T]) Remove(v T) {
	delete((*s), v)
}

func (s *Set[T]) RemoveAll(v []T) {
	for _, value := range v {
		delete((*s), value)
	}
}

func (s Set[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) Len() int {
	return len(s)
}
