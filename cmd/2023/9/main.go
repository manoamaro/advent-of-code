package main

import (
	"strings"

	"manoamaro.github.com/advent-of-code/pkg/aoc"
	"manoamaro.github.com/advent-of-code/pkg/strings2"
)

var challenge = aoc.New(2023, 9, aoc.LinesProcessor(), part1, part2)

func main() {
	challenge.Run()
}

func part1(input []string) int {
	sum := 0
	for _, line := range input {
		if line == "" {
			continue
		}
		i := strings2.MapToInt(strings.Split(line, " "))
		reduced := reduceToZeros(i)
		extrapolated := extrapolateRight(reduced)
		sum += extrapolated[0][len(extrapolated[0])-1]
	}
	return sum
}

func part2(input []string) int {
	sum := 0
	for _, line := range input {
		if line == "" {
			continue
		}
		i := strings2.MapToInt(strings.Split(line, " "))
		reduced := reduceToZeros(i)
		extrapolated := extrapolateLeft(reduced)
		sum += extrapolated[0][0]
	}
	return sum
}

func reduceToZeros(input []int) [][]int {
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

func extrapolateRight(input [][]int) [][]int {
	input[len(input)-1] = append(input[len(input)-1], 0)
	for i := len(input) - 2; i >= 0; i-- {
		lastNewValue := input[i+1][len(input[i+1])-1]
		input[i] = append(input[i], input[i][len(input[i])-1]+lastNewValue)
	}
	return input
}

func extrapolateLeft(input [][]int) [][]int {
	// prepend 0
	input[len(input)-1] = append([]int{0}, input[len(input)-1]...)
	for i := len(input) - 2; i >= 0; i-- {
		firstNewValueBelow := input[i+1][0]
		firstNewValue := input[i][0] - firstNewValueBelow
		input[i] = append([]int{firstNewValue}, input[i]...)
	}
	return input
}
