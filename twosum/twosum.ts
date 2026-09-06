function twoSum(nums: number[], target: number): number[] {

	const mp = new Map<number, number>()
	for (let i = 0; i < nums.length; i++) {
		const surplus = target - nums[i]
		if (mp.has(nums[i])) {
			return [i, mp.get(nums[i]) as number]
		} else {
			mp.set(surplus, i)

		}
	}

	return []



};

console.log(twoSum([1, 2, 4, 6], 3))
