package roman

func ConvertToRoman(num int) string {
	var remainder int
	remainder = num % 5

	var result string

	if num >= 5 {
		result = "V"
	}

	if remainder > 0 && remainder < 4 {
		for i := 0; i < remainder; i++ {
			result += "I"
		}
	} else if remainder == 4 {
		result = "IV"
	}
	return result
}
