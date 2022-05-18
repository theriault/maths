package statistics

import (
	"fmt"
	"math"
	"testing"
)

func ExampleMax() {
	vals := []uint8{8, 7, 3, 2, 6, 11, 6, 7, 2, 1, 7}
	max := Max(vals...)
	fmt.Printf("The max of %v is %v.\n", vals, max)

	// Output:
	// The max of [8 7 3 2 6 11 6 7 2 1 7] is 11.
}

func TestMax(t *testing.T) {
	cases := []testArrayCase{
		{
			name:     "no numbers",
			input:    []uint{},
			expected: "0",
		},
		{
			name:     "negatives",
			input:    []int{-10, -20},
			expected: "-10",
		},
		{
			name:     "mixed",
			input:    []int{-10, -20, 5, 15},
			expected: "15",
		},
		{
			name:     "float +INF start",
			input:    []float64{math.Inf(1), 5, 10},
			expected: "+Inf",
		},
		{
			name:     "float -INF end",
			input:    []float64{math.Inf(-1), 5, 10},
			expected: "10",
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
				actual = Max(c.input.([]uint)...)
			case []uint64:
				actual = Max(c.input.([]uint64)...)
			case []int:
				actual = Max(c.input.([]int)...)
			case []float64:
				actual = Max(c.input.([]float64)...)
			}
			if fmt.Sprintf("%v", actual) != c.expected {
				t.Logf("expected %v, got %v", c.expected, actual)
				t.FailNow()
			}
		})
	}
}

var testMaxVals []uint = make([]uint, 1000)

func BenchmarkMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Max(testMaxVals...)
	}
}
