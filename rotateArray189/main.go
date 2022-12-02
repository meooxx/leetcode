package main

// O(n*k)
func rotate1(nums []int, k int) {
	for ; k > 0; k-- {
		size := len(nums) - 1
		tail := nums[size]
		for i := size; i >= 1; i-- {
			nums[i] = nums[i-1]
		}
		nums[0] = tail
	}

}
// 1234 567
// 1 copy [1234] ==> 1234
// 2 567 4567 => 567 1234
func rotate2(nums []int, k int) {
	if len(nums) <= 1 {
		return
	}
	k = k % len(nums)
	cp := make([]int, len(nums)-k)
	copy(cp, nums[:len(nums)-k])

	for i, j := 0, len(nums)-k; j < len(nums); {
		nums[i] = nums[j]
		i++
		j++
	}

	for i, j := k, 0; j < len(cp); {
		nums[i] = cp[j]
		i++
		j++
	}

}

// space O(1)
func rotate(nums []int, k int) {
	if len(nums) <= 1 {
		return
	}
	k = k % len(nums)

	reverse := func(nums []int, start, end int) {
		for start < end {
			nums[start], nums[end] = nums[end], nums[start]
			start++
			end--
		}
	}
	// 1234567
	// 7654321 reverse1
	// 5674321 reverse2
	// 5671234 reverse3
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)

}
