package statistics

import (
	"fmt"
	"testing"
)

func ExampleMedian() {
	median := Median(7, 5, 1, 4, 3, 6, 2)
	fmt.Printf("The median of [7 5 1 4 3 6 2] is %v.\n", median)

	// Output:
	// The median of [7 5 1 4 3 6 2] is 4.
}

func TestMedian(t *testing.T) {
	cases := []testArrayCase{
		{
			name:     "no input",
			input:    []uint{},
			expected: "float64 NaN",
		},
		{
			name:     "uints",
			input:    []uint{5, 15},
			expected: "float64 10",
		},
		{
			name:     "float32s",
			input:    []float32{5, 15, 7.5, 11},
			expected: "float64 9.25",
		},
		{
			name:     "ints",
			input:    []int{2, 1, 3},
			expected: "float64 2",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			var actual float64
			switch c.input.(type) {
			case []uint:
				actual = Median(c.input.([]uint)...)
			case []int:
				actual = Median(c.input.([]int)...)
			case []float32:
				actual = Median(c.input.([]float32)...)
			}
			r := fmt.Sprintf("%T %v", actual, actual)
			if r != c.expected {
				t.Logf("expected %v, got %v", c.expected, r)
				t.FailNow()
			}
		})
	}
}

func BenchmarkMedian(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Median(1, -5, 8, 3, -13)
	}
}
