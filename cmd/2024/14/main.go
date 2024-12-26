package main

import (
	"regexp"
	"slices"
	"strings"

	"github.com/rs/zerolog/log"
	"manoamaro.github.com/advent-of-code/pkg/aoc"
	"manoamaro.github.com/advent-of-code/pkg/collections"
	"manoamaro.github.com/advent-of-code/pkg/set"
	"manoamaro.github.com/advent-of-code/pkg/strings2"
)

var challenge = aoc.New(2024, 14, parseInput, part1, part2)

type robot struct {
	x, y   int
	vx, vy int
}

func main() {
	challenge.Run()
}

var width, height = 0, 0

func parseInput(input string) []robot {
	regex := regexp.MustCompile(`^p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)$`)
	res := []robot{}
	for _, line := range strings.Split(input, "\n") {
		parts := regex.FindAllStringSubmatch(line, -1)
		robot := robot{
			x:  strings2.Atoi[int](parts[0][1]),
			y:  strings2.Atoi[int](parts[0][2]),
			vx: strings2.Atoi[int](parts[0][3]),
			vy: strings2.Atoi[int](parts[0][4]),
		}
		res = append(res, robot)
	}

	for _, r := range res {
		width, height = max(width, r.x+1), max(height, r.y+1)
	}

	return res
}

func part1(input []robot) int {
	robots := slices.Clone(input)
	for range 100 {
		for j := range robots {
			move(&robots[j])
		}
	}
	count := []int{0, 0, 0, 0}
	for _, r := range robots {
		isLeft := r.x < width/2
		isRight := r.x > width/2
		isTop := r.y < height/2
		isBottom := r.y > height/2
		switch {
		case isLeft && isTop:
			count[0]++
		case isRight && isTop:
			count[1]++
		case isLeft && isBottom:
			count[2]++
		case isRight && isBottom:
			count[3]++
		}
	}
	return count[0] * count[1] * count[2] * count[3]
}

func part2(input []robot) int {
	robots := slices.Clone(input)
	seen := set.New[string]()
	count := 0
	halfTree := []int{1, 2, 3, 4, 8, 9, 13, 14, 18, 22, 23}
	treeBranches := collections.Map(halfTree, func(i int) string { return strings.Repeat("#", i) })

	for {
		count++
		places := set.New[[2]int]()
		for j := range robots {
			move(&robots[j])
			places.Add([2]int{robots[j].x, robots[j].y})
		}
		// check if robots (places) made a christmas tree
		buffer := strings.Builder{}
		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				if places.Contains([2]int{x, y}) {
					buffer.WriteRune('#')
				} else {
					buffer.WriteRune(' ')
				}
			}
			buffer.WriteRune('\n')
		}
		r := buffer.String()

		// Tries to find a christmas tree in the places.
		// After visually inspecting the output, the tree looks like this:
		// #################################
		// #                               #
		// #                               #
		// #                               #
		// #                               #
		// #                       #       #
		// #                      ##       #
		// #                  #  ###       #
		// #                 ## ####       #
		// #             #  ########       #
		// #            ## #########       #
		// #        #  #############       #
		// #       ## ##############       #
		// #      ##################       #
		// #     ######################    #
		// #    #######################    #
		// #     ######################    #
		// #      ##################       #
		// #       ## ##############       #
		// #        #  #############       #
		// #            ## #########       #
		// #             #  ########       #
		// #                 ## ####       #
		// #                  #  ###       #
		// #                      ##       #
		// #                       #       #
		// #                               #
		// #                               #
		// #                               #
		// #                               #
		// #################################
		//

		found := true
		for _, b := range treeBranches {
			if !strings.Contains(r, b) {
				found = false
				break
			}
		}

		if found {
			log.Debug().Msg("Found a tree")
			log.Debug().Msg(r)
			break
		}

		if seen.Contains(r) {
			break
		}
		seen.Add(r)
	}
	return count
}

func move(robot *robot) {
	robot.x = (robot.x + robot.vx + width) % width
	robot.y = (robot.y + robot.vy + height) % height
}
