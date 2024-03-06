func combinationSum3(k int, n int) [][]int {

	// [1-9]
	candidates := make([]int, 9)

	for i := range candidates {
		candidates[i] = i + 1
	}
	result := [][]int{}
	next(candidates, k, n, &result, []int{})
	return result
}

func next(candidates []int, k int, n int, result *[][]int, current []int) {
	if k == 0 || n <= 0 || len(candidates) == 0 {
		if k == 0 && n == 0 && len(current) > 0 {
			temp := make([]int, len(current))
			copy(temp, current)
			*result = append(*result, temp)

		}
		return
	}

	for i := range candidates {
		current = append(current, candidates[i])
		next(candidates[i+1:], k-1, n-candidates[i], result, current)
		current = current[:len(current)-1]
	}

}