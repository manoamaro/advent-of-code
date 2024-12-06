package grid

import "iter"

type Grid[T comparable] [][]T

func New[T comparable](rows, cols int) Grid[T] {
	grid := make([][]T, rows)
	for i := range rows {
		grid[i] = make([]T, cols)
	}
	return grid
}

func (g Grid[T]) Get(row, col int) *T {
	if row < 0 || row >= len(g) || col < 0 || col >= len(g[0]) {
		return nil
	}
	return &g[row][col]
}

func (g Grid[T]) Set(row, col int, value T) {
	if row < 0 || row >= len(g) || col < 0 || col >= len(g[0]) {
		return
	}
	g[row][col] = value
}

func (g Grid[T]) Rows() int {
	return len(g)
}

func (g Grid[T]) Cols() int {
	return len(g[0])
}

func (g *Grid[T]) ItemsSeq() iter.Seq[*T] {
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
