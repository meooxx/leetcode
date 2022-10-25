package main

func isPalindrome(s string) bool {
	isChar := func(a byte) bool {
		if a >= 'a' && a <= 'z' || a >= 'A' && a <= 'Z' {
			return true
		}
		return false
	}
	isNumber := func(a byte) bool {
		if a >= '0' && a <= '9' {
			return true
		}
		return false
	}
	toSlowerCase := func(a byte) byte {
		if a >= 'A' && a <= 'Z' {
			return a + 'a' - 'A'
		}
		return a
	}
	for left, right := 0, len(s)-1; left < right; {
		if !isChar(s[left]) && !isNumber(s[left]) {
			left++
			continue
		}
		if !isChar(s[right]) && !isNumber(s[right]) {
			right--
			continue
		}
		if toSlowerCase(s[left]) != toSlowerCase(s[right]) {
			return false
		}
		left++
		right--

	}
	return true
}