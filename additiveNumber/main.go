package main

import "strconv"
import "strings"
import "fmt"

func main() {
	fmt.Println(isAdditiveNumber("0112"))
	fmt.Println(isAdditiveNumber("01124"))

}

func isAdditiveNumber(num string) bool {

	max := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	// i, j is length of n1, n2

	// 1 2 3  3 / 2 == 1
	for i := 1; i <= len(num)/2; i++ {
		// the len of sum is great or equal max(i, j)
		// 11 + 1, <= len(11)
		for j := 1; max(i, j) <= len(num)-j-i; j++ {
			if isValid(i, j, num) {
				return true
			}

		}
	}
	return false

}

func isValid(i, j int, num string) bool {
	// 000, 1 02 3, 101
	if num[0] == '0' && i > 1 {
		return false
	}
	if num[i] == '0' && j > 1 {
		return false
	}

	x1, _ := strconv.Atoi(num[0:i])
	x2, _ := strconv.Atoi(num[i : i+j])
	for start := i + j; start != len(num); {
		// 0112
		sum := x1 + x2
		x1 = x2
		x2 = sum

		sumStr := strconv.Itoa(sum)
		if !strings.HasPrefix(num[start:], sumStr) {
			return false
		}
		start += len(sumStr)

	}
	return true
}
