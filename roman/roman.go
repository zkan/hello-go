package roman

func ConvertToRoman(num int) string {
	if num < 4 {
		var result string
		for i := 0; i < num; i++ {
			result += "I"
		}
		return result
	} else {
		return "IV"
	}
}
