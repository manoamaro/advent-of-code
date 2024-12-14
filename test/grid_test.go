package test

import (
	"testing"

	"manoamaro.github.com/advent-of-code/pkg/grid"
)

func TestGrid(t *testing.T) {
	g := grid.New[int](3, 3)
	if g.Rows() != 3 {
		t.Errorf("expected 3 rows, got %d", g.Rows())
	}
	if g.Cols() != 3 {
		t.Errorf("expected 3 cols, got %d", g.Cols())
	}
	g.Set(0, 0, 1)
	g.Set(0, 1, 2)
	g.Set(0, 2, 3)
	g.Set(1, 0, 4)
	g.Set(1, 1, 5)
	g.Set(1, 2, 6)
	g.Set(2, 0, 7)
	g.Set(2, 1, 8)
	g.Set(2, 2, 9)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if *g.Get(i, j) != i*3+j+1 {
				t.Errorf("expected %d, got %d", i*3+j+1, *g.Get(i, j))
			}
		}
	}
	g2 := g.Copy()
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if *g2.Get(i, j) != i*3+j+1 {
				t.Errorf("expected %d, got %d", i*3+j+1, *g2.Get(i, j))
			}
		}
	}
}

func TestGridInBounds(t *testing.T) {
	g := grid.New[int](3, 3)
	if !g.InBounds(grid.Cell{0, 0}) {
		t.Errorf("expected true, got false")
	}
	if !g.InBounds(grid.Cell{2, 2}) {
		t.Errorf("expected true, got false")
	}
	if g.InBounds(grid.Cell{3, 3}) {
		t.Errorf("expected false, got true")
	}
	if g.InBounds(grid.Cell{-1, -1}) {
		t.Errorf("expected false, got true")
	}
}

func TestGridItemsSeq(t *testing.T) {
	g := grid.New[int](3, 3)
	g.Set(0, 0, 1)
	g.Set(0, 1, 2)
	g.Set(0, 2, 3)
	g.Set(1, 0, 4)
	g.Set(1, 1, 5)
	g.Set(1, 2, 6)
	g.Set(2, 0, 7)
	g.Set(2, 1, 8)
	g.Set(2, 2, 9)
	var items []int
	for item := range g.ItemsSeq() {
		items = append(items, *item)
	}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := range expected {
		if items[i] != expected[i] {
			t.Errorf("expected %d, got %d", expected[i], items[i])
		}
	}
}

func TestGridFindFunc(t *testing.T) {
	g := grid.New[int](3, 3)
	g.Set(0, 0, 1)
	g.Set(0, 1, 2)
	g.Set(0, 2, 3)
	g.Set(1, 0, 4)
	g.Set(1, 1, 5)
	g.Set(1, 2, 6)
	g.Set(2, 0, 7)
	g.Set(2, 1, 8)
	g.Set(2, 2, 9)
	cell := g.FindFunc(func(v int) bool {
		return v == 5
	})
	if cell == nil {
		t.Errorf("expected cell, got nil")
	}
	if cell != nil && (cell[0] != 1 || cell[1] != 1) {
		t.Errorf("expected [1, 1], got %v", cell)
	}
}

func TestGridFindAllFunc(t *testing.T) {
	g := grid.New[int](3, 3)
	g.Set(0, 0, 1)
	g.Set(0, 1, 2) // even
	g.Set(0, 2, 3)
	g.Set(1, 0, 4) // even
	g.Set(1, 1, 5)
	g.Set(1, 2, 6) // even
	g.Set(2, 0, 7)
	g.Set(2, 1, 8) // even
	g.Set(2, 2, 9)
	cells := g.FindAllFunc(func(v int) bool {
		return v%2 == 0
	})
	expected := []grid.Cell{{0, 1}, {1, 0}, {1, 2}, {2, 1}}
	for i := range expected {
		if cells[i] != expected[i] {
			t.Errorf("expected %v, got %v", expected[i], cells[i])
		}
	}
}

func TestGridSet(t *testing.T) {
	g := grid.New[int](3, 3)
	g.Set(0, 0, 1)
	if *g.Get(0, 0) != 1 {
		t.Errorf("expected 1, got %d", *g.Get(0, 0))
	}
	g.Set(0, 0, 2)
	if *g.Get(0, 0) != 2 {
		t.Errorf("expected 2, got %d", *g.Get(0, 0))
	}
}

func TestGridAdjacents(t *testing.T) {
	g := grid.New[int](3, 3)
	g.Set(0, 0, 1)
	g.Set(0, 1, 2)
	g.Set(0, 2, 3)
	g.Set(1, 0, 4)
	g.Set(1, 1, 5)
	g.Set(1, 2, 6)
	g.Set(2, 0, 7)
	g.Set(2, 1, 8)
	g.Set(2, 2, 9)
	adjacents := g.Adjacents(grid.Cell{1, 1})
	expected := map[grid.Cell]*int{
		{0, 1}: constToPtr(2),
		{1, 0}: constToPtr(4),
		{1, 2}: constToPtr(6),
		{2, 1}: constToPtr(8),
	}
	for cell, value := range adjacents {
		if *value != *expected[cell] {
			t.Errorf("expected %d, got %d", *expected[cell], *value)
		}
	}
}

func constToPtr[T any](v T) *T {
	return &v
}
