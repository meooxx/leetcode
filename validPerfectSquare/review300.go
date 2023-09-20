package main


func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	start := 1
	end  := num / 2 // 4
	for start <= end {
		mid := (end - start) / 2 + start
		if mid * mid == num {
			return true
		} else if mid * mid > num {
				end = mid  - 1 
		} else {
			start = mid + 1
		}
	} 
	return false
}