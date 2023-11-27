package internal

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Void struct{}

func ReadInputFromCache(day int) string {
	data, err := os.ReadFile(fmt.Sprintf(".cache/input_%d", day))
	if err != nil {
		return ""
	}
	return string(data)
}

func SaveToCache(day int, input string) {
	data := []byte(input)
	os.Mkdir(".cache", 0755)
	os.WriteFile(fmt.Sprintf(".cache/input_%d", day), data, 0644)
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
