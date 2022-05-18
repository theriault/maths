package statistics

import (
	"fmt"
	"testing"
)

func ExampleQuantile() {
	vals := []float64{3, 6, 7, 8, 8, 10, 13, 15, 16, 20}
	quantiles, _ := Quantile(vals, 4)
	fmt.Printf("The quantiles of %v are %v.\n", vals, quantiles)

	// Output:
	// The quantiles of [3 6 7 8 8 10 13 15 16 20] are [7.25 9 14.5].
}

func TestQuantile(t *testing.T) {
	cases := []struct {
		name      string
		quantiles int
		input     []float64
		expected  string
		err       string
	}{
		{
			name:      "test with no numbers",
			quantiles: 4,
			input:     []float64{},
			expected:  "[]",
			err:       "insufficient data to produce 4-iles - need at least 1 number",
		},
		{
			name:      "test with no quantile",
			quantiles: 0,
			input:     []float64{},
			expected:  "[]",
			err:       "q must be >= 1: 0",
		},
		{
			name:      "test quartile with odd number of items",
			quantiles: 4,
			input:     []float64{1, 1, 9, 10},
			expected:  "[1 5 9.25]",
			err:       "<nil>",
		},
		{
			name:      "test quartile with odd number of items",
			quantiles: 4,
			input:     []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected:  "[3 5 7]",
			err:       "<nil>",
		},
		{
			name:      "test quintile with odd number of items",
			quantiles: 5,
			input:     []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected:  "[2.6 4.2 5.8 7.4]",
			err:       "<nil>",
		},
		{
			name:      "test median with even number of items",
			quantiles: 2,
			input:     []float64{3, 6},
			expected:  "[4.5]",
			err:       "<nil>",
		},
		{
			name:      "test median with odd number of items",
			quantiles: 2,
			input:     []float64{3, 6, 10},
			expected:  "[6]",
			err:       "<nil>",
		},
		{
			name:      "test quartiles with less than four items",
			quantiles: 4,
			input:     []float64{3, 6, 7},
			expected:  "[4.5 6 6.5]",
			err:       "<nil>",
		},
		{
			name:      "test quartiles with exactly 4 items",
			quantiles: 4,
			input:     []float64{3, 6, 7, 8},
			expected:  "[5.25 6.5 7.25]",
			err:       "<nil>",
		},
		{
			name:      "test quartiles with even number of items",
			quantiles: 4,
			input:     []float64{3, 6, 7, 8, 8, 10, 13, 15, 16, 20},
			expected:  "[7.25 9 14.5]",
			err:       "<nil>",
		},
		{
			name:      "test quartiles with odd number of items",
			quantiles: 4,
			input:     []float64{3, 6, 7, 8, 8, 9, 10, 13, 15, 16, 20},
			expected:  "[7.5 9 14]",
			err:       "<nil>",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual, err := Quantile(c.input, c.quantiles)
			if fmt.Sprintf("%v", actual) != c.expected {
				t.Logf("expected %v, got %v", c.expected, actual)
				t.FailNow()
			}
			if fmt.Sprintf("%v", err) != c.err {
				t.Logf("expected error %v, got %v", c.err, err)
				t.FailNow()
			}
		})
	}
}

func BenchmarkQuantile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Quantile([]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 4)
	}
}
