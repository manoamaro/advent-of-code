package main

import (
	"manoamaro.github.com/advent-of-code/pkg/aoc"
	"manoamaro.github.com/advent-of-code/pkg/maps2"
	"manoamaro.github.com/advent-of-code/pkg/set"
	"slices"
)

var challenge = aoc.New(2024, 22, aoc.IntLinesProcessor, part1, part2)

func main() {
	challenge.TestPart1(aoc.GetSourceFilePath("test_input.txt"), 37327623)
	challenge.TestPart2(aoc.GetSourceFilePath("test_input_2.txt"), 23)
	challenge.Run()
}

func getNextSecret(secret int) int {
	secret = (secret ^ (secret * 64)) % 16777216
	secret = (secret ^ (secret / 32)) % 16777216
	secret = (secret ^ (secret * 2048)) % 16777216
	return secret
}

func part1(input []int) int {
	sum := 0
	for _, secret := range input {
		for i := 0; i < 2000; i++ {
			secret = getNextSecret(secret)
		}
		sum += secret
	}
	return sum
}

func part2(input []int) int {
	pricesDiffTotal := maps2.New[[4]int, int]()
	for _, secret := range input {
		prices := make([]int, 2000)
		prices[0] = secret % 10
		for i := 1; i < 2000; i++ {
			secret = getNextSecret(secret)
			prices[i] = secret % 10
		}
		seen := set.New[[4]int]()
		for i := range len(prices) - 4 {
			values := prices[i : i+5]
			seq := [4]int{values[1] - values[0], values[2] - values[1], values[3] - values[2], values[4] - values[3]}
			if seen.Contains(seq) {
				continue
			}
			seen.Add(seq)
			pricesDiffTotal[seq] += values[4]
		}
	}
	return slices.Max(pricesDiffTotal.Values())
}
