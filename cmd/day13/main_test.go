package main

import (
	"slices"
	"strings"
	"testing"
)

func TestFindFold(t *testing.T) {
	cases := []struct {
		pattern string
		want    []int
	}{
		{"#.##..##.", []int{5, 7}},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			got := FindFold(strings.Split(c.pattern, ""))
			if slices.Compare(got, c.want) != 0 {
				t.Errorf("FindFold(%q) == %d, want %d", c.pattern, got, c.want)
			}
		})
	}
}
