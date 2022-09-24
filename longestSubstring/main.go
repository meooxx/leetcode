package main

import "fmt"

// aaabc, position 上一个重复位置
// |          无重复, [a]0
//  |         找到重复的index 0, 1-0 = 1 [a]1 position 1
//   |        找到重复的index 1, 2-1 = 1 [a]2 position 2
//    |       无重复, [b]3,
// ...
// len(aaabc) - position = 3, return 3
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	m := map[rune]int{}
	position, r := 0, 0
	for i, v := range s {
		if lastPos, ok := m[v]; ok {
			if i-position > r {
				r = i - position
			}
			// case:abba, position后退
			// = for aaab
			if lastPos >= position {
				position = lastPos + 1
			}
		}
		m[v] = i
	}
	// case abcd, or aaabcde 连续不重复到结束, 成最长subStr
	if len(s)-position > r {
		r = len(s) - position
	}

	return r
}

func main() {
	// abba
	// abcdedf
	b := lengthOfLongestSubstring("aaaabcd")
	fmt.Println(b)
}
