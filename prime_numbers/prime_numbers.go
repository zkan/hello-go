package main

func GetPrimeNumbers(number int) []int {
	var results []int

	for i := 2; i <= number; i++ {
		count := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				count++
			}
		}
		if count == 2 {
			results = append(results, i)
		}
	}

	return results
}
