package main

func canConstruct(ransomNote string, magazine string) bool {
	mp := [26]int{}
	for i := 0; i < len(magazine); i++ {
		index := int(magazine[i] - 'a')
		mp[index]++
	}
	for i := range ransomNote {
		index := int(ransomNote[i] - 'a')
		if mp[index] <= 0 {
			return false
		}
		mp[index]--
	}
	return true
}
