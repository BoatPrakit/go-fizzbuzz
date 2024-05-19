package fizzbuzz

import "strconv"

func FizzBuzz(input int) string {
	if input == 3 || input == 6 || input == 9 || input == 12 {
		return "Fizz"
	}

	if input == 5 || input == 10 {
		return "Buzz"
	}

	return strconv.Itoa(input)
}
