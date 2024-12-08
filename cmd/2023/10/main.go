package main

import (
	"fmt"
	"slices"

	"manoamaro.github.com/advent-of-code/pkg/aoc"
	"manoamaro.github.com/advent-of-code/pkg/utils"
)

var challenge = aoc.New(2023, 10, aoc.GridStringProcessor(), part1, part2)

func main() {
	input, err := utils.ReadInputSlice2d(2023, 10)
	if err != nil {
		panic(err)
	}

	pipes := pipes_T(input)
	part1(pipes)
	part2(pipes)
}

func part1(input [][]string) int {
	i := pipes_T(input)
	return i.navigateLoop()
}

func isOnLoop(x, y int, loop [][]int) bool {
	for _, l := range loop {
		if l[0] == x && l[1] == y {
			return true
		}
	}
	return false
}

func part2(i [][]string) int {
	input := pipes_T(i)
	loop := input.getLoop()
	floodMap := make([][]bool, len(input))
	for i := 0; i < len(input); i++ {
		floodMap[i] = make([]bool, len(input[i]))
	}

	count := 0
	isUp := false
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if isOnLoop(i, j, loop) {
				c := input[i][j]
				if c == "S" && isUp {
					isUp = false
				} else if c == "S" && !isUp {
					isUp = true
				} else if c == PIPE_F && !isUp {
					isUp = true
				} else if c == PIPE_F && isUp {
					isUp = false
				} else if c == PIPE_7 && isUp {
					isUp = false
				} else if c == PIPE_7 && !isUp {
					isUp = true
				} else if c == PIPE_VERTICAL && isUp {
					isUp = false
				} else if c == PIPE_VERTICAL && !isUp {
					isUp = true
				}
			} else {
				floodMap[i][j] = isUp
				if isUp {
					count++
				}
			}
		}
	}

	//print floodMap
	for i := 0; i < len(floodMap); i++ {
		for j := 0; j < len(floodMap[i]); j++ {
			if floodMap[i][j] {
				fmt.Print("+")
			} else {
				fmt.Print("-")
			}
		}
		fmt.Println()
	}

	return count
}

var PIPE_VERTICAL = "|"
var PIPE_HORIZONTAL = "-"
var PIPE_J = "J"
var PIPE_L = "L"
var PIPE_7 = "7"
var PIPE_F = "F"
var UP_START = []string{PIPE_VERTICAL, PIPE_7, PIPE_F}
var DOWN_START = []string{PIPE_VERTICAL, PIPE_J, PIPE_L}
var LEFT_START = []string{PIPE_HORIZONTAL, PIPE_L, PIPE_F}
var RIGHT_START = []string{PIPE_HORIZONTAL, PIPE_J, PIPE_7}

type pipes_T [][]string

func (p pipes_T) start() (x, y int) {
	for x, line := range p {
		for y, pipe := range line {
			if pipe == "S" {
				return x, y
			}
		}
	}
	return -1, -1
}

func (p pipes_T) startAndTerminations() (start []int, terminations [][]int) {
	x, y := p.start()
	start = append(start, x, y)
	if x-1 >= 0 && slices.Contains(UP_START, p[x-1][y]) {
		terminations = append(terminations, []int{x - 1, y})
	}
	if x+1 < len(p) && slices.Contains(DOWN_START, p[x+1][y]) {
		terminations = append(terminations, []int{x + 1, y})
	}
	if y-1 >= 0 && slices.Contains(LEFT_START, p[x][y-1]) {
		terminations = append(terminations, []int{x, y - 1})
	}
	if y+1 < len(p[x]) && slices.Contains(RIGHT_START, p[x][y+1]) {
		terminations = append(terminations, []int{x, y + 1})
	}

	return start, terminations
}

func (p pipes_T) getLoop() [][]int {
	start, terminations := p.startAndTerminations()
	loop := [][]int{start}
	px, py := start[0], start[1]
	nx, ny := terminations[0][0], terminations[0][1]
	for {
		x, y := p.navigateNext(px, py, nx, ny)
		if nx == -1 || ny == -1 {
			return loop
		}
		if nx == start[0] && ny == start[1] {
			return loop
		}
		loop = append(loop, []int{nx, ny})
		px, py, nx, ny = nx, ny, x, y
	}
}

func (p pipes_T) navigateLoop() int {
	start, loops := p.startAndTerminations()
	ploop1 := []int{start[0], start[1]}
	ploop2 := []int{start[0], start[1]}
	loop1 := loops[0]
	loop2 := loops[1]

	count := 1

	for {
		nx1, ny1 := p.navigateNext(ploop1[0], ploop1[1], loop1[0], loop1[1])
		if nx1 == -1 || ny1 == -1 {
			return 0
		}
		nx2, ny2 := p.navigateNext(ploop2[0], ploop2[1], loop2[0], loop2[1])
		if nx1 == -1 || ny1 == -1 {
			return 0
		}
		count += 1

		if nx1 == nx2 && ny1 == ny2 {
			return count
		}

		ploop1[0], ploop1[1], loop1[0], loop1[1] = loop1[0], loop1[1], nx1, ny1
		ploop2[0], ploop2[1], loop2[0], loop2[1] = loop2[0], loop2[1], nx2, ny2
	}
}

func (p pipes_T) navigateNext(px, py, nx, ny int) (int, int) {
	// 0: up, 1: down, 2: left, 3: right
	comingFrom := 0
	if px == nx {
		if ny > py {
			comingFrom = 2
		} else {
			comingFrom = 3
		}
	} else if py == ny {
		if nx > px {
			comingFrom = 0
		} else {
			comingFrom = 1
		}
	}

	switch p[nx][ny] {
	case PIPE_VERTICAL:
		if comingFrom == 0 {
			return nx + 1, ny
		}
		if comingFrom == 1 {
			return nx - 1, ny
		}
	case PIPE_HORIZONTAL:
		if comingFrom == 2 {
			return nx, ny + 1
		}
		if comingFrom == 3 {
			return nx, ny - 1
		}
	case PIPE_J:
		if comingFrom == 0 {
			return nx, ny - 1
		}
		if comingFrom == 2 {
			return nx - 1, ny
		}
	case PIPE_L:
		if comingFrom == 0 {
			return nx, ny + 1
		}
		if comingFrom == 3 {
			return nx - 1, ny
		}
	case PIPE_7:
		if comingFrom == 1 {
			return nx, ny - 1
		}
		if comingFrom == 2 {
			return nx + 1, ny
		}
	case PIPE_F:
		if comingFrom == 1 {
			return nx, ny + 1
		}
		if comingFrom == 3 {
			return nx + 1, ny
		}
	case "S":
		return -1, -1
	}
	return nx, ny
}
