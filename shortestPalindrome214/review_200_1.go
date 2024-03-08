package main

func main() {
	shortestPalindrome("aacecaaa")
}
func shortestPalindrome(s string) string {
	reverse := func(s string) string {
		rs := []rune(s)
		for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
			rs[i], rs[j] = rs[j], rs[i]
		}
		return string(rs)
	}
	sNew := s + "#" + reverse(s)
	table := make([]int, len(sNew))
	i, j := 0, 1

	// ab a  b    ab c
	// 00 1  2    34
	// 01 2  3    45
	//    i' i-1  i  j
	for j < len(sNew) {
        
		if sNew[i] == sNew[j] {
			i++
			table[j] = i
			j++
		} else {
			if i != 0 {
				i = table[i-1]
			} else {
				j++

			}
		}
	}
	rightIndex := table[len(sNew)-1]
	return reverse(s)[:len(s)-rightIndex] + s
}
