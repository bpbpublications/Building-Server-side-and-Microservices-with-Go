package test

import "testing"

func TestSum(t *testing.T) {
	result := sum(2, 5)
	if result != 7 {
		t.Errorf("Incorrect result, got: %v, want: %v.", result, 7)
	}
}
