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

	input := parseMatrixRaw(i)
	expected := e

	actual := findGalaxies(input)
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

	input := parseMatrixRaw(i)
	actual := findGalaxies(input)

	expectedMatrix := expandBruteForce(input, 100)
	expected := findGalaxies(expectedMatrix)

	expanded := expandGalaxySpaces(input, 100)
	fmt.Println(actual)
	fmt.Println(expected)
	fmt.Println(expanded)
	if !reflect.DeepEqual(expanded, expected) {
		t.Errorf("Expected %v, got %v", expected, expanded)
	}
}
