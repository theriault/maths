package statistics

import (
	"fmt"
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
			name:     "no numbers",
			input:    []uint{},
			expected: "NaN",
		},
		{
			name:     "negatives",
			input:    []int{-10, -20},
			expected: "-13.333333333333332",
		},
		{
			name:     "two numbers",
			input:    []uint{5, 15},
			expected: "7.5",
		},
		{
			name:     "zero",
			input:    []int{-5, 5},
			expected: "NaN",
		},
		{
			name:     "floats",
			input:    []float32{2.5, 4.5},
			expected: "3.2142857142857144",
		},
		{
			name:     "large numbers",
			input:    []uint64{1, 4294967295, 18446744073709551615},
			expected: "2.999999999301508",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			var actual float64
			switch c.input.(type) {
			case []uint:
				actual = HarmonicMean(c.input.([]uint)...)
			case []uint64:
				actual = HarmonicMean(c.input.([]uint64)...)
			case []int:
				actual = HarmonicMean(c.input.([]int)...)
			case []float32:
				actual = HarmonicMean(c.input.([]float32)...)
			}
			if fmt.Sprintf("%v", actual) != c.expected {
				t.Logf("expected %v, got %v", c.expected, actual)
				t.FailNow()
			}
		})
	}
}

func BenchmarkHarmonicMean(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HarmonicMean(i, i-5, i+3, i-8)
	}
}
