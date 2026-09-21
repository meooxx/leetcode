


function countAndSay(n: number): string {
	//         		 1   1
	// 		    		11   2
	// 		    		21   3
	// 	    		1211   4
	let answer = '1'
	if (n <= 1) {
		return answer
	}
	for (let i = 1; i < n; i++) {
		let temStr = answer[0]
		let count = 1
		for (let start = 1; start < answer.length; start++) {
			if (answer[start] != answer[start - 1]) {
				temStr += `${count}${answer[start - 1]}`
				count = 1
			} else {
				count++
			}
		}
		temStr += `${count}${answer[answer.length - 1]}`
		answer = temStr
	}
	return answer
};