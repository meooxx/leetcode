package main

func findRepeatedDnaSequences(s string) []string {
	mp := map[string]int{}
	result := []string{}
	for i := 0; i+9 < len(s); i++ {
		seq := s[i : i+10]
		mp[seq]++
		// in case of duplication
		if mp[seq] == 2 {
			result = append(result, seq)
		}
	}
	return result
}
