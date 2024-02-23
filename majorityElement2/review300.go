package main

func main() {
	majorityElement([]int{2, 1, 1, 3, 1, 4, 5, 6})
}
func majorityElement(nums []int) []int {
	count1, count2 := 0, 0
	X, Y := -1, -1

	for i := range nums {
		// 先判断是否与X, Y 相同 再判断count是否为0
		// 否则 2 1 1 3  1
		// 此时 count2 应该++,
		// 如果先判断count 则 会讲X赋值为1, 此时X, Y 都为1 是错误的❌
		if X == nums[i] {
			count1++
		} else if Y == nums[i] {
			count2++
		} else if count1 == 0 {
			X = nums[i]
			count1++
		} else if count2 == 0 {
			Y = nums[i]
			count2++
		} else {
			count1--
			count2--
		}
	}
	result := []int{}

	count1 = 0
	count2 = 0
	for i := range nums {
		if nums[i] == X {
			count1++
		// why else if ?
		// de-dupcate
		// e.g. [-1, -1, -1]
		}else if nums[i] == Y {
			count2++
		}
	}
	if count1 > len(nums)/3 {
		result = append(result, X)
	}
	if count2 > len(nums)/3 {
		result = append(result, Y)
	}
	return result
}
