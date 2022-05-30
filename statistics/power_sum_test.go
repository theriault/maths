package statistics

import (
	"fmt"
	"testing"
)

func ExamplePowerSum() {
	sum := PowerSum([]float64{2, 3, 4}, 2)
	fmt.Printf("The power sum of [2 3 4] to the 2nd power is %0.0f.\n", sum)

	// Output:
	// The power sum of [2 3 4] to the 2nd power is 29.
}

func TestPowerSum(t *testing.T) {
	cases := []struct {
		name     string
		input    any
		expected string
		exponent float64
	}{
		{
			name:     "test 1",
			input:    []float64{2, 3, 4},
			expected: "float64 29.00000",
			exponent: 2,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			var actual float64
			switch c.input.(type) {
			case []uint:
				actual = PowerSum(c.input.([]uint), c.exponent)
			case []int:
				actual = PowerSum(c.input.([]int), c.exponent)
			case []float64:
				actual = PowerSum(c.input.([]float64), c.exponent)
			}
			r := fmt.Sprintf("%T %0.5f", actual, actual)
			if r != c.expected {
				t.Logf("expected %v, got %v", c.expected, r)
				t.FailNow()
			}
		})
	}
}

func BenchmarkPowerSum(b *testing.B) {
	X := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	for i := 0; i < b.N; i++ {
		PowerSum(X, 2)
	}
}
