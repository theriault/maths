# Maths

Maths includes mathematical functions not defined in the standard Go
math package.

## Number Theory

```go
// prime factorize any uint64
numbertheory.PrimeFactorization(184756) // will return []uint64{2, 2, 11, 13, 17, 19}

// Euler's totient
numbertheory.Totient(68) // will return uint64(32)

// Jordan's totient
numbertheory.TotientK(60, 2) // will return uint64(2304)

// Number of divisors
numbertheory.NumberOfDivisors(48) // will return uint64(10)
```
