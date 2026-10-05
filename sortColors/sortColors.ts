/**
 Do not return anything, modify nums in-place instead.
 */
function sortColors(nums: number[]): void {

	let i = 0
	let start = 0
	let end = nums.length - 1
	while (start < end) {
		if (nums[start] === 0) {
			[nums[i], nums[start]] = [nums[start], nums[i]]
			i++
			start++
		} else if (nums[start] === 2) {
			[nums[start], nums[end]] = [nums[end], nums[start]]
			end--
		} else {
			start++
		}
	}


};