package strings2

import (
	"strconv"
	"strings"

	"manoamaro.github.com/advent-of-code/pkg/collections"
	"manoamaro.github.com/advent-of-code/pkg/math2"
)

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func MapToInt(in []string) []int {
	return collections.MapNotError(in, func(s string) (int, error) {
		return strconv.Atoi(strings.TrimSpace(s))
	})
}

func Atoi[T math2.Number](str string) T {
	var zero T
	switch any(zero).(type) {
	case int:
		num, _ := strconv.Atoi(str)
		return T(num)
	case int64:
		num, _ := strconv.ParseInt(str, 10, 64)
		return T(num)
	case float32:
		num, _ := strconv.ParseFloat(str, 32)
		return T(num)
	case float64:
		num, _ := strconv.ParseFloat(str, 64)
		return T(num)
	default:
		return zero
	}
}
