function compareVersion(version1: string, version2: string): number {
	const v1 = version1.split('.')
	const v2 = version2.split('.')
	let i = 0
	let j = 0
	while (i < v1.length && j < v2.length) {
		const n1 = parseInt(v1[i])
		const n2 = parseInt(v2[i])
		if (n1 < n2) return -1
		if (n1 > n2) return 1
		i++
		j++
	}
	while (i < v1.length) {
		const n1 = parseInt(v1[i])
		i++
		if (n1 == 0) continue
		if (n1 > 0) return 1
	}
	while (j < v2.length) {
		const n1 = parseInt(v2[j])
		j++
		if (n1 == 0) continue
		if (n1 > 0) return -1
	}
	return 0

};