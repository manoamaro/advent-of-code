package main

import (
	"slices"

	"manoamaro.github.com/advent-of-code/pkg/aoc"
)

var challenge = aoc.New(2024, 9, parseInput, part1_2, part2_2)

func main() {
	challenge.Run()
}

type block struct {
	id     int
	length int
}

func parseInput(input string) []block {
	blocks := []block{}
	for i, c := range input {
		id := -1
		if i%2 == 0 {
			id = i / 2
		}
		l := int(c - 48) // simple conversion from char [0-9] to int
		blocks = append(blocks, block{id, l})
	}
	return blocks
}

func part1_2(blocks []block) int {
	compacted := construct(blocks)

	for i, j := 0, len(compacted)-1; i < j; {
		if compacted[i] != -1 {
			i++
			continue
		}
		if compacted[j] == -1 {
			j--
			continue
		}
		compacted[i], compacted[j] = compacted[j], compacted[i]
	}
	return calculateChecksum(compacted)
}

func part2_2(blocks []block) int {
	newBlocks := slices.Clone(blocks)

	for i := len(newBlocks) - 1; i >= 0; i-- {
		if newBlocks[i].id == -1 {
			continue
		}
		for j := 0; j < i; j++ {
			if newBlocks[j].id != -1 {
				continue
			}
			if newBlocks[j].length < newBlocks[i].length {
				continue
			}
			lenDiff := newBlocks[j].length - newBlocks[i].length
			newBlocks[j].id, newBlocks[i].id = newBlocks[i].id, newBlocks[j].id
			newBlocks[j].length -= lenDiff
			if lenDiff > 0 {
				newBlocks = slices.Insert(newBlocks, j+1, block{-1, lenDiff})
				i++
			}
			break
		}
	}
	compressed := construct(newBlocks)
	return calculateChecksum(compressed)
}

func construct(blocks []block) []int {
	memory := []int{}
	for i := 0; i < len(blocks); i++ {
		for j := 0; j < blocks[i].length; j++ {
			memory = append(memory, blocks[i].id)
		}
	}
	return memory
}

func calculateChecksum(memory []int) int {
	checksum := 0
	for i := 0; i < len(memory); i++ {
		if memory[i] == -1 {
			continue
		}
		checksum += i * memory[i]
	}
	return checksum
}
