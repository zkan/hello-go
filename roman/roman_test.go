package roman

import "testing"

func TestOneShouldGetI(t *testing.T) {
	expected := "I"
	result := ConvertToRoman(1)

	if result != expected {
		t.Errorf("%s is expected result but got %s", expected, result)
	}
}
