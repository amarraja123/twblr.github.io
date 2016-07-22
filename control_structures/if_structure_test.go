package control_structures

import "testing"

var fizzBuzzTests = []struct {
	input  int32
	output string
}{
	{2, "2"},
	{4, "4"},
	{100, "100"},
	{3, "Fizz"},
	{5, "Buzz"},
	{15, "FizzBuzz"},
	{999, "Fizz"},
}

func TestFizzBuzz(t *testing.T) {
	for _, fb := range fizzBuzzTests {
		actualOutput := fizzBuzz(fb.input)
		if fb.output != actualOutput {
			t.Errorf("For input %d, expected output was %s, But got: %s", fb.input, fb.output, actualOutput)
		}
	}
}
