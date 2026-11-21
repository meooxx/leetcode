function findPeakElement(nums: number[]): number {
	// [0]
	if (nums.length === 1) return 0
	// [3,1]
	if (nums[0] > nums[1]) return 0
	// [1,3]
	if (nums[nums.length - 1] > nums[nums.length - 2]) return nums.length - 1

	let l = 0
	let r = nums.length

	while (l <= r) {
		const mid = (l + r) >> 1
		if (nums[mid] > nums[mid - 1] && nums[mid] > nums[mid + 1]) {
			return mid

			//   1 3 [2]  4 5 
		} else if (nums[mid] <= nums[mid - 1]) {
			r = mid - 1
			//  1 [2] 3 
		} else if (nums[mid] <= nums[mid + 1]) {
			l = mid + 1
		}
	}
	return -1
};