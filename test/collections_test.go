package test

import (
	"fmt"
	"slices"
	"testing"

	"manoamaro.github.com/advent-of-code/pkg/collections"
)

func TestDiff(t *testing.T) {
	cases := []struct {
		a, b []int
		want []int
	}{
		{
			a:    []int{1, 2, 3},
			b:    []int{2, 3, 4},
			want: []int{1},
		},
		{
			a:    []int{1},
			b:    []int{6},
			want: []int{1},
		},
	}

	for _, c := range cases {
		got := collections.Diff(c.a, c.b)
		fmt.Println(got)
		slices.Sort(got)
		slices.Sort(c.want)
		if slices.Compare(got, c.want) != 0 {
			t.Errorf("Diff(%v, %v) == %v, want %v", c.a, c.b, got, c.want)
		}
	}
}

func TestCombinations(t *testing.T) {
	cases := []struct {
		input    []int
		size     int
		expected [][]int
	}{
		{
			input: []int{1, 2, 3},
			size:  2,
			expected: [][]int{
				{1, 2},
				{1, 3},
				{2, 3},
			},
		},
		{
			input: []int{1, 2, 3, 4},
			size:  2,
			expected: [][]int{
				{1, 2},
				{1, 3},
				{1, 4},
				{2, 3},
				{2, 4},
				{3, 4},
			},
		},
		{
			input: []int{1, 2, 3, 4},
			size:  3,
			expected: [][]int{
				{1, 2, 3},
				{1, 2, 4},
				{1, 3, 4},
				{2, 3, 4},
			},
		},
	}

	for _, c := range cases {
		got := collections.Combinations(c.input, c.size)
		for i, g := range got {
			if slices.Compare(g, c.expected[i]) != 0 {
				t.Errorf("Combinations(%v, %v) == %v, want %v", c.input, c.size, got, c.expected)
			}
		}
	}
}
