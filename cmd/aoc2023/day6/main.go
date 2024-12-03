package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"manoamaro.github.com/advent-of-code/pkg/collections"
	"manoamaro.github.com/advent-of-code/pkg/utils"
)

func main() {
	input, err := utils.ReadInputLines(2023, 6)
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
	races := parseMultiple(input)
	sums := make([]int, 0)
	for _, race := range races {
		sum := 0
		for d := 0; d <= race.duration; d++ {
			remaining := race.duration - d
			speed := d
			totalDistance := speed * remaining
			if totalDistance > race.distance {
				sum += 1
			}
		}
		sums = append(sums, sum)
	}
	t := 1
	for _, sum := range sums {
		t *= sum
	}
	fmt.Println(t)
}

func part2(input []string) {
	fmt.Println("Part 2")
	race := parse(input)
	sum := 0
	for d := 0; d <= race.duration; d++ {
		remaining := race.duration - d
		speed := d
		totalDistance := speed * remaining
		if totalDistance > race.distance {
			sum += 1
		}
	}
	fmt.Println(sum)
}

func parseMultiple(input []string) []Race {
	durationsRaw := strings.Split(strings.Split(input[0], ":")[1], " ")
	durations := collections.MapToInt(durationsRaw)

	distancesRaw := strings.Split(strings.Split(input[1], ":")[1], " ")
	distances := collections.MapToInt(distancesRaw)

	races := make([]Race, 0)
	for i := 0; i < len(durations); i++ {
		races = append(races, Race{duration: durations[i], distance: distances[i]})
	}
	return races
}

func parse(input []string) Race {
	durationRaw := strings.ReplaceAll(strings.Split(input[0], ":")[1], " ", "")
	duration, err := strconv.Atoi(durationRaw)
	if err != nil {
		panic(err)
	}

	distanceRaw := strings.ReplaceAll(strings.Split(input[1], ":")[1], " ", "")
	distance, err := strconv.Atoi(distanceRaw)
	if err != nil {
		panic(err)
	}

	return Race{duration: duration, distance: distance}
}

type Race struct {
	duration int
	distance int
}
