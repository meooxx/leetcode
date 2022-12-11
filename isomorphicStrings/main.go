package main

func isIsomorphic1(s string, t string) bool {
	mp1 := map[byte]byte{}
	mp2 := map[byte]byte{}

	for i := range s {
		if v, ok := mp1[s[i]]; ok && v != t[i] {
			return false
		}
		if v, ok := mp2[t[i]]; ok && v != s[i] {
			return false
		}
		mp1[s[i]] = t[i]
		mp2[t[i]] = s[i]
	}
	return true
}


//   a     c    b
//   1     2    0
//   1     2    2
//   b     d    d
//   T     T    F
func isIsomorphic(s string, t string) bool {
	mp1 := [256]int{}
	mp2 := [256]int{}
	for i := range s {
		if mp1[s[i]] != mp2[t[i]] {
			return false
		}
		mp1[s[i]] = i + 1
		mp2[t[i]] = i + 1
	}
	return true
}
