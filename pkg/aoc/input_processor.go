package aoc

import (
	"manoamaro.github.com/advent-of-code/pkg/collections"
	"strings"

	"manoamaro.github.com/advent-of-code/pkg/grid"
	"manoamaro.github.com/advent-of-code/pkg/strings2"
)

type Splitter[T comparable] func(string) []T

type InputProcessor[T any] func(string) T

func StringProcessor(input string) string {
	return strings.TrimSpace(input)
}

func LinesProcessor() InputProcessor[[]string] {
	return func(input string) []string {
		return strings.Split(input, "\n")
	}
}

func GridProcessor[T comparable](f func(rune) T) InputProcessor[grid.Grid[T]] {
	return func(input string) grid.Grid[T] {
		lines := strings.Split(input, "\n")
		var g grid.Grid[T]
		for _, line := range lines {
			g = append(g, collections.Map([]rune(line), f))
		}
		return g
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

func RuneGridProcessor(input string) grid.Grid[rune] {
	lines := strings.Split(input, "\n")
	g := grid.New[rune](len(lines), len(lines[0]))
	for r, line := range lines {
		for c, char := range line {
			g.Set(r, c, char)
		}
	}
	return g
}

func IntGridProcessor(input string) grid.Grid[int] {
	lines := strings.Split(input, "\n")
	var grid grid.Grid[int]
	for _, line := range lines {
		line = strings.TrimSpace(line)
		grid = append(grid, strings2.MapCharsToInts([]rune(line)))
	}
	return grid
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
