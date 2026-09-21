function combinationSum2(candidates: number[], target: number): number[][] {
	candidates.sort((a, b) => a - b);
	const answers: number[][] = []
	next(candidates, 0, target, [], answers)
	return answers
};

function next(candidates: number[], position: number, reminder: number, nums: number[], result: number[][]) {
	if (reminder === 0) {
		result.push([...nums])
		return
	}
	for (let i = position; i < candidates.length; i++) {

		const n = candidates[i]
		if (reminder - n < 0) return
		if (i > position && i < candidates.length) {
			if (candidates[i] === candidates[i - 1]) continue
		}


		nums.push(n)
		next(candidates, i + 1, reminder - n, nums, result)
		nums.pop()

	}


}