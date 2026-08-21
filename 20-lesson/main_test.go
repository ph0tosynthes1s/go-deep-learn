package main

import "testing"

func TestNPlus(t *testing.T) {
	result := nPlus(1)
	if result != 2 {
		t.Errorf("nPlus(%d) = %d; want 2", nPlus(1), result)
	}
}
