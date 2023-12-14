generate:
	go run cmd/generator/main.go $(day)

run:
	go run cmd/day$(day)/main.go
