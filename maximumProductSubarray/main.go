package main

import "fmt"

func main() {
	fmt.Println(maxProduct([]int{-2, 0, -1}))
	fmt.Println(maxProduct([]int{0, 2}))
	fmt.Println(maxProduct1([]int{-2, 0, -1}))
	fmt.Println(maxProduct1([]int{0, 2}))
}

// awsome!
// -1, -2  1, 2, 3
// If nums of negative n is even and all on '0' side, then
// the max is n1 * n2 * n, sum(start, end) == sum(end, start)
// or 1, 2,3,-1, 4, 6, 7
// so the max is max(sum[index(0), index(-1)), sum(index(-1),length])
// 即如果负数个数为偶数从左到右, 从右到左顺序遍历结果一致
// 如果负数个数为基数, 那么结果出现在
// 负数左边  1 * 2 *3
// 或者负数右边 4 * 6 * 7
// case 1 1 1 -2 -1 2 2 -6==> -2 的右边,因为-2右侧俩个-2数,相当于+数
func maxProduct(nums []int) int {
	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}
	prod := 1
	result := nums[0]

	for i := 0; i < len(nums); i++ {
		prod = prod * nums[i]
		result = max(prod, result)
		if prod == 0 {
			prod = 1
		}
	}
	prod = 1
	for i := len(nums) - 1; i >= 0; i-- {
		prod = prod * nums[i]
		result = max(prod, result)
		if prod == 0 {
			prod = 1
		}
	}
	return result
}

func maxProduct1(nums []int) int {
	maxSum := nums[0]
	minSum := nums[0]
	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maxSum, minSum = minSum, maxSum
		}
		// case 1, 2   0  3 4
		// 1 * 2 * 0 	=> maxSum = 0
		// 3 > 0(maxSum) => maxSum = 3 
		// =>3 * 4
		maxSum = max(maxSum*nums[i], nums[i])
		minSum = min(minSum*nums[i], nums[i])
		result = max(result, maxSum)
	}
	return result
}
