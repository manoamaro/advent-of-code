package main

import (
	"reflect"
	"testing"
)

func TestReduceToZeros(t *testing.T) {
	input := []int{0, 3, 6, 9, 12, 15}
	expected := [][]int{
		{0, 3, 6, 9, 12, 15},
		{3, 3, 3, 3, 3},
		{0, 0, 0, 0},
	}
	actual := reduceToZeros(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestExtrapolateRight(t *testing.T) {
	input := [][]int{
		{0, 3, 6, 9, 12, 15},
		{3, 3, 3, 3, 3},
		{0, 0, 0, 0},
	}
	expected := [][]int{
		{0, 3, 6, 9, 12, 15, 18},
		{3, 3, 3, 3, 3, 3},
		{0, 0, 0, 0, 0},
	}
	actual := extrapolateRight(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestExtrapolateLeft(t *testing.T) {
	input := [][]int{
		{10, 13, 16, 21, 30, 45},
		{3, 3, 5, 9, 15},
		{0, 2, 4, 6},
		{2, 2, 2},
		{0, 0},
	}
	expected := [][]int{
		{5, 10, 13, 16, 21, 30, 45},
		{5, 3, 3, 5, 9, 15},
		{-2, 0, 2, 4, 6},
		{2, 2, 2, 2},
		{0, 0, 0},
	}
	actual := extrapolateLeft(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
