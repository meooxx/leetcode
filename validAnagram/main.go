package main

func isAnagram(s string, t string) bool {
	if s == t {
		return true
	}
	mp := map[rune]int{}
	for _, b := range t {
		mp[b]++
	}
	for _, b := range s {
		if _, ok := mp[b]; ok {
			mp[b]--
		} else {
			return false
		}
	}
	for i := range mp {
		if mp[i] > 0 || mp[i] < 0 {
			return false
		}
	}
	return true
}
