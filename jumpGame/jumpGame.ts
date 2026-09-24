


function canJump(nums: number[]): boolean {
	if(nums.length === 1) return true
	let farRight = nums[0]
	let start = 0
	while (start < nums.length) {
		farRight = Math.max(farRight, start + nums[start])
		if (start >= farRight) {
			return false
		}

		start++
	}
	return true
};