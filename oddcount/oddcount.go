package main

func OddCount(number int) []int {
	var results []int

	for i := 1; i < number; i++ {
		if i%2 != 0 {
			results = append(results, i)
		}
	}

	return results
}
