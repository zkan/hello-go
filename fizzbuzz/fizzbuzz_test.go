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
		test{number: 5, result: "Buzz"},
		test{number: 10, result: "Buzz"},
		test{number: 15, result: "FizzBuzz"},
		test{number: 30, result: "FizzBuzz"},
	}

	for _, each := range tests {
		var result string = FizzBuzz(each.number)

		if result != each.result {
			t.Errorf("It should get %s but got %s", each.result, result)
		}
	}
}
