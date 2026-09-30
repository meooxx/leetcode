



function isNumber(s: string): boolean {
	// (+-)111.222e(-+)2
	let seenSign = false
	let seenNumber = false
	let seenE = false
	let seenNumberAfterE = false
	let seenPoint = false
	s = s.toLowerCase()
	for (let i = 0; i < s.length; i++) {
		switch (true) {
			case s[i] >= '0' && s[i] <= '9':
				seenNumber = true
				seenNumberAfterE = true
				break
			case s[i] == '+' || s[i] == '-':
				if (i != 0 && s[i - 1] != 'e') {
					return false
				}
				seenSign = true
				seenNumberAfterE = false
				break
			case s[i] == '.': {
				if (seenE || seenPoint) return false
				seenPoint = true
				break
			}
			case s[i] === 'e': {
				if (!seenNumber || seenE) return false
				seenE = true
				seenNumberAfterE = false
				break
			}
			default: {
				return false
			}
		}
	}
	return seenNumber && seenNumberAfterE

};