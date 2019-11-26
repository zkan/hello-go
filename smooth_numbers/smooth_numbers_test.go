package main

import "testing"

func TestGetSmoothNumbers(t *testing.T) {
	t.Run("Input16ShouldGetPowerOf2", func(t *testing.T) {
		expected := "power of 2"
		result := GetSmoothNumber(16)

		if result != expected {
			t.Errorf("It should get %s but got %s", expected, result)
		}
	})
}
