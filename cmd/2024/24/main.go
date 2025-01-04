package main

import (
	"regexp"
	"slices"
	"strings"

	"manoamaro.github.com/advent-of-code/pkg/aoc"
	"manoamaro.github.com/advent-of-code/pkg/collections"
	"manoamaro.github.com/advent-of-code/pkg/maps2"
	"manoamaro.github.com/advent-of-code/pkg/set"
	"manoamaro.github.com/advent-of-code/pkg/strings2"
)

func main() {
	challenge := aoc.New(2024, 24, parseInput, part1, part2)
	challenge.TestPart1(aoc.GetSourceFilePath("input_test_1.txt"), 2024)
	challenge.Run()
}

type i struct {
	wires maps2.Map[string, bool]
	ops   [][4]string
}

func parseInput(input string) *i {
	parts := strings.Split(input, "\n\n")
	wires := maps2.New[string, bool]()
	for _, line := range strings.Split(parts[0], "\n") {
		wv := strings.Split(line, ": ")
		wires.Set(wv[0], wv[1] == "1")
	}

	ops := [][4]string{}
	opsReg := regexp.MustCompile(`^(.{3})\s(AND|OR|XOR)\s(.{3})\s->\s(.{3})$`)
	for _, line := range strings.Split(parts[1], "\n") {
		opsParts := opsReg.FindStringSubmatch(line)
		if len(opsParts) == 5 {
			ops = append(ops, [4]string{opsParts[1], opsParts[2], opsParts[3], opsParts[4]})
		}
	}

	return &i{
		wires: wires,
		ops:   ops,
	}
}

func part1(input *i) int {
	processed := set.New[[4]string]()
	for {
		if processed.Len() == len(input.ops) {
			break
		}
		for _, op := range input.ops {
			if processed.Contains(op) {
				continue
			}
			w1, w2, w3 := op[0], op[2], op[3]
			if input.wires.Has(w1) && input.wires.Has(w2) {
				processed.Add(op)
				w1v, w2v := *input.wires.GetOrPanic(w1), *input.wires.GetOrPanic(w2)
				switch op[1] {
				case "AND":
					input.wires.Set(w3, w1v && w2v)
				case "OR":
					input.wires.Set(op[3], w1v || w2v)
				case "XOR":
					input.wires.Set(op[3], w1v != w2v)
				}
			}
		}
	}
	zWires := collections.FilterFunc(input.wires.Keys(), func(k string) bool { return strings.HasPrefix(k, "z") })
	slices.SortFunc(zWires, func(i, j string) int { return strings2.Atoi[int](j[1:]) - strings2.Atoi[int](i[1:]) })
	zValues := collections.Map(zWires, func(k string) int {
		if *input.wires.GetOrPanic(k) {
			return 1
		} else {
			return 0
		}
	})
	return collections.Fold(zValues, 0, func(acc int, v int) int {
		return (acc << 1) | v
	})
}

func part2(input *i) int {
	return 0
}
