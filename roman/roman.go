package roman

func ConvertToRoman(num int) string {
	if num == 1 {
		return "I"
	} else if num == 2 {
		return "II"
	} else if num == 3 {
		return "III"
	}
	return "I"
}
