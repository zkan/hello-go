package main

func isPrime(number int) bool {
	count := 0
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			count++
		}
	}
	if count == 2 {
		return true
	}

	return false
}

func GetPrimeNumbers(number int) []int {
	var results []int

	for i := 2; i <= number; i++ {
		if isPrime(i) {
			results = append(results, i)
		}
	}

	return results
}
