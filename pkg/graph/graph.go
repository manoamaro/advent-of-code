package graph

import (
	"fmt"
	"manoamaro.github.com/advent-of-code/pkg/queue"
	"manoamaro.github.com/advent-of-code/pkg/set"
	"math"
	"slices"
	"strings"
)

type Edge[T comparable, V any] struct {
	To     *T
	Value  V
	Weight int
}

type NodeValue[T comparable, V any] struct {
	Node  T
	Value V
}

type Path[T comparable] []T

type Graph[T comparable, V any] struct {
	Edges map[T][]Edge[T, V]
}

func New[T comparable, V any]() *Graph[T, V] {
	return &Graph[T, V]{Edges: make(map[T][]Edge[T, V])}
}

func (g *Graph[T, V]) AddTwoWayEdge(from, to T, weight1, weight2 int, value1, value2 V) {
	g.AddOneWayEdge(from, to, weight1, value1)
	g.AddOneWayEdge(to, from, weight2, value2)
}

func (g *Graph[T, V]) AddOneWayEdge(from, to T, weight int, value V) {
	g.Edges[from] = append(g.Edges[from], Edge[T, V]{To: &to, Weight: weight, Value: value})
}

func (g *Graph[T, V]) FindShortestPathBetween(start, end T) Path[T] {
	var result Path[T]
	pq := queue.NewPriorityQueue[Path[T]]()
	pq.PushValue(Path[T]{start}, 0)
	seen := set.New[T]()
	best := math.MaxInt
	for currPath, prior := range pq.SeqPriority() {
		currIndex := len(currPath) - 1
		currNode := currPath[currIndex]
		seen.Add(currNode)
		if currNode == end && prior < best {
			result = currPath
			best = prior
			continue
		}
		for _, edge := range g.Edges[currNode] {
			if seen.Contains(*edge.To) || prior+edge.Weight >= best {
				continue
			}
			newPath := slices.Clone(currPath)
			newPath = append(newPath, *edge.To)
			pq.PushValue(newPath, prior+edge.Weight)
		}
	}
	return result
}

func (g *Graph[T, V]) FindShortestPathsBetween(start, end T) []Path[T] {
	var results []Path[T]
	pq := queue.NewPriorityQueue[Path[T]]()
	pq.PushValue(Path[T]{start}, 0)
	seen := set.New[T]()
	best := math.MaxInt
	for currPath, prior := range pq.SeqPriority() {
		currIndex := len(currPath) - 1
		currNode := currPath[currIndex]
		seen.Add(currNode)
		if currNode == end && prior <= best {
			results = append(results, currPath)
			best = prior
			continue
		}
		for _, edge := range g.Edges[currNode] {
			if seen.Contains(*edge.To) || prior+edge.Weight >= best {
				continue
			}
			newPath := slices.Clone(currPath)
			newPath = append(newPath, *edge.To)
			pq.PushValue(newPath, prior+edge.Weight)
		}
	}
	return results
}

func (g *Graph[T, V]) String() string {
	var sb strings.Builder
	for k, v := range g.Edges {
		sb.WriteString(fmt.Sprintf("%v -> ", k))
		for _, e := range v {
			sb.WriteString(fmt.Sprintf("%s ", &e))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (e *Edge[T, V]) String() string {
	return fmt.Sprintf("%v (%v,%v)", *e.To, e.Weight, e.Value)
}
