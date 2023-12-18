package main

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"manoamaro.github.com/aoc-2023/internal"
)

func main() {
	input, err := internal.ReadInputLines(13)
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
	patterns := parseInput(input)

	sum := 0

	for _, pattern := range patterns {
		vh, v := FindFirstFoldVH(pattern)
		if vh == 0 {
			sum += v
		} else if vh == 1 {
			sum += 100 * v
		}
	}
	fmt.Println(sum)
}

func part2(input []string) {
	fmt.Println("Part 2")
	patterns := parseInput(input)

	sum := 0
	for _, pattern := range patterns {
		vh, v := FindSmudge(pattern)
		if vh == 0 {
			sum += v
		} else if vh == 1 {
			sum += 100 * v
		}
		fmt.Println()
	}

	fmt.Println(sum)
}

func toggle(s string) string {
	if s == "." {
		return "#"
	}
	return "."
}

func parseInput(input []string) [][]string {
	patterns := [][]string{{}}
	p := 0
	for _, line := range input {
		if line == "" {
			p++
			patterns = append(patterns, []string{})
			continue
		}
		patterns[p] = append(patterns[p], line)
	}
	return patterns
}

func FindSmudge(pattern []string) (int, int) {
	ovm := FindFoldV(pattern)
	ohm := FindFoldH(pattern)

	fmt.Println("Old:", ovm, ohm)

	for i := 0; i < len(pattern); i++ {
		p := strings.Split(pattern[i], "")
		for j := 0; j < len(pattern[i]); j++ {
			p[j] = toggle(p[j])
			pattern[i] = strings.Join(p, "")

			v := FindFoldV(pattern)
			h := FindFoldH(pattern)

			p[j] = toggle(p[j])
			pattern[i] = strings.Join(p, "")

			dv := internal.Diff(v, ovm)
			dh := internal.Diff(h, ohm)

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

func FindFoldV(pattern []string) []int {
	vertical := make([]string, len(pattern[0]))
	for i := 0; i < len(pattern[0]); i++ {
		for _, line := range pattern {
			vertical[i] += string(line[i])
		}
	}
	return FindFold(vertical)
}

func FindFoldH(pattern []string) []int {
	return FindFold(pattern)
}

func FindFirstFoldVH(pattern []string) (int, int) {
	verticalFold := FindFoldV(pattern)
	horizontalFold := FindFoldH(pattern)
	if len(verticalFold) > 0 {
		return 0, verticalFold[0]
	} else if len(horizontalFold) > 0 {
		return 1, horizontalFold[0]
	}
	return -1, -1
}

func FindFold(pattern []string) []int {
	if len(pattern) <= 1 {
		return nil
	}

	if len(pattern) == 2 && pattern[0] == pattern[1] {
		return []int{0}
	}

	l := len(pattern)
	acc := []int{}
	for i := 1; i < l; i++ {
		pb := internal.Min(i+i, l)
		pa := internal.Max(0, i-(pb-i))
		p1 := pattern[pa:i]
		p2 := pattern[i:pb]
		p1 = internal.Reverse(p1)
		if slices.Compare(p1, p2) == 0 {
			acc = append(acc, i)
		}
	}
	return acc
}
