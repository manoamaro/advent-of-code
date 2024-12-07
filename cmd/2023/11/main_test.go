package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindGalaxies(t *testing.T) {
	i := []string{
		"...#......",
		".......#..",
		"#.........",
		"..........",
		"......#...",
		".#........",
		".........#",
		"..........",
		".......#..",
		"#...#.....",
	}
	e := 9

	input := ParseMatrixRaw(i)
	expected := e

	actual := FindGalaxies(input)
	fmt.Println(actual)

	if len(actual) != expected {
		t.Errorf("Expected length %d, got %d", expected, len(actual))
	}
}

func TestExpandGalaxies(t *testing.T) {
	i := []string{
		"...#......",
		".......#..",
		"#.........",
		"..........",
		"......#...",
		".#........",
		".........#",
		"..........",
		".......#..",
		"#...#.....",
	}

	input := ParseMatrixRaw(i)
	actual := FindGalaxies(input)

	expectedMatrix := ExpandBruteForce(input, 100)
	expected := FindGalaxies(expectedMatrix)

	expanded := ExpandGalaxySpaces(input, 100)
	fmt.Println(actual)
	fmt.Println(expected)
	fmt.Println(expanded)
	if !reflect.DeepEqual(expanded, expected) {
		t.Errorf("Expected %v, got %v", expected, expanded)
	}
}
