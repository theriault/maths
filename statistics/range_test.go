package statistics

import (
	"fmt"
	"math"
	"testing"
)

func ExampleRange() {
	vals := []uint8{8, 7, 3, 2, 6, 11, 6, 7, 2, 1, 7}
	max := Range(vals...)
	fmt.Printf("The range of %v is %v.\n", vals, max)

	// Output:
	// The range of [8 7 3 2 6 11 6 7 2 1 7] is 10.
}

func TestRange(t *testing.T) {
	cases := []testArrayCase{
		{
			name:     "no numbers",
			input:    []uint{},
			expected: "0",
		},
		{
			name:     "negatives",
			input:    []int{-10, -20},
			expected: "10",
		},
		{
			name:     "mixed",
			input:    []int{-10, -20, 5, 15},
			expected: "35",
		},
		{
			name:     "float +INF start",
			input:    []float64{math.Inf(1), 5, 10},
			expected: "+Inf",
		},
		{
			name:     "float -INF start",
			input:    []float64{math.Inf(-1), 5, 10},
			expected: "+Inf",
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
				actual = Range(c.input.([]uint)...)
			case []uint64:
				actual = Range(c.input.([]uint64)...)
			case []int:
				actual = Range(c.input.([]int)...)
			case []float64:
				actual = Range(c.input.([]float64)...)
			}
			if fmt.Sprintf("%v", actual) != c.expected {
				t.Logf("expected %v, got %v", c.expected, actual)
				t.FailNow()
			}
		})
	}
}

func BenchmarkRange(b *testing.B) {
	var vals []uint = make([]uint, 10000)
	for i := 0; i < b.N; i++ {
		vals[0] = uint(i)
		Range(vals...)
	}
}
