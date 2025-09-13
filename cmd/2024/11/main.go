package main

import (
	"strings"

	"manoamaro.github.com/advent-of-code/pkg/aoc"
	"manoamaro.github.com/advent-of-code/pkg/mathx"
	"manoamaro.github.com/advent-of-code/pkg/strutil"
)

var challenge = aoc.New(2024, 11, parseInput, part1, part2)

func main() {
	challenge.Run()
}

func parseInput(input string) []int {
	input = strings.TrimSpace(input)
	return strutil.MapToInt(strings.Fields(input))
}

func part1(input []int) int {
	for i := 0; i < 25; i++ {
		newInput := []int{}
		for j := 0; j < len(input); j++ {
			switch {
			case input[j] == 0:
				newInput = append(newInput, 1)
			case numDigits(input[j])%2 == 0:
				p := mathx.Power(10, numDigits(input[j])/2)
				newInput = append(newInput, input[j]/p)
				newInput = append(newInput, input[j]%p)
			default:
				newInput = append(newInput, input[j]*2024)
			}
		}
		input = newInput
	}
	return len(input)
}

func part2(input []int) int {
	count := map[int]int{}
	for _, v := range input {
		count[v]++
	}
	for i := 0; i < 75; i++ {
		newCount := map[int]int{}
		for i, c := range count {
			switch {
			case i == 0:
				newCount[1] += c
			case numDigits(i)%2 == 0:
				p := mathx.Power(10, numDigits(i)/2)
				l, r := i/p, i%p
				newCount[l] += c
				newCount[r] += c
			default:
				newCount[i*2024] += c
			}
		}
		count = newCount
	}
	total := 0
	for _, c := range count {
		total += c
	}
	return total
}

func numDigits(n int) int {
	if n == 0 {
		return 1
	}
	return mathx.Floor(mathx.Log10(mathx.Abs(n))) + 1
}
