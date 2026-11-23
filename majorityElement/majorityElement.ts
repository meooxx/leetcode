function majorityElement(nums: number[]): number {

	let count = 0
	let pre = nums[0]
	for (let i = 0; i < nums.length; i++) {
		const n = nums[i]
		if (pre == n) count++
		else count--
		if (count == 0) pre = nums[i + 1]
	}
	return pre
};