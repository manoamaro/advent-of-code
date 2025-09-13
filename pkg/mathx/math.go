package mathx

import (
	"math"
)

type Integer interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

type Float interface {
	float32 | float64
}

type Number interface {
	Integer | Float
}

// LowestCommonMultiple returns the lowest common multiple of a and b.
func LCM[N Integer](a, b N) N {
	return (a / GCD(a, b)) * b
}

// GreatestCommonDivisor returns the greatest common divisor of a and b.
func GCD[N Integer](a, b N) N {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Array2D[T Number](x, y int) [][]T {
	s := make([][]T, x)
	for i := range s {
		s[i] = make([]T, y)
	}
	return s
}

func Array3D[T Number](x, y, z int) [][][]T {
	s := make([][][]T, x)
	for i := range s {
		s[i] = Array2D[T](y, z)
	}
	return s
}

func RotateMatrix[T Number](matrix [][]T) [][]T {

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

func PrintMatrix[T Number](matrix [][]T) {
	for _, row := range matrix {
		for _, v := range row {
			print(v)
		}
		println()
	}
}

func Summation[T Number](n T) T {
	return n * (n + 1) / 2
}

// ManhattanDistance calculates the Manhattan distance between two points in a 2D space.
// It takes two integer slices, 'a' and 'b', representing the coordinates of the two points.
// The function returns the Manhattan distance as an integer.
func ManhattanDistance[T Integer](a, b [2]T) T {
	return Abs(a[0]-b[0]) + Abs(a[1]-b[1])
}

func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Abs[T Number](a T) T {
	if a < 0 {
		return -a
	}
	return a
}

func Floor[T Number](a T) T {
	return T(math.Floor(float64(a)))
}

func Log10[T Number](a T) T {
	return T(math.Log10(float64(a)))
}

func Power[T Number](a, b T) T {
	return T(math.Pow(float64(a), float64(b)))
}
