package main

func findDuplicate1(nums []int) int {
	anwser := 0
	bits := 1
	for 1<<bits < len(nums) {
		bits++
	}
	for index := 0; index < bits; index++ {

		a := 0
		b := 0

		for i := 0; i < len(nums); i++ {
			if nums[i]>>index&1 == 1 {
				a++
			}
			if i > 0 {
				if i>>index&1 == 1 {
					b++
				}
			}
		}
		// case: 2,2,2,2
		// 0 4  a
		// 2 2  b
		if a > b {
			anwser += 1 << index
		}
	}

	return anwser
}

// binary search 1,n
// if count of nums[i] <= mid, search mid+1, right
// else search left, mid
// 1,2,2, 3,4
// r,l
// 4,1   mid == 2 == (4-1)/2+1,  total nums[i] == 3 <= 2 
// 2,1   mid == 1 == (2-1)/2+1,  total         == 1 <= 1
//       mid = 1+1
//  anwser = 2
func findDuplicate0(nums []int) int {
	left := 1
	right := len(nums) - 1
	for left < right {
		mid := (right-left)/2 + left
		count := 0
		for i := range nums {
			if nums[i] <= mid {
				count++
			}
		}
		//
		// 123 3 4
		if count > mid {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// 1 2 3 4 2
// 1 -> 2 ->3 -4
//      \______|
func findDuplicate(nums []int) int {

	slow := nums[0]
	fast := nums[nums[0]]
	for fast != slow {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
