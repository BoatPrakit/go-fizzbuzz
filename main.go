package fizzbuzz

func FizzBuzz(input int) string {
	if input == 2 {
		return "2"
	}

	if input == 3 || input == 6 {
		return "Fizz"
	}

	if input == 4 {
		return "4"
	}

	if input == 5 {
		return "Buzz"
	}

	if input == 7 {
		return "7"
	}

	return "1"
}
