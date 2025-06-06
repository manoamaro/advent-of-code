package fn

import "testing"

func TestIdentity(t *testing.T) {
	if Identity(42) != 42 {
		t.Errorf("expected 42")
	}
}

func TestEq(t *testing.T) {
	eq := Eq(5)
	if !eq(5) || eq(4) {
		t.Errorf("eq function failed")
	}
}
