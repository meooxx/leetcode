function longestConsecutive(nums: number[]): number {
	const nSet = new Set(nums)
	let answer = 0
	// iterate through set, avoiding duplication
	for (const n of nSet) {

		// skip repetition
		// 1 -> 5, then 2 , skip2 because of nSet.has(2-1)
		if (!nSet.has(n - 1)) {
			let step = 1
			while (nSet.has(n + step)) {
				step++
			}
			answer = Math.max(step, answer)
		}

	}
	return answer
};