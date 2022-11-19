package main

import "fmt"

func main() {
	fmt.Println(findPeakElement([]int{1,2,3,1}))
}
// case: 1: mid  peak在右;   2: peak mid peak再左;   3: mid == peak
// 
// ❌       ❌    先判断前后,不是peak的话, 则left++, right--
// 1, 2, 3, 1  
//    |  | 		left     right
//    2  3       2   -   3 => 2 < 3 
//       |		left/right
//     	 3       2<3>1 		return index(3)
func findPeakElement(nums []int) int {
	left := 0
	right := len(nums) - 1
	if len(nums) == 1 {
		return 0
	}
	if nums[0] > nums[1] {
		return 0
	}
	if nums[right-1] < nums[right] {
		return right
	}
	left = 1
	right = right - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			return mid
		} else if nums[mid] < nums[mid-1] {
			right = mid - 1
		} else if nums[mid] < nums[mid+1] {
			left = mid + 1
		}
	}
	return -1
}
