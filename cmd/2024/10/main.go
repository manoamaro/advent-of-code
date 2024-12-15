package main

import (
	"manoamaro.github.com/advent-of-code/pkg/aoc"
	"manoamaro.github.com/advent-of-code/pkg/grid"
)

var challenge = aoc.New(2024, 10, aoc.IntGridProcessor, part1, part2)

func main() {
	challenge.Run()
}

func part1(input grid.Grid[int]) int {
	trailheads := input.FindAllFunc(func(i int) bool { return i == 0 })
	count := 0
	for _, trailhead := range trailheads {
		visited := make(map[grid.Cell]bool)
		queue := []grid.Cell{trailhead}
		for {
			if len(queue) == 0 {
				break
			}
			// Pop
			current := queue[0]
			queue = queue[1:]
			if visited[current] {
				continue
			}
			visited[current] = true
			value := input.Get(current[0], current[1])
			if *value == 9 {
				// Found the end
				count++
				continue
			}
			// get up, down, left, right
			for _, dir := range [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				next := input.Get(current[0]+dir[0], current[1]+dir[1])
				if next == nil {
					continue
				}
				if *next == *value+1 {
					queue = append(queue, grid.Cell{current[0] + dir[0], current[1] + dir[1]})
				}
			}
		}
	}
	return count
}

func part2(input grid.Grid[int]) int {
	trailheads := input.FindAllFunc(func(i int) bool { return i == 0 })
	count := 0
	for _, trailhead := range trailheads {
		queue := []grid.Cell{trailhead}
		for {
			if len(queue) == 0 {
				break
			}
			// Pop
			current := queue[0]
			queue = queue[1:]
			value := input.Get(current[0], current[1])
			if *value == 9 {
				// Found the end
				count++
				continue
			}
			// get up, down, left, right
			for _, dir := range [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				next := input.Get(current[0]+dir[0], current[1]+dir[1])
				if next == nil {
					continue
				}
				if *next == *value+1 {
					queue = append(queue, grid.Cell{current[0] + dir[0], current[1] + dir[1]})
				}
			}
		}
	}
	return count
}
