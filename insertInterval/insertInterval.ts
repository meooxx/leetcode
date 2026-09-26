
//  [4, 5], [6,9], [10,13], [14, 15], [7,12](new)
//  1 append all arrs where curr[1] < newInterval[1]
//  2 first modify newInterval [6, 12] as 7<=9 && 6<=7
//  break at [10, 13] as 12 < 13, 
//  while merge this last eligible interval
//  [6, 13], also because 12 < 13
// append rest of arrs
function insert(intervals: number[][], newInterval: number[]): number[][] {
	const answer: number[][] = []
	let start = 0
	while (start < intervals.length) {
		if (intervals[start][1] < newInterval[0]) {
			answer.push(intervals[start])
		} else {
			break
		}
		start++
	}

	while (start < intervals.length) {
		newInterval[0] = Math.min(newInterval[0], intervals[start][0])
		//             last element is eligible for merging
		// 3,5    6,8, 7 10  [4,9] => [3, 10]
		if (newInterval[1] < intervals[start][1]) {
			// break at the last emelent position 
			// where newInterval[1]< curr[1]
			// while newInterval[1] > curr[0] 
			if (newInterval[1] >= intervals[start][0]) {
				newInterval[1] = intervals[start][1]
				start++
			}
			break
		}
		start++

	}
	answer.push(newInterval, ...intervals.slice(start))
	return answer
};