function groupAnagrams(strs: string[]): string[][] {

	const wordMap = new Map<string, string[]>()
	for (const str of strs) {
		const key = str.split('').sort((a, b) => a.charCodeAt(0) - b.charCodeAt(0)).join('')
		if (wordMap.has(key)) {
			wordMap.get(key)!.push(str)
		} else {
			wordMap.set(key, [str])
		}

	}
	const answer = Array.from(wordMap.values())

	return answer
};

