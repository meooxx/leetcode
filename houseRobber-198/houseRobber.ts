function rob(nums: number[]): number {
	let prepre = 0
	let pre = 0
	let maxMoney = 0
	for (const n of nums) {
		maxMoney = Math.max(pre, prepre + n)
		prepre = pre
		pre = maxMoney
	}
	return maxMoney
};