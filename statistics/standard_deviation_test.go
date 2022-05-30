package statistics

import (
	"fmt"
	"testing"
)

func ExampleStandardDeviation() {
	X := []uint8{8, 3, 6, 2, 7, 1, 8, 3, 7, 4, 8}
	stddev := StandardDeviation(X...)
	fmt.Printf("The standard deviation of %v is %v.\n", X, stddev)

	// Output:
	// The standard deviation of [8 3 6 2 7 1 8 3 7 4 8] is 2.5160640914723724.
}

func TestStandardDeviation(t *testing.T) {
	cases := []testArrayCase{
		{
			name:     "test 1",
			input:    []float64{8, 3, 6, 2, 7, 1, 8, 3, 7, 4, 8},
			expected: "float64 2.5160641",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual := StandardDeviation(c.input.([]float64)...)
			r := fmt.Sprintf("%T %.7f", actual, actual)
			if r != c.expected {
				t.Logf("expected %v, got %v", c.expected, r)
				t.FailNow()
			}
		})
	}
}
