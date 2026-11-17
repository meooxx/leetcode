function wordBreak(s: string, wordDict: string[]): string[] {
	const result: string[] = []
	const wordSet = new Set(wordDict)
	const next = (str: string, curr: string[]) => {
		if (str.length === 0) {
			result.push(curr.join(' '))
			return
		}
		for (let i = 1; i <= str.length; i++) {
			const w = str.slice(0, i)
			if (wordSet.has(w)) {
				curr.push(w)
				next(str.slice(i), curr)
				curr.pop()
			}
		}

	}
	next(s, [])
	return result
};