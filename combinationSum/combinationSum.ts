

function combinationSum(candidates: number[], target: number): number[][] {
	const answers: number[][] = []
	candidates.sort((a, b) => a - b)
	next(candidates, target, 0, [], answers)
	return answers
};

function next(candidates: number[], reminder: number, i: number, nums: number[], result: number[][]) {
	if (reminder < 0) return
	if (reminder === 0) {
		result.push([...nums])
		return
	}


	for (; i < candidates.length; i++) {
		if (reminder - candidates[i] < 0) return
		nums.push(candidates[i])
		next(candidates, reminder - candidates[i], i, nums, result)
		nums.pop()

	}

}

combinationSum([2, 3, 5], 7)