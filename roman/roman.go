package roman

func ConvertToRoman(num int) string {
	var remainder int
	remainder = num % 5

	var result string
	if remainder > 0 && remainder < 4 {
		for i := 0; i < num; i++ {
			result += "I"
		}
	} else if remainder == 4 {
		result = "IV"
	} else {
		result = "V"
	}
	return result
}
