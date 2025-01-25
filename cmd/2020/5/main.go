package main

import (
	"math"
	"sort"

	"manoamaro.github.com/advent-of-code/pkg/aoc"
)

var challenge = aoc.New(2020, 5, aoc.LinesProcessor(), part1, part2)

func main() {
	challenge.Run()
}

func part1(input []string) int {
	maxId := 0
	for _, line := range input {
		row, col := processPass(line, 0, 127, 0, 7)
		id := (row * 8) + col
		if id > maxId {
			maxId = id
		}
	}
	return maxId
}

func part2(input []string) int {
	seatsIds := make([]int, 0)
	for _, line := range input {
		row, col := processPass(line, 0, 127, 0, 7)
		id := (row * 8) + col
		seatsIds = append(seatsIds, id)
	}
	sort.Ints(seatsIds)
	return findYourSeatID(seatsIds)
}

func processPass(input string, minRow int, maxRow int, minCol int, maxCol int) (row int, col int) {
	if len(input) > 0 {
		next := string(input[0])
		nextMinRow, nextMaxRow, nextMinCol, nextMaxCol := minRow, maxRow, minCol, maxCol

		switch next {
		case "F":
			nextMaxRow = nextMaxRow - int(math.Ceil(float64(maxRow-minRow)/2))

		case "B":
			nextMinRow = nextMinRow + int(math.Ceil(float64(maxRow-minRow)/2))

		case "R":
			nextMinCol = nextMinCol + int(math.Ceil(float64(maxCol-minCol)/2))

		case "L":
			nextMaxCol = nextMaxCol - int(math.Ceil(float64(maxCol-minCol)/2))

		}
		return processPass(input[1:], nextMinRow, nextMaxRow, nextMinCol, nextMaxCol)
	} else {
		return minRow, minCol
	}
}

func findYourSeatID(seats []int) int {
	for i := 1; i < len(seats)-1; i++ {
		last := seats[i]
		next := seats[i+1]
		if next-last == 2 {
			return last + 1
		}
	}
	return -1
}
