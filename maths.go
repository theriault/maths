// Package maths provides shared interfaces and functions across all packages in the maths module.
package maths

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Float interface {
	~float32 | ~float64
}

type Complex interface {
	~complex64 | ~complex128
}

// Mutate copies the given slice and returns a new slice suitable for mutation.
//
// Space complexity: O(n)
func Mutate[T Integer | Float | Complex](X []T) []T {
	Y := make([]T, len(X))
	copy(Y, X)
	return Y
}

// Coerce copies the given slice into a new primitive type t and returns a new slice suitable for mutation
//
// Time complexity: O(n)
// Space complexity: O(n)
func Coerce[FromType Integer | Float, ToType Integer | Float](X []FromType, t ToType) []ToType {
	coercedX := make([]ToType, len(X))
	for k, v := range X {
		coercedX[k] = ToType(v)
	}
	return coercedX
}
