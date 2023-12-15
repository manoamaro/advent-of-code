package main

import (
	"fmt"
	"strings"
	"time"

	"manoamaro.github.com/aoc-2023/internal"
)

func main() {
	input, err := internal.ReadInputLines(9)
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
	sum := 0
	for _, line := range input {
		if line == "" {
			continue
		}
		i := internal.MapToInt(strings.Split(line, " "))
		reduced := ReduceToZeros(i)
		extrapolated := ExtrapolateRight(reduced)
		sum += extrapolated[0][len(extrapolated[0])-1]
	}
	fmt.Println("Sum:", sum)
}

func part2(input []string) {
	fmt.Println("Part 2")
	sum := 0
	for _, line := range input {
		if line == "" {
			continue
		}
		i := internal.MapToInt(strings.Split(line, " "))
		reduced := ReduceToZeros(i)
		extrapolated := ExtrapolateLeft(reduced)
		sum += extrapolated[0][0]
	}
	fmt.Println("Sum:", sum)
}

func ReduceToZeros(input []int) [][]int {
	allArray := make([][]int, 0)
	allArray = append(allArray, input)
	currentArray := input
	for {
		newArray := make([]int, len(currentArray)-1)
		for i := 0; i < len(currentArray)-1; i++ {
			newArray[i] = currentArray[i+1] - currentArray[i]
		}
		allArray = append(allArray, newArray)
		currentArray = newArray
		finished := true
		for i := 0; i < len(newArray); i++ {
			if newArray[i] != 0 {
				finished = false
				break
			}
		}
		if finished {
			break
		}
	}
	return allArray
}

func ExtrapolateRight(input [][]int) [][]int {
	input[len(input)-1] = append(input[len(input)-1], 0)
	for i := len(input) - 2; i >= 0; i-- {
		lastNewValue := input[i+1][len(input[i+1])-1]
		input[i] = append(input[i], input[i][len(input[i])-1]+lastNewValue)
	}
	return input
}

func ExtrapolateLeft(input [][]int) [][]int {
	// prepend 0
	input[len(input)-1] = append([]int{0}, input[len(input)-1]...)
	for i := len(input) - 2; i >= 0; i-- {
		firstNewValueBelow := input[i+1][0]
		firstNewValue := input[i][0] - firstNewValueBelow
		input[i] = append([]int{firstNewValue}, input[i]...)
	}
	return input
}
