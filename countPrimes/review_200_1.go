package main

import "math"

func countPrimes(n int) int {
	if n < 2 {
		return 0
	}
	primes := make([]bool, n+1)
	primes[0] = true
	primes[1] = true

	sqtN := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqtN; i++ {
		for j := i * i; j <= n; j += i {
			primes[j] = true
		}
	}
	// less than n
	primes[n] = true
	count := 0
	for i := range primes {
		if !primes[i] {
			count++
		}
	}
	return count
}
