package statistics

import (
	"fmt"
	"math"
	"testing"
)

func ExampleGeometricMean() {
	mean := GeometricMean(1, 1000)
	fmt.Printf("The geometric mean of [1 1000] is %0.3f.\n", mean)

	// Output:
	// The geometric mean of [1 1000] is 31.623.
}

func TestGeometricMean(t *testing.T) {
	cases := []testArrayCase{
		{
			input:    []uint{},
			expected: math.NaN(),
		},
		{
			input:    []uint{5, 15},
			expected: 8.660254037844387,
		},
		{
			input:    []int{0},
			expected: 0,
		},
		{
			input:    []int{-5, 5},
			expected: math.NaN(),
		},
		{
			input:    []float32{2.5, 4.5},
			expected: 3.3541019662496847,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("test %v", c), func(t *testing.T) {
			var actual float64
			switch c.input.(type) {
			case []uint:
				actual = GeometricMean(c.input.([]uint)...)
			case []int:
				actual = GeometricMean(c.input.([]int)...)
			case []int32:
				actual = GeometricMean(c.input.([]int32)...)
			case []float32:
				actual = GeometricMean(c.input.([]float32)...)
			case []float64:
				actual = GeometricMean(c.input.([]float64)...)
			}
			if math.IsNaN(c.expected) {
				if !math.IsNaN(actual) {
					t.Logf("expected %v, got %v", c.expected, actual)
					t.FailNow()
				}
			} else {
				if actual != c.expected {
					t.Logf("expected %0.3f, got %v", c.expected, actual)
					t.FailNow()
				}
			}
		})
	}
}

func BenchmarkGeometricMean(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GeometricMean(i, i-5, i+3, i-8)
	}
}
