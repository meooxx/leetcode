package main

func titleToNumber(columnTitle string) int {
	i := len(columnTitle) - 1
	result := 0
	base := 1
	for i >= 0 {
		n := int(columnTitle[i]-'A'+1) * base
		result += n
		base *= 26
		i--
	}
	return result
}
