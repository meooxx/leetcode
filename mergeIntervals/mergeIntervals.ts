
function merge(intervals: number[][]): number[][] {
	if(intervals.length < 2) return intervals
	intervals.sort((a,b) => a[0] - b[0])
	const answers: number[][] = [intervals[0]]

	for (let i = 1; i < intervals.length; i++) {
		const pre = answers[answers.length - 1]
		const curr = intervals[i]
		if (pre[1] >= curr[0]) {
			pre[1] = Math.max(pre[1], curr[1])
		} else {
			answers.push(curr)
		}
	}
	return answers

};