package aoc

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"manoamaro.github.com/advent-of-code/pkg/errors"
)

func GetInput(year int, day int) string {

	if cached := getFromCache(year, day); len(cached) > 0 {
		return cached
	}

	godotenv.Load()

	session := os.Getenv("AOC_SESSION")
	if len(session) == 0 {
		panic("cannot find session")
	}

	request := errors.Must(http.NewRequest(http.MethodGet, fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day), nil))

	request.AddCookie(&http.Cookie{
		Name:  "session",
		Value: session,
	})

	response := errors.Must(http.DefaultClient.Do(request))

	defer response.Body.Close()

	b := errors.Must(io.ReadAll(response.Body))

	input := string(b)

	saveToCache(year, day, input)

	return input
}

func getCacheFilePath(year int, day int) string {
	return fmt.Sprintf(".cache/input_%d_%d.txt", year, day)
}

func getFromCache(year int, day int) string {
	data, err := os.ReadFile(getCacheFilePath(year, day))
	if err != nil {
		return ""
	}
	return string(data)
}

func saveToCache(year int, day int, input string) {
	data := []byte(input)
	os.Mkdir(".cache", 0755)
	os.WriteFile(getCacheFilePath(year, day), data, 0644)
}
