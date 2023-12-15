package internal

import "math"

// LowestCommonMultiple returns the lowest common multiple of a and b.
func LowestCommonMultiple(a, b uint64) uint64 {
	return a / GreatestCommonDivisor(a, b) * b
}

// GreatestCommonDivisor returns the greatest common divisor of a and b.
func GreatestCommonDivisor(a, b uint64) uint64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Array3D creates a 3D array of ints with dimensions x, y, z.
func Array3D(x, y, z int) [][][]int {
	s := make([][][]int, x)
	for i := range s {
		s[i] = make([][]int, y)
		for j := range s[i] {
			s[i][j] = make([]int, z)
		}
	}
	return s
}

func RotateMatrix[T any](matrix [][]T) [][]T {

	// reverse the matrix
	for i, j := 0, len(matrix)-1; i < j; i, j = i+1, j-1 {
		matrix[i], matrix[j] = matrix[j], matrix[i]
	}

	// transpose it
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return matrix
}

func PrintMatrix[T any](matrix [][]T) {
	for _, row := range matrix {
		for _, v := range row {
			print(v)
		}
		println()
	}
}

func Summation(n int) int {
	return n * (n + 1) / 2
}

// ManhattanDistance calculates the Manhattan distance between two points in a 2D space.
// It takes two integer slices, 'a' and 'b', representing the coordinates of the two points.
// The function returns the Manhattan distance as an integer.
func ManhattanDistance(a, b []int) int {
	return int(math.Abs(float64(a[0]-b[0])) + math.Abs(float64(a[1]-b[1])))
}
