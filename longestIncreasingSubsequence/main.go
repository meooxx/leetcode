package main
// 2 6 8 3 4 5
// 2 6 8  =>  2 6 8
//        =>  2 3
// 
//  2 6 8
//  2 3 4 5  ✅
//  == 2  6  8  => 2 3 8 => 2 3 4
func lengthOfLIS(nums []int) int {
	result := []int{nums[0]}
	for i := range nums {
		if nums[i] > result[len(result)-1] {
			result = append(result, nums[i])
		} else {
			l, r := 0, len(result)-1
			// 1 4 8
			for l < r {
				midIndex := (r-l)/2 + l
				if nums[i] > result[midIndex] {
					l = midIndex + 1
				} else {
					r = midIndex
				}
			}
			result[r] = nums[i]
		}
	}
	return len(result)
}
