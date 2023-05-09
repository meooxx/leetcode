package main

func findDuplicate(nums []int) int {
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
