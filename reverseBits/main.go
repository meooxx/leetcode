package main

func reverseBits(num uint32) uint32 {
	base := uint32(1)
	anwser := uint32(0)
	for i := 31; i >= 0; i-- {
		anwser += (num >> i & 1) * base
		base <<= 1
	}
	return anwser
}