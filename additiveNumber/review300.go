package main

import (
	"strconv"
	"strings"
)

func isAdditiveNumber(num string) bool {

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// why equal ?
	// 3/2 = 1
	for i := 1; i <= len(num)/2; i++ {
		// the sum is great or equal max(i,j)
		for j := 1; max(i, j) <= len(num)-i-j; j++ {
			if isValid(i, j, num) {
				return true
			}
		}
	}
	return false
}

func isValid(i, j int, num string) bool {
	// 01 + 11
	if num[0] == '0' && i > 1 {
		return false
	}
	// 11 + 02
	if num[i] == '0' && j > 1 {
		return false
	}
	x1, _ := strconv.Atoi(num[0:i])
	x2, _ := strconv.Atoi(num[i : i+j])
	// start != len(num) ?
	// case: 1, 1, 1(sum), length of sum is less than length of i+j
	// so that go into for statement and return false
	for start := i + j; start != len(num); {
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
