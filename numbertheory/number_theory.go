// Package numbertheory provides mathematical functions related to number theory.
//
// https://en.wikipedia.org/wiki/Number_theory
package numbertheory

type Integer interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

type Float interface {
	float32 | float64
}

type Number interface {
	Integer | Float
}
