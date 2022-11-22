package main

// 			 
// 1, 3,  4,  5, 8,\
// 1. two pointer
// 2. think of binary search
// 	1 => 3, 4, 5,8 => 1+4<target => 5-8 =>...
// 	3 => 4 5 8 => 3+5 > target => 3-4
func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left, right}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return []int{}
}
