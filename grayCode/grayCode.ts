function grayCode(n: number): number[] {
	//  00
	//  10
	//  11
	//  01
	// 101
	// 111
	// 110
	const list: number[] = [0]

	for (let i = 0; i < n; i++) {
		const size = list.length
		for (let j = size - 1; j >= 0; j--) {
			list.push(list[j] | (1 << i))
		}
	}
	return list

};