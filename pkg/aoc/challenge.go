package aoc

import (
	"flag"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"manoamaro.github.com/advent-of-code/pkg/errors"
)

type Solver[T any, R comparable] func(T) R

type Challenge[T any, R comparable] struct {
	year, day                int
	inputProcessor           InputProcessor[T]
	part1Solver, part2Solver Solver[T, R]
}

func New[T any, R comparable](year, day int, inputProcessor InputProcessor[T], part1Solver Solver[T, R], part2Solver Solver[T, R]) *Challenge[T, R] {
	return &Challenge[T, R]{
		year:           year,
		day:            day,
		inputProcessor: inputProcessor,
		part1Solver:    part1Solver,
		part2Solver:    part2Solver,
	}
}

func (d *Challenge[T, R]) Run() {
	defer func() {
		if r := recover(); r != nil {
			log.Error().Msgf("Recovered from panic: %v", r)
		}
	}()

	// Parse args flags
	debug := flag.Bool("debug", false, "sets log level to debug")
	part := flag.Int("p", 0, "runs only the specified part")
	flag.Parse()

	// Set up logger
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// input file if provided
	inputFile := flag.Arg(0)

	// Get input
	var rawInput string
	if inputFile != "" {
		f := errors.Must(os.ReadFile(inputFile))
		rawInput = string(f)
	} else {
		rawInput = GetInput(d.year, d.day)
	}
	rawInput = strings.TrimSpace(rawInput)
	input := d.processInput(rawInput)

	var p1, p2 R
	if *part == 1 {
		p1 = d.solvePart1(input)
	} else if *part == 2 {
		p2 = d.solvePart2(input)
	} else {
		p1, p2 = d.solvePart1(input), d.solvePart2(input)
	}

	log.Info().Msgf("%d %d Part 1: %v", d.year, d.day, p1)
	log.Info().Msgf("%d %d Part 2: %v", d.year, d.day, p2)
}

func (d *Challenge[T, R]) processInput(input string) T {
	startTime := time.Now()
	in := d.inputProcessor(input)
	log.Debug().Msgf("Input processing took %v", time.Since(startTime))
	return in
}

func (d *Challenge[T, R]) solvePart1(input T) R {
	startTime := time.Now()
	p1 := d.part1Solver(input)
	log.Debug().Msgf("Part 1 took %v", time.Since(startTime))
	return p1
}

func (d *Challenge[T, R]) solvePart2(input T) R {
	startTime := time.Now()
	p1 := d.part2Solver(input)
	log.Debug().Msgf("Part 2 took %v", time.Since(startTime))
	return p1
}
