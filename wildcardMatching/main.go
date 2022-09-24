package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(isMatch("ab", "*a"))
	fmt.Println(isMatch("aa", ""))
	fmt.Println(isMatch("aa", "aa"))
	fmt.Println(isMatch("aaa", "aa"))
	fmt.Println(isMatch("aaa", "a*"))
	fmt.Println(isMatch("acdcb", "a*c?b"))
	fmt.Println(isMatch("adceb", "*a*b"))
	fmt.Println(isMatch("aaaa", "***a"))
	fmt.Println(isMatch("", "****"))
	fmt.Println(isMatch("mississippi", "m??*ss*?i*pi"))
	pre := time.Now()
	fmt.Println(isMatch2("bbbbbbbabbaabbabbbbaaabbabbabaaabbababbbabbbabaaabaab", "b*b*ab**ba*b**b***bba"))
	fmt.Println(time.Since(pre))
}

// 遇到 * 先尝试0次, 但是记录匹配的str和*的位置,直接跳过
// 后面失败返回此处再次尝试, 即跳过一个str
// abdc  a*c
// a a  继续
// b *  跳过 *
// b c  不匹配, 尝试用* 匹配一个字符即跳过b
// d c  不匹配继续跳过d
// c c  ok

func isMatch(s string, p string) bool {
	j := 0
	sPosition := -1
	pPosition := -1
	for i := 0; i < len(s); {
		if j < len(p) && (s[i] == p[j] || p[j] == '?') {
			i++
			j++
		} else if j < len(p) && p[j] == '*' {
			j++
			pPosition = j
			sPosition = i
		} else if pPosition != -1 {
			i = sPosition + 1
			sPosition = i
			j = pPosition
		} else {
			return false
		}
	}
	// case: a, a*****
	for j < len(p) && p[j] == '*' {
		j++
	}
	return j == len(p)

}

//  aaa   a*
func isMatch2(s string, p string) bool {
	return match(s, p)
}

func match(s, p string) bool {
	if s == "" && p == "" {
		return true
	}

	if s == "" && p != "" && p[0] != '*' {
		return false
	}
	if p == "" && s != "" {
		return false
	}

	if p[0] == '*' {
		for i := 0; i <= len(s); {
			if match(s[i:], p[1:]) {
				return true
			}
			i++
		}
		return false
	}
	if p[0] == '?' || p[0] == s[0] {
		return match(s[1:], p[1:])
	}
	return false
}
