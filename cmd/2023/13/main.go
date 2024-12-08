package main

import (
	"fmt"
	"slices"
	"strings"

	"manoamaro.github.com/advent-of-code/pkg/aoc"
	"manoamaro.github.com/advent-of-code/pkg/collections"
	"manoamaro.github.com/advent-of-code/pkg/math2"
)

var challenge = aoc.New(2023, 13, parseInput, part1, part2)

func main() {
	challenge.Run()
}

func part1(patterns [][]string) int {
	sum := 0

	for _, pattern := range patterns {
		vh, v := findFirstFoldVH(pattern)
		if vh == 0 {
			sum += v
		} else if vh == 1 {
			sum += 100 * v
		}
	}
	return sum
}

func part2(patterns [][]string) int {
	sum := 0
	for _, pattern := range patterns {
		vh, v := findSmudge(pattern)
		if vh == 0 {
			sum += v
		} else if vh == 1 {
			sum += 100 * v
		}
		fmt.Println()
	}

	return sum
}

func toggle(s string) string {
	if s == "." {
		return "#"
	}
	return "."
}

func parseInput(input string) [][]string {
	lines := strings.Split(input, "\n")
	patterns := [][]string{{}}
	p := 0
	for _, line := range lines {
		if line == "" {
			p++
			patterns = append(patterns, []string{})
			continue
		}
		patterns[p] = append(patterns[p], line)
	}
	return patterns
}

func findSmudge(pattern []string) (int, int) {
	ovm := findFoldV(pattern)
	ohm := findFoldH(pattern)

	fmt.Println("Old:", ovm, ohm)

	for i := 0; i < len(pattern); i++ {
		p := strings.Split(pattern[i], "")
		for j := 0; j < len(pattern[i]); j++ {
			p[j] = toggle(p[j])
			pattern[i] = strings.Join(p, "")

			v := findFoldV(pattern)
			h := findFoldH(pattern)

			p[j] = toggle(p[j])
			pattern[i] = strings.Join(p, "")

			dv := collections.Diff(v, ovm)
			dh := collections.Diff(h, ohm)

			if len(dv) > 0 {
				fmt.Println("New V:", v, h, dv)
				return 0, dv[0]
			}

			if len(dh) > 0 {
				fmt.Println("New H:", v, h, dh)
				return 1, dh[0]
			}
		}
	}
	return 0, 0
}

func findFoldV(pattern []string) []int {
	vertical := make([]string, len(pattern[0]))
	for i := 0; i < len(pattern[0]); i++ {
		for _, line := range pattern {
			vertical[i] += string(line[i])
		}
	}
	return findFold(vertical)
}

func findFoldH(pattern []string) []int {
	return findFold(pattern)
}

func findFirstFoldVH(pattern []string) (int, int) {
	verticalFold := findFoldV(pattern)
	horizontalFold := findFoldH(pattern)
	if len(verticalFold) > 0 {
		return 0, verticalFold[0]
	} else if len(horizontalFold) > 0 {
		return 1, horizontalFold[0]
	}
	return -1, -1
}

func findFold(pattern []string) []int {
	if len(pattern) <= 1 {
		return nil
	}

	if len(pattern) == 2 && pattern[0] == pattern[1] {
		return []int{0}
	}

	l := len(pattern)
	acc := []int{}
	for i := 1; i < l; i++ {
		pb := math2.Min(i+i, l)
		pa := math2.Max(0, i-(pb-i))
		p1 := pattern[pa:i]
		p2 := pattern[i:pb]
		p1 = collections.Reverse(p1)
		if slices.Compare(p1, p2) == 0 {
			acc = append(acc, i)
		}
	}
	return acc
}
