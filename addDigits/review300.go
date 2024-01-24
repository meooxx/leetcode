package main

//	n = d0 + d1 * 10 + d2 * 100 ... dk * 1000..
//	n = d0 + d1 * (9*1+1)+ d2 * (9*11+1) ...+ dk(9*111... + 1)
//	n = (d0 + d1 + dk) + 9(d1 * 1 + d2 * 11 + dk *111.)
//
// n % 9 = d0 + d1 + dk
// sum of digitals = n % 9
func addDigits(num int) int {
	if num < 10 {
		return num
	}
	// if num == 9, return 9
	// or sum == 9, e.g. 18
	if num % 9 == 0 {
    return 9
  }
	return num % 9
}
