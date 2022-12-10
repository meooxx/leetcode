package main

import "math"

func countPrimes(n int) int {
	if n < 2 {
		return 0
	}
	primes := make([]bool, n+1)
	primes[0] = true
	primes[1] = true
	count := 0
	// 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16
	// why i * i, e.g.4
	//   2*4  3*4    4*4 5*4
	// <4被前面覆盖了  -----     
	sq := int(math.Sqrt(float64(n)))
	// i <= sq, 10: 3 => 9 
	for i := 2; i <= sq; i++ {
		if !primes[i] {
			for j := i * i; j < n; j += i {
				primes[j] = true
			}
		}
	}
	for i := 2; i < n; i++ {
		if !primes[i] {
			count++
		}
	}
	return count
}
