package main

// the first run ^ get a^b, cuz' only a and is unique
// 1, 2, 1, 3,2,5
// after first run, we can obtain 6( 3 ^5 == 6 == 110),
// which indicates that the first 2 bits are different between 3, 5
// so we can use arbitary `1` bit to partition the nums into two groups
// e.g.  110 , 010
// 10 10 11   ^ 010 == 1 => 10 ^ 10 ^11 => 11
// 1  1  3    ^ 010 == 0 => 1 ^ 1 ^ 3   => 3
// the same goes for 100        
func singleNumber(nums []int) []int {
	diff := 0
	for i := range nums {
		diff ^= nums[i]
	}
	// got last bit 1
	// diff- 1 mean last bit become 0
	// then ~(diff-1), the last bit must 1
	// ~(diff-1) == -diff
	// ~(diff-1) == ~(diff-1) + 1 -1 = -(diff-1) - 1 = -diff
	diff &= -diff
	result := []int{0, 0}
	for i := range nums {
		if diff&nums[i] == 0 {
			result[0] ^= nums[i]
		} else {
			result[1] ^= nums[i]
		}
	}
	return result
}
