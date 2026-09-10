function romanToInt(s: string): number {
	const map = new Map([
		["I", 1],
		["IV", 4],
		["V", 5],
		["IX", 9],
		["X", 10],
		["XL", 40],
		["L", 50],
		["XC", 90],
		["C", 100],
		["CD", 400],
		["D", 500],
		["CM", 900],
		["M", 1000]])
	let answer = 0
	for (let i = 0; i < s.length; i++) {
		if (i + 1 < s.length) {
			const sym = s.slice(i, i + 2)
			if (map.has(sym)) {
				answer += map.get(sym)!
				i++
				continue
			}
		}
		answer += map.get(s.slice(i, i + 1))!
	}
	return answer
};
