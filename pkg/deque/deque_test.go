package deque

import "testing"

func TestPushPop(t *testing.T) {
	d := New[int]()
	d.PushBack(1)
	d.PushBack(2)
	d.PushFront(0)
	if v := d.PopFront(); v == nil || *v != 0 {
		t.Fatalf("expected 0")
	}
	if v := d.PopBack(); v == nil || *v != 2 {
		t.Fatalf("expected 2")
	}
}

func TestSeq(t *testing.T) {
	d := New[int](1, 2, 3)
	var res []int
	for v := range d.SeqFront() {
		res = append(res, v)
	}
	expected := []int{1, 2, 3}
	for i := range expected {
		if res[i] != expected[i] {
			t.Fatalf("unexpected seq")
		}
	}
}
