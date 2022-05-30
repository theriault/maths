package statistics

import (
	"fmt"
	"testing"
)

func ExampleExcessSampleKurtosis() {
	X := []uint8{8, 3, 6, 2, 7, 1, 8, 3, 7, 4, 8}
	skewness := ExcessSampleKurtosis(X...)
	fmt.Printf("The sample kurtosis of %v is %.6f.\n", X, skewness)

	// Output:
	// The sample kurtosis of [8 3 6 2 7 1 8 3 7 4 8] is -1.644500.
}

func TestExcessSampleKurtosis(t *testing.T) {
	cases := []testArrayCase{
		{
			name:     "test 1",
			input:    []float64{8, 3, 6, 2, 7, 1, 8, 3, 7, 4, 8},
			expected: "float64 -1.6444996",
		},
		{
			name:     "test 2",
			input:    []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: "float64 -1.2000000",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual := ExcessSampleKurtosis(c.input.([]float64)...)
			r := fmt.Sprintf("%T %.7f", actual, actual)
			if r != c.expected {
				t.Logf("expected %v, got %v", c.expected, r)
				t.FailNow()
			}
		})
	}
}
