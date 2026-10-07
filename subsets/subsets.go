package subsets

func subsets(nums []int) [][]int {
	n := 0
	result := [][]int{}
	for n <= len(nums) {
		placeNum(nums, 0, []int{}, n, &result)
		n++
	}
	return result
}

func placeNum(nums []int, start int, current []int, reminder int, result *[][]int) {

	if reminder == 0 {
		c := make([]int, len(current))
		copy(c, current)
		*result = append(*result, c)
	}
	for i := start; i < len(nums); i++ {
		current = append(current, nums[i])
		placeNum(nums, i+1, current, reminder-1, result)
		current = current[:len(current)-1]
	}

}
