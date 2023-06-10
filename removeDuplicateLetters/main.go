package main

func removeDuplicateLetters(s string) string {
	char := [26]int{}
	for i := range s {
		idx := int(s[i] - 'a')
		char[idx] = i
	}
	visited := [26]bool{}
	result := []byte{}
	for i := range s {
		idx := int(s[i] - 'a')
		if len(result) == 0 {
			result = append(result, s[i])
			visited[idx] = true
		} else {
			if visited[idx] {
				continue
			}
			if s[i] > result[len(result)-1] {
				result = append(result, s[i])
				visited[idx] = true
			} else {
				for j := len(result) - 1; j >= 0; j-- {
					preIndex := int(result[j] - 'a')
					if result[j] > s[i] && char[preIndex] > i {
						result = result[:j]
						visited[preIndex] = false
					} else {
						break
					}
				}
				result = append(result, s[i])
				visited[int(s[i]-'a')] = true
			}
		}

	}
	return string(result)
}
