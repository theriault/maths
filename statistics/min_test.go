package statistics

import (
	"fmt"
	"math"
	"testing"
)

func ExampleMin() {
	vals := []uint8{8, 7, 3, 2, 6, 11, 6, 7, 2, 1, 7}
	min := Min(vals...)
	fmt.Printf("The min of %v is %v.\n", vals, min)

	// Output:
	// The min of [8 7 3 2 6 11 6 7 2 1 7] is 1.
}

func TestMin(t *testing.T) {
	cases := []testArrayCase{
		{
			name:     "no numbers",
			input:    []uint{},
			expected: "0",
		},
		{
			name:     "negatives",
			input:    []int{-10, -20},
			expected: "-20",
		},
		{
			name:     "mixed",
			input:    []int{-10, -20, 5, 15},
			expected: "-20",
		},
		{
			name:     "float +INF start",
			input:    []float64{math.Inf(1), 5, 10},
			expected: "5",
		},
		{
			name:     "float -INF end",
			input:    []float64{math.Inf(-1), 5, 10},
			expected: "-Inf",
		},
		{
			name:     "float NaN end",
			input:    []float64{5, 10, math.NaN()},
			expected: "NaN",
		},
		{
			name:     "float NaN start",
			input:    []float64{math.NaN(), 5, 10},
			expected: "NaN",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			var actual interface{}
			switch c.input.(type) {
			case []uint:
				actual = Min(c.input.([]uint)...)
			case []uint64:
				actual = Min(c.input.([]uint64)...)
			case []int:
				actual = Min(c.input.([]int)...)
			case []float64:
				actual = Min(c.input.([]float64)...)
			}
			if fmt.Sprintf("%v", actual) != c.expected {
				t.Logf("expected %v, got %v", c.expected, actual)
				t.FailNow()
			}
		})
	}
}

var testMinVals []uint = make([]uint, 1000)

func BenchmarkMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Min(testMinVals...)
	}
}
