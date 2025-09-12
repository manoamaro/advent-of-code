package maps2

import "testing"

func TestMultiMap(t *testing.T) {
	m := NewMultiMap[int, int]()
	m.Add(1, 1)
	m.Add(1, 2)
	vals := m.Get(1)
	if len(vals) != 2 || vals[0] != 1 || vals[1] != 2 {
		t.Fatalf("multimap failed")
	}
}
