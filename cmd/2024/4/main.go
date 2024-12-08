package main

import (
	"manoamaro.github.com/advent-of-code/pkg/aoc"
)

var challenge = aoc.New(2024, 4, aoc.GridStringProcessor(), part1, part2)

func main() {
	challenge.Run()
}

func part1(grid [][]string) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			for _, direction := range directions {
				if findXmas(grid, i, j, direction) {
					count++
				}
			}
		}
	}
	return count
}

func part2(grid [][]string) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if findX_Mas(grid, i, j) {
				count++
			}
		}
	}
	return count
}

var directions = [][]int{
	{0, 1},   // Horizontal right
	{0, -1},  // Horizontal left
	{1, 0},   // Vertical down
	{-1, 0},  // Vertical up
	{1, 1},   // Diagonal down-right
	{-1, -1}, // Diagonal up-left
	{1, -1},  // Diagonal down-left
	{-1, 1},  // Diagonal up-right
}

var xmas = []string{"X", "M", "A", "S"}

func findXmas(grid [][]string, x, y int, direction []int) bool {
	dx, dy := direction[0], direction[1]
	for i := 0; i < len(xmas); i++ {
		r, c := x+i*dx, y+i*dy
		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] != xmas[i] {
			return false
		}
	}
	return true
}

func findX_Mas(grid [][]string, x, y int) bool {
	defer func() bool {
		if r := recover(); r != nil {
		}
		return false
	}()

	p1 := append([]string{}, grid[x][y], grid[x+1][y+1], grid[x+2][y+2])
	p2 := append([]string{}, grid[x][y+2], grid[x+1][y+2-1], grid[x+2][y+2-2])

	if (p1[0] == "M" && p1[1] == "A" && p1[2] == "S") || (p1[0] == "S" && p1[1] == "A" && p1[2] == "M") {
		if (p2[0] == "M" && p2[1] == "A" && p2[2] == "S") || (p2[0] == "S" && p2[1] == "A" && p2[2] == "M") {
			return true
		}
	}
	return false
}
