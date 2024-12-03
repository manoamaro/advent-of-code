package main

import (
	"fmt"
	"math"
	"strings"
	"sync"
	"time"

	"manoamaro.github.com/advent-of-code/pkg/collections"
	"manoamaro.github.com/advent-of-code/pkg/utils"
)

func main() {
	input, err := utils.ReadInputLines(2023, 5)
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
	almanac := parse(input)
	waitGroup := sync.WaitGroup{}
	locations := make(chan int)
	for _, seed := range almanac.Seeds {
		waitGroup.Add(1)
		go func(seed int) {
			location := MapSeedToLocation(almanac, seed)
			locations <- location
			waitGroup.Done()
		}(seed)
	}
	go func() {
		waitGroup.Wait()
		close(locations)
	}()

	lowestLocation := math.MaxInt
	for location := range locations {
		if location < lowestLocation {
			lowestLocation = location
		}
	}
	fmt.Println(lowestLocation)
}

func part2(input []string) {
	fmt.Println("Part 2")
	almanac := parse(input)
	realSeeds := [][]int{}
	for idx, seed := range almanac.Seeds {
		if idx%2 == 0 {
			realSeeds = append(realSeeds, []int{seed, seed + almanac.Seeds[idx+1]})
		}
	}

	waitGroup := sync.WaitGroup{}
	locations := make(chan int)

	for _, seedRange := range realSeeds {
		waitGroup.Add(1)
		go func(seedRange []int) {
			fmt.Println("Starting SeedRange:", seedRange, seedRange[1]-seedRange[0])
			intLowestLocation := math.MaxInt
			for seed := seedRange[0]; seed < seedRange[1]; seed++ {
				location := MapSeedToLocation(almanac, seed)
				if location < intLowestLocation {
					intLowestLocation = location
				}
			}
			locations <- intLowestLocation
			fmt.Println("Finished SeedRange:", seedRange, intLowestLocation)
			waitGroup.Done()
		}(seedRange)
	}

	go func() {
		waitGroup.Wait()
		close(locations)
	}()

	lowestLocation := math.MaxInt
	for location := range locations {
		if location < lowestLocation {
			lowestLocation = location
		}
	}
	fmt.Println(lowestLocation)

}

type Almanac struct {
	Seeds                 []int
	SeedToSoil            []MapRange
	SoilToFertilizer      []MapRange
	FertilizerToWater     []MapRange
	WaterToLight          []MapRange
	LightToTemperature    []MapRange
	TemperatureToHumidity []MapRange
	HumidityToLocation    []MapRange
}

func (a Almanac) String() string {
	return fmt.Sprintf("Seeds: %v\nSeedToSoil: %v\nSoilToFertilizer: %v\nFertilizerToWater: %v\nWaterToLight: %v\nLightToTemperature: %v\nTemperatureToHumidity: %v\nHumidityToLocation: %v", a.Seeds, a.SeedToSoil, a.SoilToFertilizer, a.FertilizerToWater, a.WaterToLight, a.LightToTemperature, a.TemperatureToHumidity, a.HumidityToLocation)
}

func MapSeedToLocation(almanac Almanac, seed int) int {
	soil := MapRangeTo(seed, almanac.SeedToSoil)
	fertilizer := MapRangeTo(soil, almanac.SoilToFertilizer)
	water := MapRangeTo(fertilizer, almanac.FertilizerToWater)
	light := MapRangeTo(water, almanac.WaterToLight)
	temperature := MapRangeTo(light, almanac.LightToTemperature)
	humidity := MapRangeTo(temperature, almanac.TemperatureToHumidity)
	location := MapRangeTo(humidity, almanac.HumidityToLocation)
	return location
}

type MapRange struct {
	DestinationStart int
	SourceStart      int
	Range            int
}

func MapRangeTo(value int, mapRanges []MapRange) int {
	for _, m := range mapRanges {
		if value >= m.SourceStart && value < m.SourceStart+m.Range {
			offset := value - m.SourceStart
			return m.DestinationStart + offset
		}
	}
	return value
}

func (m MapRange) String() string {
	return fmt.Sprintf("Dest: %d, Src: %d, Range: %d;", m.DestinationStart, m.SourceStart, m.Range)
}

func parse(input []string) Almanac {
	almanac := Almanac{}
	parsingStep := "seeds"
	for _, line := range input {
		if len(line) == 0 {
			continue
		}
		if strings.Contains(line, "seeds:") {
			seeds := collections.MapToInt(strings.Split(strings.Split(line, ":")[1], " "))
			almanac.Seeds = seeds
		} else if strings.Contains(line, "seed-to-soil map:") {
			parsingStep = "seed-to-soil"
			continue
		} else if strings.Contains(line, "soil-to-fertilizer map:") {
			parsingStep = "soil-to-fertilizer"
			continue
		} else if strings.Contains(line, "fertilizer-to-water map:") {
			parsingStep = "fertilizer-to-water"
			continue
		} else if strings.Contains(line, "water-to-light map:") {
			parsingStep = "water-to-light"
			continue
		} else if strings.Contains(line, "light-to-temperature map:") {
			parsingStep = "light-to-temperature"
			continue
		} else if strings.Contains(line, "temperature-to-humidity map:") {
			parsingStep = "temperature-to-humidity"
			continue
		} else if strings.Contains(line, "humidity-to-location map:") {
			parsingStep = "humidity-to-location"
			continue
		}
		lineSplit := strings.Split(line, " ")
		values := collections.MapToInt(lineSplit)
		if parsingStep == "seed-to-soil" {
			almanac.SeedToSoil = append(almanac.SeedToSoil, MapRange{values[0], values[1], values[2]})
		} else if parsingStep == "soil-to-fertilizer" {
			almanac.SoilToFertilizer = append(almanac.SoilToFertilizer, MapRange{values[0], values[1], values[2]})
		} else if parsingStep == "fertilizer-to-water" {
			almanac.FertilizerToWater = append(almanac.FertilizerToWater, MapRange{values[0], values[1], values[2]})
		} else if parsingStep == "water-to-light" {
			almanac.WaterToLight = append(almanac.WaterToLight, MapRange{values[0], values[1], values[2]})
		} else if parsingStep == "light-to-temperature" {
			almanac.LightToTemperature = append(almanac.LightToTemperature, MapRange{values[0], values[1], values[2]})
		} else if parsingStep == "temperature-to-humidity" {
			almanac.TemperatureToHumidity = append(almanac.TemperatureToHumidity, MapRange{values[0], values[1], values[2]})
		} else if parsingStep == "humidity-to-location" {
			almanac.HumidityToLocation = append(almanac.HumidityToLocation, MapRange{values[0], values[1], values[2]})
		}
	}
	return almanac
}
