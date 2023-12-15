package main

import (
	"slices"
	"testing"
)

func TestFindGroupsInLines(t *testing.T) {
	cases := []struct {
		line     []int
		expected []int
	}{
		{[]int{0, 0, 0, 0, 0}, []int{}},
		{[]int{1, 1, 1, 1, 1}, []int{5}},
		{[]int{0, 1, 1, 1, 1}, []int{4}},
		{[]int{1, 1, 1, 1, 0}, []int{4}},
		{[]int{1, 1, 1, 0, 1}, []int{3, 1}},
		{[]int{1, 1, 0, 1, 1}, []int{2, 2}},
		{[]int{1, 0, 1, 1, 1}, []int{1, 3}},
		{[]int{0, 1, 1, 1, 0}, []int{3}},
		{[]int{1, 1, 1, 0, 0}, []int{3}},
		{[]int{1, 1, 0, 0, 1}, []int{2, 1}},
		{[]int{1, 0, 0, 1, 1}, []int{1, 2}},
		{[]int{0, 1, 1, 0, 0}, []int{2}},
		{[]int{1, 1, 0, 0, 0}, []int{2}},
		{[]int{1, 0, 0, 0, 1}, []int{1, 1}},
		{[]int{0, 1, 0, 0, 0}, []int{1}},
		{[]int{0, 0, 1, 0, 0}, []int{1}},
		{[]int{0, 0, 0, 1, 0}, []int{1}},
		{[]int{0, 0, 0, 0, 1}, []int{1}},
		{[]int{1, 0, 1, 0, 1, 1, 1, 0, 1}, []int{1, 1, 3, 1}},
	}
	for _, c := range cases {
		actual := findGroupsInLine(c.line)
		if slices.Compare(actual, c.expected) != 0 {
			t.Errorf("findGroupsInLine(%v) == %v, expected %v", c.line, actual, c.expected)
		}
	}
}

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

func TestPermutation2(t *testing.T) {
	l := []int{-1, -1, -1, 0, 1, 1, 1}
	g := []int{1, 1, 3}
	qIdx := []int{0, 1, 2}
	c := t2(l, qIdx, g, 0, 0)
	println(c)
}
