package main

import "fmt"

/**
 *          i ==2       <-- 5 >4 >3
 * arr := 1 2 5 4 3
 *   5 > 4 > 3 && 2 < 5
 * 右边往左 第一个降序数字, 交换右边比他大的第一个数字
 * 从i位置后面的 生序反转
 * 1 2 543  13542
 * 13542  13 245
 *
 */

func main() {
	arr := []int{1, 1, 3}
	arr1 := []int{1, 1, 5}
	arr2 := []int{1, 2, 3}
	arr3 := []int{5, 4, 3, 2, 1}

	nextPermutation(arr)
	nextPermutation(arr1)
	nextPermutation(arr2)
	nextPermutation(arr3)
	//nextPermutation(arr4)

	fmt.Println(arr, arr1, arr2, arr3)
}

//      4   7  5  3
func nextPermutation(nums []int) {
	if len(nums) < 1 {
		return
	}
	i := len(nums) - 1
	// 123
	// 132
	for i >= 0 {
		if i-1 >= 0 && nums[i-1] < nums[i] {
			i--
			break
		}
		i--
	}
	j := len(nums) - 1
	//   i  j
	// 124 5 3
	for i >= 0 && j > 0 {
		if nums[j] > nums[i] {
			nums[i], nums[j] = nums[j], nums[i]
			break
		}
		j--
	}

	//     i
	// 125 43
	// revers
	// 125 34
	i += 1
	for j := len(nums) - 1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
