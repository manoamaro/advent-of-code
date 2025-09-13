package main

import (
	"slices"
	"strings"

	"manoamaro.github.com/advent-of-code/pkg/aoc"
	"manoamaro.github.com/advent-of-code/pkg/mathx"
)

var challenge = aoc.New(2023, 11, parseInput, part1, part2)

func main() {
	challenge.Run()
}

func parseInput(input string) [][]bool {
	lines := strings.Split(input, "\n")
	return parseMatrixRaw(lines)
}

func part1(matrix [][]bool) int {
	galaxies := expandGalaxySpaces(matrix, 1)
	sum := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += mathx.ManhattanDistance(galaxies[i], galaxies[j])
		}
	}
	return sum
}

func part2(matrix [][]bool) int {
	galaxies := expandGalaxySpaces(matrix, 999999)

	sum := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += mathx.ManhattanDistance(galaxies[i], galaxies[j])
		}
	}
	return sum
}

func findGalaxies(input [][]bool) [][2]int {
	galaxies := make([][2]int, 0)
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] {
				galaxies = append(galaxies, [2]int{i, j})
			}
		}
	}
	return galaxies
}

func parseMatrixRaw(input []string) [][]bool {
	output := make([][]bool, 0)
	for _, v := range input {
		if v == "" {
			continue
		}
		l := make([]bool, 0)
		for _, v2 := range strings.Split(v, "") {
			l = append(l, v2 != ".")
		}
		output = append(output, l)
	}
	return output
}

func expandGalaxySpaces(input [][]bool, size int) [][2]int {
	galaxies := findGalaxies(input)
	newGalaxies := findGalaxies(input)
	for i := 0; i < len(input); i++ {
		if !slices.Contains(input[i], true) {
			for g, galaxy := range galaxies {
				if galaxy[0] > i {
					newGalaxies[g][0] += size
				}
			}
		}
	}
	for i := 0; i < len(input[0]); i++ {
		copy := true
		// check if column contains any #
		for j := 0; j < len(input); j++ {
			if input[j][i] {
				copy = false
				break
			}
		}
		if copy {
			for g, galaxy := range galaxies {
				if galaxy[1] > i {
					newGalaxies[g][1] += size
				}
			}
		}
	}
	return newGalaxies
}

func expandBruteForce(input [][]bool, size int) [][]bool {
	output := make([][]bool, 0)
	// copy row if row contains only .
	for i := 0; i < len(input); i++ {
		cur := input[i]
		line := make([]bool, len(cur))
		copy(line, cur)
		output = append(output, line)
		if !slices.Contains(cur, true) {
			for x := 0; x < size; x++ {
				line := make([]bool, len(cur))
				copy(line, cur)
				output = append(output, line)
			}
		}
	}

	output2 := make([][]bool, len(output))
	for i := 0; i < len(output); i++ {
		output2[i] = make([]bool, 0)
	}

	for i := 0; i < len(output[0]); i++ {
		copy := true
		for j := 0; j < len(output); j++ {
			output2[j] = append(output2[j], output[j][i])
			if output[j][i] {
				copy = false
			}
		}
		if copy {
			for x := 0; x < size; x++ {
				for j := 0; j < len(output); j++ {
					output2[j] = append(output2[j], output[j][i])
				}
			}
		}
	}
	return output2
}
