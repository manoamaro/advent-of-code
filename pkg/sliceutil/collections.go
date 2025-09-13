package sliceutil

import (
	"iter"
	"slices"

	m "manoamaro.github.com/advent-of-code/pkg/mathx"
)

// SortFunc defines a function type for sorting elements of type T.
type SortFunc[T comparable] func(a, b T) int

// Sum calculates the sum of a slice of numbers.
func Sum[N m.Number](v []N) N {
	r := N(0)
	for _, i := range v {
		r += i
	}
	return r
}

// Map applies a function to each element of a slice and returns a new slice with the results.
func Map[T any, U any](in []T, f func(T) U) []U {
	r := make([]U, len(in))
	for i, v := range in {
		r[i] = f(v)
	}
	return r
}

// MapNotError applies a function to each element of a slice and returns a new slice with the results, ignoring errors.
func MapNotError[T any, U any](in []T, f func(T) (U, error)) []U {
	r := make([]U, 0)
	for _, v := range in {
		u, err := f(v)
		if err == nil {
			r = append(r, u)
		}
	}
	return r
}

// FlatMap applies a function to each element of a slice and flattens the results into a single slice.
func FlatMap[T any, U any](in []T, f func(T) []U) []U {
	r := make([]U, 0)
	for _, v := range in {
		r = append(r, f(v)...)
	}
	return r
}

// Fold reduces a slice to a single value using a provided function and an initial value.
func Fold[T any, U any](in []T, initial U, f func(U, T) U) U {
	r := initial
	for _, v := range in {
		r = f(r, v)
	}
	return r
}

// Reverse returns a new slice with the elements of the input slice in reverse order.
func Reverse[T any](input []T) []T {
	r := make([]T, len(input))
	for i, v := range input {
		r[len(input)-i-1] = v
	}
	return r
}

// Diff returns the elements in slice a that aren't in slice b.
func Diff[T comparable](a, b []T) []T {
	mb := make(map[T]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []T
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

// SlideSeq returns a sequence of slices of a given size from the input slice.
func SlideSeq[T any](s []T, size int) iter.Seq[[]T] {
	if size <= 0 {
		return nil
	}
	return func(yield func([]T) bool) {
		for i := 0; i < len(s)-size+1; i++ {
			if !yield(s[i : i+size]) {
				return
			}
		}
	}
}

// Combinations returns all combinations of a given size from the input slice.
func Combinations[T any](s []T, size int) [][]T {
	if size == 0 {
		return [][]T{{}}
	}
	if len(s) == 0 {
		return nil
	}
	if size == 1 {
		r := make([][]T, len(s))
		for i, v := range s {
			r[i] = []T{v}
		}
		return r
	}
	var r [][]T
	for i, v := range s {
		for _, c := range Combinations[T](s[i+1:], size-1) {
			r = append(r, append([]T{v}, c...))
		}
	}
	return r
}

// Delete removes the element at the specified index from the slice and returns the new slice.
func Delete[T any](s []T, i int) []T {
	if i < 0 || i >= len(s) {
		return s
	}
	cp := make([]T, len(s)-1)
	copy(cp, s[:i])
	copy(cp[i:], s[i+1:])
	return cp
}

// FirstFunc returns a pointer to the first element in the slice that satisfies the provided function.
func FirstFunc[T any](s []T, f func(i T) bool) *T {
	for _, v := range s {
		if f(v) {
			return &v
		}
	}
	return nil
}

// FilterFunc returns a new slice containing only the elements that satisfy the provided function.
func FilterFunc[T any](s []T, f func(i T) bool) []T {
	r := make([]T, 0, len(s))
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return slices.Clip(r)
}

// ProductFunc returns a new slice containing the results of applying a function to all pairs of elements from two slices.
func ProductFunc[T any](s1, s2 []T, f func(a, b T) T) []T {
	if len(s1) == 0 {
		return s2
	}
	totalCombinations := len(s1) * len(s2)
	combinations := make([]T, 0, totalCombinations)
	for _, v1 := range s1 {
		for _, v2 := range s2 {
			combinations = append(combinations, f(v1, v2))
		}
	}
	return combinations
}
