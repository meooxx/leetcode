package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(restoreIpAddresses2("25525511135"))
	fmt.Println(restoreIpAddresses2("0000"))
	fmt.Println(restoreIpAddresses("101023"))
}

// 255 255 255 11 131
func restoreIpAddresses(s string) []string {
	result := []string{}
	putDot(&result, []string{}, s)
	return result
}

func putDot(result *[]string, cur []string, rem string) {

	if len(cur) == 4 {
		item := strings.Join(cur, ".")
		if rem != "" {
			return
		}
		*result = append(*result, item)
		return
	}
	max := 3
	if len(rem) < 3 {
		max = len(rem)
	}
	for i := 1; i <= max; i++ {

		// this part may slow the leetcode tester
		lenStr := len(rem) - i
		nBit := 4 - len(cur) - 1
		// 1.1.1. 位数不够
		if lenStr < nBit*1 {
			break
		}
		// 1.1.1.1 1 位数超过
		if lenStr > nBit*3 {
			continue
		}
		//

		curStr := rem[:i]
		n, err := strconv.ParseInt(curStr, 10, 16)
		if curStr[0] == '0' && len(curStr) > 1 {
			break
		}
		if err != nil || n > 255 {
			return
		}
		cur = append(cur, curStr)

		putDot(result, cur, rem[i:])
		cur = cur[:len(cur)-1]

	}

}

func restoreIpAddresses2(s string) []string {
	result := []string{}

	findNext(&result, []string{}, s, 0)

	return result
}

func findNext(result *[]string, cur []string, s string, index int) {
	if len(cur) == 4 {
		if index >= len(s) {
			item := strings.Join(cur, ".")
			*result = append(*result, item)
		}

		return
	}
	max := 3
	if index+max >= len(s) {
		max = len(s) - index
	}
	for i := 1; i <= max; i++ {
		if i > 1 && s[index+i] == '0' {
			return
		}
		seg := s[index : index+i]
		n, err := strconv.ParseInt(seg, 10, 32)
		if err != nil {
			return
		}
		if n <= 255 {
			cur = append(cur, seg)
			findNext(result, cur, s, index+i)
			cur = cur[:len(cur)-1]
		}

	}

}
