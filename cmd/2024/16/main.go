package main

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"manoamaro.github.com/advent-of-code/pkg/aoc"
	"manoamaro.github.com/advent-of-code/pkg/collections"
	"manoamaro.github.com/advent-of-code/pkg/deque"
	"manoamaro.github.com/advent-of-code/pkg/grid"
	"manoamaro.github.com/advent-of-code/pkg/maps"
	"manoamaro.github.com/advent-of-code/pkg/queue"
	"manoamaro.github.com/advent-of-code/pkg/set"
)

var challenge = aoc.New(2024, 16, parseInput, part1, part2)

func main() {
	challenge.Run()
}

type value struct {
	pos grid.Cell
	dir grid.Dir
}

func parseInput(input string) grid.Grid[byte] {
	lines := strings.Split(input, "\n")
	g := grid.New[byte](len(lines), len(lines[0]))
	for i, line := range lines {
		for j, c := range line {
			g[i][j] = byte(c)
		}
	}
	return g
}

func part1(input grid.Grid[byte]) int {
	pq := queue.NewPriorityQueue[value]()
	seen := set.New[value]()
	start := *input.Find('S')
	end := *input.Find('E')

	initial := value{start, grid.Right}
	pq.PushValue(initial, 0)
	seen.Add(initial)

	cost := 0
	for v, p := range pq.SeqPriority() {
		seen.Add(v)
		if v.pos == end {
			cost = p
			break
		}
		// 1. move in the direction
		nextFwd := value{v.pos.Move(v.dir), v.dir}
		// 2. rotate cw
		nextCw := value{v.pos, v.dir.RotateCW()}
		// 3. rotate ccw
		nextCcw := value{v.pos, v.dir.RotateCCW()}
		for i, next := range []value{nextFwd, nextCw, nextCcw} {
			if input[next.pos[0]][next.pos[1]] == '#' {
				continue
			}
			if seen.Contains(next) {
				continue
			}
			if i == 0 {
				pq.PushValue(next, p+1)
			} else {
				pq.PushValue(next, p+1000)
			}
		}
	}

	return cost
}

func part2(input grid.Grid[byte]) int {
	start := *input.Find('S')
	end := *input.Find('E')

	pq := queue.NewPriorityQueue[value]()
	initial := value{start, grid.Right}
	pq.PushValue(initial, 0)
	lowestCostToCell := maps.New(maps.NewEntry(initial, 0))
	backtrack := maps.New[value, set.Set[value]]()
	endStates := set.New[value]()

	bestCost := math.MaxInt

	for cell, cost := range pq.SeqPriority() {
		if cost > lowestCostToCell.GetOr(cell, math.MaxInt) {
			continue
		}
		if cell.pos == end {
			if cost > bestCost {
				break
			}
			bestCost = cost
			endStates.Add(cell)
		}
		// 1. move in the direction
		nextFwd := value{cell.pos.Move(cell.dir), cell.dir}
		// 2. rotate cw
		nextCw := value{cell.pos, cell.dir.RotateCW()}
		// 3. rotate ccw
		nextCcw := value{cell.pos, cell.dir.RotateCCW()}
		for i, next := range []value{nextFwd, nextCw, nextCcw} {
			nextCost := cost
			if i == 0 {
				nextCost += 1
			} else {
				nextCost += 1000
			}

			if input[next.pos[0]][next.pos[1]] == '#' {
				continue
			}

			lowestCost := lowestCostToCell.GetOr(next, math.MaxInt)
			if nextCost > lowestCost {
				continue
			}
			if nextCost < lowestCost {
				backtrack.Set(next, set.New[value]())
				lowestCostToCell.Set(next, nextCost)
			}
			backtrack.GetOrPanic(next).Add(cell)

			pq.PushValue(next, nextCost)
		}
	}

	q := deque.New(endStates.Slice()...)
	seen := set.New[value](endStates.Slice()...)
	for state := range q.SeqFront() {
		b := backtrack.GetOr(state, set.New[value]())
		for _, last := range b.Slice() {
			if seen.Contains(last) {
				continue
			}
			seen.Add(last)
			q.PushBack(last)
		}
	}
	seenCells := collections.Map(seen.Slice(), func(v value) grid.Cell {
		return v.pos
	})
	s := set.New(seenCells...)
	uniqueCells := s.Slice()
	slices.SortFunc(uniqueCells, func(a, b grid.Cell) int {
		return a[0] - b[0]
	})

	return len(uniqueCells)
}

func (v value) String() string {
	return fmt.Sprintf("{%d,%d,%d,%d}", v.pos[0], v.pos[1], v.dir[0], v.dir[1])
}
