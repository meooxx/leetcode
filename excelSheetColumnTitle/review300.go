package main

// for a string "ABC", and it's corresponding number n
// n = (A+1) * 26^2 + (B+1) * 26^1 + (C+1) * 26^0
// n - 1 = ...  C + 1 -1
//
// C = (n-1) % 26

func convertToTitle(columnNumber int) string {
	result := ""
	for columnNumber >= 0 {
		columnNumber -= 1
		result = string(byte(columnNumber%26)+'A') + result
		columnNumber /= 26
	}
	return result
}
