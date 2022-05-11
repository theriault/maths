package number

import (
	"math"
)

type Integer interface {
	int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64
}

type Float interface {
	float32 | float64
}

type Number interface {
	Integer | Float
}

func Abs[A Number](a A) A {
	return A(math.Abs(float64(a)))
}

func Add[A Number, B Number](a A, b B) A {
	return a + A(b)
}

func Acos[A Number](a A) A {
	return A(math.Acos(float64(a)))
}

func Acosh[A Number](a A) A {
	return A(math.Acosh(float64(a)))
}

func Asin[A Number](a A) A {
	return A(math.Asin(float64(a)))
}

func Asinh[A Number](a A) A {
	return A(math.Asinh(float64(a)))
}

func Atan[A Number](a A) A {
	return A(math.Atan(float64(a)))
}

func Atan2[A Number, B Number](a A, b B) A {
	return A(math.Atan2(float64(a), float64(b)))
}

func Atanh[A Number](a A) A {
	return A(math.Atanh(float64(a)))
}

func Cbrt[A Number](a A) A {
	return A(math.Cbrt(float64(a)))
}

func Ceil[A Number](a A) A {
	return A(math.Ceil(float64(a)))
}

func Copysign[A Number, B Number](a A, b B) A {
	return A(math.Copysign(float64(a), float64(b)))
}

func Cos[A Number](a A) A {
	return A(math.Cos(float64(a)))
}

func Cosh[A Number](a A) A {
	return A(math.Cosh(float64(a)))
}

func Dim[A Number, B Number](a A, b B) A {
	return A(math.Dim(float64(a), float64(b)))
}

func Div[A Number, B Number](a A, b B) A {
	return a / A(b)
}

func Erf[A Number](a A) A {
	return A(math.Erf(float64(a)))
}

func Erfc[A Number](a A) A {
	return A(math.Erfc(float64(a)))
}

func Erfcinv[A Number](a A) A {
	return A(math.Erfcinv(float64(a)))
}

func Erfinv[A Number](a A) A {
	return A(math.Erfinv(float64(a)))
}

func Exp[A Number](a A) A {
	return A(math.Exp(float64(a)))
}

func Exp2[A Number](a A) A {
	return A(math.Exp2(float64(a)))
}

func Expm1[A Number](a A) A {
	return A(math.Expm1(float64(a)))
}

func FMA[A Number, B Number, C Number](a A, b B, c C) A {
	return A(math.FMA(float64(a), float64(b), float64(c)))
}

func Floor[A Number](a A) A {
	return A(math.Floor(float64(a)))
}

func Gamma[A Number](a A) A {
	return A(math.Gamma(float64(a)))
}

func Hypot[A Number, B Number](a A, b B) A {
	return A(math.Hypot(float64(a), float64(b)))
}

func J0[A Number](a A) A {
	return A(math.J0(float64(a)))
}

func J1[A Number](a A) A {
	return A(math.J1(float64(a)))
}

func Log[A Number](a A) A {
	return A(math.Log(float64(a)))
}

func Log10[A Number](a A) A {
	return A(math.Log10(float64(a)))
}

func Log1p[A Number](a A) A {
	return A(math.Log1p(float64(a)))
}

func Log2[A Number](a A) A {
	return A(math.Log2(float64(a)))
}

func Logb[A Number](a A) A {
	return A(math.Logb(float64(a)))
}

func Max[A Number, B Number](a A, b B) A {
	return A(math.Max(float64(a), float64(b)))
}

func Min[A Number, B Number](a A, b B) A {
	return A(math.Min(float64(a), float64(b)))
}

func Mod[A Number, B Number](a A, b B) A {
	return A(math.Mod(float64(a), float64(b)))
}

func Mul[A Number, B Number](a A, b B) A {
	return a * A(b)
}

func Pow[A Number, B Number](a A, b B) A {
	return A(math.Pow(float64(a), float64(b)))
}

func Pow10[A Number](a A) A {
	return A(math.Pow10(int(a)))
}

func Remainder[A Number, B Number](a A, b B) A {
	return A(math.Remainder(float64(a), float64(b)))
}

func Round[A Number](a A) A {
	return A(math.Round(float64(a)))
}

func RoundToEven[A Number](a A) A {
	return A(math.RoundToEven(float64(a)))
}

func Signbit[A Number](a A) bool {
	return math.Signbit(float64(a))
}

func Sin[A Number](a A) A {
	return A(math.Sin(float64(a)))
}

func Sincos[A Number](a A) (A, A) {
	b, c := math.Sincos(float64(a))
	return A(b), A(c)
}

func Sinh[A Number](a A) A {
	return A(math.Sinh(float64(a)))
}

func Sqrt[A Number](a A) A {
	return A(math.Sqrt(float64(a)))
}

func Sub[A Number, B Number](a A, b B) A {
	return a - A(b)
}

func Tan[A Number](a A) A {
	return A(math.Tan(float64(a)))
}

func Tanh[A Number](a A) A {
	return A(math.Tanh(float64(a)))
}

func Trunc[A Number](a A) A {
	return A(math.Trunc(float64(a)))
}

func Y0[A Number](a A) A {
	return A(math.Y0(float64(a)))
}

func Y1[A Number](a A) A {
	return A(math.Y1(float64(a)))
}

func Yn[A Number, B Number](a A, b B) A {
	return A(math.Yn(int(a), float64(b)))
}
