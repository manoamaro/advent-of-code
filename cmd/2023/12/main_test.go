package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := []string{
		"???.### 1,1,3",
		".??..??...?##. 1,1,3",
		"?#?#?#?#?#?#?#? 1,3,1,6",
		"????.#...#... 4,1,1",
		"????.######..#####. 1,6,5",
		"?###???????? 3,2,1",
	}
	part1(input)
}

func TestPart2(t *testing.T) {
	input := []string{
		".??..??...?##. 1,1,3",
	}
	part2(input)
}

func TestCalculate(t *testing.T) {
	cases := []struct {
		input  string
		groups []int
		want   int
	}{
		{"???.###", []int{1, 1, 3}, 1},
		{".??..??...?##.", []int{1, 1, 3}, 4},
		{"?#?#?#?#?#?#?#?", []int{1, 3, 1, 6}, 1},
		{"????.#...#...", []int{4, 1, 1}, 1},
		{"????.######..#####.", []int{1, 6, 5}, 4},
		{"?###????????", []int{3, 2, 1}, 10},
	}

	for _, c := range cases {
		got := calculate(c.input, c.groups, map[string]int{})
		if got != c.want {
			t.Errorf("Calculate(%q, %v) == %d, want %d", c.input, c.groups, got, c.want)
		}
	}
}
