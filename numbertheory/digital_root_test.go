package numbertheory

import (
	"fmt"
	"testing"
)

func ExampleDigitalRoot() {
	n := DigitalRoot(9045, 10)
	fmt.Printf("The digital root of 9045 in base-10 is %v.\n", n)

	// Output:
	// The digital root of 9045 in base-10 is 9.
}

func TestDigitalRoot(t *testing.T) {
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
			expected: 1,
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
			name:     "0b111",
			n:        0b111,
			b:        2,
			expected: 1,
		},
		{
			name:     "9045 in base 10",
			n:        9045,
			b:        10,
			expected: 9,
		},
		{
			name:     "9045 in base 2",
			n:        9045,
			b:        2,
			expected: 1,
		},
		{
			name:     "9045 in base 31",
			n:        9045,
			b:        31,
			expected: 15,
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
			actual := DigitalRoot(c.n, c.b)
			if c.expected != actual {
				t.Logf("expected %v, got %v", c.expected, actual)
				t.FailNow()
			}
		})
	}
}

func BenchmarkDigitalRoot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DigitalRoot(i, 10)
	}
}
