package main

import (
	"fmt"
)

func main() {
	fmt.Println(isHappy(2))
}

func isHappy(n int) bool {
	cache := map[int]bool{}
	for n > 0 {
		if _, ok := cache[n]; ok {
			return false
		}
		if n == 1 {
			return true
		}
		cache[n] = false
		sum := 0
		for n > 0 {
			rest := n % 10
			sum += rest * rest
			n /= 10
		}

		n = sum
	}
	return false
}
