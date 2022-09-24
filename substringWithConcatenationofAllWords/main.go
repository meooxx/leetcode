package main

import "fmt"

func main() {
	fmt.Println(findSubstring("barfoofoofoobarman", []string{"foo", "bar"}))
}

func findSubstring(s string, words []string) []int {

	m1 := map[string]int{}
	var result []int
	for _, v := range words {
		m1[v] = m1[v] + 1
	}
	wordLen := len(words[0])
	for i := 0; i <= len(s)-wordLen; i += 1 {
		m2 := map[string]int{}
		count := 0
		for j := i; j <= len(s)-wordLen; j += wordLen {
			word := s[j : j+wordLen]
			if _, ok := m1[word]; ok {
				m2[word]++
				if m2[word] > m1[word] {
					break
				}
				count++

			} else {
				break
			}
			if count == len(words) {
				result = append(result, i)
			}
		}

	}

	return result
}
