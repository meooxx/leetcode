// https://leetcode.com/problems/bitwise-and-of-numbers-range/solutions/4759869/beat-100-bitwise-operation/

// 
//111  000 
//111  111
// nubers between 0x111 000 to 0x111 111 result in 0 by applying AND
// only 111 & 111 contributes to the result
// so we just find the common prefixes
// and make shifts
function rangeBitwiseAnd(left: number, right: number): number {
	if (right == 0) return 0
	let bits = 0
	while (left != right) {
		left >>= 1
		right >>= 1
		bits++
	}
	return left << bits
};


