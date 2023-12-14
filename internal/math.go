package internal

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
