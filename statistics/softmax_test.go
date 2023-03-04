package statistics

import (
	"fmt"
	"testing"
)

func ExampleSoftmax() {
	X := []uint8{8, 3, 6, 2, 7, 1, 8, 3, 7, 4, 8}
	softmax := Softmax(X)
	fmt.Printf("The softmax of %v is %v.\n", X, softmax)

	// Output:
	// The softmax of [8 3 6 2 7 1 8 3 7 4 8] is [0.2559982742777946 0.0017249028039411246 0.034645598957469395 0.0006345562795889148 0.0941765020821687 0.0002334402095269995 0.2559982742777946 0.0017249028039411246 0.0941765020821687 0.004688771947811214 0.2559982742777946].
}

func TestSoftmax(t *testing.T) {
	cases := []struct {
		name     string
		input    []float64
		expected string
	}{
		{
			name:     "test 1",
			input:    []float64{8, 3, 6, 2, 7, 1, 8, 3, 7, 4, 8},
			expected: "[]float64 [0.2559983 0.0017249 0.0346456 0.0006346 0.0941765 0.0002334 0.2559983 0.0017249 0.0941765 0.0046888 0.2559983]",
		},
		{
			name:     "test 2",
			input:    []float64{1, 2, 3, 4, 1, 2, 3},
			expected: "[]float64 [0.0236405 0.0642617 0.1746813 0.4748330 0.0236405 0.0642617 0.1746813]",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual := Softmax(c.input)
			r := fmt.Sprintf("%T %.7f", actual, actual)
			if r != c.expected {
				t.Logf("expected %v, got %v", c.expected, r)
				t.FailNow()
			}
		})
	}
}

func TestGeneralizedSoftmax(t *testing.T) {
	cases := []struct {
		name        string
		input       []float64
		temperature float64
		expected    string
	}{
		{
			name:        "test half temperature",
			input:       []float64{8, 3, 6, 2, 7, 1, 8, 3, 7, 4, 8},
			temperature: 0.5,
			expected:    "[]float64 [0.3040050 0.0000138 0.0055680 0.0000019 0.0411426 0.0000003 0.3040050 0.0000138 0.0411426 0.0001020 0.3040050]",
		},
		{
			name:        "test double temperature",
			input:       []float64{8, 3, 6, 2, 7, 1, 8, 3, 7, 4, 8},
			temperature: 2,
			expected:    "[]float64 [0.2015954 0.0165480 0.0741628 0.0100368 0.1222738 0.0060877 0.2015954 0.0165480 0.1222738 0.0272830 0.2015954]",
		},
		{
			name:        "test zero temperature",
			input:       []float64{8, 3, 6, 2, 7, 1, 8, 3, 7, 4, 8},
			temperature: 0,
			expected:    "[]float64 [NaN NaN NaN NaN NaN NaN NaN NaN NaN NaN NaN]",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual := GeneralizedSoftmax(c.input, c.temperature)
			r := fmt.Sprintf("%T %.7f", actual, actual)
			if r != c.expected {
				t.Logf("expected %v, got %v", c.expected, r)
				t.FailNow()
			}
		})
	}
}
