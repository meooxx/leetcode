function subsets(nums: number[]): number[][] {
	let n = 0
	const result: number[][] = []

	while (n <= nums.length) {
		subsetsImpl(nums, 0, [], n, result)
		n++
	}
	return result
};


function subsetsImpl(nums: number[], start: number, current: number[], n: number, result: number[][]) {

	if (n == 0) {
		result.push([...current])
		return
	}

	for (let i = start; i < nums.length; i++) {
		current.push(nums[i])
		subsetsImpl(nums, i + 1, current, n - 1, result)
		current.pop()
	}

}