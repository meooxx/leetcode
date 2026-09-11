function fourSum(nums: number[], target: number): number[][] {
	nums.sort((a, b) => a - b)
	const answer: number[][] = []
	for (let i = 0; i < nums.length - 3; i++) {
		if (i > 0 && nums[i] === nums[i - 1]) {
			continue
		}
		for (let j = i + 1; j < nums.length - 2; j++) {
			if (j > i + 1 && nums[j] === nums[j - 1]) {
				continue
			}
			let k = j+1;
			let l = nums.length - 1
			while (k < l) {
				const sum = nums[i] + nums[j] + nums[k] + nums[l]
				if (sum === target) {
					answer.push([nums[i], nums[j], nums[k], nums[l]])
					k++
					l--
					while (k < l && nums[k] === nums[k - 1]) {
						k++
					}
					while (l > k && nums[l] === nums[l + 1]) {
						l--
					}
				} else if (sum > target) {
					l--
				} else {
					k++
				}
			}
		}
	}
	return answer
};