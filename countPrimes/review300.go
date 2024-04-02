package main

import "math"

func countPrimes(n int) int {
  if n < 2 {
    return 0
  }
	NotPrimes := make([]bool, n)
	NotPrimes[0] = true
	NotPrimes[1] = true
	sqr := int(math.Sqrt(float64(n)))

	for i := 2; i <= sqr; i++ {
		for j := i * i; j < n; j += i {
			NotPrimes[j] = true
		}
	}
	count := 0
	for _, v :=  range NotPrimes {
		if !v {
			count++
		}
	}
	return count

}
