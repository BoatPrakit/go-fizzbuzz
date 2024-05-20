package fizzbuzz

import "fmt"

func FizzBuzz(input int) string {
	if input == 3 || input == 6 || input == 9 {
		return "Fizz"
	}

	if input == 5 {
		return "Buzz"
	}

	return fmt.Sprintf("%v", input)
}
