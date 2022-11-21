package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(compareVersion("1.0.0.1", "1.0"))
	fmt.Println(compareVersion("1.0.0.1", "1.0.1"))

}

func compareVersion1(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")

	l1 := len(v1)
	l2 := len(v2)
	start := 0
	for start < l1 && start < l2 {
		n1, _ := strconv.ParseInt(v1[start], 10, 32)
		n2, _ := strconv.ParseInt(v2[start], 10, 32)
		if n1 > n2 {
			return 1
		} else if n1 < n2 {
			return -1
		}
		start++
	}
	for start < l1 {
		n1, _ := strconv.ParseInt(v1[start], 10, 32)
		if n1 > 0 {
			return 1
		}
		start++
	}
	for start < l2 {
		n2, _ := strconv.ParseInt(v2[start], 10, 32)
		if n2 > 0 {
			return -1
		}
		start++
	}

	return 0

}

// 1.01.1.1
// 1.01.2
func compareVersion(version1 string, version2 string) int {
	l1 := len(version1)
	l2 := len(version2)
	i, j := 0, 0
	for i < l1 || j < l2 {
		num1 := 0
		for i < l1 && version1[i] != '.' {
			num1 = num1*10 + int(version1[i]-'0')
			i++
		}
		num2 := 0
		for j < l2 && version2[j] != '.' {
			num2 = num2*10 + int(version2[j]-'0')
			j++
		}
		if num1 > num2 {
			return 1
		} else if num1 < num2 {
			return -1
		}
		i++
		j++
	}
	return 0
}
