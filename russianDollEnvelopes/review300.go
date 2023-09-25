package main

import (
	"math"
	"sort"
)

func maxEnvelopes(envelopes [][]int) int {

	sort.Slice(envelopes, func(a, b int) bool {
		if envelopes[a][0] == envelopes[b][0] {
			return envelopes[a][1] > envelopes[b][1]
		} else {
			return envelopes[a][0] < envelopes[b][0]
		}
	})
	result := make([]int, len(envelopes))
	for i := range result {
		result[i] = math.MaxInt
	}
	size := 0
	for i := range envelopes {
		height := envelopes[i][1]
		left := 0
		right := len(result)
		for left < right {
			mid := (right-left)/2 + left
			if result[mid] < height {
				left = mid + 1
			} else {
				right = mid
			}

		}
		if size == left {
			size++
		}
		result[left] = height
	}
	return size
}
