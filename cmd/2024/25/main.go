package main

import (
	"strings"

	"manoamaro.github.com/advent-of-code/pkg/aoc"
	"manoamaro.github.com/advent-of-code/pkg/set"
)

func main() {
	challenge := aoc.New(2024, 25, parseInput, part1, part2)
	challenge.TestPart1(aoc.GetSourceFilePath("input_test1.txt"), 3)
	challenge.Run()
}

type in struct {
	keys  [][5]int
	locks [][5]int
}

func parseInput(input string) in {
	blocks := strings.Split(input, "\n\n")
	keys := make([][5]int, 0)
	locks := make([][5]int, 0)

	for _, block := range blocks {
		lines := strings.Split(block, "\n")
		v := [5]int{}
		if lines[0] == "#####" {
			// lock
			for i := 1; i < 7; i++ {
				for j, c := range lines[i] {
					if c == '#' {
						v[j]++
					}
				}
			}
			locks = append(locks, v)
		} else {
			for i := 5; i >= 0; i-- {
				for j, c := range lines[i] {
					if c == '#' {
						v[j]++
					}
				}
			}
			keys = append(keys, v)
		}
	}

	return in{
		keys:  keys,
		locks: locks,
	}
}

func part1(input in) int {
	pairs := set.New[[2][5]int]()
	for _, key := range input.keys {
		for _, lock := range input.locks {
			if pairs.Contains([2][5]int{key, lock}) {
				continue
			}
			fits := true
			for k := 0; k < 5; k++ {
				if key[k]+lock[k] > 5 {
					fits = false
					break
				}
			}
			if fits {
				pairs.Add([2][5]int{key, lock})
			}
		}
	}
	return pairs.Len()
}

func part2(input in) int {
	return 0
}
