// 1 - 9
// 你试着随便切这 3 段，看看最长的一段是多少：
// 切法 A（极不均匀）：第一段 1 厘米，第二段 1 厘米，第三段 7 厘米  最长的一段是 7。
// 切法 B（稍微均匀）：第一段 2 厘米，第二段 3 厘米，第三段 4 厘米  最长的一段是 4.
// 切法 C（绝对平均）：第一段 3 厘米，第二段 3 厘米，第三段 3 厘米（即数字排成[1, 4, 7, 10]）
// 不能了！也就是均匀排列下间隔最小就是3了
function maximumGap(nums: number[]): number {

	if (nums.length < 2) return 0
	const smallest = Math.min(...nums)
	const largest = Math.max(...nums)
	// gap is 0 as every num is same
	if (smallest === largest) return 0
	// bucketSize 3/2 ceil to 2
	// there is no fractional bucket size
	const bucketSize = Math.ceil((largest - smallest) / (nums.length - 1))
	// n numbers have n-1 gaps, maxGap is greater 
	// or equal to avg gap
	let minArr = Array(nums.length).fill(Infinity)
	let maxArr = Array(nums.length).fill(-Infinity)


	// each bucket only keep min and max values
	for (const n of nums) {
		// calculate the idx
		const idx = Math.floor((n - smallest) / bucketSize)
		minArr[idx] = Math.min(minArr[idx], n)
		maxArr[idx] = Math.max(maxArr[idx], n)
	}
	// must have a value after iterate through the arr
	let preMax = maxArr[0] // the first bucket

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