package main

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
}
