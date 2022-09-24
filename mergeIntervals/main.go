package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(merge([][]int{
		{1, 3},
		{4, 5},
		{5, 7},
	}))
	fmt.Println(merge([][]int{
		{1, 5},
		{5, 7},
		{0, 4},
	}))
}

func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})
	result := [][]int{intervals[0]}

	for index := 1; index < len(intervals); index++ {

		current := intervals[index]
		pre := result[len(result)-1]
		if current[0] <= pre[1] {
			// [1,5] [2,3]
			if current[1] > pre[1] {
				pre[1] = current[1]
			}

		} else {
			result = append(result, current)
		}

	}

	return result
}
