package collections

import (
	"fmt"
	"iter"

	m "manoamaro.github.com/advent-of-code/pkg/math2"
)

type SortFunc[T comparable] func(a, b T) int

func Sum[N m.Number](v []N) N {
	r := N(0)
	for _, i := range v {
		r += i
	}
	return r
}

func Map[T any, U any](in []T, f func(T) U) []U {
	r := make([]U, len(in))
	for i, v := range in {
		r[i] = f(v)
	}
	return r
}

func MapSeq[T any, U any](in []T, f func(T) U) iter.Seq[U] {
	return func(yield func(U) bool) {
		for _, v := range in {
			if !yield(f(v)) {
				return
			}
		}
	}
}

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

func FlatMap[T any, U any](in []T, f func(T) []U) []U {
	r := make([]U, 0)
	for _, v := range in {
		r = append(r, f(v)...)
	}
	return r
}

func Fold[T any, U any](in []T, initial U, f func(U, T) U) U {
	r := initial
	for _, v := range in {
		r = f(r, v)
	}
	return r
}

func Reverse[T any](input []T) []T {
	r := make([]T, len(input))
	for i, v := range input {
		r[len(input)-i-1] = v
	}
	return r
}

func PrintSlice[T any](input []T) {
	fmt.Print("[")
	for _, v := range input {
		fmt.Printf("%v,", v)
	}
	fmt.Println("]")
}

func Count[T comparable](input []T, value T) int {
	r := 0
	for _, v := range input {
		if v == value {
			r++
		}
	}
	return r
}

// Diff returns the elements in a that aren't in b.
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
		for _, c := range Combinations(s[i+1:], size-1) {
			r = append(r, append([]T{v}, c...))
		}
	}
	return r
}

func Delete[T any](s []T, i int) []T {
	if i < 0 || i >= len(s) {
		return s
	}
	cp := make([]T, len(s)-1)
	copy(cp, s[:i])
	copy(cp[i:], s[i+1:])
	return cp
}
