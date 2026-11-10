
// 1 bfs: construct structure blow
//          layer1     layer 2         layer3
//   hit     hot        [lot, dot ]    [dog, log]    cog
//

// 2 dfs: start from endWord
// find  neigbour of anyword in pre layer to get a possible answer
//   [cog]
//   [dog, cog], [log, cog]
//   [dot, dog, cog], [lot, log, cog]
//   ....

function findLadders(beginWord: string, endWord: string, wordList: string[]): string[][] {

	const map = new Map<string, boolean>()
	wordList.forEach(w => map.set(w, true))
	if (!map.has(endWord)) return []
	const nodes: string[][] = []
	map.delete(beginWord)

	let queue = [beginWord]
	let found = false
	while (queue.length > 0 && !found) {
		nodes.push([...queue])
		let size = queue.length
		while (size-- > 0) {
			const top = queue.shift()!
			for (const v of [...map.keys()]) {
				if (isNeigbour(v, top)) {
					if (v == endWord) {
						found = true
					}
					queue.push(v)
					map.delete(v)
				}
			}
		}

	}
	let result = [[endWord]]
	for (let i = nodes.length - 1; i >= 0; i--) {
		let size = result.length
		// test each word in nodes[i]
		for (const word of nodes[i]) {
			// compare with first word of preWords
			// dot vs [dog, cog]
			//     ...
			for (let i = 0; i < size; i++) {
				const preWords = result[i]
				const oneWord = preWords[0]
				if (isNeigbour(word, oneWord)) {
					result.push([word, ...preWords])

				}
			}
		}

		result = result.slice(size)
	}
	return result
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