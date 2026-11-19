//https://leetcode.com/problems/maximum-product-subarray/
// from left to right, find the maximum
// or find it from right to left
// only negative num matters
//  1,2 , 3, 4, 5, -1, 6
// or 1, -2, 3,4, 6, 6, 7
function maxProduct(nums: number[]): number {
	let prod = 1
	let answer = -Infinity
	for (const num of nums) {
		prod *= num
		if (prod == 0) {
			prod = 1
		}
		answer = Math.max(answer, prod)
	}
	prod = 1
	for (let i = nums.length - 1; i >= 0; i--) {
		prod *= nums[i]
		if (prod === 0) {
			prod = 1
		}
		answer = Math.max(answer, prod)
	}
	return answer



};