package main

func pi() float64 {
	var result float64
	numerator := 4.0
	plus := true
	for i := 1; i < 1000000; i = i + 2 {
		if plus == true {
			result += numerator / float64(i)
			plus = false
		} else {
			result -= numerator / float64(i)
			plus = true
		}
	}
	return result
}
