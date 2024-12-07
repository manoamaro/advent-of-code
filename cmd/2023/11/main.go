package main

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"manoamaro.github.com/advent-of-code/pkg/math2"
	"manoamaro.github.com/advent-of-code/pkg/utils"
)

func main() {
	input, err := utils.ReadInputLines(2023, 11)
	if err != nil {
		panic(err)
	}
	startTimePart1 := time.Now()
	part1(input)
	fmt.Println("Part 1 took:", time.Since(startTimePart1))
	startTimePart2 := time.Now()
	part2(input)
	fmt.Println("Part 2 took:", time.Since(startTimePart2))
}

func part1(input []string) {
	fmt.Println("Part 1")
	matrix := ParseMatrixRaw(input)
	galaxies := ExpandGalaxySpaces(matrix, 1)

	sum := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += math2.ManhattanDistance(galaxies[i], galaxies[j])
		}
	}
	fmt.Println(sum)
}

func part2(input []string) {
	fmt.Println("Part 2")
	matrix := ParseMatrixRaw(input)
	galaxies := ExpandGalaxySpaces(matrix, 999999)

	sum := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += math2.ManhattanDistance(galaxies[i], galaxies[j])
		}
	}
	fmt.Println(sum)
}

func FindGalaxies(input [][]bool) [][]int {
	galaxies := make([][]int, 0)
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] {
				galaxies = append(galaxies, []int{i, j})
			}
		}
	}
	return galaxies
}

func ParseMatrixRaw(input []string) [][]bool {
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

func ExpandGalaxySpaces(input [][]bool, size int) [][]int {
	galaxies := FindGalaxies(input)
	newGalaxies := FindGalaxies(input)
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

func ExpandBruteForce(input [][]bool, size int) [][]bool {
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
