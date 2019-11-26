package main

import "reflect"
import "testing"

func TestIsPrime(t *testing.T) {
	t.Run("Input2ShouldGetTrue", func(t *testing.T) {
		expected := true
		result := isPrime(2)

		if result != expected {
			t.Errorf("It should get %t but got %t", expected, result)
		}
	})

	t.Run("Input13ShouldGetTrue", func(t *testing.T) {
		expected := true
		result := isPrime(13)

		if result != expected {
			t.Errorf("It should get %t but got %t", expected, result)
		}
	})
}

func TestGetPrimeFactors(t *testing.T) {
	t.Run("Input12ShouldGet223", func(t *testing.T) {
		expected := []int{2, 2, 3}
		results := getPrimeFactors(12)

		if !reflect.DeepEqual(results, expected) {
			t.Errorf("It should get %v but got %v", expected, results)
		}
	})

	t.Run("Input60ShouldGet2235", func(t *testing.T) {
		expected := []int{2, 2, 3, 5}
		results := getPrimeFactors(60)

		if !reflect.DeepEqual(results, expected) {
			t.Errorf("It should get %v but got %v", expected, results)
		}
	})
}

func TestGetMaxIn(t *testing.T) {
	expected := 5
	result := getMaxIn([]int{2, 2, 3, 5})

	if result != expected {
		t.Errorf("It should get %v but got %v", expected, result)
	}
}

func TestGetSmoothNumbers(t *testing.T) {
	t.Run("Input16ShouldGetPowerOf2", func(t *testing.T) {
		expected := "power of 2"
		result := GetSmoothNumber(16)

		if result != expected {
			t.Errorf("It should get %s but got %s", expected, result)
		}
	})

	t.Run("Input36ShouldGet3Smooth", func(t *testing.T) {
		expected := "3-smooth"
		result := GetSmoothNumber(36)

		if result != expected {
			t.Errorf("It should get %s but got %s", expected, result)
		}
	})

	t.Run("Input60ShouldGet5Smooth", func(t *testing.T) {
		expected := "5-smooth"
		result := GetSmoothNumber(60)

		if result != expected {
			t.Errorf("It should get %s but got %s", expected, result)
		}
	})
}
