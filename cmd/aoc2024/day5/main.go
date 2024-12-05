package main

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"manoamaro.github.com/advent-of-code/pkg/strings2"
	"manoamaro.github.com/advent-of-code/pkg/utils"
)

func main() {
	input, err := utils.ReadInputLines(2024, 5)
	if err != nil {
		panic(err)
	}

	rules := [][]int{}
	updates := [][]int{}
	readingRules := true
	for _, line := range input {
		if line == "" {
			readingRules = false
			continue
		}
		if readingRules {
			rules = append(rules, strings2.MapToInt(strings.Split(line, "|")))
		} else {
			updates = append(updates, strings2.MapToInt(strings.Split(line, ",")))
		}
	}

	startTimePart1 := time.Now()
	part1(rules, updates)
	fmt.Println("Part 1 took:", time.Since(startTimePart1))
	startTimePart2 := time.Now()
	part2(rules, updates)
	fmt.Println("Part 2 took:", time.Since(startTimePart2))
}

func sortFunc(rules [][]int) func(a, b int) int {
	return func(a, b int) int {
		for _, r := range rules {
			if r[0] == a && r[1] == b {
				return -1
			} else if r[0] == b && r[1] == a {
				return 1
			}
		}
		return 0
	}
}

func part1(rules [][]int, updates [][]int) {
	sum := 0
	for _, update := range updates {
		if slices.IsSortedFunc(update, sortFunc(rules)) {
			sum += update[len(update)/2]
		}
	}
	fmt.Println("Part 1:", sum)
}

func part2(rules [][]int, updates [][]int) {
	sum := 0
	for _, update := range updates {
		if !slices.IsSortedFunc(update, sortFunc(rules)) {
			slices.SortFunc(update, sortFunc(rules))
			sum += update[len(update)/2]
		}
	}
	fmt.Println("Part 2:", sum)
}
