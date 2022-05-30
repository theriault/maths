package statistics

import (
	"fmt"
	"testing"
)

func TestLogisticFunction(t *testing.T) {
	fn := LogisticFunction(1.0, 0.0, 1.0)
	cases := []struct {
		name     string
		x        float64
		expected string
	}{
		{
			name:     "test 0.5",
			x:        0.5,
			expected: "0.62245933",
		},
		{
			name:     "test 1",
			x:        1,
			expected: "0.73105858",
		},
		{
			name:     "test 5",
			x:        5,
			expected: "0.99330715",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := fmt.Sprintf("%.8f", fn(c.x))
			if actual != c.expected {
				t.Logf("expected %v, got %v", c.expected, actual)
				t.FailNow()
			}
		})
	}
}
