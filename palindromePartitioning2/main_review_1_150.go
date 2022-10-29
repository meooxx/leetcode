package main

import "fmt"

func main() {
	fmt.Println(minCut("abbab"))
}

//   a b c b b a
// a T F F
// b   T F
// c     T
// b
func minCut1(s string) int {
	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	minCuts := make([]int, len(s)+1)
	for i := range minCuts {
		minCuts[i] = i - 1
	}
	for i := 0; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			if s[i] == s[j] && (i-j <= 2 || dp[j+1][i-1]) {
				dp[j][i] = true
				if dp[j][i] {
					minCuts[i+1] = min(minCuts[i+1], minCuts[j]+1)
				}
			}
		}
	}
	return minCuts[len(s)]
}

func minCut(s string) int {
	cuts := make([]int, len(s)+1)
	for i := range cuts {
		cuts[i] = i - 1
	}
	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}
	// 左侧切
	//    a b a  c
	// -1 0      1
	// dabbac
	for i := range s {
		left, right := i, i
		for left >= 0 && right < len(s) && s[left] == s[right] {
			cuts[right+1] = min(cuts[right+1], cuts[left]+1)
			left--
			right++
		}
		left, right = i, i+1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			cuts[right+1] = min(cuts[right+1], cuts[left]+1)
			left--
			right++
		}
	}

	return cuts[len(s)]

}
