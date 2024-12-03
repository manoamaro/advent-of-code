package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Void struct{}

func ReadInput(year int, day int) (string, error) {

	if cached := readInputFromCache(year, day); len(cached) > 0 {
		return cached, nil
	}

	godotenv.Load()

	session := os.Getenv("AOC_SESSION")
	if len(session) == 0 {
		return "", fmt.Errorf("cannot find session")
	}

	request := Must(http.NewRequest(http.MethodGet, fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day), nil))

	request.AddCookie(&http.Cookie{
		Name:  "session",
		Value: session,
	})

	response := Must(http.DefaultClient.Do(request))

	defer response.Body.Close()

	b := Must(io.ReadAll(response.Body))

	input := string(b)

	saveToCache(year, day, input)

	return input, nil
}

func ReadInputLines(year int, day int) ([]string, error) {
	input, err := ReadInput(year, day)
	input = strings.TrimSpace(input)
	if err != nil {
		return nil, err
	}
	return strings.Split(input, "\n"), nil
}

func ReadInputSlice2d(year int, day int) ([][]string, error) {
	lines, err := ReadInputLines(year, day)
	if err != nil {
		return nil, err
	}
	var slice [][]string
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		slice = append(slice, strings.Split(line, ""))
	}
	return slice, nil
}

func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

func readInputFromCache(year int, day int) string {
	data, err := os.ReadFile(fmt.Sprintf(".cache/input_%d_%d.txt", year, day))
	if err != nil {
		return ""
	}
	return string(data)
}

func saveToCache(year int, day int, input string) {
	data := []byte(input)
	os.Mkdir(".cache", 0755)
	os.WriteFile(fmt.Sprintf(".cache/input_%d_%d.txt", year, day), data, 0644)
}
