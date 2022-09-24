package main

import (
	"fmt"
)

func main() {
	a := isP(121)
	fmt.Println(a, "a")
}

func isPalindrome(n int) bool {
	s := fmt.Sprint(n)
	for j, k := 0, len(s)-1; j < k; {
		if s[j] != s[k] {
			return false
		}
		j++
		k--
	}
	return true
}

func isP(n int) bool {
	var r int
	x := n

	for x > 0 {
		//recurring operation per loop
		//  123
		//loop1: 3, 12            put last index number to first index
		//loop2: 3 * 10 +2, 1     put last index -1 number to second
		//     :  32 * 10 + 1,    ...
		r = r*10 + x%10
		x = x / 10
	}
	return r == n
}
