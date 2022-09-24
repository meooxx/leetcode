package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

func combinationSum2(candidates []int, target int) [][]int {
	var result [][]int
	sort.Ints(candidates)
	backtrack(candidates, &result, []int{}, target, 0)
	return result
}

func backtrack(candidates []int, result *[][]int, usedCandidate []int, remainder, start int) {

	if remainder < 0 {
		return
	}
	if remainder == 0 {
		r := make([]int, len(usedCandidate))
		copy(r, usedCandidate)
		*result = append(*result, r)
	} else {
		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i-1] == candidates[i] {
				continue
			}
			usedCandidate = append(usedCandidate, candidates[i])
			backtrack(candidates, result, usedCandidate, remainder-candidates[i], i+1)
			usedCandidate = usedCandidate[:len(usedCandidate)-1]
		}
	}

}
