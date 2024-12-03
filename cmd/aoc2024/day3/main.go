package main

import (
	"fmt"
	"regexp"
	"time"

	"manoamaro.github.com/advent-of-code/pkg/math2"
	"manoamaro.github.com/advent-of-code/pkg/utils"
)

func main() {
	input, err := utils.ReadInput(2024, 3)
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

func part1(input string) {
	fmt.Println("Part 1")
	regex := regexp.MustCompile(`mul\((\d+)\,(\d+)\)`)
	matches := regex.FindAllStringSubmatch(input, -1)
	sum := 0
	for _, match := range matches {
		x, y := math2.Atoi[int](match[1]), math2.Atoi[int](match[2])
		sum += x * y
	}
	fmt.Println("Sum: ", sum)
}

func part2(input string) {
	fmt.Println("Part 2")
	regex := regexp.MustCompile(`(mul\((\d+)\,(\d+)\))|(do\(\))|(don't\(\))`)
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
				x, y := math2.Atoi[int](match[2]), math2.Atoi[int](match[3])
				sum += x * y
			}
		}
	}
	fmt.Println("Sum: ", sum)
}
