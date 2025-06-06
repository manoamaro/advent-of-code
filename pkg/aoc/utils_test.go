package aoc

import (
	"path/filepath"
	"runtime"
	"testing"
)

func TestGetSourceFilePath(t *testing.T) {
	_, this, _, _ := runtime.Caller(0)
	expected := filepath.Join(filepath.Dir(this), "file")
	if GetSourceFilePath("file") != expected {
		t.Fatalf("path mismatch")
	}
}
