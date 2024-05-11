package main

import (
	"fmt"
	"sort"
	"strings"
)

func largestNumber(nums []int) string {
	numStr := []string{}
	for i := range nums {
		numStr[i] = fmt.Sprint(nums[i])
	}
	sort.Slice(numStr, func(a, b int) bool {
		str1 := numStr[a] + numStr[b]
		str2 := numStr[b] + numStr[a]
		return str1 >= str2
	})

	if numStr[0] == "0" {
		return "0"
	}
	return strings.Join(numStr, "")
}
