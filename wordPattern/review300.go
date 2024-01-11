package main
import "strings"
func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	cacheP := [26]string{}
	cacheWord := map[string]byte{}
	if len(pattern) != len(words) {
		return false
	}
	for i := range pattern {
		charIndex := int(pattern[i] - 'a')
		// the default value of byte is 0
		if cacheP[charIndex] == "" && cacheWord[words[i]] == 0 {
			cacheP[charIndex] = words[i]
			cacheWord[words[i]] = pattern[i]
		} else {
			if cacheP[charIndex] != words[i] || cacheWord[words[i]] != pattern[i] {
				return false
			}
		}

	}
	return true
}
