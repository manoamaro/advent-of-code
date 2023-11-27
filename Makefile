day.%: DAY=$*
day.%:
	go run cmd/day$(DAY)/main.go

generate:
	go run cmd/generator/main.go $(day)