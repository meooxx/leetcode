package main

func isIsomorphic(s string, t string) bool {
	// map can be replace with array
	// ascii code 0-255
	mp1 := map[byte]int{}
	mp2 := map[byte]int{}
	for i := range s {
		// a b b
		// b a b
		// b-> a , b, b => false
		if mp1[s[i]] != mp2[t[i]] {
			return false
		}
		mp1[s[i]] = i + 1
		mp2[t[i]] = i + 1
	}
	return true
}
