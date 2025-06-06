package aoc

import (
	"reflect"
	"strings"
	"testing"
)

func TestStringProcessor(t *testing.T) {
	if StringProcessor(" a \n") != "a" {
		t.Fatalf("trim failed")
	}
}

func TestLinesProcessor(t *testing.T) {
	proc := LinesProcessor()
	lines := proc("a\nb")
	if !reflect.DeepEqual(lines, []string{"a", "b"}) {
		t.Fatalf("lines processor failed")
	}
}

func TestIntLinesProcessor(t *testing.T) {
	ints := IntLinesProcessor("1\n2")
	if len(ints) != 2 || ints[0] != 1 || ints[1] != 2 {
		t.Fatalf("int lines processor failed")
	}
}

func TestGridStringProcessor(t *testing.T) {
	proc := GridStringProcessor()
	grid := proc("ab\ncd")
	if grid[1][1] != "d" {
		t.Fatalf("grid string processor failed")
	}
}

func TestRuneGridProcessor(t *testing.T) {
	g := RuneGridProcessor("ab\ncd")
	if *g.Get(1, 1) != 'd' {
		t.Fatalf("rune grid failed")
	}
}

func TestIntGridProcessor(t *testing.T) {
	g := IntGridProcessor("12\n34")
	if *g.Get(1, 0) != 3 {
		t.Fatalf("int grid failed")
	}
}

func TestInts2dProcessor(t *testing.T) {
	splitter := func(s string) []string { return strings.Split(s, ",") }
	proc := Ints2dProcessor(splitter)
	res := proc("1,2\n3,4")
	if res[1][1] != 4 {
		t.Fatalf("ints2d failed")
	}
}

func TestGridProcessor(t *testing.T) {
	proc := GridProcessor(func(r rune) int { return int(r - '0') })
	g := proc("12\n34")
	if *g.Get(1, 1) != 4 {
		t.Fatalf("grid processor failed")
	}
}
