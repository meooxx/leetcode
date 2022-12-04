package main

func hammingWeight(num uint32) int {
	count := 0
	for ; num > 0; num >>= 1 {
		if num&1 == 1 {
			count++
		}
	}
	return count
}
