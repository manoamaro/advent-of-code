package graph

import "testing"

func simpleGraph() *Graph[int, int] {
	g := New[int, int]()
	g.AddOneWayEdge(1, 2, 1, 0)
	g.AddOneWayEdge(2, 3, 1, 0)
	g.AddOneWayEdge(1, 3, 10, 0)
	return g
}

func TestHasEdge(t *testing.T) {
	g := simpleGraph()
	if !g.HasEdge(1, 2) || g.HasEdge(3, 1) {
		t.Fatalf("has edge failed")
	}
}

func TestNeighbors(t *testing.T) {
	g := simpleGraph()
	n := g.Neighbors(1)
	if len(n) != 2 {
		t.Fatalf("neighbors failed")
	}
}

func TestGetEdge(t *testing.T) {
	g := simpleGraph()
	e := g.GetEdge(1, 2)
	if e == nil || *e.To != 2 || e.Weight != 1 {
		t.Fatalf("get edge failed")
	}
}

func TestShortestPath(t *testing.T) {
	g := simpleGraph()
	p := g.FindShortestPathBetween(1, 3)
	if len(p) != 3 || p[0] != 1 || p[2] != 3 {
		t.Fatalf("shortest path failed")
	}
}

func TestShortestPaths(t *testing.T) {
	g := New[int, int]()
	g.AddOneWayEdge(1, 2, 1, 0)
	g.AddOneWayEdge(1, 3, 2, 0)
	g.AddOneWayEdge(2, 3, 1, 0)
	paths := g.FindShortestPathsBetween(1, 3)
	if len(paths) != 2 {
		t.Fatalf("expected two paths")
	}
}
