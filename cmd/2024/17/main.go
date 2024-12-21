package main

import (
	"fmt"
	"slices"
	"strings"

	"manoamaro.github.com/advent-of-code/pkg/aoc"
	"manoamaro.github.com/advent-of-code/pkg/collections"
	"manoamaro.github.com/advent-of-code/pkg/strings2"
)

func main() {
	challenge := aoc.New(2024, 17, parseInput, part1, part2)
	challenge.TestPart1(aoc.GetSourceFilePath("test_input.txt"), "4,6,3,5,6,3,5,2,1,0")
	challenge.Run()
}

var regA = 0
var regB = 0
var regC = 0

func parseInput(input string) []int {
	lines := strings.Split(input, "\n")
	lines[0] = strings.Split(lines[0], ":")[1]
	regA = strings2.Atoi[int](lines[0])
	lines[1] = strings.Split(lines[1], ":")[1]
	regB = strings2.Atoi[int](lines[1])
	lines[2] = strings.Split(lines[2], ":")[1]
	regC = strings2.Atoi[int](lines[2])
	lines[4] = strings.Split(lines[4], ":")[1]
	return strings2.MapToInt(strings.Split(lines[4], ","))
}

func combo(i, a, b, c int) int {
	switch {
	case i <= 3:
		return i
	case i == 4:
		return a
	case i == 5:
		return b
	case i == 6:
		return c
	default:
		panic("invalid combo")
	}
}

func solve(prog []int, rA, rB, rC int) []int {
	out := []int{}
	pointer := 0
	for pointer < len(prog) {
		inst := prog[pointer]
		op := prog[pointer+1]
		switch inst {
		case 0: // adv
			rA >>= combo(op, rA, rB, rC)
		case 1: // bxl
			rB ^= op
		case 2: // bst
			rB = combo(op, rA, rB, rC) % 8
		case 3: // jnz
			if rA != 0 {
				pointer = op
				continue
			}
		case 4: // bxc
			rB ^= rC
		case 5: // out
			out = append(out, combo(op, rA, rB, rC)%8)
		case 6: // bdv
			rB = rA >> combo(op, rA, rB, rC)
		case 7: // cdv
			rC = rA >> combo(op, rA, rB, rC)
		}
		pointer += 2
	}
	return out
}

func part1(prog []int) string {
	res := solve(prog, regA, regB, regC)
	out := collections.Map(res, func(i int) string {
		return fmt.Sprintf("%v", i)
	})
	return strings.Join(out, ",")
}

func part2(prog []int) string {
	rA := 1
	lenProg := len(prog)
	for {
		res := solve(prog, rA, regB, regC)
		if slices.Equal(res, prog) {
			break
		}
		if slices.Equal(res, prog[lenProg-len(res):]) {
			rA *= 8
		} else {
			if rA%8 == 7 {
				rA /= 8
			}
			rA++
		}
	}
	return fmt.Sprintf("%v", rA)
}
