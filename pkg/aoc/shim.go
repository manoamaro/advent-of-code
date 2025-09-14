package aoc

import (
    iaoc "manoamaro.github.com/advent-of-code/internal/aoc"
    "manoamaro.github.com/advent-of-code/pkg/grid"
)

// Type aliases
type Solver[T any, R comparable] = iaoc.Solver[T, R]
type Challenge[T any, R comparable] = iaoc.Challenge[T, R]
type InputProcessor[T any] = iaoc.InputProcessor[T]

// Constructors
func New[T any, R comparable](year, day int, inputProcessor InputProcessor[T], part1Solver Solver[T, R], part2Solver Solver[T, R]) *Challenge[T, R] {
    return iaoc.New(year, day, inputProcessor, part1Solver, part2Solver)
}

// Input processors
func StringProcessor(input string) string { return iaoc.StringProcessor(input) }
func LinesProcessor() InputProcessor[[]string] { return iaoc.LinesProcessor() }
func IntLinesProcessor(input string) []int { return iaoc.IntLinesProcessor(input) }
func GridProcessor[T comparable](f func(rune) T) InputProcessor[grid.Grid[T]] { return iaoc.GridProcessor(f) }
func GridStringProcessor() InputProcessor[[][]string] { return iaoc.GridStringProcessor() }
func RuneGridProcessor(input string) grid.Grid[rune] { return iaoc.RuneGridProcessor(input) }
func IntGridProcessor(input string) grid.Grid[int] { return iaoc.IntGridProcessor(input) }
func Ints2dProcessor(splitter func(string) []string) InputProcessor[[][]int] { return iaoc.Ints2dProcessor(splitter) }

// Utils
func GetSourceFilePath(file string) string { return iaoc.GetSourceFilePath(file) }
func GetInput(year int, day int) string { return iaoc.GetInput(year, day) }

