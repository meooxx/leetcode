function search1(nums: number[], target: number): number {
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


function search(nums: number[], target: number): number {
	let left = 0
	let right = nums.length - 1
  // two sorted sequence
	// 4 5 6 7 1 2 3
	// case1:  left ele is less than or equal to mid, 
	// and target is in nums[left] ~ nums[mid]
	//  4 5  1 2 3
	// case2: left is great than mid
	// and target is in nums[mid] ~ nums[right] 
	while (left <= right) {

		let mid = ~~((right - left) / 2) + left

		if (nums[mid] === target) {
			return mid
		} else if (nums[left] <= nums[mid]) {
			if (target >= nums[left] && target < nums[mid]) {
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
	return -1
};