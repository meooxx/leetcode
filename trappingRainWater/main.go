package main

import "fmt"

func main() {
	fmt.Println(trap([]int{4, 1, 3}))
}

func trap(height []int) int {
	left, right := 0, len(height)-1
	maxLeft, maxRight := 0, 0
	water := 0

	for left < right {
		if height[left] < height[right] {
			if height[left] > maxLeft {
				maxLeft = height[left]
			} else {
				water += maxLeft - height[left]
			}
			left++
		} else {
			if height[right] > maxRight {
				maxRight = height[right]
			} else {
				water += maxRight - height[right]
			}
			right--
		}
	}
	return water
}
