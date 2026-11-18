function maxPoints(points: number[][]): number {
	//  y1 = kx1 + b
	//  y2 = kx2 + b
	// k = (y1-y2)/(x1-x2)
	let maxLen = 0
	if (points.length == 1) return 1
	const gcd = (a: number, b: number): number => {
		a = Math.abs(a)
		b = Math.abs(b)
		// 9 6
		if (a == 0 || b === 0) return 0
		while (b !== 0) {
			[a, b] = [b, a % b]
		}
		return a


	}
	for (let i = 0; i < points.length; i++) {
		const map = new Map

		for (let j = i + 1; j < points.length; j++) {
			let dy = points[i][1] - points[j][1]
			let dx = points[i][0] - points[j][0]
			let k = ''
			if (dx == 0) {
				k = `dy/0`
			} else if (dy == 0) {
				k = '0/dx'
			}
			else {
				// dy = 1, dx = -1,  -1/1
				// dy = -1, dx = 1,  -1/1
				if (dx < 0) {
					dx = -dx
					dy = -dy
				}
				const g = gcd(dy, dx)
				k = `${dy / g}/${dx / g}`
			}


			map.set(k, (map.get(k) || 0) + 1)
			// 1,2 -> 2,4
			// 1,2 -> 3,6
			// (1,2), (2,4), (3,6) three points
			maxLen = Math.max(maxLen, map.get(k) + 1)
		}
	}
	return maxLen
};