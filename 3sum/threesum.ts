function threeSum(nums: number[]): number[][] {
	nums.sort((a, b) => {
		return a - b
	})
	const answer: number[][] = []
	for (let i = 0; i < nums.length - 2; i++) {
		if (nums[i] > 0) {
			return answer
		}
		if (i > 0 && nums[i] === nums[i - 1]) {
			continue
		}
		let j = i + 1
		let k = nums.length - 1;
		while (j < k) {

			if (nums[i] + nums[k] + nums[j] === 0) {
				answer.push([nums[i], nums[j], nums[k]])
				j++
				k--
				while (j < k && nums[j] === nums[j - 1]) {
					j++
				}
				while (k > j && nums[k] === nums[k + 1]) {
					k--
				}
			} else if (nums[i] + nums[j] + nums[k] > 0) {
				k--
			} else {
				j++
			}

		}

	}
	return answer

};

threeSum([-2, 0, 1, 1, 2])