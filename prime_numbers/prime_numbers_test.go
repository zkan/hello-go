package main

import "reflect"
import "testing"

type test struct {
	number  int
	results []int
}

func TestGetPrimeNumbers(t *testing.T) {
	var tests = []test{
		test{number: 7, results: []int{2, 3, 5, 7}},
		test{number: 15, results: []int{2, 3, 5, 7, 11, 13}},
		test{number: 19, results: []int{2, 3, 5, 7, 11, 13, 17, 19}},
	}

	for _, each := range tests {
		var results []int = GetPrimeNumbers(each.number)

		if !reflect.DeepEqual(results, each.results) {
			t.Errorf("It should get %v but got %v", each.results, results)
		}
	}
}
