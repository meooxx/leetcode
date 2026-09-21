

function permute(nums: number[]): number[][] {
	if (nums.length === 0) return []
	if (nums.length === 1) return [nums]
	const answers: number[][] = []
	next(nums, [], answers)
	return answers
};

function next(nums: number[], curr: number[], result: number[][]) {
	if (nums.length === 0) {
		result.push([...curr])
		return
	}
	for (let i = 0; i < nums.length; i++) {
		curr.push(nums[i])
		next([...nums.slice(0, i), ...nums.slice(i + 1)], curr, result)
		curr.pop()
	}
}