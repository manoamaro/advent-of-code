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
	inputFile                string
	debug, runTests          bool
	runPart                  int
}

func New[T any, R comparable](year, day int, inputProcessor InputProcessor[T], part1Solver Solver[T, R], part2Solver Solver[T, R]) *Challenge[T, R] {
	challenge := &Challenge[T, R]{
		year:           year,
		day:            day,
		inputProcessor: inputProcessor,
		part1Solver:    part1Solver,
		part2Solver:    part2Solver,
	}
	challenge.setup()
	return challenge
}

func (d *Challenge[T, R]) setup() {
	// Parse args flags
	debug := flag.Bool("debug", false, "sets log level to debug")
	part := flag.Int("p", 0, "runs only the specified part")
	runTests := flag.Bool("test", false, "runs tests")
	flag.Parse()
	// Set up logger
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	d.debug = *debug
	d.runPart = *part
	d.runTests = *runTests
	// input file if provided
	d.inputFile = flag.Arg(0)
}

func (d *Challenge[T, R]) Run() {
	defer func() {
		if r := recover(); r != nil {
			log.Error().Msgf("Recovered from panic: %v", r)
		}
	}()
	// Get input
	var rawInput string
	if d.inputFile != "" {
		f := errors.Must(os.ReadFile(d.inputFile))
		rawInput = string(f)
	} else {
		rawInput = GetInput(d.year, d.day)
	}
	rawInput = strings.TrimSpace(rawInput)

	d.solvePart1(rawInput)
	d.solvePart2(rawInput)
}

func (d *Challenge[T, R]) TestPart1(file string, expected R) {
	if !d.runTests {
		return
	}
	rawInput := errors.Must(os.ReadFile(file))
	p1 := d.solvePart1(string(rawInput))
	if p1 != expected {
		log.Error().Msgf("Expected %v, got %v", expected, p1)
	} else {
		log.Info().Msgf("Test passed")
	}
}

func (d *Challenge[T, R]) TestPart2(file string, expected R) {
	if !d.runTests {
		return
	}
	rawInput := errors.Must(os.ReadFile(file))
	p2 := d.solvePart2(string(rawInput))
	if p2 != expected {
		log.Error().Msgf("Expected %v, got %v", expected, p2)
	} else {
		log.Info().Msgf("Test passed")
	}
}

func (d *Challenge[T, R]) processInput(input string) T {
	startTime := time.Now()
	in := d.inputProcessor(input)
	log.Debug().Msgf("Input processing took %v", time.Since(startTime))
	return in
}

func (d *Challenge[T, R]) solvePart1(rawInput string) R {
	var r R
	if d.runPart == 0 || d.runPart == 1 {
		input := d.processInput(rawInput)
		startTime := time.Now()
		r = d.part1Solver(input)
		log.Debug().Msgf("Part 1 took %v", time.Since(startTime))
		log.Info().Msgf("%d %d Part 1: %v", d.year, d.day, r)
	}
	return r
}

func (d *Challenge[T, R]) solvePart2(rawInput string) R {
	var r R
	if d.runPart == 0 || d.runPart == 2 {
		input := d.processInput(rawInput)
		startTime := time.Now()
		r = d.part2Solver(input)
		log.Debug().Msgf("Part 2 took %v", time.Since(startTime))
		log.Info().Msgf("%d %d Part 2: %v", d.year, d.day, r)
	}
	return r
}
