package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	t.Run("should return 1 when input 1", func(t *testing.T) {
		input := 1

		got := FizzBuzz(input)
		want := "1"

		if got != want {
			t.Errorf("want %v but got %v", want, got)
		}
	})

	t.Run("should return 2 when input 2", func(t *testing.T) {
		input := 2

		got := FizzBuzz(input)
		want := "2"

		if got != want {
			t.Errorf("want %v but got %v", want, got)
		}
	})
}
