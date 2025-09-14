SHELL := /bin/bash

.PHONY: run test vet fmt

# Run a specific year/day using the meta-runner.
# Usage: make run y=2024 d=05 [p=1|2] [debug=1] [test=1] [input=path]
run:
	@ if [ -z "$(y)" ] || [ -z "$(d)" ]; then \
		echo "Usage: make run y=<year> d=<dd> [p=1|2] [debug=1] [test=1] [input=path]"; \
		exit 2; \
	fi; \
	FLAGS=""; \
	if [ "$(debug)" = "1" ]; then FLAGS="$$FLAGS --debug"; fi; \
	if [ "$(test)" = "1" ]; then FLAGS="$$FLAGS --test"; fi; \
	if [ -n "$(p)" ]; then FLAGS="$$FLAGS -p $(p)"; fi; \
	if [ -n "$(input)" ]; then FLAGS="$$FLAGS --input $(input)"; fi; \
	go run ./cmd/aoc -y $(y) -d $(d) $$FLAGS

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test ./...

