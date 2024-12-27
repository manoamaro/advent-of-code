package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"manoamaro.github.com/advent-of-code/pkg/aoc"
	"manoamaro.github.com/advent-of-code/pkg/collections"
	"manoamaro.github.com/advent-of-code/pkg/graph"
	"manoamaro.github.com/advent-of-code/pkg/strings2"
	"math"
	"strings"
)

var challenge = aoc.New(2024, 21, aoc.StringProcessor, part1, part2)

func main() {
	challenge.TestPart1(aoc.GetSourceFilePath("test_input.txt"), 68*29)
	challenge.Run()
}

var numKeypad = graph.New[rune, rune]()
var dirKeypad = graph.New[rune, rune]()

func init() {
	numKeypad.AddTwoWayEdge('A', '0', 1, 1, '<', '>')
	numKeypad.AddTwoWayEdge('A', '3', 1, 1, '^', 'v')
	numKeypad.AddTwoWayEdge('0', '2', 1, 1, '^', 'v')
	numKeypad.AddTwoWayEdge('3', '2', 1, 1, '<', '>')
	numKeypad.AddTwoWayEdge('3', '6', 1, 1, '^', 'v')
	numKeypad.AddTwoWayEdge('2', '1', 1, 1, '<', '>')
	numKeypad.AddTwoWayEdge('1', '4', 1, 1, '^', 'v')
	numKeypad.AddTwoWayEdge('2', '5', 1, 1, '^', 'v')
	numKeypad.AddTwoWayEdge('6', '5', 1, 1, '<', '>')
	numKeypad.AddTwoWayEdge('6', '9', 1, 1, '^', 'v')
	numKeypad.AddTwoWayEdge('5', '4', 1, 1, '<', '>')
	numKeypad.AddTwoWayEdge('5', '8', 1, 1, '^', 'v')
	numKeypad.AddTwoWayEdge('4', '7', 1, 1, '^', 'v')
	numKeypad.AddTwoWayEdge('7', '8', 1, 1, '>', '<')
	numKeypad.AddTwoWayEdge('8', '9', 1, 1, '>', '<')
	log.Debug().Msgf("Num Keypad:\n%s", numKeypad)

	dirKeypad.AddTwoWayEdge('A', '^', 1, 1, '<', '>')
	dirKeypad.AddTwoWayEdge('A', '>', 1, 1, 'v', '^')
	dirKeypad.AddTwoWayEdge('^', 'v', 1, 1, 'v', '^')
	dirKeypad.AddTwoWayEdge('>', 'v', 1, 1, '<', '>')
	dirKeypad.AddTwoWayEdge('v', '<', 1, 1, '<', '>')
	log.Debug().Msgf("Dir Keypad:\n%s", dirKeypad)
}

func buildPathFromValues(g *graph.Graph[rune, rune], seq []rune) string {
	sb := strings.Builder{}
	for c := range collections.SlideSeq(seq, 2) {
		edges := g.Edges[c[0]]
		edge := collections.FirstFunc(edges, func(e graph.Edge[rune, rune]) bool {
			return *e.To == c[1]
		})
		if edge != nil {
			sb.WriteRune(edge.Value)
		}
	}
	return sb.String()
}

func findAllPaths(g *graph.Graph[rune, rune], seq string) []string {
	possibilities := make([]string, 0)
	for c := range collections.SlideSeq([]rune(seq), 2) {
		paths := g.FindShortestPathsBetween(c[0], c[1])
		var p []string
		for _, path := range paths {
			dirs := buildPathFromValues(g, path) + "A"
			p = append(p, dirs)
		}
		possibilities = collections.ProductFunc(possibilities, p, func(a, b string) string {
			return a + b
		})
	}
	return possibilities
}

func part1(input string) int {
	sum := 0
	for _, code := range strings.Split(input, "\n") {
		numCode := strings2.Atoi[int](code[0:3])
		robot1 := findAllPaths(numKeypad, "A"+code)
		minLen := math.MaxInt
		for _, nextRobot := range robot1 {
			next := solve(dirKeypad, nextRobot, 2)
			minLen = min(minLen, next)
		}
		sum += minLen * numCode
	}
	return sum
}

func part2(input string) int {
	sum := 0
	for _, code := range strings.Split(input, "\n") {
		numCode := strings2.Atoi[int](code[0:3])
		robot1 := findAllPaths(numKeypad, "A"+code)
		minLen := math.MaxInt
		for _, nextRobot := range robot1 {
			next := solve(dirKeypad, nextRobot, 25)
			if next < minLen {
				minLen = next
			}
		}
		sum += minLen * numCode
	}
	return sum
}

var solveCache = map[string]int{}

func solve(g *graph.Graph[rune, rune], seq string, depth int) int {
	if depth == 0 {
		return len(seq)
	}
	cacheKey := fmt.Sprintf("%s_%d", seq, depth)
	if v, ok := solveCache[cacheKey]; ok {
		return v
	}
	total := 0
	for c := range collections.SlideSeq([]rune("A"+seq), 2) {
		paths := g.FindShortestPathsBetween(c[0], c[1])
		minLength := math.MaxInt
		for _, path := range paths {
			nextSeq := buildPathFromValues(g, path) + "A"
			minLength = min(minLength, solve(g, nextSeq, depth-1))
		}
		total += minLength
	}
	solveCache[cacheKey] = total
	return total
}
