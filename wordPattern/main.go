package main

import "strings"

func wordPattern(pattern string, s string) bool {

	sArr := strings.Split(s, " ")
	if len(sArr) != len(pattern) {
		return false
	}
	mp := [26]string{}
	mp1 := map[string]rune{}
	for i, val := range pattern {
		v := int(val - 'a')
		_, ok := mp1[sArr[i]]
		if mp[v] == "" {
			mp[v] = sArr[i]
		}
		if !ok {
			mp1[sArr[i]] = val
		}
		if mp[v] != sArr[i] || mp1[sArr[i]] != val {
			return false
		}
	}
	return true
}
