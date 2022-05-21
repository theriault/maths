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

[Source](/combinatorics/factorial.go) | [Tests](/combinatorics/factorial_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Factorial) | [OEIS](https://oeis.org/A000142)

$\displaystyle n! \\; = \\; \prod_{i=1}^{n} i$

```go
combinatorics.Factorial(10) // will return uint64(3628800)
```

#### Falling Factorial

[Source](/combinatorics/falling_factorial.go) | [Tests](/combinatorics/falling_factorial_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Falling_and_rising_factorials)

$ \displaystyle x^{\underline{n}} \\; = \\; \prod _{k=1}^{n}(x-k+1) $

```go
combinatorics.FallingFactorial(8, 3) // will return uint64(336)
combinatorics.PartialPermutations(8, 3) // will return uint64(336)
```

#### Rising Factorial

[Source](/combinatorics/rising_factorial.go) | [Tests](/combinatorics/rising_factorial_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Falling_and_rising_factorials)

$ \displaystyle x^{\overline{n}} \\; = \\; \prod _{k=1}^{n}(x+k-1) $

```go
combinatorics.RisingFactorial(2, 3) // will return uint64(24)
```

### Number Theory

```go
import "github.com/theriault/maths/numbertheory"
```

#### Aliquot Sum

[Source](/numbertheory/aliquot_sum.go) | [Tests](/numbertheory/aliquot_sum_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Aliquot_sum) | [OEIS](https://oeis.org/A001065)

$ \displaystyle s(n) = \sum_{\substack{i = 1 \\\\ i | n}}^{n-1} i $

```go
numbertheory.AliquotSum(60) // will return uint64(108)
```

#### Divisors function

$ \displaystyle \sigma_z(n) = \sum_{\substack{i = 1 \\\\ i | n}}^{n} i^{z} $

##### Number-of-divisors (z = 0)

[Source](/numbertheory/number_of_divisors.go) | [Tests](/numbertheory/number_of_divisors_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Divisor_function) | [OEIS](https://oeis.org/A000005)

```go
numbertheory.NumberOfDivisors(48) // will return uint64(10)
```

##### Sum-of-divisors (z = 1)

[Source](/numbertheory/sum_of_divisors.go) | [Tests](/numbertheory/sum_of_divisors_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Divisor_function) | [OEIS](https://oeis.org/A000203)

```go
numbertheory.SumOfDivisors(48) // will return uint64(124)
```

#### Möbius function

[Source](/numbertheory/mobius.go) | [Tests](/numbertheory/mobius_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Mobius_function) | [OEIS](https://oeis.org/A008683)

$ \displaystyle \mu(n) = \begin{cases} +1 & n \text{ is square-free with even number of prime factors} \\\\ -1 & n \text{ is square-free with odd number of prime factors} \\\\ 0 & n \text{ is not square-free} \end{cases} $

```go
numbertheory.Mobius(70) // will return int8(-1)
```

#### Politeness

[Source](/numbertheory/politeness.go) | [Tests](/numbertheory/politeness_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Polite_number) | [OEIS](https://oeis.org/A069283)

$ \displaystyle p(n) = \left( \prod_{\substack{p |n \\\\ p \neq 2}}^{n} v_p(n) + 1\right)-1 $

where $v_p(n)$ is the [$p$-adic order](https://en.wikipedia.org/wiki/P-adic_order#Integers)

```go
numbertheory.Politeness(32) // will return uint64(0)
```

#### Polygonal Numbers

[Source](/numbertheory/polygonal_number.go) | [Tests](/numbertheory/polygonal_number_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Polygonal_number)

$\displaystyle p(s, n) = \frac{(s-2)n^2-(s-4)n}{2}$ 

```go
numbertheory.PolygonalNumber(3, 4) // will return uint64(10)
```

Finding $n$:

$\displaystyle p(s, x) = \frac{\sqrt{8(s-2)x+(s-4)^2}+(s-4)}{2(s-2)}$

```go
numbertheory.PolygonalRoot(3, 10) // will return float64(4)
```

Finding $s$:

$\displaystyle p(n, x) = 2+\frac{2}{n} \cdot \frac{x-n}{n-1}$

```go
numbertheory.PolygonalSides(4, 10) // will return float64(3)
```

#### Prime Factorization

[Source](/numbertheory/prime_factorization.go) | [Tests](/numbertheory/prime_factorization_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Integer_factorization)

```go
numbertheory.PrimeFactorization(184756) // will return []uint64{2, 2, 11, 13, 17, 19}
```

#### Primorial

[Source](/numbertheory/primorial.go) | [Tests](/numbertheory/primorial_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Primorial) | [OEIS](https://oeis.org/A002110)

$ \displaystyle n\\# = \prod_{\substack{i=2 \\\\ i \in \mathbb{P}}}^{n} i $

```go
numbertheory.Primorial(30) // will return uint64(6469693230)
```

#### Radical

[Source](/numbertheory/radical.go) | [Tests](/numbertheory/radical_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Radical_of_an_integer) | [OEIS](https://oeis.org/A007947)

$ \displaystyle rad(n) = \prod_{p | n}p $

```go
numbertheory.Radical(60) // will return uint64(30)
```

#### Totient

##### Euler's Totient

[Source](/numbertheory/totient.go) | [Tests](/numbertheory/totient_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Euler's_totient_function) | [OEIS](https://oeis.org/A000010)

$\displaystyle \varphi(n) = n \prod_{p | n} \left(1 - \frac{1}{p}\right) $ 

```go
numbertheory.Totient(68) // will return uint64(32)
```

##### Jordan's Totient

[Source](/numbertheory/totient.go) | [Tests](/numbertheory/totient_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Jordan's_totient_function)

$\displaystyle J_k(n) = n^k \prod_{p | n} \left(1 - \frac{1}{p^k}\right) $ 

```go
numbertheory.TotientK(60, 2) // will return uint64(2304)
```

### Statistics

```go
import "github.com/theriault/maths/statistics"
```

#### Average/Mean

##### Arithmetic Mean

[Source](/statistics/mean.go) | [Tests](/statistics/mean_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Mean#Arithmetic_mean_(AM))

$ \displaystyle \bar{x} = \frac{1}{n}\left (\sum_{i=1}^n{x_i}\right ) = \frac{x_1+x_2+\cdots +x_n}{n} $

```go
statistics.Mean(1, 1000) // will return float64(500.5)
```

##### Geometric Mean

[Source](/statistics/geometric_mean.go) | [Tests](/statistics/geometric_mean_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Mean#Geometric_mean_(GM))

$ \displaystyle \bar{x} = \left( \prod_{i=1}^n{x_i} \right )^\frac{1}{n} = \exp{\left( {\frac{1}{n}\sum\limits_{i=1}^{n}\ln x_i} \right)} = \left(x_1 x_2 \cdots x_n \right)^\frac{1}{n} $

```go
statistics.GeometricMean(1, 1000) // will return float64(31.62...)
```

##### Harmonic Mean

[Source](/statistics/harmonic_mean.go) | [Tests](/statistics/harmonic_mean_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Mean#Harmonic_mean_(HM))

$ \displaystyle \bar{x} = n \left ( \sum_{i=1}^n \frac{1}{x_i} \right ) ^{-1} $

```go
statistics.HarmonicMean(1, 1000) // will return float64(1.99...)
```

#### Mode

[Source](/statistics/mode.go) | [Tests](/statistics/mode_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Mode_(statistics))

```go
statistics.Mode(8, 3, 6, 2, 7, 1, 8, 3, 7, 4, 8) // will return []float64{8}
```

#### Quantiles (Median/Tertile/Quartile/.../Percentile)

[Source](/statistics/quantile.go) | [Tests](/statistics/quantile_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Quantile)

```go
n := []float64{3, 6, 7, 8, 8, 10, 13, 15, 16, 20}
statistics.Quantile(n, 2) // median: will return []float64{9}
statistics.Quantile(n, 3) // tertiles: will return []float64{8, 13}
statistics.Quantile(n, 4) // quartiles: will return []float64{7.25, 9, 14.5}
statistics.Quantile(n, 100) // percentile: will return []float64{3.27, 3.54, 3.81, 4.08, ...95 other values...}
```

Median ([Source](/statistics/median.go) | [Tests](/statistics/median_test.go))

```go
statistics.Median(n) // will return float64(9)
```

#### Sample Extrema (Max/Min/Range)

[Wikipedia](https://en.wikipedia.org/wiki/Sample_maximum_and_minimum)

```go
n := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
```

##### Sample Maximum / Largest Observation

[Source](/statistics/max.go) | [Tests](/statistics/max_test.go)

```go
statistics.Max(n...) // will return uint8(10)
```

##### Sample Minimum / Smallest Observation

[Source](/statistics/min.go) | [Tests](/statistics/min_test.go)

```go
statistics.Min(n...) // will return uint8(1)
```

##### Range

[Source](/statistics/range.go) | [Tests](/statistics/range_test.go)

```go
statistics.Range(n...) // will return uint8(9)
```

#### Summation

[Source](/statistics/sum.go) | [Tests](/statistics/sum_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Summation)

```go
statistics.Sum(1.1, 1.2, 1.3) // will return float64(3.6)
```
