package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"manoamaro.github.com/advent-of-code/pkg/aoc"
	"manoamaro.github.com/advent-of-code/pkg/collections"
	"manoamaro.github.com/advent-of-code/pkg/graph"
	"manoamaro.github.com/advent-of-code/pkg/strings2"
	"math"
	"slices"
	"strings"
)

var challenge = aoc.New(2024, 21, aoc.StringProcessor, part1, part2)

func main() {
	challenge.Run()
}

var keypad = graph.NewBiGraph[rune, rune]()
var directional = graph.NewBiGraph[rune, rune]()

func init() {
	keypad.AddEdge('A', '0', 1, 1, '<', '>')
	keypad.AddEdge('A', '3', 1, 1, '^', 'v')
	keypad.AddEdge('0', '2', 1, 1, '^', 'v')
	keypad.AddEdge('3', '2', 1, 1, '<', '>')
	keypad.AddEdge('3', '6', 1, 1, '^', 'v')
	keypad.AddEdge('2', '1', 1, 1, '<', '>')
	keypad.AddEdge('1', '4', 1, 1, '^', 'v')
	keypad.AddEdge('2', '5', 1, 1, '^', 'v')
	keypad.AddEdge('6', '5', 1, 1, '<', '>')
	keypad.AddEdge('6', '9', 1, 1, '^', 'v')
	keypad.AddEdge('5', '4', 1, 1, '<', '>')
	keypad.AddEdge('5', '8', 1, 1, '^', 'v')
	keypad.AddEdge('4', '7', 1, 1, '^', 'v')
	keypad.AddEdge('7', '8', 1, 1, '>', '<')
	keypad.AddEdge('8', '9', 1, 1, '>', '<')
	printGraph(keypad)

	directional.AddEdge('A', '^', 1, 1, '<', '>')
	directional.AddEdge('A', '>', 1, 1, 'v', '^')
	directional.AddEdge('^', 'v', 1, 1, 'v', '^')
	directional.AddEdge('>', 'v', 1, 1, '<', '>')
	directional.AddEdge('v', '<', 1, 1, '<', '>')
	printGraph(directional)
}

func printGraph(g *graph.BiGraph[rune, rune]) {
	sb := strings.Builder{}
	for k, v := range g.Edges {
		sb.WriteString(fmt.Sprintf("%c -> ", k))
		for _, e := range v {
			sb.WriteString(fmt.Sprintf("%c (%c) ", *e.To, e.Value))
		}
		sb.WriteString("\n")
	}
	log.Info().Msgf("\n%s", sb.String())
}

func buildPathFromValues(g *graph.BiGraph[rune, rune], seq []rune) []rune {
	result := []rune{}
	for c := range collections.SlideSeq(seq, 2) {
		edges := g.Edges[c[0]]
		edge := collections.FirstFunc(edges, func(e graph.Edge[rune, rune]) bool {
			return *e.To == c[1]
		})
		if edge != nil {
			result = append(result, edge.Value)
		}
	}
	return result
}

// mergeLists generates all combinations of the given list of lists
func mergeLists(lists [][]string) []string {
	var result []string

	// Helper function for recursion
	var backtrack func(index int, path []string)
	backtrack = func(index int, path []string) {
		// If we've reached the end of the lists, combine the path into a string and add to the result
		if index == len(lists) {
			result = append(result, strings.Join(path, ""))
			return
		}

		// Iterate over each string in the current list
		for _, item := range lists[index] {
			backtrack(index+1, append(path, item))
		}
	}

	// Start the backtracking process
	backtrack(0, []string{})
	return result
}

var cache = map[string][]string{}

func CombineSlices(slice1, slice2 []string) []string {
	totalCombinations := len(slice1) * len(slice2)
	combinations := make([]string, 0, totalCombinations)
	if len(slice1) == 0 {
		return slice2
	}
	for _, s1 := range slice1 {
		for _, s2 := range slice2 {
			combinations = append(combinations, s1+s2)
		}
	}
	return combinations
}

func findAndMergePaths(g *graph.BiGraph[rune, rune], seq string) []string {
	possibilities := make([]string, 0)
	for c := range collections.SlideSeq([]rune(seq), 2) {
		if v, ok := cache[string(c)]; ok {
			possibilities = CombineSlices(possibilities, v)
		} else {
			paths := g.FindShortestPathsBetween(c[0], c[1])
			var p []string
			for _, path := range paths {
				dirs := buildPathFromValues(g, path)
				dirs = append(dirs, 'A')
				p = append(p, string(dirs))
			}
			cache[string(c)] = p
			possibilities = CombineSlices(possibilities, p)
		}
	}
	lens := collections.Map(possibilities, func(s string) int { return len(s) })
	slices.Sort(lens)
	return collections.FilterFunc(possibilities, func(s string) bool { return len(s) == lens[0] })
}

func sortByLength(a, b string) int {
	return len(a) - len(b)
}

func part1(input string) int {
	sum := 0
	for _, code := range strings.Split(input, "\n") {
		numCode := strings2.Atoi[int](code[0:3])
		best := math.MaxInt
		keyboardPaths := findAndMergePaths(keypad, "A"+code)
		for _, keyboardPath := range keyboardPaths {
			directionalPathsLvl1 := findAndMergePaths(directional, "A"+keyboardPath)
			for _, dp1 := range directionalPathsLvl1 {
				directionalPathsLvl2 := findAndMergePaths(directional, "A"+dp1)
				slices.SortFunc(directionalPathsLvl2, sortByLength)
				if len(directionalPathsLvl2[0]) < best {
					best = len(directionalPathsLvl2[0])
				}
			}
		}
		sum += best * numCode
	}
	return sum
}

func part2(input string) int {
	sum := 0
	for _, code := range strings.Split(input, "\n") {
		numCode := strings2.Atoi[int](code[0:3])
		var shortestPath string
		keyboardPaths := findAndMergePaths(keypad, "A"+code)
		for _, keyboardPath := range keyboardPaths {
			directionalPathsLvl1 := findAndMergePaths(directional, "A"+keyboardPath)
			for _, dp1 := range directionalPathsLvl1 {
				directionalPathsLvl2 := findAndMergePaths(directional, "A"+dp1)
				for _, dp2 := range directionalPathsLvl2 {
					directionalPathsLvl3 := findAndMergePaths(directional, "A"+dp2)
					for _, dp3 := range directionalPathsLvl3 {
						directionalPathsLvl4 := findAndMergePaths(directional, "A"+dp3)
						for _, dp4 := range directionalPathsLvl4 {
							directionalPathsLvl5 := findAndMergePaths(directional, "A"+dp4)
							for _, dp5 := range directionalPathsLvl5 {
								if len(dp5) < len(shortestPath) {
									shortestPath = dp5
								}
							}
						}
					}
				}
			}
		}
		sum += len(shortestPath) * numCode
	}
	return sum
}

func printRunes(r []rune, sep string) string {
	sb := strings.Builder{}
	for i, c := range r {
		sb.WriteRune(c)
		if i < len(r)-1 {
			sb.WriteString(sep)
		}
	}
	return sb.String()
}
