package internal

import (
	"fmt"
	"strconv"
)

func SumInts(v []int) int {
	r := 0
	for _, i := range v {
		r += i
	}
	return r
}

func MapToInt(in []string) []int {
	r := make([]int, len(in))
	for i, v := range in {
		r[i], _ = strconv.Atoi(v)
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

func PrintC[T any](input []T) {
	fmt.Print("[")
	for _, v := range input {
		fmt.Printf("%v,", v)
	}
	fmt.Println("]")
}
