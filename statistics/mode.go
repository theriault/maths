package statistics

import (
	"sort"

	"github.com/theriault/maths"
)

// Mode returns the values that appear most often in a set of data values.
//
// Nullary invocation returns an empty array.
// The result may be unimodal or multimodal depending on the distribution of the data.
//
// Time complexity: O(n log n)
// Space complexity: O(n)
//
// https://en.wikipedia.org/wiki/Mode_(statistics)
func Mode[A maths.Integer | maths.Float](numbers ...A) []float64 {
	modes := make([]float64, 0)
	l := len(numbers)
	if l == 0 {
		return modes
	}

	c := make([]float64, l)
	for k, v := range numbers {
		c[k] = float64(v)
	}
	sort.Float64s(c)

	len := 0
	longestLen := 0
	lastVal := c[0]
	for _, v := range c {
		if v == lastVal {
			len++
		} else {
			if len > longestLen {
				modes = []float64{lastVal}
				longestLen = len
			} else if len == longestLen {
				modes = append(modes, lastVal)
			}
			len = 1
			lastVal = v
		}
	}
	if len > longestLen {
		modes = []float64{lastVal}
	} else if len == longestLen {
		modes = append(modes, lastVal)
	}
	return modes
}
