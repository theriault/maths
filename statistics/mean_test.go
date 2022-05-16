package statistics

import (
	"fmt"
	"math"
	"testing"
)

func ExampleMean() {
	mean := Mean(1, 1000)
	fmt.Printf("The mean of [1 1000] is %0.1f.\n", mean)

	// Output:
	// The mean of [1 1000] is 500.5.
}

func TestMean(t *testing.T) {
	cases := []testArrayCase{
		{
			input:    []uint{},
			expected: math.NaN(),
		},
		{
			input:    []uint{5, 15},
			expected: 10,
		},
		{
			input:    []int{-5, 5},
			expected: 0,
		},
		{
			input:    []float32{2.5, 4.5},
			expected: 3.5,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("test %v", c.input), func(t *testing.T) {
			t.Parallel()
			var actual float64
			switch c.input.(type) {
			case []uint:
				actual = Mean(c.input.([]uint)...)
			case []int:
				actual = Mean(c.input.([]int)...)
			case []float32:
				actual = Mean(c.input.([]float32)...)
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

func BenchmarkMean(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mean(i, i-5, i+3, i-8)
	}
}
