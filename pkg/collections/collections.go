package collections

import (
	"fmt"
	"iter"
	"strconv"
	"strings"

	m "manoamaro.github.com/advent-of-code/pkg/math2"
)

func Sum[N m.Number](v []N) N {
	r := N(0)
	for _, i := range v {
		r += i
	}
	return r
}

func MapToInt(in []string) []int {
	r := make([]int, 0)
	for _, v := range in {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}
		_r, err := strconv.Atoi(v)
		if err == nil {
			r = append(r, _r)
		}
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

func Slide[T any](s []T, size int) iter.Seq[[]T] {
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

func Delete[T any](s []T, i int) []T {
	if i < 0 || i >= len(s) {
		return s
	}
	cp := make([]T, len(s)-1)
	copy(cp, s[:i])
	copy(cp[i:], s[i+1:])
	return cp
}
