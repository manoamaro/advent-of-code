package strings2

import "testing"

func TestReverseString(t *testing.T) {
	if ReverseString("abc") != "cba" {
		t.Fatalf("reverse failed")
	}
}

func TestMapToInt(t *testing.T) {
	res := MapToInt([]string{"1", "2", "3"})
	expected := []int{1, 2, 3}
	for i := range expected {
		if res[i] != expected[i] {
			t.Fatalf("map failed")
		}
	}
}

func TestMapCharsToInts(t *testing.T) {
	res := MapCharsToInts([]rune("12"))
	if len(res) != 2 || res[0] != 1 || res[1] != 2 {
		t.Fatalf("map chars failed")
	}
}

func TestAtoi(t *testing.T) {
	if Atoi[int]("42") != 42 {
		t.Fatalf("atoi failed")
	}
	if Atoi[int64]("42") != 42 {
		t.Fatalf("atoi int64 failed")
	}
	if Atoi[float64]("3.5") != 3.5 {
		t.Fatalf("atoi float failed")
	}
}
