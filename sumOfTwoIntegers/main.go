package main

// a 10
// b 11

// c = a & b = 1
// a = a ^ b = 1
// sum(a, b) = 1<<1 ^ 1 = 5

//	 11         c: 10,    a: 101 ,    b 100
//	110         c: 100,   a: 1        b 1000
//
// 1001         c: 0      a: 1001
func getSum(a int, b int) int {
	for b != 0 {
		carry := a & b
		a = a ^ b
		b = carry << 1
	}
	return a
}
