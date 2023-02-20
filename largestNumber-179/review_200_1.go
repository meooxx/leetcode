package main

import (
	"fmt"
	"sort"
)

func largestNumber(nums []int) string {
	strNums := make([]string, len(nums))
	for i := range nums {
		strNums[i] = fmt.Sprint(nums[i])
	}
	sort.Slice(strNums, func(a, b int) bool {
		sum1 := strNums[a] + strNums[b]
		sum2 := strNums[b] + strNums[a]
		return sum1 >= sum2
	})

	if strNums[0] == "0" {
		return "0"
	}
	return strings.Join(strNums, "")
}
