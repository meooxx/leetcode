function threeSumClosest(nums: number[], target: number): number {

	nums.sort((a, b) => {
		return a - b
	})
	let closest = Number.MAX_SAFE_INTEGER

	for (let i = 0; i < nums.length - 2; i++) {
		if (i > 0 && nums[i] === nums[i - 1]) {
			continue
		}
		let j = i + 1
		let k = nums.length - 1
		while (j < k) {
			const sum = nums[i] + nums[j] + nums[k]
			if (Math.abs(sum - target) < Math.abs(closest - target)) {
				closest = sum
			}
			if (sum < target) {
				j++
				while (j < k && nums[j] === nums[j - 1]) {
					j++
				}
			} else {
				k--
				while (k > j && nums[k] === nums[k + 1]) {
					k--
				}
			}



		}
	}
	return closest

};