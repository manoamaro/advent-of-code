package aoc

import (
	"strings"

	"manoamaro.github.com/advent-of-code/pkg/strings2"
)

type Splitter[T comparable] func(string) []T

type InputProcessor[T any] func(string) T

func NoOpProcessor() InputProcessor[string] {
	return func(input string) string {
		return input
	}
}

func LinesProcessor() InputProcessor[[]string] {
	return func(input string) []string {
		return strings.Split(input, "\n")
	}
}

func GridStringProcessor() InputProcessor[[][]string] {
	return func(input string) [][]string {
		lines := strings.Split(input, "\n")
		var slice [][]string
		for _, line := range lines {
			slice = append(slice, strings.Split(line, ""))
		}
		return slice
	}
}

func IntsProcessor() InputProcessor[[]int] {
	return func(input string) []int {
		lines := strings.Split(input, "\n")
		return strings2.MapToInt(lines)
	}
}

func Ints2dProcessor(splitter Splitter[string]) InputProcessor[[][]int] {
	return func(input string) [][]int {
		lines := strings.Split(input, "\n")
		var slice [][]int
		for _, line := range lines {
			slice = append(slice, strings2.MapToInt(splitter(line)))
		}
		return slice
	}
}
