package statistics

import (
	"fmt"
	"testing"
)

func ExampleStandardError() {
	X := []uint8{8, 3, 6, 2, 7, 1, 8, 3, 7, 4, 8}
	stddev := StandardError(X...)
	fmt.Printf("The standard error of %v is %v.\n", X, stddev)

	// Output:
	// The standard error of [8 3 6 2 7 1 8 3 7 4 8] is 0.7586218672636642.
}

func TestStandardError(t *testing.T) {
	cases := []testArrayCase{
		{
			name:     "test 1",
			input:    []float64{8, 3, 6, 2, 7, 1, 8, 3, 7, 4, 8},
			expected: "float64 0.7586219",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual := StandardError(c.input.([]float64)...)
			r := fmt.Sprintf("%T %.7f", actual, actual)
			if r != c.expected {
				t.Logf("expected %v, got %v", c.expected, r)
				t.FailNow()
			}
		})
	}
}
