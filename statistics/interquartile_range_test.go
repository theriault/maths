package statistics

import (
	"fmt"
	"testing"
)

func ExampleInterquartileRange() {
	vals := []float64{3, 6, 7, 8, 8, 10, 13, 15, 16, 20}
	r, _ := InterquartileRange(vals...)
	fmt.Printf("The interquartile range of %v is %v.\n", vals, r)

	// Output:
	// The interquartile range of [3 6 7 8 8 10 13 15 16 20] is 7.25.
}

func TestInterquartileRange(t *testing.T) {
	cases := []struct {
		name     string
		input    []float64
		expected string
		err      string
	}{
		{
			name:     "test IQR with odd number of items",
			input:    []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: "4 <nil>",
		},
		{
			name:     "test IQR with even number of items",
			input:    []float64{3, 6, 7, 8, 8, 10, 13, 15, 16, 20},
			expected: "7.25 <nil>",
		},
		{
			name:     "test IQR with no elements",
			input:    []float64{},
			expected: "0 insufficient data to produce 4-iles - need at least 1 number",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual, err := InterquartileRange(c.input...)
			if actual := fmt.Sprintf("%v %v", actual, err); actual != c.expected {
				t.Logf("expected %v, got %v", c.expected, actual)
				t.FailNow()
			}
		})
	}
}
