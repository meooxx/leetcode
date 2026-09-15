
function findSubstring(s: string, words: string[]): number[] {
	const wordsMap = new Map<string, number>()
	const result: number[] = []
	for (const word of words) {
		wordsMap.set(word, (wordsMap.get(word) ?? 0) + 1)
	}
	const wordLen = words[0].length
	const usedWord = new Map<string, number>()

	for (let i = 0; i <= s.length - wordLen; i++) {
		let count = 0
		usedWord.clear()
		for (let start = i; start <= s.length - wordLen; start += wordLen) {
			const sub = s.slice(start, start + wordLen)
			if (wordsMap.has(sub)) {
				usedWord.set(sub, (usedWord.get(sub) ?? 0) + 1)
				count++
				if (usedWord.get(sub)! > wordsMap.get(sub)!) {
					break
				}
			} else {
				break
			}
			if (count === words.length) {
				result.push(i)
				break
			}

		}

	}
	return result
};


findSubstring('a', ["a"])