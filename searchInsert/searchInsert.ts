function searchInsert(nums: number[], target: number): number {
	let left = 0
	let right = nums.length - 1
	while (left <= right) {
		const mid = ~~((right - left) / 2) + left
		if (target === nums[mid]) {
			return mid
		} else if (target < nums[mid]) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
};


