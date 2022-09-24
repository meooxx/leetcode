package main

import (
	//"strings"
	"fmt"
)

func main() {
	fmt.Print(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}

func longestCommonPrefix(strs []string) string {
	first := strs[0]
	str := ""
	for i := 0; i < len(first); i++ {
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) <= i {
				return str
			}
			if strs[j][i:i+1] != first[i:i+1] {
				return str
			}

		}
		str += first[i : i+1]

	}
	return str
}
