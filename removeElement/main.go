package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 3, 4, 5, 6}

	fmt.Println(removeElement(nums, 2), nums)

}
// 将非目标值移动到空位上
// nums[i] 不等于val移动到  nums[j]
func removeElement(nums []int, val int) int {
	j := 0
	n := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[j] = nums[i]
			n++
			j++
		}
	}
	return j
}
