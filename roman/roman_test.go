package roman

import "testing"

func TestOneShouldGetI(t *testing.T) {
	expected := "I"
	result := ConvertToRoman(1)

	if result != expected {
		t.Errorf("%s is expected result but got %s", expected, result)
	}
}

func TestTwoShouldGetII(t *testing.T) {
	expected := "II"
	result := ConvertToRoman(2)

	if result != expected {
		t.Errorf("%s is expected result but got %s", expected, result)
	}
}

func TestThreeShouldGetIII(t *testing.T) {
	expected := "III"
	result := ConvertToRoman(3)

	if result != expected {
		t.Errorf("%s is expected result but got %s", expected, result)
	}
}

func TestFourShouldGetIV(t *testing.T) {
	expected := "IV"
	result := ConvertToRoman(4)

	if result != expected {
		t.Errorf("%s is expected result but got %s", expected, result)
	}
}

func TestFiveShouldGetV(t *testing.T) {
	expected := "V"
	result := ConvertToRoman(5)

	if result != expected {
		t.Errorf("%s is expected result but got %s", expected, result)
	}
}

func TestSixShouldGetVI(t *testing.T) {
	expected := "VI"
	result := ConvertToRoman(6)

	if result != expected {
		t.Errorf("%s is expected result but got %s", expected, result)
	}
}

func TestSevenShouldGetVII(t *testing.T) {
	expected := "VII"
	result := ConvertToRoman(7)

	if result != expected {
		t.Errorf("%s is expected result but got %s", expected, result)
	}
}

func TestEightShouldGetVIII(t *testing.T) {
	expected := "VIII"
	result := ConvertToRoman(8)

	if result != expected {
		t.Errorf("%s is expected result but got %s", expected, result)
	}
}

func TestNineShouldGetIX(t *testing.T) {
	expected := "IX"
	result := ConvertToRoman(9)

	if result != expected {
		t.Errorf("%s is expected result but got %s", expected, result)
	}
}
