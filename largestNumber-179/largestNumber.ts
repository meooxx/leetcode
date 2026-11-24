function largestNumber(nums: number[]): string {
	const strNums = nums.map(String)
	strNums.sort((a, b) => {
		if (a + b == b + a) return 0
		else if (a + b > b + a) return -1
		return 1
	})
	if (strNums[0] === '0') return '0'
	return strNums.join('')
};