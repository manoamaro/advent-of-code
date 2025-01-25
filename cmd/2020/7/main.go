package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
	"manoamaro.github.com/advent-of-code/pkg/aoc"
)

var challenge = aoc.New(2020, 7, parseInput, part1, part2)

func main() {
	challenge.Run()
}

func parseInput(input string) Bags {
	bags := make(Bags, 0)
	regexInnerBags, _ := regexp.Compile(`(\d)\s?([\w|\s]+)\sbags?`)

	for _, rawBags := range strings.Split(input, "\n") {
		parts := strings.Split(rawBags, " bags contain ")
		outerBagName := strings.TrimSpace(parts[0])
		outerBag := bags.findOrCreateBag(outerBagName)
		matches := regexInnerBags.FindAllStringSubmatch(parts[1], -1)
		for _, innerMatch := range matches {
			innerBagQuantity, err := strconv.Atoi(innerMatch[1])
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to parse inner bag quantity")
			}
			innerBagName := strings.TrimSpace(innerMatch[2])
			innerBag := bags.findOrCreateBag(innerBagName)
			outerBag.children = append(outerBag.children, &BagContains{
				quantity: innerBagQuantity,
				bag:      innerBag,
			})
		}
	}
	return bags
}

func part1(bags Bags) int {
	bagsContaining := make([]*Bag, 0)

	for _, bag := range bags {
		log.Debug().Msgf("Checking bag %s", bag)
		if bag.fitsBag("shiny gold") {
			bagsContaining = append(bagsContaining, bag)
		}
	}
	return len(bagsContaining)
}

func part2(bags Bags) int {
	shinyGoldBag := bags.findOrCreateBag("shiny gold")
	return shinyGoldBag.countInsideBags()
}

type BagContains struct {
	quantity int
	bag      *Bag
}
type Bag struct {
	name     string
	children []*BagContains
}

type Bags []*Bag

func (r *Bags) findOrCreateBag(name string) *Bag {
	for _, bag := range *r {
		if bag.name == name {
			return bag
		}
	}

	newBag := &Bag{
		name: name,
	}

	*r = append(*r, newBag)

	return newBag
}

func (b *Bag) String() string {
	children := make([]string, 0)
	for _, child := range b.children {
		children = append(children, fmt.Sprintf("%d %s", child.quantity, child.bag.name))
	}
	return fmt.Sprintf("%s bag contains %s", b.name, strings.Join(children, ";"))
}

func (b *Bag) fitsBag(name string) bool {
	for _, child := range b.children {
		if child.bag.name == name || child.bag.fitsBag(name) {
			return true
		}
	}
	return false
}

func (b *Bag) countInsideBags() (count int) {
	for _, child := range b.children {
		count += child.quantity + (child.quantity * child.bag.countInsideBags())
	}
	return
}
