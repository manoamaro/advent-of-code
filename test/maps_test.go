package test

import (
	"testing"

	"manoamaro.github.com/advent-of-code/pkg/maps"
)

func TestSetGet(t *testing.T) {
	cases := []struct {
		key, value int
	}{
		{key: 1, value: 1},
		{key: 2, value: 2},
		{key: 3, value: 3},
		{key: 4, value: 4},
		{key: 5, value: 5},
		{key: 6, value: 6},
		{key: 7, value: 7},
		{key: 1, value: 8},
		{key: 2, value: 9},
	}

	m := maps.New[int, int]()
	for _, c := range cases {
		m.Set(c.key, c.value)
		if v, ok := m.Get(c.key); !ok || v != c.value {
			t.Errorf("m.Get(%d) == %d, expected %d", c.key, v, c.value)
		}
	}
}
