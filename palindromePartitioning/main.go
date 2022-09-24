package main

import (
	"fmt"
)

func main() {
	fmt.Println(partition("aabvcc"))
	fmt.Println(partition("aa"))

}
func partition(s string) [][]string {

	if len(s) == 1 {
		return [][]string{{s}}
	}
	mp := map[string]bool{}
	isPalindrome := func(s string) bool {
		if len(s) <= 1 {
			return true
		}
		if v, ok := mp[s]; ok {
			return v
		}
		for left, right := 0, len(s)-1; left < right; {
			if s[left] != s[right] {
				mp[s] = false
				return false
			}
			left++
			right--
		}
		mp[s] = true
		return true
	}
	result := [][]string{}
	var partionSub func(s string, cur []string)
	partionSub = func(s string, cur []string) {
		if len(s) <= 1 {
			if len(s) == 1 {
				cur = append(cur, s)
			}
			result = append(result, cur)
			return
		}

		for pos := 1; pos <= len(s); pos++ {
			leftStr := s[:pos]
			rightStr := s[pos:]
			newCur := append([]string{}, cur...)
			if isPalindrome(leftStr) {
				newCur = append(newCur, leftStr)
				partionSub(rightStr, newCur)
			}
		}
	}
	partionSub(s, []string{})
	return result
}
