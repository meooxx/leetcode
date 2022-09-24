function isPalindrome(x: number): boolean {
	let n = 0;
	let rest = x;
	for (; rest > 0; ) {
		n = n * 10 + (rest % 10);
		rest = Math.floor(rest / 10);
	}
	return n === x;
}
