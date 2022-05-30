package statistics

import (
	"fmt"
	"testing"
)

func ExampleSampleStandardDeviation() {
	X := []uint8{8, 3, 6, 2, 7, 1, 8, 3, 7, 4, 8}
	stddev := SampleStandardDeviation(X...)
	fmt.Printf("The sample standard deviation of %v is %v.\n", X, stddev)

	// Output:
	// The sample standard deviation of [8 3 6 2 7 1 8 3 7 4 8] is 2.638870281699418.
}

func TestSampleStandardDeviation(t *testing.T) {
	cases := []testArrayCase{
		{
			name:     "test 1",
			input:    []float64{8, 3, 6, 2, 7, 1, 8, 3, 7, 4, 8},
			expected: "float64 2.6388703",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual := SampleStandardDeviation(c.input.([]float64)...)
			r := fmt.Sprintf("%T %.7f", actual, actual)
			if r != c.expected {
				t.Logf("expected %v, got %v", c.expected, r)
				t.FailNow()
			}
		})
	}
}
