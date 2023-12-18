package internal

import (
	"fmt"
	"strconv"
	"strings"
)

func Sum[N Number](v []N) N {
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
	r := input
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
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