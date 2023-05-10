package numbertheory

import (
	"fmt"
	"testing"
)

func ExampleExtendedGCD() {
	gcd, x, y := ExtendedGCD(48, 18)
	fmt.Printf("The extended GCD of 18 and 48 is (%v, %v, %v).", gcd, x, y)

	// Output:
	// The extended GCD of 18 and 48 is (6, -1, 3).
}

func TestExtendedGCD(t *testing.T) {
	cases := []struct {
		name     string
		a        int
		b        int
		expected string
	}{
		{
			name:     "2,1",
			a:        2,
			b:        1,
			expected: "1 0 1",
		},
		{
			name:     "1,2",
			a:        1,
			b:        2,
			expected: "1 1 0",
		},
		{
			name:     "3,5",
			a:        3,
			b:        5,
			expected: "1 2 -1",
		},
		{
			name:     "5,3",
			a:        5,
			b:        3,
			expected: "1 -1 2",
		},
		{
			name:     "48,18",
			a:        48,
			b:        18,
			expected: "6 -1 3",
		},
		{
			name:     "18,48",
			a:        18,
			b:        48,
			expected: "6 3 -1",
		},
		{
			name:     "-18,48",
			a:        -18,
			b:        48,
			expected: "-6 3 1",
		},
		{
			name:     "11,13",
			a:        11,
			b:        13,
			expected: "1 6 -5",
		},
		{
			name:     "0,11",
			a:        0,
			b:        11,
			expected: "11 0 1",
		},
		{
			name:     "1,1",
			a:        1,
			b:        1,
			expected: "1 1 0",
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			gcd, x, y := ExtendedGCD(c.a, c.b)
			actual := fmt.Sprintf("%v %v %v", gcd, x, y)
			if actual != c.expected {
				t.Logf("expected %v, got %v\n", c.expected, actual)
				t.FailNow()
			}
		})
	}
}

func BenchmarkExtendedGCD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ExtendedGCD(i, i+30)
	}
}
