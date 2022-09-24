package main

import "fmt"

func main() {
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(4))
	fmt.Println(climbStairs(5))

}

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	t := []int{1, 2}
	for i := 3; i <= n; i++ {
		t0 := t[0]
		t[0] = t[1]
		t[1] = t0 + t[1]
	}
	return t[1]
}
