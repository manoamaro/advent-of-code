package main

import (
	"fmt"
	"regexp"
	"strconv"

	"manoamaro.github.com/advent-of-code/internal"
)

func testInput() []string {
	return []string{
		"sixsevenfivefourxf4mzhmkztwonepzt",
		"nineninesixskjkbhx6nineoneightj",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}
}

func main() {
	input, err := internal.ReadInputLines(2023, 1)
	if err != nil {
		panic(err)
	}
	part1(input)
	part2(input)
}

func part1(input []string) {
	re := regexp.MustCompile(`\d`)
	sum := 0
	for _, line := range input {
		digits := re.FindAllString(line, -1)
		numberStr := fmt.Sprintf("%s%s", digits[0], digits[len(digits)-1])
		number, _ := strconv.Atoi(numberStr)
		sum += number
	}
	fmt.Println(sum)
}

func part2(input []string) {
	reg := regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`)
	regRev := regexp.MustCompile(`\d|enin|thgie|neves|xis|evif|ruof|eerht|owt|eno`)
	sum := 0
	for _, line := range input {
		firstDigit := parseNameToDigit(reg.FindString(line))
		lastDigit := parseNameToDigit(regRev.FindString(internal.ReverseString(line)))
		numberStr := fmt.Sprintf("%d%d", firstDigit, lastDigit)
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			panic(err)
		}
		sum += number
	}
	fmt.Println(sum)
}

func parseNameToDigit(name string) int {
	switch name {
	case "one":
		return 1
	case internal.ReverseString("one"):
		return 1
	case "two":
		return 2
	case internal.ReverseString("two"):
		return 2
	case "three":
		return 3
	case internal.ReverseString("three"):
		return 3
	case "four":
		return 4
	case internal.ReverseString("four"):
		return 4
	case "five":
		return 5
	case internal.ReverseString("five"):
		return 5
	case "six":
		return 6
	case internal.ReverseString("six"):
		return 6
	case "seven":
		return 7
	case internal.ReverseString("seven"):
		return 7
	case "eight":
		return 8
	case internal.ReverseString("eight"):
		return 8
	case "nine":
		return 9
	case internal.ReverseString("nine"):
		return 9
	}
	n, err := strconv.Atoi(name)
	if err != nil {
		panic(err)
	}
	return n
}
