package main

func reverseBits(num uint32) uint32 {
	// base := uint32(2)
	result := uint32(0)
	for i := 0; i <= 31; i++ {
		result = result<<1 + num&1
		num >>= 1
	}
	return result
}
