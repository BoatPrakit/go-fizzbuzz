package fizzbuzz

import "strconv"

func FizzBuzz(input int) string {
	if input%3 == 0 || input%5 == 0 {
		return Fizz(input%3 == 0) + Buzz(input%5 == 0)
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

func Buzz(b bool) string {
	buzz := map[bool]string{
		true:  "Buzz",
		false: "",
	}
	return buzz[b]
}
