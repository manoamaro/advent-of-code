package main

import (
	"manoamaro.github.com/advent-of-code/pkg/aoc"
	"manoamaro.github.com/advent-of-code/pkg/fn"
	"manoamaro.github.com/advent-of-code/pkg/grid"
	"manoamaro.github.com/advent-of-code/pkg/maps2"
	"manoamaro.github.com/advent-of-code/pkg/math2"
	"manoamaro.github.com/advent-of-code/pkg/queue"
	"manoamaro.github.com/advent-of-code/pkg/set"
)

var minPicSecSaved = 50

func main() {
	challenge := aoc.New(2024, 20, aoc.GridProcessor(func(r rune) rune { return r }), part1, part2)
	minPicSecSaved = 1
	challenge.TestPart1(aoc.GetSourceFilePath("test_input.txt"), 44)
	minPicSecSaved = 50
	challenge.TestPart2(aoc.GetSourceFilePath("test_input.txt"), 285)
	minPicSecSaved = 100
	challenge.Run()
}

func part1(input grid.Grid[rune]) int {
	start := *input.FindFunc(fn.Eq('S'))

	distances := findDistances(input, start)

	count := 0
	dirs := []grid.Dir{{2, 0}, {1, 1}, {0, 2}, {-1, 1}}
	for i := 0; i < input.Rows(); i++ {
		for j := 0; j < input.Cols(); j++ {
			cell := grid.NewCell(i, j)
			if *input.GetCell(cell) == '#' {
				continue
			}
			for _, d := range dirs {
				next := cell.Move(d)
				if !input.InBounds(next) {
					continue
				}
				if *input.Get(i+d[0], j+d[1]) == '#' {
					continue
				}
				if math2.Abs(*distances.GetCell(cell)-*distances.GetCell(next)) >= minPicSecSaved+2 {
					count += 1
				}
			}
		}
	}
	return count
}

func part2(input grid.Grid[rune]) int {
	start := *input.FindFunc(fn.Eq('S'))
	end := *input.FindFunc(fn.Eq('E'))
	path := maps2.New[grid.Cell, int]()
	path.Set(start, 0)
	pos := start
	for pos != end {
		for n, nv := range input.Neighbors(pos) {
			if *nv == '#' {
				continue
			}
			if path.Has(n) {
				continue
			}
			path.Set(n, *path.GetOrPanic(pos)+1)
			pos = n
			break
		}
	}

	count := 0
	seen := set.New[[2]grid.Cell]()
	for cell, value := range path {
		for cell2, value2 := range path {
			if cell == cell2 {
				continue
			}
			if seen.Contains([2]grid.Cell{cell, cell2}) {
				continue
			}
			dist := math2.ManhattanDistance(cell, cell2)
			if dist > 20 {
				continue
			}
			seen.Add([2]grid.Cell{cell, cell2})
			seen.Add([2]grid.Cell{cell2, cell})
			if math2.Abs(value-value2) >= minPicSecSaved+dist {
				count += 1
			}
		}
	}

	return count
}

type v struct {
	cell grid.Cell
	dist int
}

// findDistances calculates the shortest distances from the start cell to all other cells in the grid.
// It returns a grid of distances
func findDistances(input grid.Grid[rune], start grid.Cell) grid.Grid[int] {
	distances := grid.New[int](input.Rows(), input.Cols())
	distances.Fill(-1)
	distances.SetCell(start, 0)
	q := queue.New(start)
	for cell := range q.Seq() {
		for n := range input.Neighbors(cell) {
			if *input.GetCell(n) == '#' {
				continue
			}
			if *distances.GetCell(n) != -1 {
				continue
			}
			distances.SetCell(n, *distances.GetCell(cell)+1)
			q.Push(n)
		}
	}
	return distances
}
