package main

func findRepeatedDnaSequences(s string) []string {
	mp := map[string]int{}
	result := []string{}
	for i := 0; i+9 < len(s); i++ {
		seg := s[i : i+10]
		mp[seg]++
		if mp[seg] == 2 {
			result = append(result, seg)
		}
	}
	return result
}
