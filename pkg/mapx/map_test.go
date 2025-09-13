package mapx

import "testing"

func TestMapBasic(t *testing.T) {
	m := New[int, int](NewEntry(1, 1))
	if v, ok := m.Get(1); !ok || v != 1 {
		t.Fatalf("get failed")
	}
	m.Set(2, 2)
	if !m.Has(2) {
		t.Fatalf("has failed")
	}
	m.Delete(2)
	if m.Has(2) {
		t.Fatalf("delete failed")
	}
	if len(m.Keys()) != 1 {
		t.Fatalf("keys failed")
	}
	if len(m.Values()) != 1 {
		t.Fatalf("values failed")
	}
	clone := m.Clone()
	if len(clone) != 1 || clone[1] != 1 {
		t.Fatalf("clone failed")
	}
}

func TestMapAddI(t *testing.T) {
	m := New[int, int]()
	m2 := m.AddI(NewEntry(1, 1))
	if m.Has(1) || !m2.Has(1) {
		t.Fatalf("AddI failed")
	}
}
