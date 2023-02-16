package main

// https://leetcode.com/problems/excel-sheet-column-title/solutions/441430/detailed-explanation-here-s-why-we-need-n-at-first-of-every-loop-java-python-c/?orderBy=most_votes
// A == 0, Z == 25
// n = 'ABC'
// 		n = (A+1)*26^2 + (B+1) * 26^1 + C+1
// 		n -1 = (A+1)*26^2 + (B+1) * 26^1 + C
// 1. C ==> C == (n-1) % 26
// 2. B? 
//   => (n-1)/26 = (A+1)*26^1 + (B+1) * 26^0
// m = (A+1)*26^1 + (B+1)
// m - 1 = (A+1)*26^1 + B
//   B == m-1 % 26
func convertToTitle(columnNumber int) string {
	anwser := []byte{}
	for columnNumber != 0 {
		columnNumber--
		anwser = append(anwser, byte('A'+columnNumber%26))
		columnNumber = columnNumber / 26
	}
	for left, right := 0, len(anwser)-1; left < right; {
		anwser[left], anwser[right] = anwser[right], anwser[left]
		left++
		right--
	}
	return string(anwser)
}

