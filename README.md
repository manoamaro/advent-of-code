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

### Running a Solution

Each day has its own command under `cmd/<YEAR>/<DAY>`. Run a specific
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
