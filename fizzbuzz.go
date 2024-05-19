package fizzbuzz

import "strconv"

func FizzBuzz(input int) string {
	if input%3 == 0 || input%5 == 0 {
		return IsFizzBuzz(input)
	}

	return strconv.Itoa(input)
}

func IsFizzBuzz(n int) string {
	isFizz := map[bool]string{
		true:  "Fizz",
		false: "",
	}

	isBuzz := map[bool]string{
		true:  "Buzz",
		false: "",
	}

	return isFizz[n%3 == 0] + isBuzz[n%5 == 0]
}
