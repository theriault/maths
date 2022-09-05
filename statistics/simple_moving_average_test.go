package statistics

import (
	"fmt"
	"testing"
)

func ExampleSimpleMovingAverage() {
	sma := SimpleMovingAverage(3, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Printf("The simple moving average over the last 3 datapoints of [1 2 3 4 5 6 7 8 9] is %v.\n", sma)

	// Output:
	// The simple moving average over the last 3 datapoints of [1 2 3 4 5 6 7 8 9] is [2 3 4 5 6 7 8].
}

func TestSimpleMovingAverage(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		k        int
		expected string
	}{
		{
			name:     "+1 repeating",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			k:        3,
			expected: "[2 3 4 5 6 7 8]",
		},
		{
			name:     "+1 +1 -2 repeating",
			input:    []int{1, 2, 3, 1, 2, 3, 1, 2, 3},
			k:        3,
			expected: "[2 2 2 2 2 2 2]",
		},
		{
			name:     "+4 -2 repeating",
			input:    []int{1, 5, 3, 7, 5, 9, 7, 11},
			k:        3,
			expected: "[3 5 5 7 7 9]",
		},
		{
			name:     "k larger than n",
			input:    []int{1, 2, 3, 4},
			k:        10,
			expected: "[2.5]",
		},
		{
			name:     "k < 1",
			input:    []int{1, 2, 3, 4},
			k:        -1,
			expected: "[]",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual := SimpleMovingAverage(c.k, c.input...)
			if fmt.Sprintf("%v", actual) != c.expected {
				t.Logf("expected %v, got %v", c.expected, actual)
				t.FailNow()
			}
		})
	}
}

func BenchmarkSimpleMovingAverage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SimpleMovingAverage(3, 3, 6, 2, 6, 2, 6, 7, 2, 5, 4, 2, 4, 6)
	}
}
