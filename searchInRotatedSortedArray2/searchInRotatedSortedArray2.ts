


function search(nums: number[], target: number): boolean {
	let left = 0
	let right = nums.length - 1
	while (left <= right) {
		while (left < right && nums[left] == nums[left + 1]) {
			left++
		}
		while (left < right && nums[right] === nums[right - 1]) {
			right--
		}
		let mid = ~~((right - left) / 2) + left
		if (nums[mid] === target) {
			return true
			//  9 10 [12] 4 6 8
		} else if (nums[mid] >= nums[left]) {
			if (target < nums[mid] && target >= nums[left]) {
				right = mid - 1
			} else {
				left = mid + 1
			}

		} else {
			if (target > nums[mid] && target <= nums[right]) {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

	}
	return false

};