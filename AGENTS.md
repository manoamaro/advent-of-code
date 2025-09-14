# Repository Guidelines

## Project Structure & Module Organization
- `cmd/<YEAR>/<DAY>`: Individual AoC solutions (each is a `main` package). Example: `cmd/2024/05`.
- `cmd/generator`: Scaffolds new days from `day.tmpl` (`go run ./cmd/generator -y 2024 -d 5`). Uses zero‑padded day folders (`05`).
- `pkg/*`: Reusable libs.
  - `sliceutil`: slice helpers (Map, Fold, Reverse, SlideSeq, Sum).
  - `strutil`: string helpers (ReverseString, Atoi, MapToInt).
  - `mapx`: typed map helpers (Map[K,V], MultiMap, GetOr, GetOrPanic).
  - `mathx`: math + generics (GCD/LCM, Abs/Min/Max, ManhattanDistance).
  - Others: `grid`, `queue`, `deque`, `set`, `graph`, `aoc`.
- `internal/aoc`: Repo‑specific challenge runner and input handling (imports changed from `pkg/aoc`).
- `test/`: Cross‑package tests only.
- `.env.example`: Template for local config. `.cache/` is created at runtime for input caching.

## Build, Test, and Development Commands
- Install deps: `go mod download`
- Format & vet: `go fmt ./... && go vet ./...`
- Run a day: `go run ./cmd/2024/05 [-debug] [-p 1|2] [<input_file>]`
  - Flags come from `internal/aoc.Challenge`: `-debug` enables verbose logs, `-p` runs a single part, optional input file overrides online input.
- Build a day: `go build ./cmd/2024/05`
- Run tests: `go test ./...` (unit tests in `pkg/*`, some day tests in `cmd/<YEAR>/<DAY>`). Coverage: `go test ./... -cover`.

## Coding Style & Naming Conventions
- Go 1.25+. Use `gofmt` defaults; keep functions small and focused.
- Package names are lower‑case (`pkg/set`, `pkg/queue`). Files use snake_case where helpful (e.g., `priority_queue.go`).
- Tests end with `_test.go`. Prefer descriptive names; avoid one‑letter identifiers outside short loops.
- Logging via `zerolog`; error handling often uses `pkg/errors.Must(...)` for brevity in tools/CLIs.

## Testing Guidelines
- Colocate unit tests with code in `pkg/*` (use `package_name_test` for black‑box; same package for white‑box where needed).
- Reserve `/test` for cross‑package/integration tests only; consider `//go:build integration` tags for slower suites.
- Use `testdata/` folders for fixtures the package should read (ignored by `go build`).
- Day solutions: keep sample inputs in the day folder (e.g., `cmd/2024/24/input_test_1.txt`) and, when helpful, call `challenge.TestPart1/2(...)` from `main.go`.
- Keep `go test ./...` fast and green; add focused tests for new utilities.

## Commit & Pull Request Guidelines
- Commits: short, imperative subject. Examples: `fix: flag parsing in aoc challenge`, `2024: day 24 final`.
- PRs: include a clear description, scope (year/day or package), any linked issues, and notes on performance or complexity when relevant. Ensure `go fmt`, `go vet`, and `go test ./...` pass.

## Security & Configuration
- Inputs may be fetched online using `AOC_SESSION` from `.env`. Copy `.env.example` to `.env` and set `AOC_SESSION=<your_session_cookie>`.
- Do not commit `.env` or `.cache/` (ignored by `.gitignore`).
