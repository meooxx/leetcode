package main

func findKthLargest(nums []int, k int) int {
	var quickSelect func(nums []int, left, right, k int) int
	quickSelect = func(nums []int, left, right, k int) int {
		i, j := left, right-1
		val := nums[right]
		for i <= j {
			for i <= j && nums[i] < val {
				i++
			}
			for i <= j && nums[j] >= val {
				j--
			}
			if i <= j {
				nums[i], nums[j] = nums[j], nums[i]
				i++
				j--
			}
		}
		nums[i], nums[right] = nums[right], nums[i]
		count := right - i + 1
		if count == k {
			// 中间是最小的, 数量又相等, 正好符合 the Kth largest
			// nums[right] >= mid
			return nums[i]
		} else if count > k {
			return quickSelect(nums, i+1, right, k)
		} else {
			// 右侧不够 [123 5], k == 2 val 取3
			// 则需要再从 左侧在找一个
			return quickSelect(nums, left, j, k-count)
		}
	}
	return quickSelect(nums, 0, len(nums)-1, k)
}
