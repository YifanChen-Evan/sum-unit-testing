package math

import (
	"testing"
)

func TestSum(t *testing.T) {
	// define the inputs and expected outputs for testing cases
	cases := []struct {
		a, b     int
		expected int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{-1, 1, 0},
	}

	// iterate testing cases
	for _, c := range cases {
		result := Sum(c.a, c.b) // !!! don't write "math.Sum()", because they are in the same package
		if result != c.expected {
			t.Errorf("Sum(%d, %d) = %d, expected %d", c.a, c.b, result, c.expected) // !!! can't miss the Error"f", because "%d"
		}
	}
}
