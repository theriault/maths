package numbertheory

import (
	"fmt"
	"testing"
)

func ExampleGCD() {
	gcd := GCD(48, 18)
	fmt.Printf("The GCD of 18 and 48 is %v.", gcd)

	// Output:
	// The GCD of 18 and 48 is 6.
}

func TestGCD(t *testing.T) {
	cases := []struct {
		name     string
		a        int
		b        int
		expected string
	}{
		{
			name:     "48,18",
			a:        48,
			b:        18,
			expected: "6",
		},
		{
			name:     "18,48",
			a:        18,
			b:        48,
			expected: "6",
		},
		{
			name:     "-18,48",
			a:        -18,
			b:        48,
			expected: "-6",
		},
		{
			name:     "11,13",
			a:        11,
			b:        13,
			expected: "1",
		},
		{
			name:     "0,11",
			a:        0,
			b:        11,
			expected: "11",
		},
		{
			name:     "0,0",
			a:        0,
			b:        0,
			expected: "1",
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual := GCD(c.a, c.b)
			if fmt.Sprintf("%v", actual) != c.expected {
				t.Logf("expected %v, got %v\n", c.expected, actual)
				t.FailNow()
			}
		})
	}
}

func BenchmarkGCD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GCD(i, i+30)
	}
}
