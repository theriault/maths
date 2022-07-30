package numbertheory

import (
	"fmt"
	"testing"
)

func ExampleCoprime() {
	a := 2 * 3 * 5
	b := 7 * 11 * 13
	r := Coprime(a, b)
	fmt.Printf("coprime(%d, %d) is %v.\n", a, b, r)

	// Output:
	// coprime(30, 1001) is true.
}

func TestCoprime(t *testing.T) {
	compare := func(t *testing.T, a bool, b bool) {
		if a != b {
			t.Logf("expected %v, got %v", a, b)
			t.FailNow()
		}
	}
	cases := []struct {
		a uint64
		b uint64
		r bool
	}{
		{
			a: 3 * 5 * 7,
			b: 7 * 11 * 13,
			r: false,
		},
		{
			a: 7 * 11 * 13,
			b: 3 * 5 * 7,
			r: false,
		},
		{
			a: 3 * 5 * 7,
			b: 11 * 13 * 17,
			r: true,
		},
		{
			a: 11 * 13 * 17,
			b: 3 * 5 * 7,
			r: true,
		},
	}

	for k, c := range cases {
		k, c := k, c
		t.Run(fmt.Sprintf("test %v", k), func(t *testing.T) {
			t.Parallel()
			compare(t, c.r, Coprime(c.a, c.b))
		})
	}
}

func BenchmarkCoprime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Coprime(2*3*5, 5*7*11)
	}
}
