package statistics

import (
	"fmt"
	"testing"
)

func ExampleMode() {
	mode := Mode(8, 3, 6, 2, 7, 1, 8, 3, 7, 4, 8)
	fmt.Printf("The mode of [8, 3, 6, 2, 7, 1, 8, 3, 7, 4, 8] is %v.\n", mode)

	// Output:
	// The mode of [8, 3, 6, 2, 7, 1, 8, 3, 7, 4, 8] is [8].
}

func TestMode(t *testing.T) {
	cases := []testArrayCase{
		{
			name:     "no values",
			input:    []uint{},
			expected: "[]",
		},
		{
			name:     "multimodal",
			input:    []uint{5, 15},
			expected: "[5 15]",
		},
		{
			name:     "mode at end",
			input:    []uint{1, 1, 1, 2, 2, 3, 3, 3, 3},
			expected: "[3]",
		},
		{
			name:     "mode at start",
			input:    []uint{1, 1, 1, 1, 2, 2, 3, 3, 3},
			expected: "[1]",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			var actual []float64
			switch c.input.(type) {
			case []uint:
				actual = Mode(c.input.([]uint)...)
			case []int:
				actual = Mode(c.input.([]int)...)
			case []float32:
				actual = Mode(c.input.([]float32)...)
			}
			actualStr := fmt.Sprintf("%v", actual)
			if actualStr != c.expected {
				t.Logf("expected %v, got %v", c.expected, actualStr)
				t.FailNow()
			}
		})
	}
}

func BenchmarkMode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mode(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 1)
	}
}
