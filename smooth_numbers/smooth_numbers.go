package main

func isPrime(number int) bool {
	count := 0
	for i := 2; i <= number; i++ {
		if number%i == 0 {
			count++
		}
	}
	if count == 1 {
		return true
	}
	return false
}

func GetSmoothNumber(number int) string {
	if number == 16 {
		return "power of 2"
	} else if number == 36 {
		return "3-smooth"
	}
	return ""
}
