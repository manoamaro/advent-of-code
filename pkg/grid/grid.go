package grid

import (
	"fmt"
	"iter"
	"strings"
)

type Grid[T comparable] [][]T

func New[T comparable](rows, cols int) Grid[T] {
	grid := make([][]T, rows)
	for i := range rows {
		grid[i] = make([]T, cols)
	}
	return grid
}

func (g *Grid[T]) Get(row, col int) *T {
	if !g.InBounds(Cell{row, col}) {
		return nil
	}
	return &(*g)[row][col]
}

func (g *Grid[T]) GetCell(c Cell) *T {
	return g.Get(c[0], c[1])
}

func (g *Grid[T]) Set(row, col int, value T) {
	if !g.InBounds(Cell{row, col}) {
		return
	}
	(*g)[row][col] = value
}

func (g *Grid[T]) Rows() int {
	return len(*g)
}

func (g *Grid[T]) Cols() int {
	return len((*g)[0])
}

func (g *Grid[T]) Copy() Grid[T] {
	c := New[T](g.Rows(), g.Cols())
	for i := range *g {
		copy(c[i], (*g)[i])
	}
	return c
}

func (g *Grid[T]) ValuesSeq() iter.Seq[*T] {
	return func(yield func(*T) bool) {
		for i := range *g {
			for j := range (*g)[i] {
				if !yield(&(*g)[i][j]) {
					return
				}
			}
		}
	}
}

func (g *Grid[T]) Cells() []Cell {
	var cells []Cell
	for i := range *g {
		for j := range (*g)[i] {
			cells = append(cells, Cell{i, j})
		}
	}
	return cells
}

func (g *Grid[T]) CellsSeq() iter.Seq[Cell] {
	return func(yield func(Cell) bool) {
		for i := range *g {
			for j := range (*g)[i] {
				if !yield(Cell{i, j}) {
					return
				}
			}
		}
	}
}

func (g *Grid[T]) Find(value T) *Cell {
	return g.FindFunc(func(v T) bool { return v == value })
}

func (g *Grid[T]) FindFunc(f func(T) bool) *Cell {
	for i := range *g {
		for j := range (*g)[i] {
			if f((*g)[i][j]) {
				return &Cell{i, j}
			}
		}
	}
	return nil
}

func (g *Grid[T]) FindAll(value T) []Cell {
	return g.FindAllFunc(func(v T) bool { return v == value })
}

func (g *Grid[T]) FindAllFunc(f func(T) bool) []Cell {
	var cells []Cell
	for i := range *g {
		for j := range (*g)[i] {
			if f((*g)[i][j]) {
				cells = append(cells, Cell{i, j})
			}
		}
	}
	return cells
}

func (g *Grid[T]) InBounds(c Cell) bool {
	return c[0] >= 0 && c[0] < g.Rows() && c[1] >= 0 && c[1] < g.Cols()
}

// Neighbors returns a map of adjacent cells and their corresponding values
// for a given cell in the grid. It checks the four possible adjacent positions
// (up, down, left, right) and returns a map with the adjacent cells as keys and
// pointers to the corresponding values in the grid as values. nil is returned
// for cells that are out of bounds.
//
// Parameters:
//   - c: The cell for which to find the neighbors.
//
// Returns:
//
//	A map where the keys are the adjacent cells and the values are pointers to
//	the corresponding values in the grid.
func (g *Grid[T]) Neighbors(c Cell) map[Cell]*T {
	neighbors := make(map[Cell]*T)
	for _, d := range CardinalDirections {
		cell := c.Move(d)
		neighbor := g.GetCell(cell)
		if neighbor != nil {
			neighbors[cell] = neighbor
		}
	}
	return neighbors
}

func (g *Grid[T]) String() string {
	var b strings.Builder
	for i := range *g {
		for j := range (*g)[i] {
			b.WriteString(fmt.Sprint((*g)[i][j]))
			b.WriteRune('\t')
		}
		b.WriteString("\n")
	}
	return b.String()
}

func (g *Grid[T]) Fill(v T) {
	for i := range *g {
		for j := range (*g)[i] {
			(*g)[i][j] = v
		}
	}
}

func (g *Grid[T]) SetCell(start Cell, v T) {
	g.Set(start[0], start[1], v)
}
