package test

import (
	"slices"
	"strconv"
	"testing"

	"manoamaro.github.com/advent-of-code/pkg/collections"
)

func TestMapFunctions(t *testing.T) {
	r := collections.Map([]int{1, 2, 3}, func(i int) int { return i * 2 })
	if !slices.Equal(r, []int{2, 4, 6}) {
		t.Fatalf("map failed")
	}
	r2 := collections.MapNotError([]string{"1", "a"}, func(s string) (int, error) { return strconv.Atoi(s) })
	if !slices.Equal(r2, []int{1}) {
		t.Fatalf("mapnoterror failed")
	}
}

func TestFlatMapFoldReverse(t *testing.T) {
	fm := collections.FlatMap([][]int{{1, 2}, {3}}, func(s []int) []int { return s })
	if !slices.Equal(fm, []int{1, 2, 3}) {
		t.Fatalf("flatmap failed")
	}
	sum := collections.Fold([]int{1, 2, 3}, 0, func(a, b int) int { return a + b })
	if sum != 6 {
		t.Fatalf("fold failed")
	}
	rev := collections.Reverse([]int{1, 2, 3})
	if !slices.Equal(rev, []int{3, 2, 1}) {
		t.Fatalf("reverse failed")
	}
}

func TestOtherCollectionFuncs(t *testing.T) {
	seq := collections.SlideSeq([]int{1, 2, 3}, 2)
	var got [][]int
	for s := range seq {
		got = append(got, s)
	}
	if len(got) != 2 {
		t.Fatalf("slideseq failed")
	}
	del := collections.Delete([]int{1, 2, 3}, 1)
	if !slices.Equal(del, []int{1, 3}) {
		t.Fatalf("delete failed")
	}
	first := collections.FirstFunc([]int{1, 2, 3}, func(i int) bool { return i == 2 })
	if first == nil || *first != 2 {
		t.Fatalf("firstfunc failed")
	}
	filt := collections.FilterFunc([]int{1, 2, 3}, func(i int) bool { return i > 1 })
	if !slices.Equal(filt, []int{2, 3}) {
		t.Fatalf("filterfunc failed")
	}
	prod := collections.ProductFunc([]int{1, 2}, []int{3}, func(a, b int) int { return a + b })
	if !slices.Equal(prod, []int{4, 5}) {
		t.Fatalf("productfunc failed")
	}
	if collections.Sum([]int{1, 2, 3}) != 6 {
		t.Fatalf("sum failed")
	}
}
