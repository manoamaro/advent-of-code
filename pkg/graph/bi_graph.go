package graph

import (
	"fmt"
	"manoamaro.github.com/advent-of-code/pkg/collections"
	"manoamaro.github.com/advent-of-code/pkg/queue"
	"manoamaro.github.com/advent-of-code/pkg/set"
	"math"
	"slices"
	"strings"
)

type Edge[T comparable, V any] struct {
	To     *T
	Weight int
	Value  V
}

type BiGraph[T comparable, V any] struct {
	Edges map[T][]Edge[T, V]
}

func NewBiGraph[T comparable, V any]() *BiGraph[T, V] {
	return &BiGraph[T, V]{Edges: make(map[T][]Edge[T, V])}
}

func (g *BiGraph[T, V]) AddEdge(from, to T, weight1, weight2 int, value1, value2 V) {
	g.AddEdgeOneWay(from, to, weight1, value1)
	g.AddEdgeOneWay(to, from, weight2, value2)
}

func (g *BiGraph[T, V]) AddEdgeOneWay(from, to T, weight int, value V) {
	g.Edges[from] = append(g.Edges[from], Edge[T, V]{To: &to, Weight: weight, Value: value})
}

type NodeValue[T comparable, V any] struct {
	Node  T
	Value V
}

type Path[T comparable] []T

func (g *BiGraph[T, V]) FindShortestPathBetween(start, end T) Path[T] {
	var result Path[T]
	pq := queue.NewPriorityQueue[Path[T]]()
	pq.PushValue(Path[T]{start}, 0)
	seen := set.New[T]()
	best := math.MaxInt
	for currPath, prio := range pq.SeqPriority() {
		currIndex := len(currPath) - 1
		currNode := currPath[currIndex]
		seen.Add(currNode)
		if currNode == end && prio < best {
			result = currPath
			best = prio
			continue
		}
		for _, edge := range g.Edges[currNode] {
			if seen.Contains(*edge.To) || prio+edge.Weight >= best {
				continue
			}
			newPath := slices.Clone(currPath)
			newPath = append(newPath, *edge.To)
			pq.PushValue(newPath, prio+edge.Weight)
		}
	}
	return result
}

func (g *BiGraph[T, V]) FindShortestPathsBetween(start, end T) []Path[T] {
	var results []Path[T]
	pq := queue.NewPriorityQueue[Path[T]]()
	pq.PushValue(Path[T]{start}, 0)
	seen := set.New[T]()
	best := math.MaxInt
	for currPath, prio := range pq.SeqPriority() {
		currIndex := len(currPath) - 1
		currNode := currPath[currIndex]
		seen.Add(currNode)
		if currNode == end && prio <= best {
			results = append(results, currPath)
			best = prio
			continue
		}
		for _, edge := range g.Edges[currNode] {
			if seen.Contains(*edge.To) || prio+edge.Weight >= best {
				continue
			}
			newPath := slices.Clone(currPath)
			newPath = append(newPath, *edge.To)
			pq.PushValue(newPath, prio+edge.Weight)
		}
	}
	return collections.FilterFunc(results, func(p Path[T]) bool { return len(p) == best+1 })
}

func (g *BiGraph[T, V]) String() string {
	var sb strings.Builder
	for k, v := range g.Edges {
		sb.WriteString(fmt.Sprintf("%v -> %v\n", k, v))
	}
	return sb.String()
}

func (e *Edge[T, V]) String() string {
	return fmt.Sprintf("%v (%v,%v)", *e.To, e.Weight, e.Value)
}
