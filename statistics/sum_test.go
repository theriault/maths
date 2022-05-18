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
			expected: "float64 0.00000",
		},
		{
			name:     "uints",
			input:    []uint{5, 15},
			expected: "float64 20.00000",
		},
		{
			name:     "float32s",
			input:    []float32{1.1, 1.2, 1.3},
			expected: "float64 3.60000",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			var actual float64
			switch c.input.(type) {
			case []uint:
				actual = Sum(c.input.([]uint)...)
			case []int:
				actual = Sum(c.input.([]int)...)
			case []float32:
				actual = Sum(c.input.([]float32)...)
			}
			r := fmt.Sprintf("%T %0.5f", actual, actual)
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

/*
func testFloatEqual(actual float64, expected float64) bool {
	return actual == expected || actual > math.Nextafter(expected, expected+1) || actual < math.Nextafter(expected, expected-1)
}
*/
