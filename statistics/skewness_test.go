package statistics

import (
	"fmt"
	"testing"
)

func ExampleSkewness() {
	X := []uint8{8, 3, 6, 2, 7, 1, 8, 3, 7, 4, 8}
	skewness := Skewness(X...)
	fmt.Printf("The skewness of %v is %.6f.\n", X, skewness)

	// Output:
	// The skewness of [8 3 6 2 7 1 8 3 7 4 8] is -0.274241.
}

func TestSkewness(t *testing.T) {
	cases := []testArrayCase{
		{
			name:     "test 1",
			input:    []float64{8, 3, 6, 2, 7, 1, 8, 3, 7, 4, 8},
			expected: "float64 -0.2742407",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual := Skewness(c.input.([]float64)...)
			r := fmt.Sprintf("%T %.7f", actual, actual)
			if r != c.expected {
				t.Logf("expected %v, got %v", c.expected, r)
				t.FailNow()
			}
		})
	}
}
