function singleNumber(nums: number[]): number {
	let answer = 0


	for (let i = 0; i < 32; i++) {
		let count = 0
		let mask = 1 << i
		for (const num of nums) {

			if ((num & mask) == mask) count++;

		}
		if (count % 3 == 1) {
			answer |= mask
		}

	}
	return answer
};