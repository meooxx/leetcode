package main

func minPatches(nums []int, n int) int {
	missing := 1
	added := 0
	i := 0
	// [1-8)
	//      if 7, will extend to 1-15)
	//      if 12, will not extend to 1-20 casuse [1-8), [12-20) lack of (8-12), so we should add 8 to extend to [1, 16)
	for missing < n {
		if i < len(nums) && nums[i] <= missing {
			missing += nums[i]
			i++
		} else {
			missing += missing
			added++
		}

	}

	return added

}
