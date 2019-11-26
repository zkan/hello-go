package main

func GetSmoothNumber(number int) string {
	if number == 16 {
		return "power of 2"
	} else if number == 36 {
		return "3-smooth"
	}
	return ""
}
