package main

import (
	"math"
	"regexp"
	"strings"

	"manoamaro.github.com/advent-of-code/pkg/aoc"
	"manoamaro.github.com/advent-of-code/pkg/strings2"
)

type machine struct {
	a     []int
	b     []int
	prize []int
}

var challenge = aoc.New(2024, 13, parseInput, part1, part2)

func main() {
	challenge.Run()
}

func parseInput(input string) []machine {
	btnRegexp := regexp.MustCompile(`^Button [AB]: X\+(\d+), Y\+(\d+)$`)
	priceRegexp := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)$`)
	machines := strings.Split(input, "\n\n")
	var res []machine
	for _, m := range machines {
		lines := strings.Split(m, "\n")
		btnA := btnRegexp.FindStringSubmatch(lines[0])
		btnB := btnRegexp.FindStringSubmatch(lines[1])
		prize := priceRegexp.FindStringSubmatch(lines[2])
		m := machine{
			a:     []int{strings2.Atoi[int](btnA[1]), strings2.Atoi[int](btnA[2])},
			b:     []int{strings2.Atoi[int](btnB[1]), strings2.Atoi[int](btnB[2])},
			prize: []int{strings2.Atoi[int](prize[1]), strings2.Atoi[int](prize[2])},
		}
		res = append(res, m)
	}
	return res
}

func part1(input []machine) int {
	tokens := 0
	for _, m := range input {
		tokens += calculateTokens(m)
	}
	return tokens
}

func part2(input []machine) int {
	tokens := 0
	for _, m := range input {
		m.prize[0] += 10000000000000
		m.prize[1] += 10000000000000
		tokens += calculateTokens(m)
	}
	return tokens
}

func calculateTokens(m machine) int {
	a := float64(m.b[0]*m.prize[1]-m.prize[0]*m.b[1]) / float64(m.b[0]*m.a[1]-m.a[0]*m.b[1])
	b := float64(m.a[0]*m.prize[1]-m.prize[0]*m.a[1]) / float64(m.a[0]*m.b[1]-m.b[0]*m.a[1])
	if math.Mod(a, 1) == 0 && math.Mod(b, 1) == 0 {
		return int(a)*3 + int(b)
	}
	return 0
}
