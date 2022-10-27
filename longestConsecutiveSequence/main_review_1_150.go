package main

func longestConsecutive(nums []int) int {
	mp := map[int]bool{}

	for i := range nums {
		mp[nums[i]] = true
	}
	maxLen := 0
	for i := range nums {
		if !mp[nums[i-1]] {
			min := nums[i]
			len := 0
			for mp[min+1] {
				len++
				min++
			}
			if len > maxLen {
				maxLen = len
			}
		}
	}
	return maxLen

}
func longestConsecutive1(nums []int) int {
	mp := map[int]bool{}
	for i := range nums {
		mp[nums[i]] = true
	}
	maxLen := 0
	for i := range nums {
		minV := nums[i]
		maxV := nums[i]
		for mp[minV-1] {
			delete(mp, minV-1)
			minV--
		}
		for mp[maxV+1] {
			delete(mp, maxV)
			maxV++
		}
		if maxV-minV+1 > maxLen {
			maxLen = maxV - minV + 1
		}
	}
	return maxLen
}
