package main

import (
	"fmt"
	"strings"

	"manoamaro.github.com/advent-of-code/pkg/aoc"
	"manoamaro.github.com/advent-of-code/pkg/mathx"
	"manoamaro.github.com/advent-of-code/pkg/strutil"
)

var challenge = aoc.New(2023, 12, aoc.LinesProcessor(), part1, part2)

func main() {
	challenge.Run()
}

func part1(input []string) int {
	count := 0

	for _, line := range input {
		line, groups := parseLine(line)
		count += calculate(line, groups, map[string]int{})
	}
	return count
}

func part2(input []string) int {
	count := 0

	for _, line := range input {
		line, groups := parseLine(line)
		newLine, newGroups := "", []int{}
		for i := 0; i < 5; i++ {
			newLine += line
			if i < 4 {
				newLine += "?"
			}
			newGroups = append(newGroups, groups...)
		}

		count += calculate(newLine, newGroups, map[string]int{})
	}

	return count
}

func calculate(input string, groups []int, memo map[string]int) int {
	if len(input) == 0 {
		if len(groups) == 0 {
			return 1
		} else {
			return 0
		}
	}
	if len(groups) == 0 {
		if strings.ContainsRune(input, '#') {
			return 0
		} else {
			return 1
		}
	}
	key := fmt.Sprintf("%s-%v", input, groups)
	if val, ok := memo[key]; ok {
		return val
	}

	result := 0
	if input[0] == '.' || input[0] == '?' {
		result += calculate(input[1:], groups, memo)
	}

	if input[0] == '#' || input[0] == '?' {
		if len(input) >= groups[0] && !strings.ContainsRune(input[:groups[0]], '.') && (len(input) == groups[0] || input[groups[0]] != '#') {
			n := mathx.Min(groups[0]+1, len(input))
			result += calculate(input[n:], groups[1:], memo)
		}
	}

	memo[key] = result
	return result
}

func parseLine(line string) (string, []int) {
	parts := strings.Split(line, " ")
	groups := strutil.MapToInt(strings.Split(parts[1], ","))
	return parts[0], groups
}
