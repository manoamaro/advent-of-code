package main

import (
	"manoamaro.github.com/advent-of-code/pkg/aoc"
	"manoamaro.github.com/advent-of-code/pkg/maps2"
	"manoamaro.github.com/advent-of-code/pkg/set"
	"strings"
)

func main() {
	challenge := aoc.New(2024, 19, parseInput, part1, part2)
	challenge.TestPart1(aoc.GetSourceFilePath("test_input.txt"), 6)
	challenge.Run()
}

type in struct {
	patterns      set.Set[string]
	designs       []string
	maxLenPattern int
}

func parseInput(input string) in {
	lines := strings.Split(input, "\n")
	patterns := set.New(strings.FieldsFunc(lines[0], func(r rune) bool {
		return r == ' ' || r == ','
	})...)
	maxLenPattern := 0
	for p := range patterns.Seq() {
		if len(p) > maxLenPattern {
			maxLenPattern = len(p)
		}
	}
	designs := lines[2:]
	return in{
		patterns:      patterns,
		maxLenPattern: maxLenPattern,
		designs:       designs,
	}
}

var cache = maps2.New[string, bool]()

func isPossible(patterns set.Set[string], design string, maxLenPattern int) bool {
	if len(design) == 0 {
		return true
	}
	if cache.Has(design) {
		return *cache.GetOrPanic(design)
	}
	for i := 1; i < min(len(design), maxLenPattern)+1; i++ {
		if patterns.Contains(design[:i]) && isPossible(patterns, design[i:], maxLenPattern) {
			cache.Set(design, true)
			return true
		}
	}
	cache.Set(design, false)
	return false
}

var cacheCount = maps2.New[string, int]()

func countPossibilities(patterns set.Set[string], design string, maxLenPattern int) int {
	if len(design) == 0 {
		return 1
	}
	total := 0
	if cacheCount.Has(design) {
		return *cacheCount.GetOrPanic(design)
	}
	for i := 1; i < min(len(design), maxLenPattern)+1; i++ {
		if patterns.Contains(design[:i]) {
			total += countPossibilities(patterns, design[i:], maxLenPattern)
			cacheCount.Set(design, total)
		}
	}
	cacheCount.Set(design, total)
	return total
}

func part1(input in) int {
	count := 0
	cache.Clear()
	for _, design := range input.designs {
		if isPossible(input.patterns, design, input.maxLenPattern) {
			count++
		}
	}
	return count
}

func part2(input in) int {
	count := 0
	cacheCount.Clear()
	for _, design := range input.designs {
		count += countPossibilities(input.patterns, design, input.maxLenPattern)
	}
	return count
}
