package aoc

import (
	"flag"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"manoamaro.github.com/advent-of-code/pkg/errors"
	"manoamaro.github.com/advent-of-code/pkg/utils"
)

type Challenge[T any, R comparable] struct {
	year, day                int
	inputProcessor           InputProcessor[T]
	part1Solver, part2Solver func(T) (R, error)
}

func New[T any, R comparable](year, day int, inputProcessor InputProcessor[T], part1Solver func(T) (R, error), part2Solver func(T) (R, error)) *Challenge[T, R] {
	return &Challenge[T, R]{
		year:           year,
		day:            day,
		inputProcessor: inputProcessor,
		part1Solver:    part1Solver,
		part2Solver:    part2Solver,
	}
}

func (d *Challenge[T, R]) Run() {
	debug := flag.Bool("debug", false, "sets log level to debug")
	part := flag.Int("p", 0, "runs only the specified part")
	flag.Parse()
	inputFile := flag.Arg(0)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	var rawInput string
	if inputFile != "" {
		f := errors.Must(os.ReadFile(inputFile))
		rawInput = string(f)
	} else {
		rawInput = errors.Must(utils.ReadInput(d.year, d.day))
	}
	input := d.processInput(rawInput)
	var p1, p2 R
	if *part == 1 {
		p1 = d.solvePart1(input)
	} else if *part == 2 {
		p2 = d.solvePart2(input)
	} else {
		p1, p2 = d.Solve(input)
	}
	log.Info().Msgf("%d %d Part 1: %v", d.year, d.day, p1)
	log.Info().Msgf("%d %d Part 2: %v", d.year, d.day, p2)
}

func (d *Challenge[T, R]) processInput(input string) T {
	startTime := time.Now()
	in := errors.Must(d.inputProcessor(input))
	log.Debug().Msgf("Input processing took %v", time.Since(startTime))
	return in
}

func (d *Challenge[T, R]) Solve(input T) (R, R) {
	startTimeAll := time.Now()
	p1 := d.solvePart1(input)
	p2 := d.solvePart2(input)
	log.Debug().Msgf("All took %v", time.Since(startTimeAll))
	return p1, p2
}

func (d *Challenge[T, R]) solvePart1(input T) R {
	startTime := time.Now()
	p1 := errors.Must(d.part1Solver(input))
	log.Debug().Msgf("Part 1 took %v", time.Since(startTime))
	return p1
}

func (d *Challenge[T, R]) solvePart2(input T) R {
	startTime := time.Now()
	p1 := errors.Must(d.part2Solver(input))
	log.Debug().Msgf("Part 1 took %v", time.Since(startTime))
	return p1
}
