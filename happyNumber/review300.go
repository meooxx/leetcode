package main

func isHappy(n int) bool {

	// incase of the loop
	cache := map[int]bool{}
	for n > 0 && !cache[n] {
		if n == 1 {
			return true
		}
		result := 0
		cache[n] = true
		for n > 0 {
			rest := n % 10
			result += rest * rest
			n /= 10
		}

		n = result
	}
	return false
}
