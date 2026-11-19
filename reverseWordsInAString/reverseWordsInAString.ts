function reverseWords(s: string): string {
	if (s.length === 0) return ''
	return s.split(' ').map(w => w.trim()).filter(Boolean).reverse().join(' ')
};