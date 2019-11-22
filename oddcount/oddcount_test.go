package main

import "reflect"
import "testing"

type test struct {
	number  int
	results []int
}

func TestOddCount(t *testing.T) {
	var tests = []test{
		test{number: 7, results: []int{1, 3, 5}},
		test{number: 15, results: []int{1, 3, 5, 7, 9, 11, 13}},
	}

	for _, each := range tests {
		var results []int = OddCount(each.number)

		if !reflect.DeepEqual(results, each.results) {
			t.Errorf("It should get %v but got %v", each.results, results)
		}
	}
}
