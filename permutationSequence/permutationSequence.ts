

function getPermutation(n: number, k: number): string {

 //        0 1 2 3......
	// fact: 1 1 2 6 24 
	// (4,  9)
	// there are (4-1)! permutations in each block
	// 9 / fact[n-1] to got the block  index
	// k % fact[n-1], moving forward to rest of candidatews

	//  1      2        3   4
	// 1234    2134
	// 1243    2143
	// 1324    2314
	// 1342    2341
	// 1423    2413
	// 1432    2431
	const fact = Array(n + 1)
	// candidates
	let nums = Array(n)
	fact[0] = 1
	for (let i = 1; i <= n; i++) {
		nums[i - 1] = i
		fact[i] = fact[i - 1] * i
	}
	let answer = ""
	// convert k to 0 based
	k--
	while (nums.length > 0) {
		const blockIndex = ~~(k / fact[n - 1])
		answer += nums[blockIndex]
		// remove used element
		nums = [...nums.slice(0, blockIndex), ...nums.slice(blockIndex + 1)]
		k %= fact[n-1]
		n -= 1

	}

	return answer
};

getPermutation(3, 3)