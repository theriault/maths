package numbertheory

import (
	"fmt"
	"testing"
)

func ExampleLCM() {
	lcm := LCM(48, 18)
	fmt.Printf("The LCM of 18 and 48 is %v.", lcm)

	// Output:
	// The LCM of 18 and 48 is 144.
}

func TestLCM(t *testing.T) {
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
			expected: "144",
		},
		{
			name:     "18,48",
			a:        18,
			b:        48,
			expected: "144",
		},
		{
			name:     "-18,48",
			a:        -18,
			b:        48,
			expected: "-144",
		},
		{
			name:     "11,13",
			a:        11,
			b:        13,
			expected: "143",
		},
		{
			name:     "0,11",
			a:        0,
			b:        11,
			expected: "0",
		},
		{
			name:     "0,0",
			a:        0,
			b:        0,
			expected: "0",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual := LCM(c.a, c.b)
			if fmt.Sprintf("%v", actual) != c.expected {
				t.Logf("expected %v, got %v\n", c.expected, actual)
				t.FailNow()
			}
		})
	}
}

func BenchmarkLCM(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LCM(i, i+30)
	}
}
