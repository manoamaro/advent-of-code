package queue

import "testing"

func TestQueuePushPop(t *testing.T) {
	q := New[int]()
	q.Push(1)
	q.Push(2)
	if v := q.Pop(); v == nil || *v != 1 {
		t.Fatalf("expected 1")
	}
	if v := q.Pop(); v == nil || *v != 2 {
		t.Fatalf("expected 2")
	}
}

func TestQueueSeq(t *testing.T) {
	q := New[int](1, 2, 3)
	var res []int
	for v := range q.Seq() {
		res = append(res, v)
	}
	expected := []int{1, 2, 3}
	for i := range expected {
		if res[i] != expected[i] {
			t.Fatalf("seq mismatch")
		}
	}
}
