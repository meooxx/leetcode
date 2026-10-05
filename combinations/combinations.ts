

function combine(n: number, k: number): number[][] {
	const condidates = []
	for (let i = 1; i <= n; i++) {
		condidates.push(i)
	}
	const result: number[][] = []

	next(condidates, [], 0, k, result)
	return result

};

function next(candidates: number[], current: number[], start: number, reminder: number, result: number[][]) {
	if (reminder === 0) {
		result.push([...current])
	}
	for (let i = start; i < candidates.length; i++) {
		current.push(candidates[i])
		next(candidates, current, i + 1, reminder - 1, result)
		current.pop()
	}
}