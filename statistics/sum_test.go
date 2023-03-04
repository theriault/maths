package statistics

import (
	"fmt"
	"testing"
)

func ExampleSum() {
	sum := Sum(1.1, 1.2, 1.3)
	fmt.Printf("The sum of [1.1 1.2 1.3] is %0.1f.\n", sum)

	// Output:
	// The sum of [1.1 1.2 1.3] is 3.6.
}

func TestSum(t *testing.T) {
	cases := []testArrayCase{
		{
			name:     "no input",
			input:    []uint{},
			expected: "uint 0",
		},
		{
			name:     "uints",
			input:    []uint{5, 15},
			expected: "uint 20",
		},
		{
			name:     "uint8 overflow",
			input:    []uint8{245, 10, 2},
			expected: "uint8 1",
		},
		{
			name:     "float32s",
			input:    []float32{1.1, 1.2, 1.3},
			expected: "float32 3.6000001",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			var actual any
			switch c.input.(type) {
			case []uint8:
				actual = Sum(c.input.([]uint8)...)
			case []uint:
				actual = Sum(c.input.([]uint)...)
			case []int:
				actual = Sum(c.input.([]int)...)
			case []float32:
				actual = Sum(c.input.([]float32)...)
			}
			r := fmt.Sprintf("%T %v", actual, actual)
			if r != c.expected {
				t.Logf("expected %v, got %v", c.expected, r)
				t.FailNow()
			}
		})
	}
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)
	}
}
