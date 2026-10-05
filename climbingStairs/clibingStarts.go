package main

func climbStairs(n int) int {
	if n <= 3 {
		return n
	}
	n1 := 3
	n2 := 2

	for i := 4; i <= n; i++ {
		n2, n1 = n1, n2+n1
	}
	return n1

}
