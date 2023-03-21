package main

import "fmt"

func summaryRanges(nums []int) []string {
	index := 0
	result := []string{}
	startIndex := 0
	for i := 0; i < len(nums); i++ {
		start := nums[startIndex]
		if i != len(nums)-1 && nums[i+1] == start+index+1 {
			index++
		} else {
			if index == 0 {
				result = append(result, fmt.Sprint(start))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", start, start+index))
			}
			startIndex = i + 1
			index = 0
		}
	}

	return result
}
