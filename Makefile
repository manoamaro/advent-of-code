YEARS = 2023 2024
DAYS = $(shell seq 1 31)

generate:
	go run cmd/generator/main.go $(year) $(day)

run:
	go run cmd/aoc$(year)/day$(day)/main.go

run-all:
	$(foreach year,$(YEARS),$(foreach day,$(DAYS), if [ -d cmd/aoc$(year)/day$(day) ]; then echo "Running year $(year) day $(day)"; go run cmd/aoc$(year)/day$(day)/main.go; fi;))