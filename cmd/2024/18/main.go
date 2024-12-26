package main

import (
	"strings"

	"manoamaro.github.com/advent-of-code/pkg/aoc"
	"manoamaro.github.com/advent-of-code/pkg/grid"
	"manoamaro.github.com/advent-of-code/pkg/queue"
	"manoamaro.github.com/advent-of-code/pkg/set"
	"manoamaro.github.com/advent-of-code/pkg/strings2"
)

var size = 6
var fallen = 12

func main() {
	challenge := aoc.New(2024, 18, aoc.LinesProcessor(), part1, part2)
	challenge.TestPart1(aoc.GetSourceFilePath("test_input.txt"), [2]int{22})
	challenge.TestPart2(aoc.GetSourceFilePath("test_input.txt"), [2]int{6, 1})
	size = 70
	fallen = 1024
	challenge.Run()
}

func fallBytes(input []string, amount int) grid.Grid[bool] {
	g := grid.New[bool](size+1, size+1)
	for _, line := range input[:amount] {
		coords := strings.Split(line, ",")
		x, y := strings2.Atoi[int](coords[0]), strings2.Atoi[int](coords[1])
		g.Set(y, x, true)
	}
	return g
}

type value struct {
	cell grid.Cell
	dist int
}

func findPath(memory grid.Grid[bool], start grid.Cell, end grid.Cell) int {
	q := queue.New(value{cell: start, dist: 0})
	seen := set.New(start)

	for v := range q.Seq() {
		for next, nextValue := range memory.Neighbors(v.cell) {
			if *nextValue {
				continue
			}
			if seen.Contains(next) {
				continue
			}
			if next == end {
				return v.dist + 1
			}
			seen.Add(next)
			q.Push(value{cell: next, dist: v.dist + 1})
		}
	}
	return 0
}

func part1(input []string) [2]int {
	memory := fallBytes(input, fallen)
	start := grid.Cell{0, 0}
	end := grid.Cell{size, size}
	return [2]int{findPath(memory, start, end)}
}

func part2(input []string) [2]int {
	start := grid.Cell{0, 0}
	end := grid.Cell{size, size}

	// Binary search for the first byte position that blocks the path
	lo := 0
	hi := len(input) - 1
	for lo < hi {
		mid := (lo + hi) / 2
		memory := fallBytes(input, mid+1)
		if dist := findPath(memory, start, end); dist > 0 {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return [2]int(strings2.MapToInt(strings.Split(input[lo], ",")))
}
