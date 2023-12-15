package internal

import "testing"

func TestGreatestCommonDivisor(t *testing.T) {
	cases := []struct {
		a, b, expected uint64
	}{
		{a: 1, b: 1, expected: 1},
		{a: 2, b: 1, expected: 1},
		{a: 1, b: 2, expected: 1},
		{a: 2, b: 2, expected: 2},
		{a: 2, b: 3, expected: 1},
		{a: 3, b: 2, expected: 1},
		{a: 2, b: 4, expected: 2},
		{a: 4, b: 2, expected: 2},
		{a: 4, b: 6, expected: 2},
		{a: 6, b: 4, expected: 2},
		{a: 6, b: 9, expected: 3},
		{a: 9, b: 6, expected: 3},
		{a: 6, b: 12, expected: 6},
		{a: 12, b: 6, expected: 6},
		{a: 12, b: 18, expected: 6},
		{a: 18, b: 12, expected: 6},
		{a: 12, b: 24, expected: 12},
		{a: 24, b: 12, expected: 12},
		{a: 12, b: 36, expected: 12},
		{a: 36, b: 12, expected: 12},
		{a: 36, b: 48, expected: 12},
		{a: 48, b: 36, expected: 12},
		{a: 48, b: 60, expected: 12},
		{a: 60, b: 48, expected: 12},
		{a: 48, b: 72, expected: 24},
		{a: 72, b: 48, expected: 24},
		{a: 48, b: 84, expected: 12},
		{a: 84, b: 48, expected: 12},
		{a: 48, b: 96, expected: 48},
	}

	for _, c := range cases {
		actual := GreatestCommonDivisor(c.a, c.b)
		if actual != c.expected {
			t.Errorf("GreatestCommonDivisor(%d, %d) == %d, expected %d", c.a, c.b, actual, c.expected)
		}
	}
}

func TestLowestCommonMultiple(t *testing.T) {
	cases := []struct {
		a, b, expected uint64
	}{
		{a: 1, b: 1, expected: 1},
		{a: 2, b: 1, expected: 2},
		{a: 1, b: 2, expected: 2},
		{a: 2, b: 2, expected: 2},
		{a: 2, b: 3, expected: 6},
		{a: 3, b: 2, expected: 6},
		{a: 2, b: 4, expected: 4},
		{a: 4, b: 2, expected: 4},
		{a: 4, b: 6, expected: 12},
		{a: 6, b: 4, expected: 12},
		{a: 6, b: 9, expected: 18},
		{a: 9, b: 6, expected: 18},
		{a: 6, b: 12, expected: 12},
		{a: 12, b: 6, expected: 12},
		{a: 12, b: 18, expected: 36},
		{a: 18, b: 12, expected: 36},
		{a: 12, b: 24, expected: 24},
		{a: 24, b: 12, expected: 24},
		{a: 12, b: 36, expected: 36},
		{a: 36, b: 12, expected: 36},
		{a: 36, b: 48, expected: 144},
		{a: 48, b: 36, expected: 144},
		{a: 48, b: 60, expected: 240},
		{a: 60, b: 48, expected: 240},
		{a: 48, b: 72, expected: 144},
		{a: 72, b: 48, expected: 144},
		{a: 48, b: 84, expected: 336},
	}

	for _, c := range cases {
		actual := LowestCommonMultiple(c.a, c.b)
		if actual != c.expected {
			t.Errorf("LowestCommonMultiple(%d, %d) == %d, expected %d", c.a, c.b, actual, c.expected)
		}
	}
}

func TestSummation(t *testing.T) {
	cases := []struct {
		n        int
		expected int
	}{
		{n: 1, expected: 1},
		{n: 2, expected: 3},
		{n: 3, expected: 6},
		{n: 4, expected: 10},
		{n: 5, expected: 15},
		{n: 6, expected: 21},
		{n: 7, expected: 28},
		{n: 8, expected: 36},
		{n: 9, expected: 45},
		{n: 10, expected: 55},
	}

	for _, c := range cases {
		actual := Summation(c.n)
		if actual != c.expected {
			t.Errorf("Summation(%d) == %d, expected %d", c.n, actual, c.expected)
		}
	}
}

func TestManhattanDistance(t *testing.T) {
	cases := []struct {
		a, b     []int
		expected int
	}{
		{a: []int{0, 0}, b: []int{0, 0}, expected: 0},
		{a: []int{0, 0}, b: []int{1, 0}, expected: 1},
		{a: []int{0, 0}, b: []int{0, 1}, expected: 1},
		{a: []int{0, 0}, b: []int{1, 1}, expected: 2},
		{a: []int{1, 1}, b: []int{0, 0}, expected: 2},
		{a: []int{-1, -1}, b: []int{0, 0}, expected: 2},
		{a: []int{0, 0}, b: []int{-1, -1}, expected: 2},
		{a: []int{1, 1}, b: []int{-1, -1}, expected: 4},
		{a: []int{-1, -1}, b: []int{1, 1}, expected: 4},
	}

	for _, c := range cases {
		actual := ManhattanDistance(c.a, c.b)
		if actual != c.expected {
			t.Errorf("ManhattanDistance(%v, %v) == %d, expected %d", c.a, c.b, actual, c.expected)
		}
	}
}
