package main

import "fmt"

func main() {
	fmt.Println(combinationSum([]int{1, 2}, 8))
}

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	usedCandidate := []int{}
	backtrack(candidates, &result, usedCandidate, target, 0)
	return result
}

func backtrack(candidates []int, result *[][]int, usedCandidate []int, remainder, start int) {

	if remainder < 0 {
		return
	}
	if remainder == 0 {
		i := make([]int, len(usedCandidate))
		copy(i, usedCandidate)
		*result = append(*result, i)
	} else {
		for i := start; i < len(candidates); i++ {
			usedCandidate = append(usedCandidate, candidates[i])
			backtrack(candidates, result, usedCandidate, remainder-candidates[i], i)
			usedCandidate = usedCandidate[:len(usedCandidate)-1]
		}
	}

}
