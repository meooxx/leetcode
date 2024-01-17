package main

import (
	"strconv"
	"strings"
	"fmt"
)
func main() {
	
	fmt.Println(addOperators("123456789", 45))

}

func addOperators(num string, target int) []string {

	result := []string{}
	if num == "" {
		return result
	}
	next(num, 0, &result, nil, 0, 0, target)
	return result
}

// start: n start index
// p: "2+1"
// pre: pre num, e.g. 2+11, pre num 11
// cur: the result of curr expression, e.g. "2+1" == 3 == curr
func next(num string, start int, result *[]string, p []string, pre, curr, target int) {
	if start == len(num) {
		if curr == target {
			*result = append(*result, strings.Join(p, ""))
		}
		return
	}

	for end := start; end < len(num); end++ {
		if end != start && num[start] == '0' {
			return
		}
		istr := num[start : end+1]
		n, _ := strconv.Atoi(istr)
		if start == 0 {
			next(num, end+1, result, []string{istr}, n, n, target)
		} else {
			tempPath := append(p, "+", istr)
			next(num, end+1, result, tempPath, n, curr+n, target)
			tempPath[len(p)] = "-"
			next(num, end+1, result, tempPath, -n, curr-n, target)
			tempPath[len(p)] = "*"
			next(num, end+1, result, tempPath, n*pre, (curr-pre)+pre*n, target)
		}

	}

}
