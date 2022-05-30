package statistics

import (
	"fmt"
	"testing"
)

func ExampleVariance() {
	X := []uint8{8, 3, 6, 2, 7, 1, 8, 3, 7, 4, 8}
	variance := Variance(X...)
	fmt.Printf("The variance of %v is %v.\n", X, variance)

	// Output:
	// The variance of [8 3 6 2 7 1 8 3 7 4 8] is 6.330578512396695.
}

func TestVariance(t *testing.T) {
	cases := []testArrayCase{
		{
			name:     "test 1",
			input:    []float64{8, 3, 6, 2, 7, 1, 8, 3, 7, 4, 8},
			expected: "float64 6.3305785",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual := Variance(c.input.([]float64)...)
			r := fmt.Sprintf("%T %.7f", actual, actual)
			if r != c.expected {
				t.Logf("expected %v, got %v", c.expected, r)
				t.FailNow()
			}
		})
	}
}
