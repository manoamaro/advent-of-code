package main

import (
	"manoamaro.github.com/advent-of-code/pkg/aoc"
)

var challenge = aoc.New(2020, 1, aoc.IntLinesProcessor, part1, part2)

func main() {
	challenge.Run()
}

func part1(input []int) int {
	for _, v1 := range input {
		for _, v2 := range input {
			if v1+v2 == 2020 {
				return v1 * v2
			}
		}
	}
	return 0
}

func part2(input []int) int {
	for _, v1 := range input {
		for _, v2 := range input {
			for _, v3 := range input {
				if v1+v2+v3 == 2020 {
					return v1 * v2 * v3
				}
			}
		}
	}
	return 0
}
