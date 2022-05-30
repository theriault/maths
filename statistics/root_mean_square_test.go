package statistics

import (
	"fmt"
	"testing"
)

func ExampleRootMeanSquare() {
	vals := []uint8{8, 7, 3, 2, 6, 11, 6, 7, 2, 1, 7}
	max := RootMeanSquare(vals...)
	fmt.Printf("The root mean square of %v is %.8f.\n", vals, max)

	// Output:
	// The root mean square of [8 7 3 2 6 11 6 7 2 1 7] is 6.19383858.
}

func TestRootMeanSquare(t *testing.T) {
	cases := []testArrayCase{
		{
			name:     "test 1",
			input:    []float64{8, 7, 3, 2, 6, 11, 6, 7, 2, 1, 7},
			expected: "float64 6.193838580689391",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			var actual interface{}
			switch c.input.(type) {
			case []uint:
				actual = RootMeanSquare(c.input.([]uint)...)
			case []uint64:
				actual = RootMeanSquare(c.input.([]uint64)...)
			case []int:
				actual = RootMeanSquare(c.input.([]int)...)
			case []float64:
				actual = RootMeanSquare(c.input.([]float64)...)
			}
			if fmt.Sprintf("%T %v", actual, actual) != c.expected {
				t.Logf("expected %v, got %v", c.expected, actual)
				t.FailNow()
			}
		})
	}
}
