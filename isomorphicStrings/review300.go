package main

func isIsomorphic0(s string, t string) bool {
	mp1 := map[byte]byte{}
	mp2 := map[byte]byte{}

	for i := range s {
		if s[i] == t[i] {
			continue
		} else if mp1[s[i]] == 0 && mp2[t[i]] == 0 {
			mp1[s[i]] = t[i]
			mp2[t[i]] = s[i]
		} else {
			if mp1[s[i]] != t[i] || mp2[t[i]] != s[i] {
				return false
			}
		}

	}
	return true
}

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
