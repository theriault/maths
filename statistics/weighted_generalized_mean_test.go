package statistics

import (
	"fmt"
	"math"
	"testing"
)

func ExampleWeightedGeneralizedMean() {
	X := []uint8{8, 7, 3, 2, 6, 11, 6, 7, 2, 1, 7}
	W := []uint8{1, 2, 1, 1, 2, 1, 2, 1, 2, 1, 2}
	mean, _ := WeightedGeneralizedMean(X, W, 1)
	fmt.Printf("The 1st weighted generalized mean of %v with weights of %v is %.1f.\n", X, W, mean)

	// Output:
	// The 1st weighted generalized mean of [8 7 3 2 6 11 6 7 2 1 7] with weights of [1 2 1 1 2 1 2 1 2 1 2] is 5.5.
}

func TestWeightedGeneralizedMean(t *testing.T) {
	cases := []struct {
		name     string
		input    any
		weights  any
		exponent float64
		expected string
	}{
		{
			name:     "not same length inputs",
			input:    []float64{1000},
			weights:  []float64{1, 0.5},
			exponent: 1,
			expected: "float64 0 x and w are not the same lengths",
		},
		{
			name:     "empty inputs",
			input:    []float64{},
			weights:  []float64{},
			exponent: 1,
			expected: "float64 0 x must have at least 1 element",
		},
		{
			name:     "minimum",
			input:    []float64{1000, 1},
			weights:  []float64{1, 0.5},
			exponent: math.Inf(-1),
			expected: "float64 0.5 <nil>",
		},
		{
			name:     "minimum with NaN",
			input:    []float64{1, math.NaN()},
			weights:  []float64{0.5, 1},
			exponent: math.Inf(-1),
			expected: "float64 NaN <nil>",
		},
		{
			name:     "harmonic mean",
			input:    []float64{1, 1000},
			weights:  []float64{1, 1},
			exponent: -1,
			expected: "float64 1.9980019980019983 <nil>",
		},
		{
			name:     "geometric mean",
			input:    []float64{1, 1000},
			weights:  []float64{1, 1},
			exponent: 0,
			expected: "float64 31.622776601683793 <nil>",
		},
		{
			name:     "geometric mean with NaN",
			input:    []float64{1, math.NaN()},
			weights:  []float64{1, 1},
			exponent: 0,
			expected: "float64 NaN <nil>",
		},
		{
			name:     "quadratic mean",
			input:    []float64{1, 1000},
			weights:  []float64{1, 1},
			exponent: 2,
			expected: "float64 707.1071347398497 <nil>",
		},
		{
			name:     "cubic mean",
			input:    []float64{1, 1000},
			weights:  []float64{1, 1},
			exponent: 3,
			expected: "float64 793.7005262486666 <nil>",
		},
		{
			name:     "maximum",
			input:    []float64{1, 1000},
			weights:  []float64{1, 2},
			exponent: math.Inf(+1),
			expected: "float64 2000 <nil>",
		},
		{
			name:     "maximum with NaN",
			input:    []float64{1, math.NaN()},
			weights:  []float64{1, 2},
			exponent: math.Inf(+1),
			expected: "float64 NaN <nil>",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual, err := WeightedGeneralizedMean(c.input.([]float64), c.weights.([]float64), c.exponent)
			if actual := fmt.Sprintf("%T %v %v", actual, actual, err); actual != c.expected {
				t.Logf("expected %v, got %v", c.expected, actual)
				t.FailNow()
			}
		})
	}
}

func TestWeightedPowerMean(t *testing.T) {
	cases := []struct {
		name     string
		input    any
		weights  any
		exponent float64
		expected string
	}{
		{
			name:     "quadratic mean",
			input:    []float64{1, 1000},
			weights:  []float64{1, 1},
			exponent: 2,
			expected: "float64 707.1071347398497 <nil>",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual, err := WeightedGeneralizedMean(c.input.([]float64), c.weights.([]float64), c.exponent)
			if actual := fmt.Sprintf("%T %v %v", actual, actual, err); actual != c.expected {
				t.Logf("expected %v, got %v", c.expected, actual)
				t.FailNow()
			}
		})
	}
}

