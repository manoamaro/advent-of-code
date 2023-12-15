package main

import (
	"fmt"
	"regexp"
	"time"

	"manoamaro.github.com/aoc-2023/internal"
)

func main() {
	input, err := internal.ReadInputLines(8)
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
	m := ParseMap(input)
	curr := "AAA"
	instrPos := 0
	steps := 0
	for {
		inst := m.instructions[instrPos%len(m.instructions)]
		node := m.nodes[curr]
		switch inst {
		case 'L':
			curr = node.left
		case 'R':
			curr = node.right
		}
		instrPos += 1
		steps += 1
		if curr == "ZZZ" {
			break
		}
	}
	fmt.Println("Steps:", steps)
}

func part2(input []string) {
	fmt.Println("Part 2")
	m := ParseMap(input)

	starts := make([]string, 0)
	for k := range m.nodes {
		if k[2] == 'A' {
			starts = append(starts, k)
		}
	}

	loopsSize := make(map[string]uint64)

	for _, start := range starts {
		loopsSize[start] = 0
		curr := start
		pos := 0
		for {
			inst := m.instructions[pos%len(m.instructions)]
			node := m.nodes[curr]
			next := ""
			switch inst {
			case 'L':
				next = node.left
			case 'R':
				next = node.right
			}
			pos += 1
			loopsSize[start] += 1
			curr = next
			if next[2] == 'Z' {
				break
			}
		}
	}

	steps := uint64(1)

	for _, v := range loopsSize {
		steps = internal.LowestCommonMultiple(steps, v)
	}

	fmt.Println("Steps:", steps)
}

type Map struct {
	instructions string
	nodes        map[string]Node
}

func ParseMap(input []string) Map {
	reg := regexp.MustCompile(`^(\w{3}) = \((\w{3}), (\w{3})\)$`)
	m := Map{
		nodes: make(map[string]Node),
	}
	m.instructions = input[0]
	for _, line := range input[2:] {
		if line == "" {
			continue
		}
		match := reg.FindStringSubmatch(line)
		node := Node{
			left:  match[2],
			right: match[3],
		}
		m.nodes[match[1]] = node
	}
	return m
}

type Node struct {
	left  string
	right string
}
