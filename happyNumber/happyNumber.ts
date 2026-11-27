function isHappy(n: number): boolean {
	const visited: Set<number> = new Set()
	while (n != 1) {
		let result = 0
		let cur = n
		if (visited.has(n)) return false
		visited.add(n)
		while (cur > 0) {
			result += (cur % 10) * (cur % 10)
			cur = Math.floor(cur / 10)
		}

		n = result


	}
	return true
};