# Advent of Code
This repository contains my solutions to the [Advent of Code](https://adventofcode.com/) challenges.

## 2024
![2024](https://img.shields.io/badge/stars%20⭐-50-yellow)
![2024](https://img.shields.io/badge/days%20completed-25-red)

## 2023
![2023](https://img.shields.io/badge/stars%20⭐-38-yellow)
![2023](https://img.shields.io/badge/days%20completed-19-red)

## 2022
![2022](https://img.shields.io/badge/stars%20⭐-15-yellow)
![2022](https://img.shields.io/badge/days%20completed-7-red)

## 2021
![2021](https://img.shields.io/badge/stars%20⭐-50-yellow)
![2021](https://img.shields.io/badge/days%20completed-25-red)

## Getting Started

### Prerequisites

- Go 1.25 or later
- Make

Note on Go toolchain

- The repo pins `toolchain go1.25.0` in `go.mod`. If prompted, install it with:
  `go toolchain install go1.25.0`
- Alternatively ensure a Go 1.25 binary is on PATH (e.g., macOS: `brew install go@1.25 && brew link --overwrite go@1.25`).

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/advent-of-code.git
    cd advent-of-code
    ```

2. Install dependencies:
    ```sh
    go mod download
    ```

## Usage

### Meta-Runner (recommended)

Run any day with a single command. Days are zero-padded (e.g., `05`).

```sh
go run ./cmd/aoc -y 2024 -d 05              # run both parts
go run ./cmd/aoc -y 2024 -d 05 -p 1 --debug # only part 1, with debug logs
go run ./cmd/aoc -y 2024 -d 05 --input path/to/input.txt
```

Using Make:

```sh
make run y=2024 d=05            # run both parts
make run y=2024 d=05 p=2        # only part 2
make run y=2024 d=05 debug=1    # enable debug logs
make run y=2024 d=05 test=1     # enable embedded sample tests
```

### Running a Solution

Each day has its own command under `cmd/<YEAR>/<DAY>` (day zero‑padded, e.g., `05`). Run a specific
solution directly with `go run`:

```sh
go run ./cmd/2023/08
```

### Running All Solutions

To execute every solution for a given year, iterate over the directories in
`cmd/<YEAR>`:

```sh
for dir in cmd/2023/*; do
    go run "./$dir"
done
```
