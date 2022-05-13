package numbertheory

import (
	"fmt"
	"testing"
)

func TestPolygonalNumber(t *testing.T) {
	compare := func(t *testing.T, expected uint64, actual uint64) {
		if expected != actual {
			t.Logf("expected %v, got %v", expected, actual)
			t.FailNow()
		}
	}
	cases := []struct {
		s        uint64
		n        uint64
		expected uint64
		err      error
	}{
		{2, 0, 0, fmt.Errorf("s must be >= 3")},
		{3, 0, 0, nil},
		{3, 1, 1, nil},
		{3, 2, 3, nil},
		{3, 3, 6, nil},
		{4, 4, 16, nil},
		{5, 5, 35, nil},
		{6, 7, 91, nil},
	}
	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("test s=%v, n=%v", c.s, c.n), func(t *testing.T) {
			t.Parallel()
			a, err := PolygonalNumber(c.s, c.n)
			if c.err != nil {
				if err == nil {
					t.FailNow()
				}
			} else {
				if err != nil {
					t.FailNow()
				}
				compare(t, c.expected, a)
			}
		})
	}
}

func TestPolygonalRoot(t *testing.T) {
	compare := func(t *testing.T, expected float64, actual float64) {
		if expected != actual {
			t.Logf("expected %v, got %v", expected, actual)
			t.FailNow()
		}
	}
	cases := []struct {
		s        uint64
		x        uint64
		expected float64
		err      error
	}{
		{2, 0, 0, fmt.Errorf("s must be >= 3")},
		{3, 6, 3, nil},
		{4, 16, 4, nil},
		{5, 35, 5, nil},
		{6, 91, 7, nil},
	}
	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("test s=%v, x=%v", c.s, c.x), func(t *testing.T) {
			t.Parallel()
			a, err := PolygonalRoot(c.s, c.x)
			if c.err != nil {
				if err == nil {
					t.FailNow()
				}
			} else {
				if err != nil {
					t.FailNow()
				}
				compare(t, c.expected, a)
			}
		})
	}
}

func TestPolygonalSides(t *testing.T) {
	compare := func(t *testing.T, expected float64, actual float64) {
		if expected != actual {
			t.Logf("expected %v, got %v", expected, actual)
			t.FailNow()
		}
	}
	cases := []struct {
		n        uint64
		x        uint64
		expected float64
		err      error
	}{
		{3, 6, 3, nil},
		{4, 16, 4, nil},
		{5, 35, 5, nil},
		{7, 91, 6, nil},
		{5, 115, 13, nil},
	}
	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("testn=%v, x=%v", c.n, c.x), func(t *testing.T) {
			t.Parallel()
			a := PolygonalSides(c.n, c.x)
			compare(t, c.expected, a)
		})
	}
}
