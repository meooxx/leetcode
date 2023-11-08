package main

func removeDuplicateLetters(s string) string {
	lookBack := [26]int{}
	for i := range s {
		lookBack[int(s[i]-'a')] = i
	}
	result := []byte{}
	visited := [26]bool{}
	for i := range s {
		if len(result) == 0 {
			result = append(result, s[i])
			visited[int(s[i]-'a')] = true
		} else {
			if visited[int(s[i]-'a')] {
				continue
			}
			if s[i] > result[len(result)-1] {
				result = append(result, s[i])
				visited[int(s[i]-'a')] = true
			} else {
				for j := len(result) - 1; j >= 0; j-- {
					// e.g.
					// bcabc   abc
					// bcac    bac 
					// 确保后面存在要去掉的元素
					// bc a bc, 可以去掉bc
					// bcac.  只能去掉c
					if result[j] > s[i] && lookBack[int(result[j]-'a')] > i {
						visited[int(result[j]-'a')] = false
						result = result[:j]
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
