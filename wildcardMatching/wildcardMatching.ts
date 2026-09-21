// star matches any characters 0 - many times
// abc   a*c
// a vs a
// b vs * skip b while recording pointers of s, p
// b vs  c mismatch , 
// back to star, skipping b and trying to match c vs *
// 
function isMatch(s: string, p: string): boolean {

	let i = 0;
	let j = 0;
	let sPre = -1
	let starPre = -1
	for (; i < s.length;) {
		if (j < p.length && (s[i] === p[j] || p[j] === '?')) {
			i++
			j++
		} else if (j < p.length && p[j] == '*') {
			sPre = i
			starPre = j
			j++
		} else if (starPre !== -1) {
			i = sPre + 1
			j = starPre + 1
			sPre = i

		} else {
			return false
		}
	}
	while (j < p.length && p[j] === '*') j++


	return j === p.length
};

isMatch('abc', 'a*c')