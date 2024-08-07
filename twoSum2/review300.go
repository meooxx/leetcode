package main

func twoSum(numbers []int, target int) []int {
	for left, right := 0, len(numbers)-1; left < right; {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		}
		if numbers[left]+numbers[right] < target {
			left++
		} else {
			right--
		}

	}
	return []int{}
}
