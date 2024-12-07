package main

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"manoamaro.github.com/advent-of-code/pkg/collections"
	"manoamaro.github.com/advent-of-code/pkg/math2"
	"manoamaro.github.com/advent-of-code/pkg/strings2"
	"manoamaro.github.com/advent-of-code/pkg/utils"
)

func main() {
	input, err := utils.ReadInputLines(2024, 2)
	if err != nil {
		panic(err)
	}
	reports := make([][]int, 0)
	for _, line := range input {
		fields := strings2.MapToInt(strings.Fields(line))
		reports = append(reports, fields)
	}

	startTimePart1 := time.Now()
	part1(reports)
	fmt.Println("Part 1 took:", time.Since(startTimePart1))
	startTimePart2 := time.Now()
	part2(reports)
	fmt.Println("Part 2 took:", time.Since(startTimePart2))
}

func part1(input [][]int) {
	fmt.Println("Part 1")
	valid := 0
	for _, report := range input {
		if checkReport(report) {
			valid++
		}
	}
	fmt.Println("Valid reports:", valid)
}

func part2(input [][]int) {
	fmt.Println("Part 2")
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
	fmt.Println("Valid reports:", valid)

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
