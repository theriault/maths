package numbertheory

import (
	"fmt"
	"testing"
)

func ExamplePrimorial() {
	n, _ := Primorial(30)
	fmt.Printf("The primorial of 30 is %v.\n", n)

	// Output:
	// The primorial of 30 is 6469693230.
}

func TestPrimorial(t *testing.T) {
	compare := func(t *testing.T, n int64, a uint64, b uint64) {
		if a != b {
			t.Logf("using %v, expected %v, got %v", n, a, b)
			t.FailNow()
		}
	}
	cases := map[int64]struct {
		expected uint64
		err      bool
	}{
		-1: {0, true},
		0:  {1, false},
		5:  {2 * 3 * 5, false},
		12: {2 * 3 * 5 * 7 * 11, false},
		13: {2 * 3 * 5 * 7 * 11 * 13, false},
		52: {614889782588491410, false},
		53: {0, true},
	}
	for n, c := range cases {
		n, c := n, c
		t.Run(fmt.Sprintf("test %v", n), func(t *testing.T) {
			t.Parallel()
			actual, err := Primorial(n)
			compare(t, n, c.expected, actual)
			if c.err != (err != nil) {
				t.Logf("expected error %v, got error %v", c.err, err != nil)
				t.FailNow()
			}
		})
	}
}

func BenchmarkPrimorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Primorial(i)
	}
}
