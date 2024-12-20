package main

import (
	"fmt"
	"strings"

	"manoamaro.github.com/advent-of-code/pkg/aoc"
	"manoamaro.github.com/advent-of-code/pkg/grid"
	"manoamaro.github.com/advent-of-code/pkg/strings2"
)

func main() {
	challenge := aoc.New(2024, 18, aoc.LinesProcessor(), part1, part2)
	challenge.TestPart1(aoc.GetSourceFilePath("test_input.txt"), 22)
	challenge.TestPart1(aoc.GetSourceFilePath("test_input.txt"), 0)
	//challenge.Run()
}

func fallBytes(input []string, amount int) grid.Grid[bool] {
	g := grid.New[bool](7, 7)
	for _, line := range input[:amount] {
		coords := strings.Split(line, ",")
		x, y := strings2.Atoi[int](coords[0]), strings2.Atoi[int](coords[1])
		g.Set(y, x, true)
	}
	return g
}

func part1(input []string) int {
	memory := fallBytes(input, 12)
	fmt.Println(memory)
	return 0
}

func part2(input []string) int {
	return 0
}
