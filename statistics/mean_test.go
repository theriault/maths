package statistics

import (
	"fmt"
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
			name:     "no numbers",
			input:    []uint{},
			expected: "NaN",
		},
		{
			name:     "two numbers",
			input:    []uint{5, 15},
			expected: "10",
		},
		{
			name:     "zero",
			input:    []int{-5, 5},
			expected: "0",
		},
		{
			name:     "floats",
			input:    []float32{2.5, 4.5},
			expected: "3.5",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
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
			if fmt.Sprintf("%v", actual) != c.expected {
				t.Logf("expected %v, got %v", c.expected, actual)
				t.FailNow()
			}
		})
	}
}

var testMeanVals []uint = make([]uint, 1000)

func BenchmarkMean(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mean(testMeanVals...)
	}
}
