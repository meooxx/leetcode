

function jump(nums: number[]): number {
	let maxRight = 0
	let end = 0
	let step = 0
	for (let i = 0; i < nums.length - 1; i++) {
		// find the most far right boundary before reaching end
		// add one steps when reaching the end.
		if (nums[i] + i > maxRight) {
			maxRight = nums[i] + i
		}
		// at the first, end is 0, meaning start the first step
		if (i === end) {
			step++
			end = maxRight
		}
	}
	return step

};