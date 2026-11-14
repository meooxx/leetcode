

function partition(s: string): string[][] {
	const isPalindrom = (str: string): boolean => {
		if (str.length <= 1) return true
		let i = 0
		let j = str.length - 1
		while (i < j) {
			if (str[i] !== str[j]) return false
			i++
			j--
		}
		return true

	}
	const result: string[][] = []
	const nextPartition = (str: string, curr: string[]) => {
		if (str.length == 0) {
			result.push([...curr])
			return
		}
		for (let i = 1; i <= str.length; i++) {
			const left = str.slice(0, i)
			if (isPalindrom(left)) {
				curr.push(left)
				nextPartition(str.slice(i), curr)
				curr.pop()
			}
		}
	}
	nextPartition(s, [])
	return result


};
