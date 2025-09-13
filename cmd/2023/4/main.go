package main

import (
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"manoamaro.github.com/advent-of-code/pkg/aoc"
	"manoamaro.github.com/advent-of-code/pkg/strutil"
)

var challenge = aoc.New(2023, 4, aoc.LinesProcessor(), part1, part2)

func main() {
	challenge.Run()
}

func part1(input []string) int {
	sum := 0
	for _, line := range input {
		if strings.TrimSpace(line) == "" {
			continue
		}
		parts := strings.Split(strings.Split(line, ":")[1], "|")
		winingNumbers := strutil.MapToInt(strings.Split(parts[0], " "))
		playedNumbers := strutil.MapToInt(strings.Split(parts[1], " "))
		matched := checkNumbers(winingNumbers, playedNumbers)
		value := math.Floor(math.Pow(2, float64(len(matched)-1)))
		sum += int(value)
	}
	return sum
}

func part2(input []string) int {
	cardRgx := regexp.MustCompile(`Card\s+(\d+):`)
	cardsStack := make([][]Card, 0)
	for _, line := range input {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		cardId, _ := strconv.Atoi(cardRgx.FindStringSubmatch(line)[1])
		parts := strings.Split(strings.Split(line, ":")[1], "|")
		winingNumbers := strutil.MapToInt(strings.Split(parts[0], " "))
		playedNumbers := strutil.MapToInt(strings.Split(parts[1], " "))
		newCard := Card{cardId, winingNumbers, playedNumbers}
		cardsStack = append(cardsStack, []Card{newCard})
	}

	for idx, cards := range cardsStack {
		for _, card := range cards {
			matched := card.checkNumbers()
			for i := 1; i <= matched; i++ {
				duplicatingCard := cardsStack[idx+i][0]
				cardsStack[idx+i] = append(cardsStack[idx+i], duplicatingCard)
			}
		}
	}

	sum := 0
	for _, cards := range cardsStack {
		sum += len(cards)
	}
	return sum
}

type Card struct {
	id            int
	winingNumbers []int
	playedNumbers []int
}

func (c Card) checkNumbers() int {
	return len(checkNumbers(c.winingNumbers, c.playedNumbers))
}

func checkNumbers(winingNumbers, playedNumbers []int) []int {
	matched := make([]int, 0)
	for _, n := range winingNumbers {
		if slices.Contains(playedNumbers, n) {
			matched = append(matched, n)
		}
	}
	return matched
}
