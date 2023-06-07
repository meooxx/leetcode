package main

import "math"

func nthSuperUglyNumber(n int, primes []int) int {
	uglys := make([]int, n)
	uglys[0] = 1
	pIndex := make([]int, len(primes))
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := 1; i < n; i++ {
		uglys[i] = math.MaxInt32
		for j := 0; j < len(primes); j++ {
			// 1 [2] [4] 7 8 13 14 16 19 26 28 32
			//        min(7,[8]) 
			// 1 ?  uglys
			//   2  * uglys[pIndex] == 0, pIndex++
			//   7  * uglys[pIndex] == 0
			uglys[i] = min(uglys[i], primes[j]*uglys[pIndex[j]])
		}
		for k := 0; k < len(pIndex); k++ {
			if uglys[pIndex[k]]*primes[k] == uglys[i] {
				pIndex[k]++
			}
		}
	}
	return uglys[n-1]
}
