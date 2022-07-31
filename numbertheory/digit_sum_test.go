package numbertheory

import (
	"fmt"
	"testing"
)

func ExampleDigitSum() {
	n := DigitSum(9045, 10)
	fmt.Printf("The digit sum of 9045 in base-10 is %v.\n", n)

	// Output:
	// The digit sum of 9045 in base-10 is 18.
}

func TestDigitSum(t *testing.T) {
	cases := []struct {
		name     string
		n        int
		b        int
		expected int
	}{
		{
			name:     "9045 in base 1",
			n:        9045,
			b:        1,
			expected: 9045,
		},
		{
			name:     "9045 in base 0 (error)",
			n:        9045,
			b:        0,
			expected: 0,
		},
		{
			name:     "-9045 in base 10",
			n:        -9045,
			b:        0,
			expected: 0,
		},
		{
			name:     "9045 in base 10",
			n:        9045,
			b:        10,
			expected: 18,
		},
		{
			name:     "9045 in base 2",
			n:        9045,
			b:        2,
			expected: 7,
		},
		{
			name:     "9045 in base 31",
			n:        9045,
			b:        31,
			expected: 45,
		},
		{
			name:     "9045 in base 9045",
			n:        9045,
			b:        9044,
			expected: 2,
		},
		{
			name:     "9045 in base 9045",
			n:        9045,
			b:        9045,
			expected: 1,
		},
		{
			name:     "9045 in base 9046",
			n:        9045,
			b:        9046,
			expected: 9045,
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual := DigitSum(c.n, c.b)
			if c.expected != actual {
				t.Logf("expected %v, got %v", c.expected, actual)
				t.FailNow()
			}
		})
	}
}

func BenchmarkDigitSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DigitSum(i, 10)
	}
}
