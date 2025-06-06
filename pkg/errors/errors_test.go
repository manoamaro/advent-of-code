package errors

import "fmt"
import "testing"

func TestMustReturnsValue(t *testing.T) {
	if Must(1, nil) != 1 {
		t.Errorf("expected 1")
	}
}

func TestMustPanics(t *testing.T) {
	defer func() { recover() }()
	Must(0, fmt.Errorf("err"))
}
