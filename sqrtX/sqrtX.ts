
// search 1    9, 5
//        1    5,  3
//        3    5, 4
//        3    4, 3
//        4    4, 
// move left = mid + 1 when mid <= x/mid
// to make sure left always keep the exact or nearest value +1


function mySqrt(x: number): number {
	let left = 1
	let right = x
	while (left <= right) {
		const mid = ~~((right - left) / 2) + left
		if (mid <= x / mid) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right


};

