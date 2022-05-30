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
			expected:  "[] insufficient data to produce 4-iles - need at least 1 number",
		},
		{
			name:      "test with no quantile",
			quantiles: 0,
			input:     []float64{},
			expected:  "[] q must be >= 1: 0",
		},
		{
			name:      "test quartile with odd number of items",
			quantiles: 4,
			input:     []float64{1, 1, 9, 10},
			expected:  "[1 5 9.25] <nil>",
		},
		{
			name:      "test quartile with odd number of items",
			quantiles: 4,
			input:     []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected:  "[3 5 7] <nil>",
		},
		{
			name:      "test quintile with odd number of items",
			quantiles: 5,
			input:     []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected:  "[2.6 4.2 5.8 7.4] <nil>",
		},
		{
			name:      "test median with even number of items",
			quantiles: 2,
			input:     []float64{3, 6},
			expected:  "[4.5] <nil>",
		},
		{
			name:      "test median with odd number of items",
			quantiles: 2,
			input:     []float64{3, 6, 10},
			expected:  "[6] <nil>",
		},
		{
			name:      "test quartiles with less than four items",
			quantiles: 4,
			input:     []float64{3, 6, 7},
			expected:  "[4.5 6 6.5] <nil>",
		},
		{
			name:      "test quartiles with exactly 4 items",
			quantiles: 4,
			input:     []float64{3, 6, 7, 8},
			expected:  "[5.25 6.5 7.25] <nil>",
		},
		{
			name:      "test quartiles with even number of items",
			quantiles: 4,
			input:     []float64{3, 6, 7, 8, 8, 10, 13, 15, 16, 20},
			expected:  "[7.25 9 14.5] <nil>",
		},
		{
			name:      "test quartiles with odd number of items",
			quantiles: 4,
			input:     []float64{3, 6, 7, 8, 8, 9, 10, 13, 15, 16, 20},
			expected:  "[7.5 9 14] <nil>",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual, err := Quantile(c.input, c.quantiles)
			if actual := fmt.Sprintf("%v %v", actual, err); actual != c.expected {
				t.Logf("expected %v, got %v", c.expected, actual)
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

func TestTertile(t *testing.T) {
	cases := []struct {
		name     string
		input    []float64
		expected string
		err      string
	}{
		{
			name:     "test tertile with odd number of items",
			input:    []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: "[3.6666666666666665 6.333333333333333] <nil>",
		},
		{
			name:     "test tertiles with even number of items",
			input:    []float64{3, 6, 7, 8, 8, 10, 13, 15, 16, 20},
			expected: "[8 13] <nil>",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual, err := Tertile(c.input)
			if actual := fmt.Sprintf("%v %v", actual, err); actual != c.expected {
				t.Logf("expected %v, got %v", c.expected, actual)
				t.FailNow()
			}
		})
	}
}

func TestQuartile(t *testing.T) {
	cases := []struct {
		name     string
		input    []float64
		expected string
		err      string
	}{
		{
			name:     "test quartiles with odd number of items",
			input:    []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: "[3 5 7] <nil>",
		},
		{
			name:     "test quartiles with even number of items",
			input:    []float64{3, 6, 7, 8, 8, 10, 13, 15, 16, 20},
			expected: "[7.25 9 14.5] <nil>",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual, err := Quartile(c.input)
			if actual := fmt.Sprintf("%v %v", actual, err); actual != c.expected {
				t.Logf("expected %v, got %v", c.expected, actual)
				t.FailNow()
			}
		})
	}
}

func TestPercentile(t *testing.T) {
	cases := []struct {
		name     string
		input    []float64
		expected string
		err      string
	}{
		{
			name:     "test percentiles with odd number of items",
			input:    []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: "[1.08 1.16 1.24 1.32 1.4 1.48 1.56 1.6400000000000001 1.72 1.8 1.88 1.96 2.04 2.12 2.2 2.2800000000000002 2.3600000000000003 2.44 2.52 2.6 2.6799999999999997 2.76 2.84 2.92 3 3.08 3.16 3.24 3.32 3.4 3.48 3.56 3.64 3.72 3.8 3.88 3.96 4.04 4.12 4.2 4.279999999999999 4.359999999999999 4.4399999999999995 4.52 4.6 4.68 4.76 4.84 4.92 5 5.08 5.16 5.24 5.32 5.4 5.48 5.56 5.64 5.72 5.8 5.88 5.96 6.04 6.12 6.2 6.28 6.36 6.44 6.52 6.6 6.68 6.76 6.84 6.92 7 7.08 7.16 7.24 7.32 7.4 7.48 7.56 7.64 7.72 7.8 7.88 7.96 8.04 8.120000000000001 8.2 8.280000000000001 8.36 8.440000000000001 8.52 8.6 8.68 8.76 8.84 8.92] <nil>",
		},
		{
			name:     "test percentiles with even number of items",
			input:    []float64{3, 6, 7, 8, 8, 10, 13, 15, 16, 20},
			expected: "[3.27 3.54 3.81 4.08 4.35 4.62 4.890000000000001 5.16 5.43 5.7 5.97 6.08 6.17 6.26 6.35 6.4399999999999995 6.53 6.62 6.71 6.8 6.89 6.98 7.07 7.16 7.25 7.34 7.43 7.5200000000000005 7.609999999999999 7.699999999999999 7.79 7.88 7.970000000000001 8 8 8 8 8 8 8 8 8 8 8 8.1 8.280000000000001 8.459999999999999 8.64 8.82 9 9.18 9.36 9.540000000000001 9.72 9.9 10.120000000000003 10.39 10.66 10.93 11.2 11.47 11.74 12.01 12.28 12.55 12.82 13.06 13.24 13.419999999999998 13.6 13.78 13.959999999999999 14.14 14.32 14.5 14.68 14.86 15.02 15.11 15.2 15.290000000000001 15.379999999999999 15.469999999999999 15.559999999999999 15.649999999999999 15.74 15.83 15.92 16.04 16.4 16.759999999999998 17.120000000000005 17.480000000000004 17.839999999999996 18.199999999999996 18.560000000000002 18.92 19.28 19.64] <nil>",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual, err := Percentile(c.input)
			if actual := fmt.Sprintf("%v %v", actual, err); actual != c.expected {
				t.Logf("expected %v, got %v", c.expected, actual)
				t.FailNow()
			}
		})
	}
}
