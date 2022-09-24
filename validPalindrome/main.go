package main

import (
	"fmt"
)

func main() {
	// fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	// fmt.Println(isPalindrome("0P"))

	fmt.Println(isPalindrome("ab_a"))

}
func isPalindrome(s string) bool {
	isLetter := func(b byte) bool {
		return b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z'
	}
	isNumber := func(b byte) bool {
		return b >= '0' && b <= '9'
	}
	toLowerCase := func(b byte) byte {
		if b >= 'A' && b <= 'Z' {
			return b + 32
		}
		return b
	}
	for left, right := 0, len(s)-1; left < right; {
		if !isLetter(s[left]) && !isNumber(s[left]) {
			left++
			continue
		}
		if !isLetter(s[right]) && !isNumber(s[right]) {
			right--
			continue
		}
		if toLowerCase(s[left]) != toLowerCase(s[right]) {
			return false
		}
		left++
		right--

	}
	return true
}

func isPalindrome1(s string) bool {
	if len(s) <= 1 {
		return true
	}
	for left, right := 0, len(s)-1; left < right; {
		leftChar := s[left]
		rightChar := s[right]
		if !(leftChar >= '0' && leftChar <= '9' || leftChar >= 'a' && leftChar <= 'z' || leftChar >= 'A' && leftChar <= 'Z') {
			left++
			continue
		}
		if !(rightChar >= '0' && rightChar <= '9' || rightChar >= 'a' && rightChar <= 'z' || rightChar >= 'A' && rightChar <= 'Z') {
			right--
			continue
		}
		diff := byte('a') - 'A'
		if leftChar == rightChar {
			left++
			right--
		} else if leftChar >= '0' && leftChar <= '9' {
			if leftChar == rightChar {
				left++
				right--
			} else {
				return false
			}

		} else if leftChar+diff == rightChar || leftChar-diff == rightChar {
			left++
			right--
		} else {
			return false
		}
	}
	return true

}
