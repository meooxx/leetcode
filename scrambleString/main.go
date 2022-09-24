package main

import "fmt"

func main() {
	fmt.Println(isScramble("great", "aterg"))
	fmt.Println(isScramble("great", "gteat"))
	fmt.Println(isScramble("eebaacbcbcadaaedceaaacadccd", "eadcaacabaddaceacbceaabeccd"))
}

// 1 true, -1 false
var cache map[string]int = map[string]int{}

// gre    at
// g/r/e  a/t
func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}
	s1s2 := s1 + s2
	if cache[s1s2] == 1 {
		return true
	} else if cache[s1s2] == -1 {
		return false
	}
	char := [26]int{}
	for index := 0; index < len(s1); index++ {
		char[s1[index]-'a']++
		char[s2[index]-'a']--
	}
	for _, n := range char {
		if n != 0 {
			cache[s1s2] = -1
			return false
		}
	}
	for i := 1; i < len(s1); i++ {
		if isScramble(s1[0:i], s2[0:i]) && isScramble(s1[i:], s2[i:]) {
			cache[s1s2] = 1
			return true
		}
		// a bcd
		// bcd a
		if isScramble(s1[0:i], s2[len(s2)-i:]) && isScramble(s1[i:], s2[:len(s2)-i]) {
			cache[s1s2] = 1
			return true
		}
	}
	cache[s1s2] = -1
	return false

}