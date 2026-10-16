function subsetsWithDup(nums: number[]): number[][] {
	nums.sort()
	const result: number[][] = []
	next(nums, 0, [], result)
	return result
};

function next(nums: number[], start: number, current: number[], result: number[][]) {
	if (start > nums.length) {
		return
	} else {
		result.push([...current])
	}
	for (let i = start; i < nums.length; i++) {
		if (i > start && nums[i] == nums[i - 1]) continue
		current.push(nums[i])
		next(nums, i + 1, current, result)
		current.pop()
	}
}