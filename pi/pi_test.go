package main

import "testing"

func TestPiShouldGetPiValue(t *testing.T) {
	expected := 3.141590653589692
	result := pi()
	if result != expected {
		t.Errorf("It should get %s but got %s", expected, result)
	}
}
