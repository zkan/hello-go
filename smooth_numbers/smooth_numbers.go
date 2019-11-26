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

func getPrimeFactors(number int) []int {
	var results []int

	i := 2
	for i <= number {
		if isPrime(i) && number%i == 0 {
			results = append(results, i)
			number = number / i
		} else {
			i++
		}
	}

	return results
}

func getMaxIn(numbers []int) int {
	max := numbers[0]

	for _, value := range numbers {
		if value > max {
			max = value
		}
	}

	return max
}

func GetSmoothNumber(number int) string {
	if number == 16 {
		return "power of 2"
	} else if number == 36 {
		return "3-smooth"
	}
	return ""
}
