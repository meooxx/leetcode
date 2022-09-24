package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(minWindow("hElloe", "El"))
	fmt.Println(minWindow("a", "a"))
	fmt.Println(minWindow("acbbaca", "aba"))
	fmt.Println(minWindow("bdab", "ab"))
	fmt.Println(minWindow("bdab", "C"))

}

// ababacabc  bc
// ababac  		右指针找到完全包含bc的substr
// 	  bac 	 	左指针去掉不必要, 结果bac
//     acab 	直到一个b(属于bc),count+1, 此时右指针继续直到 count==0
//      cab 	左指针去掉不必要的值
//       abc ...
//        bc
func minWindow1(s string, t string) string {
	left := 0
	right := 0
	m := map[byte]int{}
	count := len(t)
	minStrings := ""
	for index := range t {
		m[t[index]] += 1
	}
	for right < len(s) {
		if _, ok := m[s[right]]; ok {
			// 防止重复
			if m[s[right]] > 0 {
				count--
			}
			m[s[right]]--
		}
		// why it need not left < right
		// 1 bbabbb a,  left 从0-a的index一定会count再次>0
		// 2 如果aaaa b, 不会进入这里
		for count == 0 {
			if _, ok := m[s[left]]; ok {
				m[s[left]]++
				// 防止重复
				if m[s[left]] > 0 {
					count++
				}
			}
			if len(minStrings) > len(s[left:right+1]) || minStrings == "" {
				minStrings = s[left : right+1]
			}
			left++
		}
		right++

	}
	return minStrings
}

func minWindow2(s string, t string) string {
	sUpper := s
	tUpper := t
	if len(t) == 1 {
		if strings.Contains(s, t) {
			return t
		}
	}
	m1 := map[byte]int{}
	var m2 map[byte]int
	for l := range t {
		m1[tUpper[l]]++
	}
	count := 0
	minString := ""
	for index := 0; index < len(sUpper); index++ {
		letter := sUpper[index]
		count = 0
		m2 = map[byte]int{}
		if m1[letter] > 0 {
			m2[letter]++
			count++
			for start := index + 1; start < len(sUpper); start++ {
				if m1[sUpper[start]] > 0 {
					m2[sUpper[start]]++
					if m2[sUpper[start]] <= m1[sUpper[start]] {
						count++
						if count == len(t) {
							if len(sUpper[index:start+1]) < len(minString) || minString == "" {
								minString = s[index : start+1]
							}
							break
						}
					}

				}

			}
		}

	}
	return minString
}

func minWindow(s string, t string) string {

	mp := map[byte]int{}
	count := len(t)
	start := 0
	end := 0
	str := ""

	for i := range t {
		mp[t[i]]++
	}

	for end < len(s) {
		if n, ok := mp[s[end]]; ok {
			mp[s[end]]--
			if n > 0 {
				count--
			}
		}
		for count == 0 {
			if str == "" || len(s[start:end+1]) < len(str) {
				str = s[start : end+1]
			}
			if n, ok := mp[s[start]]; ok {
				mp[s[start]]++
				if n == 0 {
					count++
				}
			}
			start++
		}
		end++
	}
	return str
}
