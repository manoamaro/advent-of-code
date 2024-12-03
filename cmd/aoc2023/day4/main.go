package main

import (
	"fmt"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"manoamaro.github.com/advent-of-code/pkg/collections"
	"manoamaro.github.com/advent-of-code/pkg/utils"
)

func main() {
	input, err := utils.ReadInputLines(2023, 4)
	if err != nil {
		panic(err)
	}
	startTimePart1 := time.Now()
	part1(input)
	fmt.Println("Part 1 took:", time.Since(startTimePart1))
	startTimePart2 := time.Now()
	part2(input)
	fmt.Println("Part 2 took:", time.Since(startTimePart2))
}

func part1(input []string) {
	fmt.Println("Part 1")
	//cardRgx := regexp.MustCompile(`Card\s(\d+):`)
	sum := 0
	for _, line := range input {
		if strings.TrimSpace(line) == "" {
			continue
		}
		//cardId := cardRgx.FindStringSubmatch(line)[1]
		parts := strings.Split(strings.Split(line, ":")[1], "|")
		winingNumbers := collections.MapToInt(strings.Split(parts[0], " "))
		playedNumbers := collections.MapToInt(strings.Split(parts[1], " "))
		matched := checkNumbers(winingNumbers, playedNumbers)
		value := math.Floor(math.Pow(2, float64(len(matched)-1)))
		sum += int(value)
	}
	fmt.Println(sum)
}

func part2(input []string) {
	cardRgx := regexp.MustCompile(`Card\s+(\d+):`)
	fmt.Println("Part 2")
	cardsStack := make([][]Card, 0)
	for _, line := range input {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		cardId, _ := strconv.Atoi(cardRgx.FindStringSubmatch(line)[1])
		parts := strings.Split(strings.Split(line, ":")[1], "|")
		winingNumbers := collections.MapToInt(strings.Split(parts[0], " "))
		playedNumbers := collections.MapToInt(strings.Split(parts[1], " "))
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
	fmt.Println("Total", sum)
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
