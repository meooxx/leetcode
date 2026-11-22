function maximumGap(nums: number[]): number {

	if (nums.length < 2) return 0
	const smallest = Math.min(...nums)
	const largest = Math.max(...nums)
	if (smallest === largest) return 0
	// bucketSize 3/2 ceil to 2
	const bucketSize = Math.ceil((largest - smallest) / (nums.length - 1))
	let minArr = Array(nums.length).fill(Infinity)
	let maxArr = Array(nums.length).fill(-Infinity)

	for (const n of nums) {
		const idx = Math.floor((n - smallest) / bucketSize)
		minArr[idx] = Math.min(minArr[idx], n)
		maxArr[idx] = Math.max(maxArr[idx], n)
	}
	// must have a value after iterate through the arr
	let preMax = maxArr[0]

	let answer = -Infinity

	//  3-9     10-16
	// use min num in next bucket - last maxVal
	for (let i = 1; i < nums.length; i++) {
		if (minArr[i] === Infinity) continue
		answer = Math.max(answer, minArr[i] - preMax)
		preMax = maxArr[i]
	}
	return answer
};