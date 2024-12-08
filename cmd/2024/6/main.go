package main

import (
	"strings"
	"sync"

	"manoamaro.github.com/advent-of-code/pkg/aoc"
	"manoamaro.github.com/advent-of-code/pkg/grid"
)

var day = aoc.New(2024, 6, parseInput, part1, part2)

func main() {
	day.Run()
}

type Dir [3]int

var (
	up    = Dir{-1, 0, 0}
	down  = Dir{1, 0, 1}
	left  = Dir{0, -1, 2}
	right = Dir{0, 1, 3}
)

type Cell [2]int

func (c Cell) move(dir Dir) Cell {
	return Cell{c[0] + dir[0], c[1] + dir[1]}
}

type guard struct {
	cell Cell
	dir  Dir
}

func parseInput(input string) grid.Grid[byte] {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	grid := make(grid.Grid[byte], len(lines))
	for i, line := range lines {
		grid[i] = make([]byte, len(line))
		for j, char := range line {
			grid[i][j] = byte(char)
		}
	}
	return grid
}

func findGuard(grid grid.Grid[byte]) guard {
	for i, row := range grid {
		for j, cell := range row {
			if cell == '^' {
				return guard{cell: Cell{i, j}, dir: up}
			}
		}
	}
	panic("guard not found")
}

func hasObstacleInFront(grid grid.Grid[byte], guard guard) bool {
	nextCell := guard.cell.move(guard.dir)
	v := grid.Get(nextCell[0], nextCell[1])
	return v != nil && *v == '#'
}

func part1(grid grid.Grid[byte]) int {
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
	return len(visited)
}

func findLoop(grid grid.Grid[byte]) bool {
	guard := findGuard(grid)
	visited := make([]bool, grid.Rows()*grid.Cols()*4)
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
		h := (guard.cell[0]*grid.Rows()+guard.cell[1])*4 + guard.dir[2]
		if visited[h] {
			return true
		}
		visited[h] = true
	}
	return false
}

func part2(grid grid.Grid[byte]) int {
	var wg sync.WaitGroup
	count := make(chan int, 17190)
	for i := 0; i < grid.Rows(); i++ {
		for j := 0; j < grid.Cols(); j++ {
			if grid[i][j] == '^' {
				continue
			}
			g := grid.Copy()
			wg.Add(1)
			go func() {
				defer wg.Done()
				g[i][j] = '#'
				if findLoop(g) {
					count <- 1
				}
			}()
		}
	}
	go func() {
		wg.Wait()
		close(count)
	}()
	c := 0
	for v := range count {
		c += v
	}
	return c
}
