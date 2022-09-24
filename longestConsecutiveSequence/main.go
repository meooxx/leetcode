package main

import "fmt"

func main() {
	fmt.Println(longestConsecutive([]int{
		100, 1, 2, 3, 200,
	}))
	fmt.Println(longestConsecutive1([]int{
		100, 1, 2, 3, 200,
	}))
}

func longestConsecutive(nums []int) int {
	mp := map[int]bool{}
	for _, n := range nums {
		mp[n] = true
	}
	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}
	maxLen := 0
	for _, n := range nums {
		if !mp[n-1] {
			y := n + 1
			for {
				if _, ok := mp[y]; ok {
					y++
				} else {
					break
				}

			}
			maxLen = max(maxLen, y-n)
		}
	}
	return maxLen
}

// less runtime leetcode
func longestConsecutive1(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	mp := map[int]bool{}
	maxLen := 0
	for _, n := range nums {
		mp[n] = true
	}

	for _, n := range nums {
		min := n
		max := n
		for mp[min-1] {
			delete(mp, min)
			min--
		}
		for mp[max+1] {
			delete(mp, max)
			max++
		}
		if max-min+1 > maxLen {
			maxLen = max - min + 1
		}
	}
	return maxLen
}
