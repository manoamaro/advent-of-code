package main

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/rs/zerolog/log"
	"manoamaro.github.com/advent-of-code/pkg/aoc"
	"manoamaro.github.com/advent-of-code/pkg/mapx"
	"manoamaro.github.com/advent-of-code/pkg/sliceutil"
	"manoamaro.github.com/advent-of-code/pkg/strutil"
)

func main() {
	challenge := aoc.New(2024, 24, parseInput, part1, part2)
	challenge.TestPart1(aoc.GetSourceFilePath("input_test_1.txt"), "2024")
	challenge.Run()
}

type i struct {
	wires mapx.Map[string, int]
	ops   mapx.Map[string, [3]string] // left, op, right
}

func parseInput(input string) *i {
	parts := strings.Split(input, "\n\n")
	wires := mapx.New[string, int]()
	for _, line := range strings.Split(parts[0], "\n") {
		wv := strings.Split(line, ": ")
		wires.Set(wv[0], strutil.Atoi[int](wv[1]))
	}

	ops := mapx.New[string, [3]string]()
	opsReg := regexp.MustCompile(`^(.{3})\s(AND|OR|XOR)\s(.{3})\s->\s(.{3})$`)
	for _, line := range strings.Split(parts[1], "\n") {
		opsParts := opsReg.FindStringSubmatch(line)
		if len(opsParts) == 5 {
			ops[opsParts[4]] = [3]string{opsParts[1], opsParts[2], opsParts[3]}
		}
	}

	return &i{
		wires: wires,
		ops:   ops,
	}
}

var oprs = map[string]func(int, int) int{
	"AND": func(a, b int) int { return a & b },
	"OR":  func(a, b int) int { return a | b },
	"XOR": func(a, b int) int { return a ^ b },
}

func calculateWire(input *i, wire string) int {
	if v, ok := input.wires[wire]; ok {
		return v
	}
	op := input.ops[wire]
	r := oprs[op[1]](calculateWire(input, op[0]), calculateWire(input, op[2]))
	input.wires[wire] = r
	return r
}

func calculateWires(input *i, wires []string) int {
	return sliceutil.Fold(wires, 0, func(acc int, e string) int {
		return acc<<1 | calculateWire(input, e)
	})
}

func part1(input *i) string {
	zWires := sliceutil.FilterFunc(input.ops.Keys(), func(s string) bool { return strings.HasPrefix(s, "z") })
	slices.SortFunc(zWires, func(a, b string) int { return strutil.Atoi[int](b[1:]) - strutil.Atoi[int](a[1:]) })
	result := calculateWires(input, zWires)
	return fmt.Sprintf("%d", result)
}

func part2(input *i) string {
	pairs := []string{}
	for range 4 {
		base := findFirstInvalid(input)
		found := false
		for a := range input.ops {
			if found {
				break
			}
			for b := range input.ops {
				if a == b {
					continue
				}
				input.ops[a], input.ops[b] = input.ops[b], input.ops[a]
				if findFirstInvalid(input) > base {
					log.Debug().Str("a", a).Str("b", b).Msg("Found")
					pairs = append(pairs, a, b)
					found = true
					break
				}
				input.ops[a], input.ops[b] = input.ops[b], input.ops[a]
			}
		}
	}
	slices.Sort(pairs)
	return strings.Join(pairs, ",")
}

func findFirstInvalid(input *i) int {
	zWires := sliceutil.FilterFunc(input.ops.Keys(), func(s string) bool { return strings.HasPrefix(s, "z") })
	slices.SortFunc(zWires, func(a, b string) int { return strutil.Atoi[int](a[1:]) - strutil.Atoi[int](b[1:]) })
	for i, wire := range zWires {
		if !verify(input, wire, i) {
			return i
		}
	}
	return -1
}

func verify(input *i, wire string, num int) bool {
	log.Trace().Str("wire", wire).Int("num", num).Msg("verify")
	op := input.ops[wire]
	// Z level is always a XOR
	if op[1] != "XOR" {
		return false
	}
	// the first level of Z is always x00 XOR y00
	if num == 0 {
		return op[0] == "x00" && op[2] == "y00" || op[0] == "y00" && op[2] == "x00"
	}
	return verifyXOR(input, op[0], num) && verifyCB(input, op[2], num) || verifyXOR(input, op[2], num) && verifyCB(input, op[0], num)
}

func verifyXOR(input *i, wire string, num int) bool {
	log.Trace().Str("wire", wire).Int("num", num).Msg("verifyXOR")
	op := input.ops[wire]
	if op[1] != "XOR" {
		return false
	}
	expectedWires := [2]string{fmt.Sprintf("x%02d", num), fmt.Sprintf("y%02d", num)}
	return [2]string{op[0], op[2]} == expectedWires || [2]string{op[2], op[0]} == expectedWires
}

func verifyCB(input *i, wire string, num int) bool {
	log.Trace().Str("wire", wire).Int("num", num).Msg("verifyCB")
	op := input.ops[wire]
	if num == 1 {
		return op[1] == "AND" && (op[0] == "x00" && op[2] == "y00" || op[0] == "y00" && op[2] == "x00")
	}

	return op[1] == "OR" && (verifyAND(input, op[2], num-1) && verifyCBR(input, op[0], num-1) || verifyAND(input, op[0], num-1) && verifyCBR(input, op[2], num-1))
}

func verifyAND(input *i, wire string, num int) bool {
	log.Trace().Str("wire", wire).Int("num", num).Msg("verifyAND")
	op := input.ops[wire]
	expectedWires := [2]string{fmt.Sprintf("x%02d", num), fmt.Sprintf("y%02d", num)}
	return op[1] == "AND" && ([2]string{op[0], op[2]} == expectedWires || [2]string{op[2], op[0]} == expectedWires)
}

func verifyCBR(input *i, wire string, num int) bool {
	log.Trace().Str("wire", wire).Int("num", num).Msg("verifyCBR")
	op := input.ops[wire]
	return op[1] == "AND" && (verifyXOR(input, op[0], num) && verifyCB(input, op[2], num) || verifyXOR(input, op[2], num) && verifyCB(input, op[0], num))
}
