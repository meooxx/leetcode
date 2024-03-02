package main

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	windowWidth := valueDiff + 1 // incase of valueDiff == 0
	bucket := map[int]int{}
	abs := func(a int) int {
		if a > 0 {
			return a
		}
		return -a
	}
	for i := range nums {
		bucketIndex := nums[i] / windowWidth
		// [-3, 3], 2 , 4; if -3/4 and 3 /4 both map  to the zero index
		// and it returns the wrong anwser true.
		// we  should detect the valDiff if nums[i] is negative
		// so we subtract 1 to go into else statement to check valDiff
		if nums[i] < 0 {
			bucketIndex--
		}
		if _, ok := bucket[bucketIndex]; ok {
			return true
		} else {
			if next, ok := bucket[bucketIndex+1]; ok && abs(next-nums[i]) <= valueDiff {
				return true
			}
			if pre, ok := bucket[bucketIndex-1]; ok && abs(nums[i]-pre) <= valueDiff {
				return true
			}
		}
		bucket[bucketIndex] = nums[i]

		if i >= indexDiff {
			index := nums[i-indexDiff] / windowWidth
			if nums[i-indexDiff] < 0 {
				index--
			}
			delete(bucket, index)
		}
	}
	return false
}
