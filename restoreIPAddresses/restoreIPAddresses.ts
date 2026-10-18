function restoreIpAddresses(s: string): string[] {
	const result: string[] = []
	parse(s, [], result)
	return result
};

function isValid(s: string): boolean {
	if (s[0] == '0') {
		return s.length === 1
	}
	if (s.length == 3) {
		return !isNaN(Number(s)) ? Number(s) < 256 : false
	}
	return true
}

function parse(s: string, current: string[], result: string[]) {
	if (s === '') {
		if (current.length === 4) {
			result.push(current.join('.'))
			return
		} else {
			return
		}
	}
	if (isValid(s[0])) {
		current.push(s[0])
		parse(s.slice(1), current, result)
		current.pop()
	}
	if (s.length >= 2 && isValid(s.slice(0, 2))) {
		current.push(s.slice(0, 2))
		parse(s.slice(2), current, result)
		current.pop()
	}
	if (s.length >= 3 && isValid(s.slice(0, 3))) {
		current.push(s.slice(0, 3))
		parse(s.slice(3), current, result)
		current.pop()
	}


}
