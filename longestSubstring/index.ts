function lengthOfLongestSubstring(s: string): number {
	if (s.length <= 1) {
		return s.length;
	}
	let position = 0;
	let r = 0;
	const m = new Map();
	for (let i = 0; i < s.length; i++) {
		if (m.has(s[i])) {
			if (i - position > r) {
				r = i - position;
			}
			// = for aaaab
			if (m.get(s[i]) >= position) {
				position = m.get(s[i]) + 1;
			}
		}
		m.set(s[i], i);
	}
	if (s.length - position > r) {
		r = s.length - position;
	}
	return r;
}

lengthOfLongestSubstring('aaaaabbbbbb');

