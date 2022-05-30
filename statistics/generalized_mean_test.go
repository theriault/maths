package statistics

import (
	"fmt"
	"math"
	"testing"
)

func ExampleGeneralizedMean() {
	vals := []uint8{8, 7, 3, 2, 6, 11, 6, 7, 2, 1, 7}
	max := GeneralizedMean(vals, 2)
	fmt.Printf("The 2nd generalized mean of %v is %.8f.\n", vals, max)

	// Output:
	// The 2nd generalized mean of [8 7 3 2 6 11 6 7 2 1 7] is 6.19383858.
}

func TestGeneralizedMean(t *testing.T) {
	cases := []struct {
		name     string
		input    any
		exponent float64
		expected string
	}{
		{
			name:     "minimum",
			input:    []float64{1, 1000},
			exponent: math.Inf(-1),
			expected: "float64 1",
		},
		{
			name:     "harmonic mean",
			input:    []float64{1, 1000},
			exponent: -1,
			expected: "float64 1.9980019980019983",
		},
		{
			name:     "geometric mean",
			input:    []float64{1, 1000},
			exponent: 0,
			expected: "float64 31.622776601683793",
		},
		{
			name:     "quadratic mean",
			input:    []float64{1, 1000},
			exponent: 2,
			expected: "float64 707.1071347398497",
		},
		{
			name:     "cubic mean",
			input:    []float64{1, 1000},
			exponent: 3,
			expected: "float64 793.7005262486666",
		},
		{
			name:     "maximum",
			input:    []float64{1, 1000},
			exponent: math.Inf(+1),
			expected: "float64 1000",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual := GeneralizedMean(c.input.([]float64), c.exponent)
			if fmt.Sprintf("%T %v", actual, actual) != c.expected {
				t.Logf("expected %v, got %v", c.expected, actual)
				t.FailNow()
			}
		})
	}
}

func TestPowerMean(t *testing.T) {
	cases := []struct {
		name     string
		input    any
		exponent float64
		expected string
	}{
		{
			name:     "quadratic mean",
			input:    []float64{1, 1000},
			exponent: 2,
			expected: "float64 707.1071347398497",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual := PowerMean(c.input.([]float64), c.exponent)
			if fmt.Sprintf("%T %v", actual, actual) != c.expected {
				t.Logf("expected %v, got %v", c.expected, actual)
				t.FailNow()
			}
		})
	}
}
