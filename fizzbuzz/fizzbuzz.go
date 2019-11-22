package main

func FizzBuzz(number int) string {
	if number%3 == 0 {
		return "Fizz"
	}
	if number%5 == 0 {
		return "Buzz"
	}
	return ""
}
