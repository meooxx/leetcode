package main

func isAnagram(s string, t string) bool {

	c := [26]int{}

	for i := range s {
		c[s[i]-'a']++
	}
	for i := range t {
		c[t[i]-'a']--
	}
	for i := range c {
		if c[i] != 0 {
			return false
		}
	}
	return true
}
