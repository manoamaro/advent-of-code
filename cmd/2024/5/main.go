package main

import (
	"slices"
	"strings"

	"manoamaro.github.com/advent-of-code/pkg/aoc"
	"manoamaro.github.com/advent-of-code/pkg/maps2"
	"manoamaro.github.com/advent-of-code/pkg/strings2"
)

var day = aoc.New(2024, 5, parseInput, solvePt1, solvePt2)

func main() {
	day.Run()
}

type Rule struct {
	l, r int
}

type Input struct {
	rules                                   maps2.Map[Rule, bool]
	updates, sortedUpdates, unsortedUpdates [][]int
}

func sortFunc(rules maps2.Map[Rule, bool]) func(a, b int) int {
	return func(a, b int) int {
		if rules.Has(Rule{a, b}) {
			return -1
		} else if rules.Has(Rule{b, a}) {
			return 1
		}
		return 0
	}
}

func parseInput(input string) Input {
	rules := maps2.New[Rule, bool]()
	updates := [][]int{}
	readingRules := true
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			readingRules = false
			continue
		}
		if readingRules {
			r := strings2.MapToInt(strings.Split(line, "|"))
			rk := Rule{r[0], r[1]}
			rules.Set(rk, true)
		} else {
			updates = append(updates, strings2.MapToInt(strings.Split(line, ",")))
		}
	}
	sortedUpdates := make([][]int, len(updates))
	unsortedUpdates := make([][]int, len(updates))

	sFunc := sortFunc(rules)

	for i, update := range updates {
		if slices.IsSortedFunc(update, sFunc) {
			sortedUpdates[i] = update
		} else {
			unsortedUpdates[i] = update
		}
	}

	return Input{rules, updates, sortedUpdates, unsortedUpdates}
}

func solvePt1(input Input) int {
	sum := 0
	for _, update := range input.sortedUpdates {
		if len(update) > 0 {
			sum += update[len(update)/2]
		}
	}
	return sum
}

func solvePt2(input Input) int {
	sum := 0
	sFunc := sortFunc(input.rules)
	for _, update := range input.unsortedUpdates {
		if len(update) == 0 {
			continue
		}
		slices.SortFunc(update, sFunc)
		sum += update[len(update)/2]
	}
	return sum
}
