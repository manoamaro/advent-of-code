package main

import (
	"fmt"
	"time"

	"manoamaro.github.com/advent-of-code/internal"
)

func main() {
	input, err := internal.ReadInputLines(2024, 1)
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
}

func part2(input []string) {
	fmt.Println("Part 2")
}
