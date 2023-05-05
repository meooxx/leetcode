package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(addOperators("123", 6))
	fmt.Println(addOperators("130", 30))
	fmt.Println(addOperators("1023", 6))
	fmt.Println(addOperators("123456789", 45))

}

// 1234
func addOperators(num string, target int) []string {

	result := []string{}
	if num == "" {
		return result
	}
	next(num, 0, target, nil, 0, 0, &result)
	return result
}

func next(num string, index int, target int, path []string, pre int, resultSofar int, result *[]string) {
	if index == len(num) {
		if target == resultSofar {
			*result = append(*result, strings.Join(path, ""))
			return
		}
	}

	//  num[index] == 0, num[i] != 0
	// 1024
	for i := index; i < len(num); i++ {
		if i != index && num[index] == '0' {
			return
		}
		cur := num[index : i+1]
		n64, _ := strconv.ParseInt(cur, 10, 32)
		n := int(n64)
		if index == 0 {
			next(num, i+1, target, []string{cur}, n, n, result)
		} else {
			signIndex := len(path)
			p := append(path, "+", cur)
			next(num, i+1, target, p, n, resultSofar+n, result)
			p[signIndex] = "-"
			next(num, i+1, target, p, -n, resultSofar-n, result)
			p[signIndex] = "*"
			next(num, i+1, target, p, pre*n, resultSofar-pre+pre*n, result)
		}

	}
	return
}
