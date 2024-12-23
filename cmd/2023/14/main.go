package main

import (
	"fmt"
	"strings"

	"manoamaro.github.com/advent-of-code/pkg/aoc"
	"manoamaro.github.com/advent-of-code/pkg/grid"
)

var challenge = aoc.New(2023, 14, parseInput, part1, part2)

func main() {
	challenge.Run()
}

func parseInput(input string) grid.Grid[rock] {
	lines := strings.Split(input, "\n")
	grid := grid.New[rock](len(lines), len(lines[0]))
	for j := range lines {
		for i := range lines[j] {
			grid[i][j] = rock(lines[i][j])
		}
	}
	return grid
}

type rock string

const (
	none  rock = "."
	round rock = "O"
	cube  rock = "#"
)

func part1(grid grid.Grid[rock]) int {
	total := 0
	v := len(grid)
	for j := range grid[0] {
		free := 0
		for i := range v {
			switch grid[i][j] {
			case none:
				free++
			case round:
				total += v - i + free
			case cube:
				free = 0
			}
		}
	}
	return total
}

func part2(grid grid.Grid[rock]) int {
	total := 0
	rows := len(grid)
	cols := len(grid[0])

	//visited := map[[1][1]string]bool{}
	// tilt the grid north
	for c := range cols {
		free := 0
		for r := range rows {
			switch grid[r][c] {
			case none:
				free++
			case round:
				if free > 0 {
					grid[r-free][c] = round
					grid[r][c] = none
				}
			case cube:
				free = 0
			}
		}
	}
	for _, row := range grid {
		fmt.Println(row)
	}

	fmt.Println("")

	// tilt the grid west
	for r := range rows {
		free := 0
		for c := range cols {
			switch grid[r][c] {
			case none:
				free++
			case round:
				if free > 0 {
					grid[r][c-free] = round
					grid[r][c] = none
				}
			case cube:
				free = 0
			}
		}
	}

	for _, row := range grid {
		fmt.Println(row)
	}

	fmt.Println("")

	// tilt the grid south
	for c := cols - 1; c >= 0; c-- {
		free := 0
		for r := rows - 1; r >= 0; r-- {
			switch grid[r][c] {
			case none:
				free++
			case round:
				if free > 0 {
					grid[r+free][c] = round
					grid[r][c] = none
				}
			case cube:
				free = 0
			}
		}
	}

	for _, row := range grid {
		fmt.Println(row)
	}

	// tilt the grid east
	for r := rows - 1; r >= 0; r-- {
		free := 0
		for c := cols - 1; c >= 0; c-- {
			switch grid[r][c] {
			case none:
				free++
			case round:
				if free > 0 {
					grid[r][c+free] = round
					grid[r][c] = none
				}
			case cube:
				free = 0
			}
		}
	}

	fmt.Println("")

	return total
}
