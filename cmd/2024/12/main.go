package main

import (
	"manoamaro.github.com/advent-of-code/pkg/aoc"
	"manoamaro.github.com/advent-of-code/pkg/collections"
	"manoamaro.github.com/advent-of-code/pkg/grid"
	"manoamaro.github.com/advent-of-code/pkg/queue"
	"manoamaro.github.com/advent-of-code/pkg/set"
)

var challenge = aoc.New(2024, 12, aoc.RuneGridProcessor, part1, part2)

func main() {
	challenge.Run()
}

func part1(input grid.Grid[rune]) int {
	fields := findFields(input)
	// calculate perimeter of each field
	perimeters := []int{}
	for _, field := range fields {
		value := input.Get(field[0][0], field[0][1])
		perimeter := 0
		for _, cell := range field {
			for _, nValue := range input.Neighbors(cell) {
				if nValue == nil || *value != *nValue {
					perimeter++
				}
			}
		}
		perimeter *= len(field)
		perimeters = append(perimeters, perimeter)
	}
	return collections.Sum(perimeters)
}

func part2(input grid.Grid[rune]) int {
	fields := findFields(input)
	totalPerimeter := 0
	for _, field := range fields {
		perimeter := 0
		for _, cell := range field {
			value := input.Get(cell[0], cell[1])
			// check each neighbor of the current cell (left, right, up, down)
			for _, dir := range [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} {
				ccwDir := []int{-dir[1], dir[0]}
				// Get values of the corner cells:
				// c = cell, n = neighbor, 1 = ccw neighbor, 2 = corner cell
				// .....
				// .....
				// .nc..
				// .21..
				nVal := input.Get(cell[0]+dir[0], cell[1]+dir[1])
				ccwVal := input.Get(cell[0]+ccwDir[0], cell[1]+ccwDir[1])
				cornerVal := input.Get(cell[0]+dir[0]+ccwDir[0], cell[1]+dir[1]+ccwDir[1])

				// check if the current cell is a convex corner
				if (ccwVal == nil || *ccwVal != *value) && (nVal == nil || *nVal != *value) {
					perimeter++
					continue
				}

				// check if the current cell is a concave corner
				if ccwVal != nil && nVal != nil && *value == *ccwVal && *value == *nVal && *value != *cornerVal {
					perimeter++
					continue
				}
			}
		}
		perimeter *= len(field)
		totalPerimeter += perimeter
	}

	return totalPerimeter
}

// findFields identifies and returns all distinct fields of connected cells in the given grid.
// Each field is represented as a slice of grid.Cell, and the function returns a slice of these fields.
//
// Parameters:
//   - input: A grid.Grid[rune] representing the input grid.
//
// Returns:
//   - A slice of slices of grid.Cell, where each inner slice represents a distinct field of connected cells.
func findFields(input grid.Grid[rune]) [][]grid.Cell {
	cells := set.New(input.Cells()...)
	fields := [][]grid.Cell{}
	for cells.Len() > 0 {
		cell := *cells.First()
		field := exploreField(input, cell)
		fields = append(fields, field)
		cells.RemoveAll(field)
	}
	return fields
}

// exploreField explores the given grid starting from the specified cell.
// It returns a slice of cells that have the same value as the starting cell.
//
// Parameters:
//   - input: A grid.Grid[rune] representing the grid to explore.
//   - start: A grid.Cell representing the starting cell for exploration.
//
// Returns:
//   - A slice of grid.Cell containing all cells with the same value as the starting cell.
func exploreField(input grid.Grid[rune], start grid.Cell) []grid.Cell {
	value := input.Get(start[0], start[1])
	seen := set.New[grid.Cell](start)
	q := queue.New[grid.Cell](start)
	for n := range q.Seq() {
		for nCell, nValue := range input.Neighbors(n) {
			if nValue != nil && *value == *nValue {
				if seen.Contains(nCell) {
					q.Push(nCell)
					seen.Add(nCell)
				}
			}
		}
	}
	return seen.Slice()
}
