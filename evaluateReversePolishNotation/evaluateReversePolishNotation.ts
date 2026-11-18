function evalRPN(tokens: string[]): number {
	const stack: number[] = []

	for (const str of tokens) {
		switch (true) {

			case str === "+" || str === '-' || str === '/' || str === '*': {
				const x = stack.pop()!
				const y = stack.pop()!
				if (str === '+') stack.push(x + y)
				if (str === '-') stack.push(y - x)
				if (str === '/') stack.push(~~(y / x))
				if (str === '*') stack.push(x * y)
				break
			}

			default:
				stack.push(parseInt(str))

		}
	}
	return stack[0]
};