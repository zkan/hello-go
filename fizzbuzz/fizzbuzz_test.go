package main

import "testing"

func TestInputThreeShouldGetFizz(t *testing.T) {
	expected := "Fizz"
	result := FizzBuzz(3)
	if result != expected {
		t.Errorf("It should get %s but got %s", expected, result)
	}
}
