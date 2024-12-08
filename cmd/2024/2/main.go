package main

import (
	"slices"
	"strings"

	"manoamaro.github.com/advent-of-code/pkg/aoc"
	"manoamaro.github.com/advent-of-code/pkg/collections"
	"manoamaro.github.com/advent-of-code/pkg/math2"
)

var challenge = aoc.New(2024, 2, aoc.Ints2dProcessor(strings.Fields), part1, part2)

func main() {
	challenge.Run()
}

func part1(input [][]int) int {
	valid := 0
	for _, report := range input {
		if checkReport(report) {
			valid++
		}
	}
	return valid
}

func part2(input [][]int) int {
	valid := 0
	for _, report := range input {
		if checkReport(report) {
			valid++
			continue
		}
		for i := 0; i < len(report); i++ {
			newReport := collections.Delete(report, i)
			if checkReport(newReport) {
				valid++
				break
			}
		}
	}
	return valid
}

func checkReport(report []int) bool {
	if !slices.IsSorted(report) && !slices.IsSorted(collections.Reverse(report)) {
		return false
	}
	for s := range collections.SlideSeq(report, 2) {
		if len(s) == 2 && (math2.Abs(s[0]-s[1]) <= 0 || math2.Abs(s[0]-s[1]) > 3) {
			return false
		}
	}
	return true
}
