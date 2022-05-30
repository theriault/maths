package statistics

import (
	"fmt"
	"testing"
)

func ExampleCentralMoment() {
	X := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	moment := CentralMoment(X, 2)
	fmt.Printf("The 2nd central moment of %v is %.2f.\n", X, moment)

	// Output:
	// The 2nd central moment of [1 2 3 4 5 6 7 8 9 10] is 8.25.
}

func TestCentralMoment(t *testing.T) {
	cases := []struct {
		name     string
		power    int
		input    []float64
		expected string
	}{
		{
			name:     "test 1",
			power:    2,
			input:    []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: "float64 8.25000",
		},
		{
			name:     "test 2",
			power:    4,
			input:    []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: "float64 120.86250",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual := CentralMoment(c.input, c.power)
			r := fmt.Sprintf("%T %.5f", actual, actual)
			if r != c.expected {
				t.Logf("expected %v, got %v", c.expected, r)
				t.FailNow()
			}
		})
	}
}
