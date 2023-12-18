package internal

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Void struct{}

func ReadInputFromCache(day int) string {
	data, err := os.ReadFile(fmt.Sprintf(".cache/input_%d.txt", day))
	if err != nil {
		return ""
	}
	return string(data)
}

func SaveToCache(day int, input string) {
	data := []byte(input)
	os.Mkdir(".cache", 0755)
	os.WriteFile(fmt.Sprintf(".cache/input_%d.txt", day), data, 0644)
}

func ReadInput(day int) (string, error) {

	if cached := ReadInputFromCache(day); len(cached) > 0 {
		return cached, nil
	}

	godotenv.Load()

	session := os.Getenv("AOC_SESSION")
	if len(session) == 0 {
		return "", fmt.Errorf("cannot find session")
	}

	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://adventofcode.com/2023/day/%d/input", day), nil)
	if err != nil {
		return "", err
	}

	request.AddCookie(&http.Cookie{
		Name:  "session",
		Value: session,
	})

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	b, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	input := string(b)

	SaveToCache(day, input)

	return input, nil
}

func ReadInputLines(day int) ([]string, error) {
	input, err := ReadInput(day)
	input = strings.TrimSpace(input)
	if err != nil {
		return nil, err
	}
	return strings.Split(input, "\n"), nil
}

func ReadInputSlice2d(day int) ([][]string, error) {
	lines, err := ReadInputLines(day)
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

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
