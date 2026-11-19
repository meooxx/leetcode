function findMin(nums: number[]): number {
	let l = 0
	let r = nums.length - 1
	// 3 1 2  or
	// 1 2 3   
	while (l < r) {
		const mid = (r + l) >> 1
		if (nums[mid] > nums[r]) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return nums[l]
}