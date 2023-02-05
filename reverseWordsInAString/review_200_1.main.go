package main

func reverseWords(s string) string {

	right := len(s) - 1

	result := ""
	for right >= 0 {
		if s[right] == ' ' {
			right--
			continue
		}
		left := right
		for left >= 0 && s[left] != ' ' {
			left--
		}
		if result != "" {
			result += " "
		}
		result += s[left+1 : right+1]

		right = left
	}
	return result
}
