package main

import (
	"manoamaro.github.com/advent-of-code/pkg/aoc"
)

type inp string
type res int

var challenge = aoc.New(2023, 18, parseInput, part1, part2)

func main() {
	challenge.Run()
}

func parseInput(input string) inp {
	return inp(input)
}

func part1(input inp) res {
	return res(0)
}

func part2(input inp) res {
	return res(0)
}
