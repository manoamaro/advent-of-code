package main

import (
	"sort"
	"strconv"
	"strings"

	"manoamaro.github.com/advent-of-code/pkg/aoc"
)

var challenge = aoc.New(2023, 7, parseHands, part1, part2)

func main() {
	challenge.Run()
}

func part1(hands []*hand) int {

	sort.Slice(hands, func(i, j int) bool {
		return compareToPart1(*hands[i], *hands[j])
	})

	maxRank := len(hands)
	total := 0
	for i, h := range hands {
		total += h.bid * (maxRank - i)
	}
	return total
}

func part2(hands []*hand) int {
	for _, h := range hands {
		h.value = valueWithJoker(*h)
	}

	sort.Slice(hands, func(i, j int) bool {
		return compareToPart2(*hands[i], *hands[j])
	})

	maxRank := len(hands)
	total := 0
	for i, h := range hands {
		total += h.bid * (maxRank - i)
	}
	return total
}

var cardValuesPart1 = map[rune]int{
	'2': 1,
	'3': 2,
	'4': 3,
	'5': 4,
	'6': 5,
	'7': 6,
	'8': 7,
	'9': 8,
	'T': 9,
	'J': 10,
	'Q': 11,
	'K': 12,
	'A': 13,
}

var cardValuesPart2 = map[rune]int{
	'J': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'Q': 11,
	'K': 12,
	'A': 13,
}

func parseHands(input string) []*hand {
	hands := make([]*hand, 0)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		bid, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		hands = append(hands, newHand(parts[0], bid))
	}
	return hands
}

type hand struct {
	cardsRaw string
	bid      int
	value    int
}

func newHand(cardsRaw string, bid int) *hand {
	return &hand{
		cardsRaw: cardsRaw,
		bid:      bid,
	}
}

func (h hand) Cards(values map[rune]int) []int {
	cards := make([]int, 0)
	for _, c := range h.cardsRaw {
		cards = append(cards, values[c])
	}
	return cards
}

func (h hand) HasJoker() bool {
	return strings.Contains(h.cardsRaw, "J")
}

func compareToPart1(hand1, hand2 hand) bool {
	h1v := hand1.Value(cardValuesPart1)
	h2v := hand2.Value(cardValuesPart1)
	if h1v == h2v {
		for k := 0; k < 5; k++ {
			h1c := hand1.cardsRaw[k]
			h2c := hand2.cardsRaw[k]
			if h1c != h2c {
				h1v += cardValuesPart1[rune(h1c)]
				h2v += cardValuesPart1[rune(h2c)]
				break
			}
		}
	}
	return h1v > h2v
}

func compareToPart2(hand1, hand2 hand) bool {
	h1v := hand1.value
	h2v := hand2.value
	if h1v == h2v {
		for k := 0; k < 5; k++ {
			h1c := hand1.cardsRaw[k]
			h2c := hand2.cardsRaw[k]
			if h1c != h2c {
				h1v += cardValuesPart2[rune(h1c)]
				h2v += cardValuesPart2[rune(h2c)]
				break
			}
		}
	}
	return h1v > h2v
}
func valueWithJoker(h hand) int {
	if !h.HasJoker() {
		return h.Value(cardValuesPart2)
	}
	biggest := h.Value(cardValuesPart2)
	for i := 0; i < 5; i++ {
		if h.cardsRaw[i] == 'J' {
			for k := range cardValuesPart2 {
				if k == 'J' {
					continue
				}
				newHand := newHand(h.cardsRaw, h.bid)
				newHand.cardsRaw = newHand.cardsRaw[:i] + string(k) + newHand.cardsRaw[i+1:]
				v := valueWithJoker(*newHand)
				if v > biggest {
					biggest = v
				}
			}
		}
	}
	return biggest
}

func (h hand) Value(values map[rune]int) int {
	cards := h.Cards(values)
	sort.Ints(cards)
	// Five of a kind
	if cards[0] == cards[4] {
		return 1000000
	}
	// Four of a kind
	if cards[0] == cards[3] || cards[1] == cards[4] {
		return 900000
	}
	// Full house
	if cards[0] == cards[2] && cards[3] == cards[4] {
		return 800000
	}
	if cards[0] == cards[1] && cards[2] == cards[4] {
		return 800000
	}
	// three of a kind
	if cards[0] == cards[2] && cards[3] != cards[4] {
		return 700000
	}
	if cards[1] == cards[3] && cards[0] != cards[4] {
		return 700000
	}
	if cards[2] == cards[4] && cards[0] != cards[1] {
		return 700000
	}
	// Two pair
	if cards[0] == cards[1] && cards[2] == cards[3] {
		return 600000
	}
	if cards[1] == cards[2] && cards[3] == cards[4] {
		return 600000
	}
	if cards[0] == cards[1] && cards[3] == cards[4] {
		return 600000
	}
	// One pair
	for i := 0; i < 4; i++ {
		if cards[i] == cards[i+1] {
			return 500000
		}
	}
	// High card
	return 400000
}
