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
