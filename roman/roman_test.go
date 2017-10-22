package roman

import "testing"

func TestNumberLessThanFourShouldGetIAndIIAndIII(t *testing.T) {
	m := map[int]string{
		1: "I",
		2: "II",
		3: "III",
	}

	for k, v := range m {
		expected := v
		result := ConvertToRoman(k)

		if result != expected {
			t.Errorf("%s is expected result but got %s", expected, result)
		}
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

func TestNumberLessThanNineShouldGetVIAndVIIAndVIII(t *testing.T) {
	m := map[int]string{
		6: "VI",
		7: "VII",
		8: "VIII",
	}

	for k, v := range m {
		expected := v
		result := ConvertToRoman(k)

		if result != expected {
			t.Errorf("%s is expected result but got %s", expected, result)
		}
	}
}

func TestNineShouldGetIX(t *testing.T) {
	expected := "IX"
	result := ConvertToRoman(9)

	if result != expected {
		t.Errorf("%s is expected result but got %s", expected, result)
	}
}

func TestTenShouldGetX(t *testing.T) {
	expected := "X"
	result := ConvertToRoman(10)

	if result != expected {
		t.Errorf("%s is expected result but got %s", expected, result)
	}
}

func TestNumberLessThanFourTeenShouldGetXIAndXIIAndXIII(t *testing.T) {
	m := map[int]string{
		11: "XI",
		12: "XII",
		13: "XIII",
	}

	for k, v := range m {
		expected := v
		result := ConvertToRoman(k)

		if result != expected {
			t.Errorf("%s is expected result but got %s", expected, result)
		}
	}
}
