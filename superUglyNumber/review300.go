package main

import "math"

func nthSuperUglyNumber(n int, primes []int) int {

	uglys := make([]int, n)
	// prime * uglys[startIndex]
	startIndex := make([]int, len(primes))

	uglys[0] = 1

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	//                  i = prime * startIndex-th ugly ->
	// uglys      	1,  2,  4,  7, 8,  13
	// primes       2,  7, 13,
	// startIndex   [       ]

	// i-th ugly number
	for i := 1; i < n; i++ {
		uglys[i] = math.MaxInt32
		// j-th prime number
		for j := 0; j < len(primes); j++ {
			uglys[i] = min(uglys[i], primes[j]*uglys[startIndex[j]])
		}

		// 如果使用了小值的那位index +1, 和去重
		// e.g. uglys[i] == 26
		// 2 和 13, 同时满足, 所以不可break
		// v =  min(p1 * 2, p2 * 2, p3 * 2)
		// if v == p1 *2, index of p1 += 1
		for k := range startIndex {
			if uglys[startIndex[k]]*primes[k] == uglys[i] {
				startIndex[k]++
				// break
			}
		}
	}
	return uglys[n-1]

}
