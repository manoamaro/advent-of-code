package main

import (
	"sort"
	"strings"

	"manoamaro.github.com/advent-of-code/pkg/aoc"
	"manoamaro.github.com/advent-of-code/pkg/mathx"
	"manoamaro.github.com/advent-of-code/pkg/strutil"
)

type INPUT struct {
	left  []int
	right []int
}

type OUTPUT int

var challenge = aoc.New(2024, 1, parseInput, part1, part2)

func main() {
	challenge.Run()
}

func parseInput(input string) INPUT {
	left, right := []int{}, []int{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		left = append(left, strutil.Atoi[int](fields[0]))
		right = append(right, strutil.Atoi[int](fields[1]))
	}
	return INPUT{left, right}
}

func part1(input INPUT) OUTPUT {
	sort.Ints(input.left)
	sort.Ints(input.right)
	sumDist := 0
	for i := 0; i < len(input.left); i++ {
		sumDist += mathx.Abs(input.left[i] - input.right[i])
	}
	return OUTPUT(sumDist)
}

func part2(input INPUT) OUTPUT {
	sort.Ints(input.left)
	sort.Ints(input.right)

	indexMap := make(map[int]int)
	for _, v := range input.right {
		indexMap[v]++
	}

	sumSimilarityScores := 0
	for _, v := range input.left {
		if indexMap[v] > 0 {
			sumSimilarityScores += v * indexMap[v]
		}
	}
	return OUTPUT(sumSimilarityScores)
}
