package fizzbuzz

import (
	"fmt"
	"testing"
)

func TestFizzBuzz(t *testing.T) {

	testcases := []struct {
		input    int
		expected string
	}{
		{1, "1"},
		{2, "2"},
		{3, "Fizz"},
		{4, "4"},
		{5, "Buzz"},
		{6, "Fizz"},
		{7, "7"},
		{8, "8"},
		{9, "Fizz"},
		{10, "Buzz"},
		{11, "11"},
		{12, "Fizz"},
		{13, "13"},
		{14, "14"},
		{15, "FizzBuzz"},
	}

	for _, testcase := range testcases {
		t.Run(fmt.Sprintf("should return %v when input %v", testcase.expected, testcase.input), func(t *testing.T) {
			input := testcase.input

			got := FizzBuzz(input)
			want := testcase.expected

			if got != want {
				t.Errorf("want %v but got %v", want, got)
			}
		})
	}
}
