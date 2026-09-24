function maxSubArray(nums: number[]): number {
	let sum = 0
	let maxSum = nums[0]
	for (const n of nums) {
		sum += n
		maxSum = Math.max(sum, maxSum)
		sum = Math.max(0, sum)
	}
	return maxSum
};