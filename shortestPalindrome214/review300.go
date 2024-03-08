package main

// http://www.btechsmartclass.com/data_structures/knuth-morris-pratt-algorithm.html
// generate table from start and the match is continual
//
//	abcd abc abce
//	0000 123 0120

func getLPSTable(pattern []byte) []int {
	i, j := 0, 1
	table := make([]int, len(pattern))
	for j < len(pattern) {
		if pattern[i] == pattern[j] {
			i++
			table[j] = i
			j++
		} else {

			if i == 0 {
				j++
			} else {
				//  0 1 2 3 ....  i
				// i !=0 即j前面(i,j同步++)已经有部分匹配了, 再判断 j 位置与 i-1 位相同时,
				// 直接基于 i位置 +1

				// [a a c a a a c ]
				// [0 1 0 1 2 2 3 ]
				i = table[i-1]
			}

		}
	}
	return table

}

// c a t a c b # b c a t a c
// 0 0 0 0 1 0 0 0 1 2 3 4 5
func shortestPalindrome(s string) string {
	reversed := []byte(s)
	for i, j := 0, len(s)-1; i < j; {
		reversed[i], reversed[j] = reversed[j], reversed[i]
		i++
		j--
	}
	pattern := append([]byte{}, s...)
	pattern = append(pattern, '#')
	pattern = append(pattern, reversed...)

	table := getLPSTable(pattern)

	mostRight := table[len(table)-1]
	return string(reversed[:len(table)-mostRight]) + s
}
