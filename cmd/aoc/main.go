package main

import (
    "flag"
    "fmt"
    "os"
    "os/exec"
)

func main() {
    year := flag.Int("y", 0, "Year (e.g., 2024)")
    day := flag.Int("d", 0, "Day (1-25)")
    part := flag.Int("p", 0, "Optional: run only part 1 or 2")
    debug := flag.Bool("debug", false, "Optional: enable debug logs")
    runTests := flag.Bool("test", false, "Optional: run embedded tests for the day")
    input := flag.String("input", "", "Optional: path to input file (overrides network fetch)")
    flag.Parse()

    if *year == 0 || *day == 0 {
        fmt.Fprintln(os.Stderr, "Usage: aoc -y <year> -d <day> [-p 1|2] [--debug] [--test] [--input file]")
        os.Exit(2)
    }

    dayFolder := fmt.Sprintf("%02d", *day)
    target := fmt.Sprintf("./cmd/%d/%s", *year, dayFolder)

    // Build args for the subcommand (pass through flags understood by the day runner)
    args := []string{"run", target}
    if *debug {
        args = append(args, "-debug")
    }
    if *runTests {
        args = append(args, "-test")
    }
    if *part == 1 || *part == 2 {
        args = append(args, "-p", fmt.Sprintf("%d", *part))
    }
    if *input != "" {
        args = append(args, *input)
    }

    cmd := exec.Command("go", args...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Stdin = os.Stdin
    if err := cmd.Run(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
        os.Exit(1)
    }
}

