package primefactorization

import (
	"fmt"
	"testing"
)

func ExamplePrimeFactorization() {
	var n uint64 = 51561510
	pf := NewPrimeFactorization(n)
	fmt.Printf("The prime factorization of %v is %v.\n", n, pf.Primes())

	// Output:
	// The prime factorization of 51561510 is [2 3 5 7 11 13 17 101].
}

func TestPrimeFactorization(t *testing.T) {
	cases := []struct {
		name                     string
		n                        int64
		expectedAliquotSum       string
		expectedMobius           string
		expectedNumberOfDivisors string
		expectedPoliteness       string
		expectedPrimes           string
		expectedRadical          string
		expectedSumOfDivisors    string
		expectedTotient          string
	}{
		{
			name:                     "1",
			n:                        1,
			expectedAliquotSum:       "0",
			expectedMobius:           "1",
			expectedNumberOfDivisors: "1",
			expectedPoliteness:       "0",
			expectedPrimes:           "[]",
			expectedRadical:          "1",
			expectedSumOfDivisors:    "1",
			expectedTotient:          "1",
		},
		{
			name:                     "2*3*5",
			n:                        2 * 3 * 5,
			expectedAliquotSum:       "42",
			expectedMobius:           "-1",
			expectedNumberOfDivisors: "8",
			expectedPoliteness:       "3",
			expectedPrimes:           "[2 3 5]",
			expectedRadical:          "30",
			expectedSumOfDivisors:    "72",
			expectedTotient:          "8",
		},
		{
			name:                     "2*2*3*3",
			n:                        2 * 2 * 3 * 3,
			expectedAliquotSum:       "55",
			expectedMobius:           "0",
			expectedNumberOfDivisors: "9",
			expectedPoliteness:       "2",
			expectedPrimes:           "[2 2 3 3]",
			expectedRadical:          "6",
			expectedSumOfDivisors:    "91",
			expectedTotient:          "12",
		},
		{
			name:                     "151*142_395_347",
			n:                        151 * 142_395_347,
			expectedAliquotSum:       "142395499",
			expectedMobius:           "1",
			expectedNumberOfDivisors: "4",
			expectedPoliteness:       "3",
			expectedPrimes:           "[151 142395347]",
			expectedRadical:          "21501697397",
			expectedSumOfDivisors:    "21644092896",
			expectedTotient:          "21359301900",
		},
	}
	for n, c := range cases {
		n, c := n, c
		t.Run(fmt.Sprintf("test %v: %v", n, c.name), func(t *testing.T) {
			var actual string
			t.Parallel()
			pf := NewPrimeFactorization(c.n)

			actual = fmt.Sprintf("%v", pf.AliquotSum())
			if actual != c.expectedAliquotSum {
				t.Errorf("AliquotSum() - got %v, expected %v", actual, c.expectedAliquotSum)
			}

			actual = fmt.Sprintf("%v", pf.Mobius())
			if actual != c.expectedMobius {
				t.Errorf("Mobius() - got %v, expected %v", actual, c.expectedMobius)
			}

			actual = fmt.Sprintf("%v", pf.NumberOfDivisors())
			if actual != c.expectedNumberOfDivisors {
				t.Errorf("NumberOfDivisors() - got %v, expected %v", actual, c.expectedNumberOfDivisors)
			}

			actual = fmt.Sprintf("%v", pf.Politeness())
			if actual != c.expectedPoliteness {
				t.Errorf("Politeness() - got %v, expected %v", actual, c.expectedPoliteness)
			}

			actual = fmt.Sprintf("%v", pf.Primes())
			if actual != c.expectedPrimes {
				t.Errorf("Primes() - got %v, expected %v", actual, c.expectedPrimes)
			}

			actual = fmt.Sprintf("%v", pf.Radical())
			if actual != c.expectedRadical {
				t.Errorf("Radical() - got %v, expected %v", actual, c.expectedRadical)
			}

			actual = fmt.Sprintf("%v", pf.SumOfDivisors())
			if actual != c.expectedSumOfDivisors {
				t.Errorf("SumOfDivisors() - got %v, expected %v", actual, c.expectedSumOfDivisors)
			}

			actual = fmt.Sprintf("%v", pf.Totient())
			if actual != c.expectedTotient {
				t.Errorf("Totient() - got %v, expected %v", actual, c.expectedTotient)
			}
		})
	}
}

func BenchmarkPrimeFactorization(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewPrimeFactorization(uint64(i))
	}
}
