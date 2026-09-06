

function lengthOfLongestSubstring(s: string): number {
	const mp = new Map<string, number>()
	let left = 0
	let right = 0
	let ml = 0
	// aabacd
	// bbbb
	// baaabca
	// baaabcdea

	for (; right < s.length;right++) {
		
		// If the character has appeared before, move left only when its previous
		// position is still inside the current window.
		// For example, in "bccddadcb", the b at index 0 is outside the window
		// when we encounter the final b, so it does not affect the left boundary.
		if (mp.has(s[right]) && mp.get(s[right])! + 1 > left) {
			left = mp.get(s[right])! + 1
			
		} else {
			ml = Math.max(ml, right - left + 1)
			
		}
		mp.set(s[right], right)

	}
	return ml
};

// console.log(lengthOfLongestSubstring('bbbb'))
console.log(lengthOfLongestSubstring('bccddadcb'))