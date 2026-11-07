

//   [1,2,1]
//        ^
// 1 from right to left, each element+=preElement
// 2 push 1 after each loop
function getRow(rowIndex: number): number[] {
	const pre = [1]
	if (rowIndex <= 0) return pre

	for (let i = 0; i < rowIndex; i++) {
		for (let j = pre.length - 1; j >= 1; j--) {
			pre[j] += pre[j - 1]
		}
		pre.push(1)
	}
	return pre
};