package main

//  pow(2, 12) % k
//          =   (pow(2, 1)^10 % k  *     pow(2,2) %k)  %k
// 

func superPow(a int, b []int) int {

	if len(b) == 0 {
		return 1
	}

	last := b[len(b)-1]
	rest := b[:len(b)-1]

	return powmod(superPow(a, rest), 10) * powmod(a, last) % 1337

}

func powmod(a, k int) int {

	result := 1
	a %= 1337
	for i := 1; i <= k; i++ {
		result = result * a % 1337
	}
	return result
}