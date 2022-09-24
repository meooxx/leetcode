package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{9, 9}))
	fmt.Println(plusOne([]int{1, 0, 9}))
}

func plusOne(digits []int) []int {
	l := len(digits)

	for i := l - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] == 10 {
			digits[i] = 0
		} else {
			break
		}
	}
	if digits[0] == 0 {

		digits = append([]int{1}, digits...)

	}
	return digits

}
