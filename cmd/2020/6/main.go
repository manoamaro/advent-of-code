package main

import (
	"strings"

	"manoamaro.github.com/advent-of-code/pkg/aoc"
)

var challenge = aoc.New(2020, 6, parseInput, part1, part2)

func main() {
	challenge.Run()
}

func parseInput(input string) []string {
	return strings.Split(input, "\n\n")
}

func part1(input []string) int {
	countAnyone := 0
	for _, group := range input {
		eachPeople := strings.Split(group, "\n")
		allTogether := strings.Join(eachPeople, "")

		answers := make(map[string]int)

		for _, answer := range allTogether {
			answers[string(answer)]++
		}

		countAnyone += len(answers)
	}
	return countAnyone
}

func part2(input []string) int {
	countEveryone := 0
	for _, group := range input {
		eachPeople := strings.Split(group, "\n")
		allTogether := strings.Join(eachPeople, "")

		answers := make(map[string]int)

		for _, answer := range allTogether {
			answers[string(answer)]++
		}

		for _, count := range answers {
			if count == len(eachPeople) {
				countEveryone++
			}
		}
	}
	return countEveryone
}
