package main

import "testing"

type test struct {
	number int
	result string
}

func TestFizzBuzz(t *testing.T) {
	var tests = []test{
		test{number: 3, result: "Fizz"},
		test{number: 6, result: "Fizz"},
	}

	for _, each := range tests {
		var result string = FizzBuzz(each.number)

		if result != each.result {
			t.Errorf("It should get %s but got %s", each.result, result)
		}
	}
}
