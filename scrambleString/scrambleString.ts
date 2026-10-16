const cache: Record<string, number> = {}

function isScramble(s1: string, s2: string): boolean {



	if (s1.length != s2.length) return false
	if (s1 === s2) return true
	if (cache[s1 + s2] === 1) return true
	else if (cache[s1 + s2] === -1) return false
	const mp: number[] = Array(26).fill(0)
	for (let i = 0; i < s1.length; i++) {
		mp[s1[i].charCodeAt(0) - 97]++
		mp[s2[i].charCodeAt(0) - 97]--
	}
	for (const n of mp) {
		if (n !== 0) {
			cache[s1 + s2] = -1
			return false
		}
	}

	// abc
	for (let i = 1; i < s1.length; i++) {
		if (isScramble(s1.slice(0, i), s2.slice(0, i))
			&& isScramble(s1.slice(i), s2.slice(i))
		) {
			cache[s1 + s2] = 1
			return true
		}
		// abc
		if (isScramble(s1.slice(0, i), s2.slice(s2.length - i))
			&& isScramble(s1.slice(i), s2.slice(0, s2.length-i))
		) {

			cache[s1 + s2] = 1
			return true
		}
	}
	cache[s1 + s2] = -1
	return false
};

isScramble('abc', 'bca')