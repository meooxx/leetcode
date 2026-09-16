

// at the first:  push -1, for the () vail case
// if the first element is ), 
// pop stack and push this index,  which becomes the pre position of the first valid (
// -1 (
// -1,   )((( => [0, 1, 2,3 ] 
//       ))))) => [1] => 3-1
// if the number of ) is greater than that of (
// then stack if only one index left 
// which is the last index of )

function longestValidParentheses(s: string): number {
	const stack: number[] = [-1]
	let maxLen = 0
	// -1 )
	for (let i = 0; i < s.length; i++) {
		if (s[i] === '(') {
			stack.push(i)
		} else {
			stack.pop()
			if (stack.length === 0) {
				stack.push(i)
				continue
			}
			const lastIndex = stack[stack.length - 1]
			maxLen = Math.max(maxLen, i - lastIndex)
		}
	}
	return maxLen



};