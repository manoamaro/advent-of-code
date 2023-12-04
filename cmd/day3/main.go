package main

import (
	"fmt"
	"regexp"
	"strconv"
	"unicode"

	"manoamaro.github.com/aoc-2023/internal"
)

func main() {
	input, err := internal.ReadInputLines(3)
	if err != nil {
		panic(err)
	}

	part1(input)
	part2(input)
}

func part1(input []string) {
	fmt.Println("Part 1")
	numberReg := regexp.MustCompile(`(\d+)`)
	engines := make([]int, 0)
	for idx, currentLine := range input {
		if len(currentLine) == 0 {
			continue
		}
		previousLine := ""
		if idx > 0 {
			previousLine = input[idx-1]
		}
		nextLine := ""
		if idx < len(input)-1 {
			nextLine = input[idx+1]
		}
		numbers := numberReg.FindAllStringIndex(currentLine, -1)
		for _, number := range numbers {
			engine, err := strconv.Atoi(currentLine[number[0]:number[1]])
			if err != nil {
				panic(err)
			}
			start := number[0]
			end := number[1]
			isEngine := checkEngine(previousLine, currentLine, nextLine, start, end)
			if isEngine {
				engines = append(engines, engine)
			} else {
			}
		}
	}

	sum := int64(0)
	for _, engine := range engines {
		sum += int64(engine)
	}
	fmt.Println("Sum", sum)
}

func part2(input []string) {
	fmt.Println("Part 2")
	sum := int64(0)
	for lineIdx, currentLine := range input {
		if len(currentLine) == 0 {
			continue
		}
		previousLine := ""
		if lineIdx > 0 {
			previousLine = input[lineIdx-1]
		}
		nextLine := ""
		if lineIdx < len(input)-1 {
			nextLine = input[lineIdx+1]
		}
		for i := 0; i < len(currentLine); i++ {
			if currentLine[i] == '*' {
				engines := findAllNumbersAroundADigit(currentLine, previousLine, nextLine, i)
				if len(engines) == 2 {
					fmt.Println("Found * at", lineIdx+1, engines)
					sum += int64(engines[0] * engines[1])
				}
			}
		}
	}
	fmt.Println("Sum", sum)
}

func findAllNumbersAroundADigit(currentLine string, previousLine string, nextLine string, idx int) []int {
	numbers := make([]int, 0)
	start := idx - 1
	end := idx + 1

	// search on the left
	numberStr := ""
	for i := start; i >= 0 && unicode.IsDigit(rune(currentLine[i])); i-- {
		numberStr = string(currentLine[i]) + numberStr
	}
	n, err := strconv.Atoi(numberStr)
	if err == nil {
		numbers = append(numbers, n)
	}

	// search on the right
	numberStr = ""
	for i := end; i < len(currentLine) && unicode.IsDigit(rune(currentLine[i])); i++ {
		numberStr += string(currentLine[i])
	}
	n, err = strconv.Atoi(numberStr)
	if err == nil {
		numbers = append(numbers, n)
	}

	// search on previous line
	pln := findNumbers(previousLine, idx)
	if len(pln) > 0 {
		numbers = append(numbers, pln...)
	}

	nln := findNumbers(nextLine, idx)
	if len(nln) > 0 {
		numbers = append(numbers, nln...)
	}

	return numbers
}

func findNumbers(line string, idx int) []int {
	numberStr := ""
	start := idx - 1
	end := idx + 1
	r := []int{}
	if len(line) > 0 {
		numberStr = ""
		if unicode.IsDigit(rune(line[idx])) {
			// search left
			for i := idx; i >= 0 && unicode.IsDigit(rune(line[i])); i-- {
				numberStr = string(line[i]) + numberStr
			}
			// search right
			for i := idx + 1; i < len(line) && unicode.IsDigit(rune(line[i])); i++ {
				numberStr += string(line[i])
			}
		}
		n, err := strconv.Atoi(numberStr)
		if err == nil {
			return []int{n}
		}

		if start >= 0 && unicode.IsDigit(rune(line[start])) {
			// search left
			for i := start; i >= 0 && unicode.IsDigit(rune(line[i])); i-- {
				numberStr = string(line[i]) + numberStr
			}
			// search right
			for i := start + 1; i < len(line) && unicode.IsDigit(rune(line[i])); i++ {
				numberStr += string(line[i])
			}
		}
		n, err = strconv.Atoi(numberStr)
		if err == nil {
			r = append(r, n)
		}
		numberStr = ""

		if end < len(line) && unicode.IsDigit(rune(line[end])) {
			// search right
			for i := end; i < len(line) && unicode.IsDigit(rune(line[i])); i++ {
				numberStr += string(line[i])
			}
		}
		n, err = strconv.Atoi(numberStr)
		if err == nil {
			r = append(r, n)
		}
	}
	return r
}

func checkEngine(previousLine string, currentLine string, nextLine string, start int, end int) bool {
	startIdx := start - 1
	if startIdx < 0 {
		startIdx = start
	}
	endIdx := end
	if end == len(currentLine) {
		endIdx = end - 1
	}

	// search on the left
	if start != 0 && currentLine[startIdx] != '.' {
		return true
	}
	// search on the right
	if end != len(currentLine) && currentLine[endIdx] != '.' {
		return true
	}

	// search on previous and next line
	for i := startIdx; i <= endIdx; i++ {
		if len(previousLine) > 0 && previousLine[i] != '.' {
			return true
		}
		if len(nextLine) > 0 && nextLine[i] != '.' {
			return true
		}
	}

	return false
}
