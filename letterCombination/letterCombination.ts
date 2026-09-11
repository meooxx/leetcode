function letterCombinations1(digits: string): string[] {
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


function letterCombinations(digits: string): string[] {
	const map = new Map([
		["2", ["a", "b", "c"]],
		["3", ["d", "e", "f"]],
		["4", ["g", "h", "i"]],
		["5", ["j", "k", "l"]],
		["6", ["m", "n", "o"]],
		["7", ["p", "q", "r", "s"]],
		["8", ["t", "u", "v"]],
		["9", ["w", "x", "y", "z"]]
	])
	const candidates = digits.split('')
	return candidates.reduce<string[]>((pre: string[], curr: string) => {
		const currArr = map.get(curr)!
		return pre.flatMap((v: string) => {
			return currArr.map((nextV: string) => {
				return v + nextV
			})
		}
		)
	}, [''])
}
