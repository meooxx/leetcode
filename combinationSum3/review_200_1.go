package main

func combinationSum3(k int, n int) [][]int {
	result := [][]int{}

	var next func(k int, curr []int, start, n int)
	next = func(k int, curr []int, start, n int) {
		if n < 0 {
			return
		}
		if k == 0 || n == 0 {
			if n == 0 && k == 0 {
				item := make([]int, len(curr))
				copy(item, curr)
				result = append(result, item)
			}

			return
		}
		for ; start <= 9; start++ {
			curr = append(curr, start)
			next(k-1, curr, start+1, n-start)
			curr = curr[:len(curr)-1]
		}

	}
	next(k, []int{}, 1, n)
	return result
}
