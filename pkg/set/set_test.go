package set

import "testing"

func TestSetBasic(t *testing.T) {
	s := New[int]()
	s.Add(1)
	if !s.Contains(1) {
		t.Fatalf("set should contain 1")
	}
	if s.Len() != 1 {
		t.Fatalf("len expected 1")
	}
	s.Remove(1)
	if s.Contains(1) {
		t.Fatalf("remove failed")
	}
}

func TestSetUnionIntersection(t *testing.T) {
	s1 := New(1, 2)
	s2 := New(2, 3)
	u := s1.Union(s2)
	if !u.Equals(New(1, 2, 3)) {
		t.Fatalf("union failed")
	}
	i := s1.Intersection(s2)
	if !i.Equals(New(2)) {
		t.Fatalf("intersection failed")
	}
	d := s1.Difference(s2)
	if !d.Equals(New(1)) {
		t.Fatalf("difference failed")
	}
	sym := s1.SymmetricDifference(s2)
	if !sym.Equals(New(1, 3)) {
		t.Fatalf("symmetric difference failed")
	}
}
