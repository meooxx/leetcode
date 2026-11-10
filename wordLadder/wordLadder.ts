function ladderLength(beginWord: string, endWord: string, wordList: string[]): number {

	const map = new Map<string, boolean>()
	wordList.forEach(w => map.set(w, true))
	let answer = 0
	if (!map.has(endWord)) return 0
	map.delete(beginWord)
	const queue = [beginWord]
	while (queue.length > 0) {
		let size = queue.length
		answer++

		while (size-- > 0) {
			const top = queue.shift()!
			for (const v of [...map.keys()]) {
				if (isNeigbour(v, top)) {
					if (v == endWord) {
						answer++

						return answer
					}
					queue.push(v)
					map.delete(v)
				}
			}

		}

	}
	return 0
};

function isNeigbour(a: string, b: string): boolean {
	if (a == '') {
		if (a == b) return true
		return false
	}
	let diff = 0
	for (let i = 0; i < a.length; i++) {
		if (a[i] !== b[i]) {
			diff++
			if (diff > 1) {
				return false
			}
		}
	}
	return diff == 1

}