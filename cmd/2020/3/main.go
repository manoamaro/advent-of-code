package main

import (
	"manoamaro.github.com/advent-of-code/pkg/aoc"
)

var challenge = aoc.New(2020, 3, aoc.LinesProcessor(), part1, part2)

func main() {
	challenge.Run()
}

func part1(input []string) int {
	return countTrees(input, 3, 1)
}

func part2(lines []string) int {
	slope1Trees := countTrees(lines, 1, 1)
	slope2Trees := countTrees(lines, 3, 1)
	slope3Trees := countTrees(lines, 5, 1)
	slope4Trees := countTrees(lines, 7, 1)
	slope5Trees := countTrees(lines, 1, 2)

	return slope1Trees * slope2Trees * slope3Trees * slope4Trees * slope5Trees
}

func countTrees(input []string, jumpRight int, jumpDown int) (result int) {
	for i := 0; i < len(input); i += jumpDown {
		line := input[i]
		lineIdx := (i / jumpDown) * jumpRight
		realLineIdx := lineIdx % len(line)
		pos := string(line[realLineIdx])
		if pos == "#" {
			result++
		}
	}
	return
}
