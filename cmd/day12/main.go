package main

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"manoamaro.github.com/aoc-2023/internal"
)

func main() {
	input, err := internal.ReadInputLines(12)
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
	count := 0

	for _, line := range input {
		lineInt, groups := parseLine(line)
		qIdx := make([]int, 0)
		for i := 0; i < len(lineInt); i++ {
			if lineInt[i] == -1 {
				qIdx = append(qIdx, i)
			}
		}

		for i := 0; i < 1<<len(qIdx); i++ {
			for j := 0; j < len(qIdx); j++ {
				if i&(1<<j) != 0 {
					lineInt[qIdx[j]] = 1
				} else {
					lineInt[qIdx[j]] = 0
				}
			}
			if checkLineValid(lineInt, groups) {
				count++
			}
		}
	}
	fmt.Println(count)
}

func part2(input []string) {
	fmt.Println("Part 2")
	count := 0

	for _, line := range input {
		initLine, initGroups := parseLine(line)
		lineInt := make([]int, 0)
		groups := make([]int, 0)

		for i := 0; i < 5; i++ {
			lineInt = append(lineInt, initLine...)
			lineInt = append(lineInt, -1)
			groups = append(groups, initGroups...)
		}
		qIdx := make([]int, 0)
		for i := 0; i < len(lineInt); i++ {
			if lineInt[i] == -1 {
				qIdx = append(qIdx, i)
			}
		}
		count += t2(lineInt, qIdx, groups, 0, 0)
	}

	fmt.Println(count)
}

var memo = make(map[string]int)

func t2(line []int, qidx []int, groups []int, pos int, count int) int {
	lineStr := fmt.Sprintf("%v", line)
	if val, ok := memo[lineStr]; ok {
		return val
	}

	if pos == len(qidx) {
		// bottom of the tree
		if checkLineValid(line, groups) {
			return count + 1
		} else {
			return count
		}
	}
	line[qidx[pos]] = 1
	left := t2(line, qidx, groups, pos+1, count)
	line[qidx[pos]] = 0
	right := t2(line, qidx, groups, pos+1, count)
	memo[lineStr] = left + right
	return memo[lineStr]
}

func findGroupsInLine(line []int) []int {
	lineGroups := make([]int, 0)

	for i := 0; i < len(line); {
		for ; i < len(line) && line[i] == 0; i++ {
		}
		if i == len(line) {
			break
		}
		start := i
		count := 0
		for ; start < len(line) && line[start] == 1; start++ {
			count++
		}
		lineGroups = append(lineGroups, count)
		i = start
	}
	return lineGroups
}

func checkLineValid(line []int, groups []int) bool {
	lineGroups := findGroupsInLine(line)
	return slices.Compare(lineGroups, groups) == 0
}

func parseLine(line string) ([]int, []int) {
	parts := strings.Split(line, " ")
	groups := internal.MapToInt(strings.Split(parts[1], ","))
	lineInt := make([]int, 0)
	for _, c := range parts[0] {
		switch c {
		case '.':
			lineInt = append(lineInt, 0)
		case '#':
			lineInt = append(lineInt, 1)
		case '?':
			lineInt = append(lineInt, -1)
		}
	}
	return lineInt, groups
}
