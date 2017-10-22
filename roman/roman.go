package roman

func ConvertToRoman(num int) string {
	var result string

	if num >= 9 {
		result = "X"
	} else if num >= 4 {
		result = "V"
	}

	remainder := num % 5
	if remainder > 0 && remainder < 4 {
		for i := 0; i < remainder; i++ {
			result += "I"
		}
	} else if remainder == 4 {
		result = "I" + result
	}
	return result
}
