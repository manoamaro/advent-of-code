package math2

import "testing"

func TestArray2D3D(t *testing.T) {
	a := Array2D[int](2, 3)
	if len(a) != 2 || len(a[0]) != 3 {
		t.Fatalf("array2d failed")
	}
	b := Array3D[int](2, 2, 2)
	if len(b) != 2 || len(b[0]) != 2 || len(b[0][0]) != 2 {
		t.Fatalf("array3d failed")
	}
}

func TestRotateMatrix(t *testing.T) {
	m := [][]int{{1, 2}, {3, 4}}
	r := RotateMatrix(m)
	if r[0][0] != 3 || r[1][0] != 4 {
		t.Fatalf("rotate failed")
	}
}

func TestMathHelpers(t *testing.T) {
	if Max(2, 3) != 3 || Min(2, 3) != 2 {
		t.Fatalf("max/min failed")
	}
	if Abs(-2) != 2 {
		t.Fatalf("abs failed")
	}
	if Floor(3.7) != 3 {
		t.Fatalf("floor failed")
	}
	if Log10(100.0) != 2 {
		t.Fatalf("log10 failed")
	}
	if Power(2, 3) != 8 {
		t.Fatalf("power failed")
	}
}
