package sliceutil_test

import (
	"fmt"
	"slices"
	"strconv"
	"testing"

	"manoamaro.github.com/advent-of-code/pkg/sliceutil"
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
		got := sliceutil.Diff(c.a, c.b)
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
		got := sliceutil.Combinations(c.input, c.size)
		for i, g := range got {
			if slices.Compare(g, c.expected[i]) != 0 {
				t.Errorf("Combinations(%v, %v) == %v, want %v", c.input, c.size, got, c.expected)
			}
		}
	}
}

func TestMapFunctions(t *testing.T) {
	r := sliceutil.Map([]int{1, 2, 3}, func(i int) int { return i * 2 })
	if !slices.Equal(r, []int{2, 4, 6}) {
		t.Fatalf("map failed")
	}
	r2 := sliceutil.MapNotError([]string{"1", "a"}, func(s string) (int, error) { return strconv.Atoi(s) })
	if !slices.Equal(r2, []int{1}) {
		t.Fatalf("mapnoterror failed")
	}
}

func TestFlatMapFoldReverse(t *testing.T) {
	fm := sliceutil.FlatMap([][]int{{1, 2}, {3}}, func(s []int) []int { return s })
	if !slices.Equal(fm, []int{1, 2, 3}) {
		t.Fatalf("flatmap failed")
	}
	sum := sliceutil.Fold([]int{1, 2, 3}, 0, func(a, b int) int { return a + b })
	if sum != 6 {
		t.Fatalf("fold failed")
	}
	rev := sliceutil.Reverse([]int{1, 2, 3})
	if !slices.Equal(rev, []int{3, 2, 1}) {
		t.Fatalf("reverse failed")
	}
}

func TestOtherCollectionFuncs(t *testing.T) {
	seq := sliceutil.SlideSeq([]int{1, 2, 3}, 2)
	var got [][]int
	for s := range seq {
		got = append(got, s)
	}
	if len(got) != 2 {
		t.Fatalf("slideseq failed")
	}
	del := sliceutil.Delete([]int{1, 2, 3}, 1)
	if !slices.Equal(del, []int{1, 3}) {
		t.Fatalf("delete failed")
	}
	first := sliceutil.FirstFunc([]int{1, 2, 3}, func(i int) bool { return i == 2 })
	if first == nil || *first != 2 {
		t.Fatalf("firstfunc failed")
	}
	filt := sliceutil.FilterFunc([]int{1, 2, 3}, func(i int) bool { return i > 1 })
	if !slices.Equal(filt, []int{2, 3}) {
		t.Fatalf("filterfunc failed")
	}
	prod := sliceutil.ProductFunc([]int{1, 2}, []int{3}, func(a, b int) int { return a + b })
	if !slices.Equal(prod, []int{4, 5}) {
		t.Fatalf("productfunc failed")
	}
	if sliceutil.Sum([]int{1, 2, 3}) != 6 {
		t.Fatalf("sum failed")
	}
}
