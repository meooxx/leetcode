package main

// binary
func findDuplicate0(nums []int) int {
	high := len(nums) - 1
	low := 0
	for low < high {
		mid := (high-low)/2 + low
		count := 0
		for i := range nums {
			if nums[i] <= mid {
				count++
			}
		}
		if count <= mid {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}

// m: stand for loop to meet point 
// h: stand for head to meet
// r: stand for meet to loop point(rest of fast)
// total slow = h + m
//               x2
// total fast = h + 2(m) + rest
// h = rest
func findDuplicate(nums []int) int {
	fast := nums[nums[0]]
	slow := nums[0]

	// 1 2 3 4 2
  //                m    loop point
	// 1 -> 2 -> 3 -> 4 -> 2 -> 3
	// 2 -> 4 -> 2 -> 4
	for fast != slow {
		fast = nums[nums[fast]]
		slow = nums[slow]
	}
	slow = 0
	for fast != slow {
		fast = nums[fast]
		slow = nums[slow]
	}
	return slow
}
