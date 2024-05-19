package fizzbuzz

import "strconv"

func FizzBuzz(input int) string {
	return IsFizzBuzz(input)
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

	isNormal := map[bool]string{
		true:  strconv.Itoa(n),
		false: "",
	}

	return isFizz[n%3 == 0] + isBuzz[n%5 == 0] + isNormal[n%3 != 0 && n%5 != 0]
}
