package main

import (
	"strconv"
	"strings"

	"manoamaro.github.com/advent-of-code/pkg/aoc"
)

var challenge = aoc.New(2020, 2, aoc.LinesProcessor(), part1, part2)

func main() {
	challenge.Run()
}

func part1(input []string) int {
	count := 0
	for _, i := range input {
		fields := strings.Fields(i)
		amount := strings.Split(fields[0], "-")
		amountMin, _ := strconv.Atoi(amount[0])
		amountMax, _ := strconv.Atoi(amount[1])
		letter := strings.Trim(fields[1], ":")
		password := strings.TrimSpace(fields[2])
		letterCount := strings.Count(password, letter)
		if letterCount >= amountMin && letterCount <= amountMax {
			count++
		}
	}
	return count
}

func part2(input []string) int {
	count := 0
	for _, i := range input {
		fields := strings.Fields(i)
		positions := strings.Split(fields[0], "-")
		pos1, _ := strconv.Atoi(positions[0])
		pos2, _ := strconv.Atoi(positions[1])
		letter := strings.Trim(fields[1], ":")
		password := strings.TrimSpace(fields[2])
		password1, password2 := string(password[pos1-1]), string(password[pos2-1])
		if (password1 == letter || password2 == letter) && password1 != password2 {
			count++
		}
	}
	return count
}
