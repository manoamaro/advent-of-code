package queue

import "testing"

func TestPriorityQueuePushPop(t *testing.T) {
	pq := NewPriorityQueue[int]()
	pq.PushValue(3, 3)
	pq.PushValue(1, 1)
	pq.PushValue(2, 2)
	if v := pq.PopValue(); v != 1 {
		t.Fatalf("expected 1")
	}
	if v := pq.PopValue(); v != 2 {
		t.Fatalf("expected 2")
	}
	if v := pq.PopValue(); v != 3 {
		t.Fatalf("expected 3")
	}
}

func TestPriorityQueueSeq(t *testing.T) {
	pq := NewPriorityQueue[int]()
	pq.PushValue(1, 1)
	pq.PushValue(2, 2)
	pq.PushValue(3, 3)
	var res []int
	for v := range pq.Seq() {
		res = append(res, v)
	}
	expected := []int{1, 2, 3}
	for i := range expected {
		if res[i] != expected[i] {
			t.Fatalf("seq mismatch")
		}
	}
}

func TestPriorityQueueSeqPriority(t *testing.T) {
	pq := NewPriorityQueue[int]()
	pq.PushValue(1, 1)
	pq.PushValue(2, 2)
	pq.PushValue(3, 3)
	var res []int
	for v, p := range pq.SeqPriority() {
		res = append(res, v+p)
	}
	if len(res) != 3 {
		t.Fatalf("expected 3 elements")
	}
}
