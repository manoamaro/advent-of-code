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

func MapCharsToInts(in []rune) []int {
	return collections.MapNotError(in, func(r rune) (int, error) {
		return int(r - '0'), nil
	})
}

func Atoi[T math2.Number](str string) T {
	var zero T
	str = strings.TrimSpace(str)
	switch any(zero).(type) {
	case int, int8, int16, int32, byte:
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
