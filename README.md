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

$\displaystyle x^{\underline{n}} \\; = \\; \prod _{k=1}^{n}(x-k+1)$

```go
combinatorics.FallingFactorial(8, 3) // will return uint64(336)
combinatorics.PartialPermutations(8, 3) // will return uint64(336)
```

#### Rising Factorial

[Source](/combinatorics/rising_factorial.go) | [Tests](/combinatorics/rising_factorial_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Falling_and_rising_factorials)

$\displaystyle x^{\overline{n}} \\; = \\; \prod _{k=1}^{n}(x+k-1)$

```go
combinatorics.RisingFactorial(2, 3) // will return uint64(24)
```

### Number Theory

```go
import "github.com/theriault/maths/numbertheory"
```

#### Aliquot Sum

[Source](/numbertheory/aliquot_sum.go) | [Tests](/numbertheory/aliquot_sum_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Aliquot_sum) | [OEIS](https://oeis.org/A001065)

$\displaystyle s(n) = \sigma_1(n) - n = \sum_{\substack{i = 1 \\\\ i | n}}^{n-1} i$

```go
numbertheory.AliquotSum(60) // will return uint64(108)
```

#### Divisors function

$\displaystyle \sigma_z(n) = \sum_{\substack{i = 1 \\\\ i | n}}^{n} i^{z}$

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

$\displaystyle \mu(n) = \begin{cases} +1 & n \text{ is square-free with even number of prime factors} \\\\ -1 & n \text{ is square-free with odd number of prime factors} \\\\ 0 & n \text{ is not square-free} \end{cases}$

```go
numbertheory.Mobius(70) // will return int8(-1)
```

#### Politeness

[Source](/numbertheory/politeness.go) | [Tests](/numbertheory/politeness_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Polite_number) | [OEIS](https://oeis.org/A069283)

$\displaystyle p(n) = \left( \prod_{\substack{p |n \\\\ p \neq 2}}^{n} v_p(n) + 1\right)-1$

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

$\displaystyle n\\# = \prod_{\substack{i=2 \\\\ i \in \mathbb{P}}}^{n} i$

```go
numbertheory.Primorial(30) // will return uint64(6469693230)
```

#### Radical

[Source](/numbertheory/radical.go) | [Tests](/numbertheory/radical_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Radical_of_an_integer) | [OEIS](https://oeis.org/A007947)

$\displaystyle rad(n) = \prod_{p | n}p$

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

##### Generalized Mean

[Source](/statistics/generalized_mean.go) | [Tests](/statistics/generalized_mean_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Generalized_mean)

```go
statistics.GeneralizedMean([]float64{1, 1000}, 2) // will return float64(707.1071347398497)
statistics.RootMeanSquare(1, 1000)  // will return float64(707.1071347398497)
```

##### Arithmetic Mean

[Source](/statistics/mean.go) | [Tests](/statistics/mean_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Mean#Arithmetic_mean_(AM))

$\displaystyle \bar{x} = \frac{1}{n}\left (\sum_{i=1}^n{x_i}\right ) = \frac{x_1+x_2+\cdots +x_n}{n}$

```go
statistics.Mean(1, 1000) // will return float64(500.5)
```

##### Geometric Mean

[Source](/statistics/geometric_mean.go) | [Tests](/statistics/geometric_mean_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Mean#Geometric_mean_(GM))

$\displaystyle \bar{x} = \left( \prod_{i=1}^n{x_i} \right )^\frac{1}{n} = \exp{\left( {\frac{1}{n}\sum\limits_{i=1}^{n}\ln x_i} \right)} = \left(x_1 x_2 \cdots x_n \right)^\frac{1}{n}$

```go
statistics.GeometricMean(1, 1000) // will return float64(31.62...)
```

##### Harmonic Mean

[Source](/statistics/harmonic_mean.go) | [Tests](/statistics/harmonic_mean_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Mean#Harmonic_mean_(HM))

$\displaystyle \bar{x} = n \left ( \sum_{i=1}^n \frac{1}{x_i} \right ) ^{-1}$

```go
statistics.HarmonicMean(1, 1000) // will return float64(1.99...)
```

#### Central Moment

[Source](/statistics/central_moment.go) | [Tests](/statistics/central_moment_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Central_moment)

```go
statistics.CentralMoment([]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2) // returns float64(8.25)
```

#### Interquartile Range (IQR)

[Source](/statistics/interquartile_range.go) | [Tests](/statistics/interquartile_range_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Interquartile_range)

```go
statistics.InterquartileRange(3, 6, 7, 8, 8, 10, 13, 15, 16, 20) // returns float64(7.25)
```

#### Kurtosis

##### Population

[Source](/statistics/kurtosis.go) | [Tests](/statistics/kurtosis_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Kurtosis)

```go
statistics.Kurtosis(8, 3, 6, 2, 7, 1, 8, 3, 7, 4, 8) // returns float64(1.5133)
```

##### Sample

[Source](/statistics/sample_kurtosis.go) | [Tests](/statistics/sample_kurtosis_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Kurtosis)

```go
statistics.SampleKurtosis(8, 3, 6, 2, 7, 1, 8, 3, 7, 4, 8) // returns float64(2.522167)
```

##### Excess Sample Kurtosis

[Source](/statistics/excess_sample_kurtosis.go) | [Tests](/statistics/excess_sample_kurtosis_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Kurtosis)

```go
statistics.ExcessSampleKurtosis([]uint8{8, 3, 6, 2, 7, 1, 8, 3, 7, 4, 8}) // returns float64(-1.6445)
```

#### Logistic Function

[Source](/statistics/logistic_function.go) | [Tests](/statistics/logistic_function_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Logistic_function)

$\displaystyle f(x) = \frac{L}{1+e^{-k(x-x_0)}}$

- $L$ - curve's max value
- $x_0$ - sigmoid's midpoint
- $k$ - logistic growth rate

```go
maxValue := 1.0
midpoint := 0.0
growthRate := 1.0
fx := statistics.LogisticFunction(maxValue, midpoint, growthRate) // will return func (x float64) float64
```

#### Mode

[Source](/statistics/mode.go) | [Tests](/statistics/mode_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Mode_(statistics))

```go
statistics.Mode(8, 3, 6, 2, 7, 1, 8, 3, 7, 4, 8) // will return []float64{8}
```

#### Power Sum

[Source](/statistics/power_sum.go) | [Tests](/statistics/power_sum_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Sums_of_powers)

$\displaystyle S_p(x_1, x_2, ..., x_n) = \sum_{i=1}^{n} x_i^p$

```go
statistics.PowerSum([]float64{2, 3, 4}, 2) // will return float64(29)
```

#### Power Sum Around

[Source](/statistics/power_sum_around.go) | [Wikipedia](https://en.wikipedia.org/wiki/Sums_of_powers)

$\displaystyle S_{p,y}(x_1, x_2, ..., x_n) = \sum_{i=1}^{n} (x_i - y)^p$

```go
statistics.PowerSumAround([]float64{2, 3, 4}, 3, 2) // will return float64(29)
```

#### Quantiles (Median/Tertile/Quartile/.../Percentile)

[Source](/statistics/quantile.go) | [Tests](/statistics/quantile_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Quantile)

```go
n := []float64{3, 6, 7, 8, 8, 10, 13, 15, 16, 20}
statistics.Quantile(n, 2) // median: will return []float64{9}
statistics.Quantile(n, 3) // tertiles: will return []float64{8, 13}
statistics.Quantile(n, 4) // quartiles: will return []float64{7.25, 9, 14.5}
statistics.Quantile(n, 100) // percentile: will return []float64{3.27, 3.54, 3.81, 4.08, ...95 other values...}

// aliases
statistics.Tertile(n) // will return []float64{8, 13}
statistics.Quartile(n) // will return []float64{7.25, 9, 14.5}
statistics.Percentile(n) // will return []float64{3.27, 3.54, 3.81, 4.08, ...95 other values...}
```

Median ([Source](/statistics/median.go) | [Tests](/statistics/median_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Median))

```go
n := []float64{3, 6, 7, 8, 8, 10, 13, 15, 16, 20}
statistics.Median(n) // will return float64(9)
```

#### Sample Extrema (Max/Min/Range)

##### Sample Maximum / Largest Observation

[Source](/statistics/max.go) | [Tests](/statistics/max_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Sample_maximum_and_minimum)

```go
n := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
statistics.Max(n...) // will return uint8(10)
```

##### Sample Minimum / Smallest Observation

[Source](/statistics/min.go) | [Tests](/statistics/min_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Sample_maximum_and_minimum)

```go
n := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
statistics.Min(n...) // will return uint8(1)
```

##### Range

[Source](/statistics/range.go) | [Tests](/statistics/range_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Range_(statistics))

```go
n := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
statistics.Range(n...) // will return uint8(9)
```

#### Skewness

##### Population

[Source](/statistics/skewness.go) | [Tests](/statistics/skewness_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Skewness)

```go
statistics.Skewness(8, 3, 6, 2, 7, 1, 8, 3, 7, 4, 8) // returns float64(-0.274241)
```

##### Sample

[Source](/statistics/sample_skewness.go) | [Tests](/statistics/sample_skewness_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Skewness)

```go
statistics.SampleSkewness(8, 3, 6, 2, 7, 1, 8, 3, 7, 4, 8) // returns float64(-0.319584)
```

#### Standard Deviation

##### Population

[Source](/statistics/standard_deviation.go) | [Tests](/statistics/standard_deviation_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Standard_deviation)

$\displaystyle \sigma = \sqrt{\left(\frac{1}{N} \sum_{i=1}^{N} x_{i}^{2} \right) - \mu^2}$

```go
statistics.StandardDeviation(8, 3, 6, 2, 7, 1, 8, 3, 7, 4, 8) // will return []float64{8}
```

##### Sample

[Source](/statistics/sample_standard_deviation.go) | [Tests](/statistics/sample_standard_deviation_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Standard_deviation)

$\displaystyle s = \sqrt{\sigma^2 \frac{N}{N-1}}$

```go
statistics.SampleStandardDeviation(8, 3, 6, 2, 7, 1, 8, 3, 7, 4, 8) // will return []float64{8}
```

#### Standard Error

##### Population

[Source](/statistics/standard_error.go) | [Tests](/statistics/standard_error_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Standard_error)

$\displaystyle \sigma_{\bar{x}} = \frac{\sigma}{\sqrt{n}}$

```go
statistics.StandardError(8, 3, 6, 2, 7, 1, 8, 3, 7, 4, 8) // will return []float64{8}
```

##### Sample

[Source](/statistics/sample_standard_error.go) | [Tests](/statistics/sample_standard_error_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Standard_error)

$\displaystyle s_{\bar{x}} = \frac{s}{\sqrt{n}}$

```go
statistics.SampleStandardError(8, 3, 6, 2, 7, 1, 8, 3, 7, 4, 8) // will return []float64{8}
```

#### Sum/Summation

[Source](/statistics/sum.go) | [Tests](/statistics/sum_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Summation)

$\displaystyle S(x_1, x_2, ..., x_n) = \sum_{i=1}^{n} x_i $

```go
statistics.Sum(1.1, 1.2, 1.3) // will return float64(3.6)
```

#### Variance

##### Population

[Source](/statistics/variance.go) | [Tests](/statistics/variance_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Variance)

$\displaystyle \sigma^2 = \left(\frac{1}{N} \sum_{i=1}^{N} x_{i}^{2} \right) - \mu^2$

```go
statistics.Variance(8, 3, 6, 2, 7, 1, 8, 3, 7, 4, 8) // will return []float64{8}
```

##### Sample

[Source](/statistics/sample_variance.go) | [Tests](/statistics/sample_variance_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Variance)

$\displaystyle s^2 = \sigma^2 \frac{N}{N-1}$

```go
statistics.SampleVariance(8, 3, 6, 2, 7, 1, 8, 3, 7, 4, 8) // will return []float64{8}
```

#### Weighted Average/Mean

##### Weighted Generalized Mean

[Source](/statistics/weighted_generalized_mean.go) | [Tests](/statistics/weighted_generalized_mean_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Generalized_mean)

```go
X := []uint8{8, 7, 3, 2, 6, 11, 6, 7, 2, 1, 7}
W := []uint8{1, 2, 1, 1, 2, 1, 2, 1, 2, 1, 2}
mean, err := statistics.WeightedGeneralizedMean(X, Y, 1) // will return float64(5.5)
```

##### Weighted Arithmetic Mean

[Source](/statistics/weighted_generalized_mean.go) | [Tests](/statistics/weighted_generalized_mean_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Generalized_mean)

```go
X := []uint8{8, 7, 3, 2, 6, 11, 6, 7, 2, 1, 7}
W := []uint8{1, 2, 1, 1, 2, 1, 2, 1, 2, 1, 2}
mean, err := statistics.WeightedMean(X, Y) // will return float64(5.5)
```

##### Weighted Geometric Mean

[Source](/statistics/weighted_generalized_mean.go) | [Tests](/statistics/weighted_generalized_mean_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Generalized_mean)

```go
X := []uint8{8, 7, 3, 2, 6, 11, 6, 7, 2, 1, 7}
W := []uint8{1, 2, 1, 1, 2, 1, 2, 1, 2, 1, 2}
mean, err := statistics.WeightedGeometricMean(X, Y) // will return float64(4.6)
```

##### Weighted Harmonic Mean

[Source](/statistics/weighted_generalized_mean.go) | [Tests](/statistics/weighted_generalized_mean_test.go) | [Wikipedia](https://en.wikipedia.org/wiki/Generalized_mean)

```go
X := []uint8{8, 7, 3, 2, 6, 11, 6, 7, 2, 1, 7}
W := []uint8{1, 2, 1, 1, 2, 1, 2, 1, 2, 1, 2}
mean, err := statistics.WeightedHarmonicMean(X, Y) // will return float64(5.8)
```
