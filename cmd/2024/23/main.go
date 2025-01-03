package main

import (
	"fmt"
	"slices"
	"strings"

	"manoamaro.github.com/advent-of-code/pkg/aoc"
	"manoamaro.github.com/advent-of-code/pkg/graph"
	"manoamaro.github.com/advent-of-code/pkg/set"
)

var challenge = aoc.New(2024, 23, parseInput, part1, part2)

func main() {
	challenge.TestPart1(aoc.GetSourceFilePath("input_test.txt"), "7")
	challenge.TestPart2(aoc.GetSourceFilePath("input_test.txt"), "co,de,ka,ta")
	challenge.Run()
}

func parseInput(input string) *graph.Graph[string, int] {
	g := graph.New[string, int]()
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		parts := strings.Split(line, "-")
		g.AddTwoWayEdge(parts[0], parts[1], 1, 1, 0, 0)
	}
	return g
}

func part1(input *graph.Graph[string, int]) string {
	triangles := set.New[[3]string]()

	for node, neighbors := range input.Edges() {
		for _, i := range neighbors {
			for _, j := range neighbors {
				if i == j {
					continue
				}
				if input.HasEdge(i, j) {
					triangle := []string{node, i, j}
					slices.Sort(triangle)
					triangles.Add([3]string(triangle))
				}
			}
		}
	}
	count := triangles.CountFunc(func(triangle [3]string) bool {
		return triangle[0][0] == 't' || triangle[1][0] == 't' || triangle[2][0] == 't'
	})
	return fmt.Sprintf("%d", count)
}

func search(g *graph.Graph[string, int], node string, seq set.Set[string], sets *set.Set[string]) {
	seqSlice := seq.Slice()
	slices.Sort(seqSlice)
	key := strings.Join(seqSlice, ",")
	if sets.Contains(key) {
		return
	}
	sets.Add(key)
	for _, neighbor := range g.Neighbors(node) {
		if seq.Contains(neighbor) {
			continue
		}
		containsAll := true
		for _, q := range seq.Slice() {
			if !slices.Contains(g.Neighbors(q), neighbor) {
				containsAll = false
				break
			}
		}
		if !containsAll {
			continue
		}
		c := seq.Copy()
		c.Add(neighbor)
		search(g, neighbor, c, sets)
	}
}

func part2(input *graph.Graph[string, int]) string {
	sets := set.New[string]()
	for node := range input.Edges() {
		search(input, node, set.New(node), &sets)
	}
	biggest := sets.MaxFunc(func(s string) int {
		return len(strings.Split(s, ","))
	})
	return biggest
}
