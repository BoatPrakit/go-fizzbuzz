package fizzbuzz

import "strconv"

func FizzBuzz(input int) string {
	if input%3 == 0 || input == 5 {
		return Fizz(input%3 == 0) + Buzz(input)
	}
	return strconv.Itoa(input)
}

func Fizz(b bool) string {
	fizz := map[bool]string{
		true:  "Fizz",
		false: "",
	}
	return fizz[b]
}

func Buzz(input int) string {
	buzz := map[bool]string{
		true:  "Buzz",
		false: "",
	}
	return buzz[input == 5]
}
