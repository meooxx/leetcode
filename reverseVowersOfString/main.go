package main

func reverseVowels(s string) string {
	sByte := []byte(s)
	mp := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}
	for i, j := 0, len(s)-1; i < j; {
		if !mp[s[i]] {
			i++
			continue
		}
		if !mp[s[j]] {
			j--
			continue
		}
		sByte[i], sByte[j] = sByte[j], sByte[i]
		i++
		j--
	}
	return string(sByte)
}
