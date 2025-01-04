package main

import (
	"github.com/rs/zerolog/log"
	"manoamaro.github.com/advent-of-code/pkg/aoc"
	"manoamaro.github.com/advent-of-code/pkg/maps2"
)

var challenge = aoc.New(2024, 8, aoc.GridStringProcessor(), part1, part2)

func main() {
	challenge.Run()
}

func part1(input [][]string) int {
	antennas := maps2.New[string, [][2]int]()
	for r, row := range input {
		for c, cell := range row {
			if cell != "." {
				antennas[cell] = append(antennas[cell], [2]int{r, c})
			}
		}
	}

	antinodes := maps2.New[[2]int, bool]()
	for _, coords := range antennas {
		// Combinations
		for i := range coords {
			for j := range coords {
				if i == j {
					continue
				}
				dir := [2]int{coords[i][0] - coords[j][0], coords[i][1] - coords[j][1]}
				antinode := [2]int{coords[i][0] + dir[0], coords[i][1] + dir[1]}
				antinodes[antinode] = true
			}
		}
	}

	count := 0
	for coord := range antinodes {
		if coord[0] >= 0 && coord[0] < len(input) && coord[1] >= 0 && coord[1] < len(input[0]) {
			count++
		}
	}
	p(input, antinodes)
	return count
}

func part2(input [][]string) int {
	antennas := maps2.New[string, [][2]int]()
	antinodes := maps2.New[[2]int, bool]()
	for r, row := range input {
		for c, cell := range row {
			if cell != "." {
				antennas[cell] = append(antennas[cell], [2]int{r, c})
				antinodes[[2]int{r, c}] = true
			}
		}
	}

	for _, coords := range antennas {
		// Combinations
		for i := range coords {
			for j := range coords {
				if i == j {
					continue
				}
				dir := [2]int{coords[i][0] - coords[j][0], coords[i][1] - coords[j][1]}
				startCord := coords[i]
				for {
					antinode := [2]int{startCord[0] + dir[0], startCord[1] + dir[1]}
					if antinode[0] < 0 || antinode[0] >= len(input) || antinode[1] < 0 || antinode[1] >= len(input[0]) {
						break
					}
					antinodes[antinode] = true
					startCord = antinode
				}
			}
		}
	}

	count := 0
	for coord := range antinodes {
		if coord[0] >= 0 && coord[0] < len(input) && coord[1] >= 0 && coord[1] < len(input[0]) {
			count++
		}
	}
	p(input, antinodes)
	return count
}

func p(input [][]string, antinodes maps2.Map[[2]int, bool]) {
	for r, row := range input {
		s := ""
		for c, cell := range row {
			if _, ok := antinodes[[2]int{r, c}]; ok {
				s += "#"
			} else {
				s += cell
			}
		}
		log.Info().Msgf("%v", s)
	}
}
