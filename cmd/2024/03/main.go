package main

import (
	"regexp"

	"manoamaro.github.com/advent-of-code/pkg/aoc"
	"manoamaro.github.com/advent-of-code/pkg/strutil"
)

var challenge = aoc.New(2024, 3, aoc.StringProcessor, part1, part2)

func main() {
	challenge.Run()
}

func part1(input string) int {
	regex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := regex.FindAllStringSubmatch(input, -1)
	sum := 0
	for _, match := range matches {
		x, y := strutil.Atoi[int](match[1]), strutil.Atoi[int](match[2])
		sum += x * y
	}
	return sum
}

func part2(input string) int {
	regex := regexp.MustCompile(`(mul\((\d+),(\d+)\))|(do\(\))|(don't\(\))`)
	matches := regex.FindAllStringSubmatch(input, -1)
	sum := 0
	enabled := true
	for _, match := range matches {
		if match[0] == "do()" {
			enabled = true
		} else if match[0] == "don't()" {
			enabled = false
		} else {
			if enabled {
				x, y := strutil.Atoi[int](match[2]), strutil.Atoi[int](match[3])
				sum += x * y
			}
		}
	}
	return sum
}
