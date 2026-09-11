

function isValid(s: string): boolean {
	const stack: string[] = []
	const map = new Map([
		[')', '('],
		[']', '['],
		['}', '{']
	])
	for (const char of s) {
		switch (char) {
			case '(':
			case '[':
			case '{':
				stack.push(char)
				break
			case ')':
			case ']':
			case '}':
				const last = stack.pop()
				if (last !== map.get(char)) {
					return false
				}
				stack.pop()
		}
	}
	return stack.length === 0

};
