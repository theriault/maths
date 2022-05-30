package statistics

import (
	"fmt"
	"testing"
)

func ExampleSampleStandardError() {
	X := []uint8{8, 3, 6, 2, 7, 1, 8, 3, 7, 4, 8}
	stddev := StandardError(X...)
	fmt.Printf("The sample standard error of %v is %v.\n", X, stddev)

	// Output:
	// The sample standard error of [8 3 6 2 7 1 8 3 7 4 8] is 0.7586218672636642.
}

func TestSampleStandardError(t *testing.T) {
	cases := []testArrayCase{
		{
			name:     "test 1",
			input:    []float64{8, 3, 6, 2, 7, 1, 8, 3, 7, 4, 8},
			expected: "float64 0.7956493",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual := SampleStandardError(c.input.([]float64)...)
			r := fmt.Sprintf("%T %.7f", actual, actual)
			if r != c.expected {
				t.Logf("expected %v, got %v", c.expected, r)
				t.FailNow()
			}
		})
	}
}
