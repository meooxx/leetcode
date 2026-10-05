///  caab,  ab   
//1 统计所有的chars in t, map[a] == 1, map[b]= 1
// count == 3
// 向右侧扫描s, 遇到a|b 或者其他char in t, map[char] -1
// 如果且此时 map[char] > 1, 此时 count -1 , 说明扫到一个有效
// 2 扫描过程中检查  count === 0, 说明包含所有的t了
// 3 收缩window, 如果 while(count === 0) 说明窗口有效
//   如果 map[char] > 0, 说明此窗口缺少char了, 应该count+1了

function minWindow(s: string, t: string): string {
	const mp = new Map<string, number>()
	for (const sub of t) {
		mp.set(sub, (mp.get(sub) ?? 0) + 1)
	}
	let count = t.length
	let end = 0
	let minStr = ''
	let start = 0
	while (end < s.length) {
		const v = mp.get(s[end])
		if (v != undefined) {
			if (v > 0) {
				count--
			}
			mp.set(s[end], v - 1)
		}

		// caba  ab
		while (count === 0) {
			if (minStr === '' || end - start < minStr.length) {
				minStr = s.slice(start, end + 1)
			}
			if (mp.has(s[start])) {
				mp.set(s[start], mp.get(s[start])! + 1)
				if (mp.get(s[start])! > 0) {
					count++
				}
			}
			start++


		}
		end++


	}
	return minStr


};
