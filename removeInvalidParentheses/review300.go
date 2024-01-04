package main

import "fmt"

func main() {
	fmt.Println(removeInvalidParentheses("()())()"))
}

func removeInvalidParentheses(s string) []string {
	anwser := []string{}
	var remove func(str string, lastIndex, lastPosition int, p0, p1 byte)
	remove = func(str string, lastIndex, lastPos int, p0, p1 byte) {
		count := 0
		for i := lastIndex; i < len(str); i++ {
			if str[i] == p0 {
				count++
			}
			if str[i] == p1 {
				count--
			}
			if count >= 0 {
				continue
			}
			for j := lastPos; j < len(str); j++ {

				// j == lastPost for str which start with')'
				// like ))()
				if str[j] == p1 && (j == lastPos || str[j-1] != p1) {
					// 自动加 1, cause j-th char has been removed.
					remove(str[:j]+str[j+1:], i, j, p0, p1)
				}
			}
			return
		}
		strBy := []byte(str)
		for i, j := 0, len(str)-1; i < j; {
			strBy[i], strBy[j] = strBy[j], strBy[i]
			i++
			j--
		}
		if p0 == '(' {
			remove(string(strBy), 0, 0, ')', '(')
		} else {
			anwser = append(anwser, string(strBy))
		}
	}
	remove(s, 0, 0, '(', ')')
	return anwser
}
