package aoc

import (
    "flag"
    "os"
    "strings"
    "sync"
    "testing"
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
    setupOnce                sync.Once
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

func (d *Challenge[T, R]) setup() {
    // Parse args flags
    debug := flag.Bool("debug", false, "sets log level to debug")
    part := flag.Int("p", 0, "runs only the specified part")
    runTests := flag.Bool("test", false, "runs tests")
    // testing.Init registers the flags used by the testing package so
    // that calling flag.Parse() doesn't error when running `go test`.
    testing.Init()
    flag.Parse()
    // Set up logger
    zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
    zerolog.DurationFieldUnit = time.Microsecond
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

func (d *Challenge[T, R]) ensureSetup() {
    d.setupOnce.Do(d.setup)
}

func (d *Challenge[T, R]) Run() {
    d.ensureSetup()
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
    d.ensureSetup()
    if !d.runTests {
        return
    }
    rawInput := errors.Must(os.ReadFile(file))
    p1 := d.solvePart1(string(rawInput))
    if p1 != expected {
        log.Error().Any("expected", expected).Any("got", p1).Msg("Test Part 1 failed.")
    } else {
        log.Info().Msg("Test Part 1 passed")
    }
}

func (d *Challenge[T, R]) TestPart2(file string, expected R) {
    d.ensureSetup()
    if !d.runTests {
        return
    }
    rawInput := errors.Must(os.ReadFile(file))
    p2 := d.solvePart2(string(rawInput))
    if p2 != expected {
        log.Error().Any("expected", expected).Any("got", p2).Msg("Test Part 2 failed.")
    } else {
        log.Info().Msg("Test Part2 passed")
    }
}

func (d *Challenge[T, R]) processInput(input string) T {
    startTime := time.Now()
    in := d.inputProcessor(input)
    log.Debug().Dur("took", time.Since(startTime)).Msgf("Finished input")
    return in
}

func (d *Challenge[T, R]) solvePart1(rawInput string) R {
    var r R
    if d.runPart == 0 || d.runPart == 1 {
        input := d.processInput(rawInput)
        startTime := time.Now()
        r = d.part1Solver(input)
        log.Debug().Dur("took", time.Since(startTime)).Msgf("Part 1")
        log.Info().Int("year", d.year).Int("day", d.day).Any("result", r).Msg("Part 1")
    }
    return r
}

func (d *Challenge[T, R]) solvePart2(rawInput string) R {
    var r R
    if d.runPart == 0 || d.runPart == 2 {
        input := d.processInput(rawInput)
        startTime := time.Now()
        r = d.part2Solver(input)
        log.Debug().Dur("took", time.Since(startTime)).Msg("Part 2")
        log.Info().Int("year", d.year).Int("day", d.day).Any("result", r).Msg("Part 2")
    }
    return r
}
