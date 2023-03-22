package main
// 2 2 333 4 2  
// []  []  |
//  2   3  |
//  2   3  |
//      3  | 
//         4  
//  -   -      => 2 33  remove from [both] stack
//           |
//  +          =>   22 33  just add 2 to [A] stack
func majorityElement(nums []int) []int {
	x, y, count1, count2 := -1, -1, 0, 0
	for i := range nums {
		if x == nums[i] {
			count1++
		} else if y == nums[i] {
			count2++
		} else if count1 == 0 {
			count1 = 1
			x = nums[i]
		} else if count2 == 0 {
			count2 = 1
			y = nums[i]
		} else {
			count1--
			count2--
		}
	}
	result := []int{}
	count1 = 0
	count2 = 0
	for i := range nums {
		if nums[i] == x {
			count1++
		} else if nums[i] == y {
			count2++
		}
	}
	if count1 > len(nums)/3 {
		result = append(result, x)
	}
	if count2 > len(nums)/3 {
		result = append(result, y)
	}
	return result
}
