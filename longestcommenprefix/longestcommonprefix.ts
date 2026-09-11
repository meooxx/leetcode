function longestCommonPrefix1(strs: string[]): string {
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

function longestCommonPrefix(strs: string[]): string {
	let prefix = strs[0]
	for (let i = 0; i < prefix.length; i++) {
		for (const str of strs) {
			if (str[i] != prefix[i]) {
				return prefix.slice(0, i)
			}
		}
	}
	return prefix
}