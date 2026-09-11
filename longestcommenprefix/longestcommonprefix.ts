function longestCommonPrefix(strs: string[]): string {
	let prefix = ''
	let answer = ''
	let len = 1
	while (len <= strs[0].length) {
		prefix = strs[0].slice(0, len)
		for (const str of strs) {
			if (prefix.length > str.length) {
				return answer
			}
			const sub = str.slice(0, prefix.length)
			if (sub != prefix) {
				return answer
			}

		}
		answer = prefix
		len++
	}
	return answer

};
