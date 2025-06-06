package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func createCache(year, day int, data string) {
	os.Mkdir(".cache", 0755)
	os.WriteFile(filepath.Join(".cache", fmt.Sprintf("input_%d_%d.txt", year, day)), []byte(data), 0644)
}

func TestReadInputFromCache(t *testing.T) {
	year, day := 9999, 1
	createCache(year, day, "data\nline2\n")
	res, err := ReadInput(year, day)
	if err != nil || res != "data\nline2\n" {
		t.Fatalf("cache read failed")
	}
}

func TestReadInputLines(t *testing.T) {
	year, day := 9999, 2
	createCache(year, day, "a\nb\n")
	lines, err := ReadInputLines(year, day)
	if err != nil || len(lines) != 2 || lines[0] != "a" || lines[1] != "b" {
		t.Fatalf("read lines failed")
	}
}

func TestReadInputSlice2d(t *testing.T) {
	year, day := 9999, 3
	createCache(year, day, "ab\ncd\n")
	slice, err := ReadInputSlice2d(year, day)
	if err != nil || len(slice) != 2 || slice[0][0] != "a" || slice[1][1] != "d" {
		t.Fatalf("slice2d failed")
	}
}
