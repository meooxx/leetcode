package main

/**
suppose:

a = Ak +B;
b = Ck + D;

so:

a * b = ACk + ADk + BCk +BD;

ab % k = (ACk + ADk + BCk +BD) % k = BD % k;

a % k = B;

b % k = D;

(a%k)(b%k)%k = BD % k;
*/

// ab % k = (a%k)*(b%k)%k
// a^1234567 = (a^1234560%k *        a^7%k)   %k
//           = ((a^123456)^10)%k *   a^7%k)   %k
//           = (a^123456 % k)^10%k * a^7%k)   %k

func superPow(a int, b []int) int {
	if len(b) == 0 {
		return 1
	}
	last := b[len(b)-1]
	b = b[:len(b)-1]
	return powmod(superPow(a, b), 10) * powmod(a, last) % 1337
}

func powmod(a, k int) int {
	result := 1
	a %= 1337
	for i := 0; i < k; i++ {
		result = (result * a) % 1337
	}
	return result
}
