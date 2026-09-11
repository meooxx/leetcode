function generateParenthesis(n: number): string[] {
	if (n === 1) {
		return ["()"]
	}
	const result: string[] = []
	generateImpl('', n, n, result)
	return result

};

function generateImpl(s: string, left: number, right: number, result: string[]) {
	if (left === 0 && right === 0) {
		result.push(s)
		return
	}
	if (left > 0) {
		generateImpl(s + '(', left - 1, right, result)
	}
	if (right > left) {
		generateImpl(s + ')', left, right - 1, result)
	}

}