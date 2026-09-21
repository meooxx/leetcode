function firstMissingPositive(nums: number[]): number {

	for (let i = 0; i < nums.length; i++) {

		// !! wathc out 
		// 																-1                 3		
		// [nums[i], nums[nums[i]-1]] = [nums[nums[i]-1], nums[i]]
		// nums[i] = -1 => nums[nums[i]-1]= nums[-2] ❌
		// 
		while (nums[i] > 0 && nums[i] < nums.length && nums[nums[i] - 1] !== nums[i]) {
			[nums[nums[i] - 1], nums[i]] = [nums[i], nums[nums[i] - 1]]
		}
	}
	let i = 0
	for (; i < nums.length; i++) {
		if (nums[i] != i + 1) return i + 1
	}
	return i + 1

};

firstMissingPositive([1, 1])