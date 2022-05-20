# Maths

[![Go Reference](https://pkg.go.dev/badge/github.com/theriault/maths.svg)](https://pkg.go.dev/github.com/theriault/maths)
[![Go](https://github.com/Theriault/maths/actions/workflows/go.yml/badge.svg)](https://github.com/Theriault/maths/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/theriault/maths)](https://goreportcard.com/report/github.com/theriault/maths)

maths includes mathematical functions not defined in the standard Go math package.

## Installation

```sh
go get github.com/theriault/maths 
```

## What's Included

### Combinatorics

```go
import "github.com/theriault/maths/combinatorics"
```

#### Factorial

[Wikipedia](https://en.wikipedia.org/wiki/Factorial) | [OEIS](https://oeis.org/A000142) | [Source](/combinatorics/factorial.go) | [Tests](/combinatorics/factorial_test.go)

$\displaystyle n! \\; = \\; \prod_{i=1}^{n} i$

```go
combinatorics.Factorial(10) // will return uint64(3628800)
```

#### Falling Factorial

[Wikipedia](https://en.wikipedia.org/wiki/Falling_and_rising_factorials) | [Source](/combinatorics/falling_factorial.go) | [Tests](/combinatorics/falling_factorial_test.go)

$ \displaystyle x^{\underline{n}} \\; = \\; \prod _{k=1}^{n}(x-k+1) $

```go
combinatorics.FallingFactorial(8, 3) // will return uint64(336)
combinatorics.PartialPermutations(8, 3) // will return uint64(336)
```

#### Rising Factorial

[Wikipedia](https://en.wikipedia.org/wiki/Falling_and_rising_factorials) | [Source](/combinatorics/rising_factorial.go) | [Tests](/combinatorics/rising_factorial_test.go)

$ \displaystyle x^{\overline{n}} \\; = \\; \prod _{k=1}^{n}(x+k-1) $

```go
combinatorics.RisingFactorial(2, 3) // will return uint64(24)
```

#### Primorial

[Wikipedia](https://en.wikipedia.org/wiki/Primorial) | [OEIS](https://oeis.org/A002110) | [Source](/combinatorics/primorial.go) | [Tests](/combinatorics/primorial_test.go)

$ \displaystyle n\\# \\; = \\; \prod_{\substack{i=2 \\\\ i \in \mathbb{P}}}^{n} i $

```go
combinatorics.Primorial(30) // will return uint64(6469693230)
```

### Number Theory

```go
import "github.com/theriault/maths/numbertheory"
```

```go
// Aliquot sum
numbertheory.AliquotSum(60) // will return uint64(108)

// Möbius function
numbertheory.Mobius(70) // will return int8(-1)

// Number of divisors
numbertheory.NumberOfDivisors(48) // will return uint64(10)

// Politeness
numbertheory.Politeness(32) // will return uint64(0)

// Polygonal numbers
numbertheory.PolygonalNumber(3, 4) // will return uint64(10)
numbertheory.PolygonalRoot(3, 10) // will return float64(4)
numbertheory.PolygonalSides(4, 10) // will return float64(3)

// Prime factorize any uint64
numbertheory.PrimeFactorization(184756) // will return []uint64{2, 2, 11, 13, 17, 19}

// Radical
numbertheory.Radical(60) // will return uint64(30)

// Sum of divisors
numbertheory.SumOfDivisors(48) // will return uint64(10)

// Euler's totient
numbertheory.Totient(68) // will return uint64(32)

// Jordan's totient
numbertheory.TotientK(60, 2) // will return uint64(2304)
```

### Statistics

```go
import "github.com/theriault/maths/statistics"
```

```go
// Min/Max/Range
n := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
statistics.Min(n...) // will return uint8(1)
statistics.Max(n...) // will return uint8(10)
statistics.Range(n...) // will return uint8(9)

// Sum
statistics.Sum(1.1, 1.2, 1.3) // will return float64(3.6)

// Mode
statistics.Mode(8, 3, 6, 2, 7, 1, 8, 3, 7, 4, 8) // will return []float64{8}

// Means
statistics.Mean(1, 1000) // will return float64(500.5)
statistics.GeometricMean(1, 1000) // will return float64(31.62...)
statistics.HarmonicMean(1, 1000) // will return float64(1.99...)

// Median/Tertile/Quantile/.../Percentile
n := []float64{3, 6, 7, 8, 8, 10, 13, 15, 16, 20}
statistics.Median(n) // will return float64(9)
statistics.Quantile(n, 2) // median: will return []float64{9}
statistics.Quantile(n, 3) // tertiles: will return []float64{8, 13}
statistics.Quantile(n, 4) // quartiles: will return []float64{7.25, 9, 14.5}
statistics.Quantile(n, 100) // percentile: will return []float64{3.27, 3.54, 3.81, 4.08, ...95 other values...}
```
