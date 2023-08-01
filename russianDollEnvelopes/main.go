package main

import "sort"
import "math"

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(a, b int) bool {
		if envelopes[a][0] == envelopes[b][0] {
			return envelopes[b][1] < envelopes[a][1]
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
		right := len(result) - 1
		// 1 4   3
		for left < right {
			midIndex := (right-left)/2 + left
			if result[midIndex] < height {
				left = midIndex + 1
			} else {
				right = midIndex
			}
		}
		result[left] = height
		if left == size {
			size++
		}
	}
	return size

}
