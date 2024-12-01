package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"manoamaro.github.com/advent-of-code/internal"
)

func main() {
	rawInput, err := internal.ReadInputLines(2024, 1)
	if err != nil {
		panic(err)
	}
	left, right := []int{}, []int{}

	for _, line := range rawInput {
		fields := strings.Fields(line)
		left = append(left, internal.Atoi[int](fields[0]))
		right = append(right, internal.Atoi[int](fields[1]))
	}

	startTimePart1 := time.Now()
	part1(left, right)
	fmt.Println("Part 1 took:", time.Since(startTimePart1))
	startTimePart2 := time.Now()
	part2(left, right)
	fmt.Println("Part 2 took:", time.Since(startTimePart2))
}

func part1(left, right []int) {
	sort.Ints(left)
	sort.Ints(right)
	sumDist := 0
	for i := 0; i < len(left); i++ {
		sumDist += internal.Abs(left[i] - right[i])
	}
	fmt.Printf("Part 1: %d\n", sumDist)
}

func part2(left, right []int) {
	fmt.Println("Part 2")
	sort.Ints(left)
	sort.Ints(right)

	indexMap := make(map[int]int)
	for _, v := range right {
		indexMap[v]++
	}

	sumSimilarityScores := 0
	for _, v := range left {
		if indexMap[v] > 0 {
			sumSimilarityScores += v * indexMap[v]
		}
	}
	fmt.Printf("Part 2: %d\n", sumSimilarityScores)
}
