package main

import (
	"strings"

	"manoamaro.github.com/advent-of-code/pkg/aoc"
	"manoamaro.github.com/advent-of-code/pkg/grid"
)

var day = aoc.New(2024, 6, parseInput, part1, part2)

func main() {
	day.Run()
}

type Dir [2]int

var (
	up    = Dir{-1, 0}
	down  = Dir{1, 0}
	left  = Dir{0, -1}
	right = Dir{0, 1}
)

type Cell [2]int

func (c Cell) move(dir Dir) Cell {
	return Cell{c[0] + dir[0], c[1] + dir[1]}
}

type Guard struct {
	cell Cell
	dir  Dir
}

func parseInput(input string) (grid.Grid[byte], error) {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	grid := make(grid.Grid[byte], len(lines))
	for i, line := range lines {
		grid[i] = make([]byte, len(line))
		for j, char := range line {
			grid[i][j] = byte(char)
		}
	}
	return grid, nil
}

func findGuard(grid grid.Grid[byte]) Guard {
	for i, row := range grid {
		for j, cell := range row {
			if cell == '^' {
				return Guard{cell: Cell{i, j}, dir: up}
			}
		}
	}
	panic("guard not found")
}

func hasObstacleInFront(grid grid.Grid[byte], guard Guard) bool {
	nextCell := guard.cell.move(guard.dir)
	v := grid.Get(nextCell[0], nextCell[1])
	return v != nil && *v == '#'
}

func part1(grid grid.Grid[byte]) (int, error) {
	guard := findGuard(grid)
	visited := make(map[Cell]bool)
	for {
		if hasObstacleInFront(grid, guard) {
			// turn right 90 degrees
			switch guard.dir {
			case up:
				guard.dir = right
			case down:
				guard.dir = left
			case left:
				guard.dir = up
			case right:
				guard.dir = down
			}
			continue
		}
		// move forward
		guard.cell = guard.cell.move(guard.dir)
		if grid.Get(guard.cell[0], guard.cell[1]) == nil {
			break
		}
		visited[guard.cell] = true
	}
	return len(visited), nil
}

func part2(grid grid.Grid[byte]) (int, error) {
	count := 0
	for cell := range grid.ItemsSeq() {
		if *cell == '^' {
			continue
		}
		prev := *cell
		*cell = '#'

		guard := findGuard(grid)
		visited := make(map[Guard]bool)
		for {
			if hasObstacleInFront(grid, guard) {
				// turn right 90 degrees
				switch guard.dir {
				case up:
					guard.dir = right
				case down:
					guard.dir = left
				case left:
					guard.dir = up
				case right:
					guard.dir = down
				}
				continue
			}
			// move forward
			guard.cell = guard.cell.move(guard.dir)
			if grid.Get(guard.cell[0], guard.cell[1]) == nil {
				break
			}
			p := Guard{cell: guard.cell, dir: guard.dir}

			if visited[p] {
				count++
				break
			}
			visited[p] = true
		}
		*cell = prev
	}
	return count, nil
}
