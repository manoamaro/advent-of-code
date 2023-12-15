package main

import (
	"fmt"
	"strings"
	"time"

	"manoamaro.github.com/aoc-2023/internal"
)

func main() {
	input, err := internal.ReadInputLines(12)
	if err != nil {
		panic(err)
	}
	startTimePart1 := time.Now()
	part1(input)
	fmt.Println("Part 1 took:", time.Since(startTimePart1))
	startTimePart2 := time.Now()
	part2(input)
	fmt.Println("Part 2 took:", time.Since(startTimePart2))
}

func part1(input []string) {
	fmt.Println("Part 1")
	count := 0

	for _, line := range input {
		line, groups := parseLine(line)
		count += Calculate(line, groups, map[string]int{})
	}
	fmt.Println(count)
}

func part2(input []string) {
	fmt.Println("Part 2")
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

		count += Calculate(newLine, newGroups, map[string]int{})
	}

	fmt.Println(count)
}

func Calculate(input string, groups []int, memo map[string]int) int {
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
		result += Calculate(input[1:], groups, memo)
	}

	if input[0] == '#' || input[0] == '?' {
		if len(input) >= groups[0] && !strings.ContainsRune(input[:groups[0]], '.') && (len(input) == groups[0] || input[groups[0]] != '#') {
			n := internal.Min(groups[0]+1, len(input))
			result += Calculate(input[n:], groups[1:], memo)
		}
	}

	memo[key] = result
	return result
}

func parseLine(line string) (string, []int) {
	parts := strings.Split(line, " ")
	groups := internal.MapToInt(strings.Split(parts[1], ","))
	return parts[0], groups
}
