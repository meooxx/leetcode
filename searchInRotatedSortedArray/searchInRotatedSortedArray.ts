function search(nums: number[], target: number): number {
	let left = 0
	let right = nums.length - 1
	let midNum
	// [3, 1, 2]
	// [5 6 7 1 2]
	while (left <= right) {
		let mid = ~~((right - left) / 2) + left
		if (nums[mid] >= nums[0] && target >= nums[0] || target < nums[0] && nums[mid] < nums[0]) {
			midNum = nums[mid]
		} else {
			if (target >= nums[0]) {
				midNum = Infinity
			} else {
				midNum = -Infinity
			}
		}
		if (target === midNum) {
			return mid
		} else if (target < midNum) {
			right = mid - 1
		} else {
			left = mid + 1
		}

	}
	return -1
};