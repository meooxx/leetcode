package main

// 1 = 1
// 4 = 1 + 3
// 9 = 1 + 3 + 5
// ....
func isPerfectSquare0(num int) bool {
	n := 1
	for n > 0 {
		num -= n
		n += 2

	}
	return n == 0
}

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	l, r := 1, num/2
	for l <= r {
		mid := (r-l)/2 + l
		squareMid := mid * mid
		if squareMid == num {
			return true
		} else if squareMid < num {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}