func ExampleWeightedMean() {
	X := []uint8{8, 7, 3, 2, 6, 11, 6, 7, 2, 1, 7}
	W := []uint8{1, 2, 1, 1, 2, 1, 2, 1, 2, 1, 2}
	mean, _ := WeightedMean(X, W)
	fmt.Printf("The weighted mean of %v with weights of %v is %.1f.\n", X, W, mean)

	// Output:
	// The weighted mean of [8 7 3 2 6 11 6 7 2 1 7] with weights of [1 2 1 1 2 1 2 1 2 1 2] is 5.5.
}

func TestWeightedMean(t *testing.T) {
	cases := []struct {
		name     string
		input    any
		weights  any
		expected string
	}{
		{
			name:     "not same length inputs",
			input:    []float64{1000},
			weights:  []float64{1, 0.5},
			expected: "float64 0 x and w are not the same lengths",
		},
		{
			name:     "empty inputs",
			input:    []float64{},
			weights:  []float64{},
			expected: "float64 0 x must have at least 1 element",
		},
		{
			name:     "test 1",
			input:    []float64{1, 1000},
			weights:  []float64{1, 2},
			expected: "float64 667 <nil>",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			var actual interface{}
			var err error
			actual, err = WeightedMean(c.input.([]float64), c.weights.([]float64))
			if actual := fmt.Sprintf("%T %v %v", actual, actual, err); actual != c.expected {
				t.Logf("expected %v, got %v", c.expected, actual)
				t.FailNow()
			}
		})
	}
}

func ExampleWeightedGeometricMean() {
	X := []uint8{8, 7, 3, 2, 6, 11, 6, 7, 2, 1, 7}
	W := []uint8{1, 2, 1, 1, 2, 1, 2, 1, 2, 1, 2}
	mean, _ := WeightedGeometricMean(X, W)
	fmt.Printf("The weighted geometric mean of %v with weights of %v is %.1f.\n", X, W, mean)

	// Output:
	// The weighted geometric mean of [8 7 3 2 6 11 6 7 2 1 7] with weights of [1 2 1 1 2 1 2 1 2 1 2] is 4.6.
}

func TestWeightedGeometricMean(t *testing.T) {
	cases := []struct {
		name     string
		input    any
		weights  any
		expected string
	}{
		{
			name:     "not same length inputs",
			input:    []float64{1000},
			weights:  []float64{1, 0.5},
			expected: "float64 0 x and w are not the same lengths",
		},
		{
			name:     "empty inputs",
			input:    []float64{},
			weights:  []float64{},
			expected: "float64 0 x must have at least 1 element",
		},
		{
			name:     "test 1",
			input:    []float64{1, 1000},
			weights:  []float64{1, 2},
			expected: "float64 99.99999999999996 <nil>",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			var actual interface{}
			var err error
			actual, err = WeightedGeometricMean(c.input.([]float64), c.weights.([]float64))
			if actual := fmt.Sprintf("%T %v %v", actual, actual, err); actual != c.expected {
				t.Logf("expected %v, got %v", c.expected, actual)
				t.FailNow()
			}
		})
	}
}

func ExampleWeightedHarmonicMean() {
	X := []uint8{8, 7, 3, 2, 6, 11, 6, 7, 2, 1, 7}
	W := []uint8{1, 2, 1, 1, 2, 1, 2, 1, 2, 1, 2}
	mean, _ := WeightedHarmonicMean(X, W)
	fmt.Printf("The weighted harmonic mean of %v with weights of %v is %.1f.\n", X, W, mean)

	// Output:
	// The weighted harmonic mean of [8 7 3 2 6 11 6 7 2 1 7] with weights of [1 2 1 1 2 1 2 1 2 1 2] is 5.8.
}
func TestWeightedHarmonicMean(t *testing.T) {
	cases := []struct {
		name     string
		input    any
		weights  any
		expected string
	}{
		{
			name:     "not same length inputs",
			input:    []float64{1000},
			weights:  []float64{1, 0.5},
			expected: "float64 0 x and w are not the same lengths",
		},
		{
			name:     "empty inputs",
			input:    []float64{},
			weights:  []float64{},
			expected: "float64 0 x must have at least 1 element",
		},
		{
			name:     "test 1",
			input:    []float64{1, 1000},
			weights:  []float64{1, 2},
			expected: "float64 2.9985007496251876 <nil>",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			var actual interface{}
			var err error
			actual, err = WeightedHarmonicMean(c.input.([]float64), c.weights.([]float64))
			if actual := fmt.Sprintf("%T %v %v", actual, actual, err); actual != c.expected {
				t.Logf("expected %v, got %v", c.expected, actual)
				t.FailNow()
			}
		})
	}
}
