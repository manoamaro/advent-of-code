package main

import (
	"testing"
)

var pipes = pipes_T{
	{"7", "-", "F", "7", "-"},
	{".", "F", "J", "|", "7"},
	{"S", "J", "L", "L", "7"},
	{"|", "F", "-", "-", "J"},
	{"L", "J", ".", "L", "J"},
}

func TestPipes_Start(t *testing.T) {
	x, y := pipes.start()
	if x != 2 || y != 0 {
		t.Errorf("Start() = %v, %v, want %v, %v", x, y, 2, 0)
	}
}

func TestPipes_StartAndTerminations(t *testing.T) {
	start, terminations := pipes.startAndTerminations()
	if start[0] != 2 || start[1] != 0 {
		t.Errorf("StartAndTerminations() start = %v, want %v", start, []int{2, 0})
	}
	if len(terminations) != 2 {
		t.Errorf("StartAndTerminations() terminations = %v, want %v", terminations, [][]int{{0, 2}, {4, 2}})
	}
}

func TestPipesNavigateFrom(t *testing.T) {
	x, y := 2, 1
	count := pipes.navigateLoop()
	if count != 8 {
		t.Errorf("NavigateFrom(%v, %v) = %v, want %v", x, y, count, 8)
	}
}

func TestGetLoop(t *testing.T) {
	loop := pipes.getLoop()
	if len(loop) != 16 {
		t.Errorf("GetLoop() = %v, want %v", loop, [][]int{{2, 0}, {0, 2}, {4, 2}, {2, 4}, {0, 2}, {4, 2}, {2, 0}, {2, 4}})
	}
}
