generate:
	go run cmd/generator/main.go $(year) $(day)

run:
	go run cmd/$(year)/$(day)/main.go
