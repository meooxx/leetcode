package main

func nthUglyNumber(n int) int {
	// don't need
	if n == 1 {
		return 1
	}
	index1 := 0
	index2 := 0
	index3 := 0

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	result := make([]int, n)
	result[0] = 1
	for i := 1; i < n; i++ {
		minV := min(min(2*result[index1], 3*result[index2]), 5*result[index3])
		result[i] = minV
		if minV == result[index1]*2 {
			index1++
		}
		if minV == result[index2]*3 {
			index2++
		}
		if minV == result[index3]*5 {
			index3++
		}

	}
	return result[n-1]
}
