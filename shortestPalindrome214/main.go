package main

import "fmt"

func main() {
	fmt.Println(getTable("abc#abd#abcd#abd"))
}

// abc abd abd abc
// 000 120 120 123
func shortestPalindrome(s string) string {

	reversed := []rune(s)
	for left, right := 0, len(reversed)-1; left < right; {
		reversed[left], reversed[right] = reversed[right], reversed[left]
		left++
		right--
	}
	reversedStr := string(reversed)
	tmpStr := s + "#" + reversedStr
	table := getTable(tmpStr)
	// abac # caba
	maxRight := table[len(table)-1]
	return reversedStr[:len(reversedStr)-maxRight] + s
}

func getTable(str string) []int {
	table := make([]int, len(str))
	for i, j := 0, 1; j < len(str); {
		if str[i] == str[j] {
			table[j] = i + 1
			i++
			j++
		} else {
			if i == 0 {
				j++
			} else {
				i = table[i-1]
			}
		}
	}
	return table
}
