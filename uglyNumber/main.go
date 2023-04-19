package main

func isUgly(n int) bool {
	if n == 0 {
		return false
	}
	primes := []int{2, 3, 5}
	for i := range primes {
		for n%primes[i] == 0 {
			n /= primes[i]
		}
	}

	return n == 1

}
