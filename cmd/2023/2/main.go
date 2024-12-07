package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"manoamaro.github.com/advent-of-code/pkg/utils"
)

func main() {
	input, err := utils.ReadInputLines(2023, 2)
	if err != nil {
		panic(err)
	}
	part1(input)
	part2(input)
}

func part1(input []string) {
	fmt.Println("Part 1")
	games := parseInput(input)
	sum := 0
	for _, game := range games {
		if game.Matches() {
			sum += game.id
		}
	}
	fmt.Println(sum)
}

func part2(input []string) {
	games := parseInput(input)
	sum := 0
	for _, game := range games {
		maxRGB := []int{0, 0, 0}
		for _, hand := range game.hands {
			if hand.red > maxRGB[0] {
				maxRGB[0] = hand.red
			}
			if hand.green > maxRGB[1] {
				maxRGB[1] = hand.green
			}
			if hand.blue > maxRGB[2] {
				maxRGB[2] = hand.blue
			}
		}
		sum += maxRGB[0] * maxRGB[1] * maxRGB[2]
	}
	fmt.Println("Part 2")
	fmt.Println(sum)
}

func parseInput(input []string) []Game {
	reg := regexp.MustCompile(`^Game\s(\d{1,})\:\s(.+)`)
	reg2 := regexp.MustCompile(`((\d{1,})\s(green|red|blue)[,\s]?)`)
	games := make([]Game, 0)
	for _, line := range input {
		if len(line) == 0 {
			continue
		}
		gameMatch := reg.FindAllStringSubmatch(line, -1)
		gameId, err := strconv.Atoi(gameMatch[0][1])
		if err != nil {
			panic(err)
		}
		handsStr := strings.Split(gameMatch[0][2], ";")
		hands := make([]Hand, 0)
		for _, handStr := range handsStr {
			hand := Hand{}
			handColors := reg2.FindAllStringSubmatch(handStr, -1)
			for _, handColor := range handColors {
				color := handColor[3]
				count, err := strconv.Atoi(handColor[2])
				if err != nil {
					panic(err)
				}
				switch color {
				case "blue":
					hand.blue = count
				case "red":
					hand.red = count
				case "green":
					hand.green = count
				}
			}
			hands = append(hands, hand)
		}
		games = append(games, Game{
			id:    gameId,
			hands: hands,
		})
	}
	return games
}

type Game struct {
	id    int
	hands []Hand
}

func (g Game) Matches() bool {
	for _, hand := range g.hands {
		if !hand.Matches() {
			return false
		}
	}
	return true
}

func (g Game) String() string {
	return fmt.Sprintf("Game{id: %d, matches: %v, hands: %v}", g.id, g.Matches(), g.hands)
}

type Hand struct {
	blue  int
	red   int
	green int
}

func (h Hand) String() string {
	return fmt.Sprintf("Hand{blue: %d, red: %d, green: %d}", h.blue, h.red, h.green)
}

func (h Hand) Matches() bool {
	return h.red <= 12 && h.green <= 13 && h.blue <= 14
}
