package aoc

import (
	"strings"

	"manoamaro.github.com/advent-of-code/pkg/strings2"
)

type InputProcessor[T any] func(string) (T, error)

func NoOpProcessor() InputProcessor[string] {
	return func(input string) (string, error) {
		return input, nil
	}
}

func LinesProcessor() InputProcessor[[]string] {
	return func(input string) ([]string, error) {
		return strings.Split(input, "\n"), nil
	}
}

func MatrixProcessor() InputProcessor[[][]string] {
	return func(input string) ([][]string, error) {
		lines := strings.Split(input, "\n")
		var slice [][]string
		for _, line := range lines {
			slice = append(slice, strings.Split(line, ""))
		}
		return slice, nil
	}
}

func IntsProcessor() InputProcessor[[]int] {
	return func(input string) ([]int, error) {
		lines := strings.Split(input, "\n")
		return strings2.MapToInt(lines), nil
	}
}

func Ints2dProcessor(sep string) InputProcessor[[][]int] {
	return func(input string) ([][]int, error) {
		lines := strings.Split(input, "\n")
		var slice [][]int
		for _, line := range lines {
			slice = append(slice, strings2.MapToInt(strings.Split(line, sep)))
		}
		return slice, nil
	}
}
