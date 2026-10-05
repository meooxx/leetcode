function simplifyPath(path: string): string {
	const pArr = path.split('/')
	const result: string[] = []

	for (const sep of pArr) {
		if (sep == '.' || sep == '') {
			continue
		}
		if (sep === '..') {
			result.pop()
			continue
		}
		result.push(sep)
	}
	return '/' + result.join('/')



};


simplifyPath("/a/b/c/../../d")