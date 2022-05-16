package statistics

import (
	"fmt"
	"math"
	"testing"
)

func ExampleHarmonicMean() {
	mean := HarmonicMean(1, 1000)
	fmt.Printf("The mean of [1 1000] is %0.3f.\n", mean)

	// Output:
	// The mean of [1 1000] is 1.998.
}

func TestHarmonicMean(t *testing.T) {
	cases := []testArrayCase{
		{
			input:    []uint{},
			expected: math.NaN(),
		},
		{
			input:    []int{-10, -20},
			expected: math.NaN(),
		},
		{
			input:    []uint{5, 15},
			expected: 7.5,
		},
		{
			input:    []int{-5, 5},
			expected: math.NaN(),
		},
		{
			input:    []float32{2.5, 4.5},
			expected: 3.2142857142857144,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("test %v", c.input), func(t *testing.T) {
			t.Parallel()
			var actual float64
			switch c.input.(type) {
			case []uint:
				actual = HarmonicMean(c.input.([]uint)...)
			case []int:
				actual = HarmonicMean(c.input.([]int)...)
			case []float32:
				actual = HarmonicMean(c.input.([]float32)...)
			}
			if math.IsNaN(c.expected) {
				if !math.IsNaN(actual) {
					t.Logf("expected %v, got %v", c.expected, actual)
					t.FailNow()
				}
			} else {
				if actual != c.expected {
					t.Logf("expected %v, got %v", c.expected, actual)
					t.FailNow()
				}
			}
		})
	}
}

func BenchmarkHarmonicMean(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HarmonicMean(i, i-5, i+3, i-8)
	}
}
