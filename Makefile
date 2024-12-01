generate:
	go run cmd/generator/main.go $(year) $(day)

run:
	go run cmd/aoc$(year)/day$(day)/main.go
