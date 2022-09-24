package main

import (
	"fmt"
)

func main() {
	//	fmt.Println(strStr("hello", "ll"))

	fmt.Println(strStr("aaa", "bba"))
	fmt.Println(strStr("", "bba"))
	fmt.Println(strStr("bbssbba", "bba"))
	fmt.Println(strStr("bb", "bba"))
	fmt.Println(strStr("mississippi", "issip"))
	fmt.Println(strStr("hello", "ll"))

}

func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			flag := true
			for j := 1; j < len(needle); j++ {
				if i+j >= len(haystack) || needle[j] != haystack[i+j] {
					flag = false
					break
				}
			}
			if flag {
				return i
			}

		}
	}
	return -1
}
