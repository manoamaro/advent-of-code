package main

import (
	"strconv"
	"strings"

	"manoamaro.github.com/advent-of-code/pkg/aoc"
	"manoamaro.github.com/advent-of-code/pkg/strutil"
)

var challenge = aoc.New(2023, 6, aoc.LinesProcessor(), part1, part2)

func main() {
	challenge.Run()
}

func part1(input []string) int {
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
	return t
}

func part2(input []string) int {
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
	return sum
}

func parseMultiple(input []string) []Race {
	durationsRaw := strings.Split(strings.Split(input[0], ":")[1], " ")
	durations := strutil.MapToInt(durationsRaw)

	distancesRaw := strings.Split(strings.Split(input[1], ":")[1], " ")
	distances := strutil.MapToInt(distancesRaw)

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
