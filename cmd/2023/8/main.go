package main

import (
	"regexp"
	"strings"

	"manoamaro.github.com/advent-of-code/pkg/aoc"
	"manoamaro.github.com/advent-of-code/pkg/math2"
)

var challenge = aoc.New(2023, 8, parseInput, part1, part2)

func main() {
	challenge.Run()
}

func part1(m inputMap) uint64 {
	curr := "AAA"
	instrPos := 0
	steps := uint64(0)
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
	return steps
}

func part2(m inputMap) uint64 {
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
		steps = math2.LCM(steps, v)
	}

	return steps
}

type inputMap struct {
	instructions string
	nodes        map[string]node
}

func parseInput(input string) inputMap {
	reg := regexp.MustCompile(`^(\w{3}) = \((\w{3}), (\w{3})\)$`)
	m := inputMap{
		nodes: make(map[string]node),
	}
	lines := strings.Split(input, "\n")
	m.instructions = lines[0]
	for _, line := range lines[2:] {
		if line == "" {
			continue
		}
		match := reg.FindStringSubmatch(line)
		node := node{
			left:  match[2],
			right: match[3],
		}
		m.nodes[match[1]] = node
	}
	return m
}

type node struct {
	left  string
	right string
}
