package main

func combinationSum3(k int, n int) [][]int {
	// k == 4
	// 1234
	if n < (1+k)*(k)/2 {
		return [][]int{}
	}
	candidates := make([]int, 9)
	for i := range candidates {
		candidates[i] = i + 1
	}
	result := [][]int{}
	next(candidates, n, k, []int{}, &result)
	return result
}

func next(candidates []int, sum, k int, curr []int, result *[][]int) {
	if sum == 0 && k == 0 {
		item := make([]int, len(curr))
		copy(item, curr)
		*result = append(*result, item)
		return
	}
	if sum < 0 || k <= 0 || len(candidates) == 0 {
		return
	}
	for i := range candidates {
		newCandidates := append([]int{}, candidates[i+1:]...)
		curr = append(curr, candidates[i])
		next(newCandidates, sum-candidates[i], k-1, curr, result)
		curr = curr[:len(curr)-1]
	}
}
