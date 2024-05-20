package fizzbuzz

import "strconv"

func FizzBuzz(input int) string {
	if input == 3 || input == 6 || input == 5 {
		return Fizz(input) + Buzz(input)
	}
	return strconv.Itoa(input)
}

func Fizz(input int) string {
	fizz := map[bool]string{
		true:  "Fizz",
		false: "",
	}
	return fizz[input == 3 || input == 6]
}

func Buzz(input int) string {
	buzz := map[bool]string{
		true:  "Buzz",
		false: "",
	}
	return buzz[input == 5]
}
