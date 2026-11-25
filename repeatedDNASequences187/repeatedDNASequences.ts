function findRepeatedDnaSequences(s: string): string[] {
	const seen = new Set()
	const result: Set<string> = new Set()

	for (let i = 0; i + 10 <= s.length; i++) {
		const seq = s.slice(i, i + 10)
		if (seen.has(seq)) {
			result.add(seq)
		} else {
			seen.add(seq)
		}
	}
	return [...result]
};