function searchRange(nums: number[], target: number): number[] {

	return [searchTarget(nums, target, false), searchTarget(nums, target, true)]

};

function searchTarget(nums: number[], target: number, forward: boolean = false): number {
	let answer = -1
	let left = 0
	let right = nums.length - 1
	while (left <= right) {
		const mid = ~~((right - left) / 2) + left

		if (nums[mid] === target) {
			answer = mid
			if (forward) {
				left = mid + 1
			} else {
				right = mid - 1
			}


		} else if (nums[mid] > target) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return answer
}