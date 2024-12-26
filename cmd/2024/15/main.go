package main

import (
	"slices"
	"strings"

	"github.com/rs/zerolog/log"
	"manoamaro.github.com/advent-of-code/pkg/aoc"
	"manoamaro.github.com/advent-of-code/pkg/collections"
	"manoamaro.github.com/advent-of-code/pkg/grid"
	"manoamaro.github.com/advent-of-code/pkg/queue"
	"manoamaro.github.com/advent-of-code/pkg/set"
)

type input struct {
	warehouse grid.Grid[byte]
	movements []byte
}

var challenge = aoc.New(2024, 15, parseInput, part1, part2)

func main() {
	challenge.Run()
}

func parseInput(in string) input {
	parts := strings.Split(in, "\n\n")
	warehouseParts := strings.Split(parts[0], "\n")
	warehouse := grid.New[byte](len(warehouseParts), len(warehouseParts[0]))
	for i, line := range warehouseParts {
		for j, c := range line {
			warehouse.Set(i, j, byte(c))
		}
	}
	movements := []byte{}
	for _, line := range strings.Split(parts[1], "\n") {
		for _, c := range line {
			movements = append(movements, byte(c))
		}
	}
	return input{
		warehouse: warehouse,
		movements: movements,
	}
}

var dirMap = map[byte]grid.Dir{
	'^': grid.Up,
	'v': grid.Down,
	'<': grid.Left,
	'>': grid.Right,
}

func part1(in input) int {
	warehouse := in.warehouse.Copy()
	robot := *warehouse.Find('@')
	for _, m := range in.movements {
		dir := dirMap[m]
		newRobot := robot.Move(dir)
		newPos := *warehouse.Get(newRobot[0], newRobot[1])
		if newPos == '.' {
			warehouse.Set(robot[0], robot[1], '.')
			warehouse.Set(newRobot[0], newRobot[1], '@')
			robot = newRobot
			continue
		}
		// Hit a box, move the box and other boxes if needed
		if newPos == 'O' {
			boxes := []grid.Cell{newRobot}
			for {
				next := boxes[len(boxes)-1].Move(dir)
				nextCell := *warehouse.Get(next[0], next[1])
				if nextCell == '#' {
					// Hit a wall, stop moving boxes
					break
				} else if nextCell == '.' || nextCell == '@' {
					// Move the boxes
					for i := len(boxes) - 1; i >= 0; i-- {
						warehouse.Set(boxes[i][0], boxes[i][1], '.')
						boxes[i] = boxes[i].Move(dir)
						warehouse.Set(boxes[i][0], boxes[i][1], 'O')
					}
					warehouse.Set(robot[0], robot[1], '.')
					warehouse.Set(newRobot[0], newRobot[1], '@')
					robot = newRobot
					break
				}
				boxes = append(boxes, next)
			}
		}
		p(warehouse)
	}
	boxes := warehouse.FindAll('O')
	return collections.Fold(boxes, 0, func(acc int, box grid.Cell) int {
		return acc + (box[0] * 100) + box[1]
	})
}

func part2(in input) int {
	// Expand the warehouse
	warehouse := grid.New[byte](in.warehouse.Rows(), in.warehouse.Cols()*2)
	for i := 0; i < in.warehouse.Rows(); i++ {
		for j := 0; j < in.warehouse.Cols(); j++ {
			v := *in.warehouse.Get(i, j)
			if v == '#' || v == '.' {
				warehouse.Set(i, j*2, v)
				warehouse.Set(i, j*2+1, v)
			} else if v == '@' {
				warehouse.Set(i, j*2, '@')
				warehouse.Set(i, j*2+1, '.')
			} else if v == 'O' {
				warehouse.Set(i, j*2, '[')
				warehouse.Set(i, j*2+1, ']')
			}
		}
	}
	robot := *warehouse.Find('@')

	for _, m := range in.movements {
		p(warehouse)

		dir := dirMap[m]
		nextRobot := robot.Move(dir)
		nextPos := *warehouse.Get(nextRobot[0], nextRobot[1])
		if nextPos == '.' {
			warehouse.Set(robot[0], robot[1], '.')
			warehouse.Set(nextRobot[0], nextRobot[1], '@')
			robot = nextRobot
			continue
		}

		if nextPos == '[' || nextPos == ']' {
			q := queue.New(nextRobot)
			seen := set.New[grid.Cell]()
			for n := range q.Seq() {
				if seen.Contains(n) {
					continue
				}
				v := *warehouse.Get(n[0], n[1])
				if v == '.' {
					continue
				}
				seen.Add(n)
				if v == '[' {
					q.Push(n.Move(grid.Right))
				} else if v == ']' {
					q.Push(n.Move(grid.Left))
				} else if v == '#' {
					// hit a wall, cannot move
					seen.Clear()
					break
				}
				q.Push(n.Move(dir))
			}
			if seen.Len() == 0 {
				continue
			}

			// Move the boxes
			toMove := seen.Slice()
			// sort the boxes to move, according to the direction, so we move the boxes from the last to the first
			slices.SortFunc(toMove, func(a, b grid.Cell) int {
				return dir[0]*(b[0]-a[0]) + dir[1]*(b[1]-a[1])
			})

			for _, c := range toMove {
				n := c.Move(dir)
				warehouse.Set(n[0], n[1], *warehouse.Get(c[0], c[1]))
				warehouse.Set(c[0], c[1], '.')
			}

			warehouse.Set(robot[0], robot[1], '.')
			warehouse.Set(nextRobot[0], nextRobot[1], '@')
			robot = nextRobot
		}
	}

	p(warehouse)
	boxes := warehouse.FindAll('[')
	return collections.Fold(boxes, 0, func(acc int, box grid.Cell) int {
		return acc + (box[0] * 100) + box[1]
	})
}

func p(warehouse grid.Grid[byte]) {
	buffer := strings.Builder{}
	buffer.WriteRune('\n')
	for _, c := range warehouse {
		buffer.WriteString(string(c))
		buffer.WriteRune('\n')
	}
	log.Debug().Msg(buffer.String())
}
