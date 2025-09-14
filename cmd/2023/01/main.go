package main

import (
	"fmt"
	"regexp"
	"strconv"

	"manoamaro.github.com/advent-of-code/pkg/aoc"
	"manoamaro.github.com/advent-of-code/pkg/strutil"
)

var challenge = aoc.New(2023, 1, aoc.LinesProcessor(), part1, part2)

func main() {
	challenge.Run()
}

func part1(input []string) int {
	re := regexp.MustCompile(`\d`)
	sum := 0
	for _, line := range input {
		digits := re.FindAllString(line, -1)
		numberStr := fmt.Sprintf("%s%s", digits[0], digits[len(digits)-1])
		number, _ := strconv.Atoi(numberStr)
		sum += number
	}
	return sum
}

func part2(input []string) int {
	reg := regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`)
	regRev := regexp.MustCompile(`\d|enin|thgie|neves|xis|evif|ruof|eerht|owt|eno`)
	sum := 0
	for _, line := range input {
		firstDigit := parseNameToDigit(reg.FindString(line))
		lastDigit := parseNameToDigit(regRev.FindString(strutil.ReverseString(line)))
		numberStr := fmt.Sprintf("%d%d", firstDigit, lastDigit)
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			panic(err)
		}
		sum += number
	}
	return sum
}

func parseNameToDigit(name string) int {
	switch name {
	case "one":
		return 1
	case strutil.ReverseString("one"):
		return 1
	case "two":
		return 2
	case strutil.ReverseString("two"):
		return 2
	case "three":
		return 3
	case strutil.ReverseString("three"):
		return 3
	case "four":
		return 4
	case strutil.ReverseString("four"):
		return 4
	case "five":
		return 5
	case strutil.ReverseString("five"):
		return 5
	case "six":
		return 6
	case strutil.ReverseString("six"):
		return 6
	case "seven":
		return 7
	case strutil.ReverseString("seven"):
		return 7
	case "eight":
		return 8
	case strutil.ReverseString("eight"):
		return 8
	case "nine":
		return 9
	case strutil.ReverseString("nine"):
		return 9
	}
	n, err := strconv.Atoi(name)
	if err != nil {
		panic(err)
	}
	return n
}
