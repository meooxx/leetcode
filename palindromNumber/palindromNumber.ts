function isPalindrome(x: number): boolean {
	if (x < 0) {
		return false
	}
	let y = 0
	while (x > 0) {
		y = y * 10 + x % 10
		x = ~~(x / 10)
	}
	return y === x
};