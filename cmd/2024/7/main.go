package main

import (
	"strings"

	"manoamaro.github.com/advent-of-code/pkg/aoc"
	"manoamaro.github.com/advent-of-code/pkg/strutil"
)

var challenge = aoc.New(2024, 7, parseInput, part1, part2)

func parseInput(input string) [][]int {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	o := make([][]int, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, ":")
		v := strutil.Atoi[int](parts[0])
		v2 := strutil.MapToInt(strings.Fields(parts[1]))
		eq := make([]int, 1)
		eq[0] = v
		eq = append(eq, v2...)
		o[i] = eq
	}
	return o
}

func calculatePart1(expected int, curr int, nums []int) bool {
	if len(nums) == 0 {
		return expected == curr
	}
	return calculatePart1(expected, curr+nums[0], nums[1:]) || calculatePart1(expected, curr*nums[0], nums[1:])
}

func part1(input [][]int) int {
	total := 0
	for _, eq := range input {
		if calculatePart1(eq[0], eq[1], eq[2:]) {
			total += eq[0]
		}
	}
	return total
}

// concat: Concatenates two integers e.g. concat(1,2) = 12
func concat(a, b int) int {
	p := 1
	for p <= b {
		p *= 10
	}
	return a*p + b
}

func calculatePart2(expected int, curr int, nums []int) bool {
	if len(nums) == 0 {
		return expected == curr
	}
	return calculatePart2(expected, curr+nums[0], nums[1:]) || calculatePart2(expected, curr*nums[0], nums[1:]) || calculatePart2(expected, concat(curr, nums[0]), nums[1:])
}

func part2(input [][]int) int {
	total := 0
	for _, eq := range input {
		if calculatePart2(eq[0], eq[1], eq[2:]) {
			total += eq[0]
		}
	}
	return total
}

func main() {
	challenge.Run()
}
