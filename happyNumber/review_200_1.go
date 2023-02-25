package main

func isHappy(n int) bool {
	if n == 1 {
		return true
	}
	cache := map[int]bool{}
	for n > 0 && !cache[n] {
		cache[n] = true
		result := 0
		for n > 0 {
			rest := n % 10
			result += rest * rest
			n /= 10
		}
		if result == 1 {
			return true
		}
		n = result

	}
	return false
}
