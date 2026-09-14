function removeDuplicates(nums: number[]): number {
	let left = 1
	let right = 1
	while (right < nums.length) {
		if (nums[right] === nums[right - 1]) {
			right++
		} else {
			nums[left] = nums[right]
			left++
			right++
		}

	}
	return left
};