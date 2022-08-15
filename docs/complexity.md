# Complexity

## Package: `bitmap`

| Function                  | Time     | Space  |
|---------------------------|----------|--------|
| `And`                     | $O(\min(n,m))$ | $O(1)$ |
| `Clear`                   | $O(1)$   | $O(1)$ |
| `ClearAll`                | $O(n)$   | $O(1)$ |
| `Flip`                    | $O(1)$   | $O(1)$ |
| `FlipAll`                 | $O(n)$   | $O(1)$ |
| `Len`                     | $O(1)$   | $O(1)$ |
| `Or`                      | $O(\min(n,m))$ | $O(1)$ |
| `Set`                     | $O(1)$   | $O(1)$ |
| `SetAll`                  | $O(n)$   | $O(1)$ |
| `Sum`                     | $O(n)$   | $O(1)$ |
| `Test`                    | $O(1)$   | $O(1)$ |
| `Valid`                   | $O(1)$   | $O(1)$ |
| `Xor`                     | $O(\min(n,m))$ | $O(1)$ |

## Package: `combinatorics`

| Function                  | Time     | Space  |
|---------------------------|----------|--------|
| `Factorial`               | $O(1)$   | $O(1)$ |
| `FallingFactorial`        | $O(n)$   | $O(1)$ |
| `PartialPermutations`     | $O(n)$   | $O(1)$ |
| `RisingFactorial`         | $O(n)$   | $O(1)$ |

## Package: `numbertheory`

| Function             | Time         | Space |
|----------------------|--------------|-------|
| `AliquotSum`         | $O(n^{0.5})$ | $O(\log{n})$ |
| `DigitSum`           | $O(\log_b n)$| $O(1)$ |
| `DigitalRoot`        | $O(1)$| $O(1)$ |
| `Coprime`            | $O(\log{\max(n,m)})$ | $O(1)$ |
| `GCD`                | $O(\log{\max(n,m)})$ | $O(1)$ |
| `LCM`                | $O(\log{\max(n,m)})$ | $O(1)$ |
| `Mobius`             | $O(n^{0.5})$ | $O(\log{n})$ |
| `NumberOfDivisors`   | $O(n^{0.5})$ | $O(\log{n})$ |
| `Politeness`         | $O(n^{0.5})$ | $O(\log{n})$ |
| `PolygonalNumber`    | $O(1)$       | $O(1)$ |
| `PolygonalRoot`      | $O(1)$       | $O(1)$ |
| `PolygonalSides`     | $O(1)$       | $O(1)$ |
| `PrimeFactorization` | $O(n^{0.5})$ | $O(\log{n})$ |
| `Radical`            | $O(n^{0.5})$ | $O(\log{n})$ |
| `SumOfDivisors`      | $O(n^{0.5})$ | $O(\log{n})$ |
| `Totient`            | $O(n^{0.5})$ | $O(\log{n})$ |
| `TotientK`           | $O(n^{0.5})$ | $O(\log{n})$ |

## Package: `statistics`

| Function                  | Time   | Space
|---------------------------|--------|-------|
| `CentralMoment`           | $O(n)$ | $O(1)$
| `ExcessSampleKurtosis`    | $O(n)$ | $O(1)$
| `GeneralizedMean`         | $O(n)$ | $O(1)$
| `GeometricMean`           | $O(n)$ | $O(1)$
| `HarmonicMean`            | $O(n)$ | $O(1)$
| `InterquartileRange`      | $O(n\log{n})$ | $O(n)$
| `Kurtosis`                | $O(n)$ | $O(1)$
| `LogisticFunction`        | $O(1)$ | $O(1)$
| `Max`                     | $O(n)$ | $O(1)$
| `Mean`                    | $O(n)$ | $O(1)$
| `Min`                     | $O(n)$ | $O(1)$
| `Mode`                    | $O(n\log{n})$ | $O(n)$
| `PowerSumAround`          | $O(n)$ | $O(1)$
| `PowerSum`                | $O(n)$ | $O(1)$
| `Quantile`                | $O(n\log{n})$ | $O(n+m)$
| `Range`                   | $O(n)$ | $O(1)$
| `RootMeanSquare`          | $O(n)$ | $O(1)$
| `SampleKurtosis`          | $O(n)$ | $O(1)$
| `SampleSkewness`          | $O(n)$ | $O(1)$
| `SampleStandardDeviation` | $O(n)$ | $O(1)$
| `SampleStandardError`     | $O(n)$ | $O(1)$
| `SampleVariance`          | $O(n)$ | $O(1)$
| `SimpleMovingAverage`     | $O(n)$ | $O(k)$
| `Skewness`                | $O(n)$ | $O(1)$
| `StandardDeviation`       | $O(n)$ | $O(1)$
| `StandardError`           | $O(n)$ | $O(1)$
| `Sum`                     | $O(n)$ | $O(1)$
| `Variance`                | $O(n)$ | $O(1)$
| `WeightedGeneralizedMean` | $O(n)$ | $O(1)$