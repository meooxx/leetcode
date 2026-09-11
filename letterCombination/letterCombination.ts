function letterCombinations(digits: string): string[] {
	const map = new Map<string, string>(
		[
			['2', 'abc'],
			['3', 'def'],
			['4', 'ghi'],
			['5', 'jkl'],
			['6', 'mno'],
			['7', 'pqrs'],
			['8', 'tuv'],
			['9', 'wxyz']]
	)
	const candidates: string[] = [

	]
	for (const digit of digits) {
		candidates.push(map.get(digit)!)
	}
	return generateCombination(candidates)
}

function generateCombination(candidates: string[]): string[] {
	const result: string[] = []
	if (candidates.length === 0) {
		return result
	}
	if (candidates.length === 1) {
		return [...candidates[0]]
	}
	const first = candidates[0]
	const others = generateCombination(candidates.slice(1))
	for (const char of first) {
		for (const str of others) {
			result.push(char + str)
		}
	}
	return result
}