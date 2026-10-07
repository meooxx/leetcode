function removeDuplicates(nums: number[]): number {
	let i = 0
	for (const n of nums) {
		// 1 1 2 2  2
		if (i < 2 || n > nums[i - 2]) {
			nums[i] = n
			i++
		}
	}
	return i
};