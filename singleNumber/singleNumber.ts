function singleNumber(nums: number[]): number {
	if (nums.length == 1) return nums[0]
	let answer = nums[0]
	for (let i = 1; i < nums.length; i++) {
		answer ^= nums[i]
	}
	return answer
};