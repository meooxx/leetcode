package main

import "fmt"

func main() {
	fmt.Println(numDecodings2("12011")) //01234
	fmt.Println(numDecodings2("1123"))
	fmt.Println(numDecodings2("2101"))
	fmt.Println(numDecodings2("2001"))
	fmt.Println(numDecodings2("230"))

}

// 1
// 1 1,    11
// 1 1 2   11 2  1 12
// dp[i-2] 每一项 + 12
// dp[i-1] 每一项 + 2
// dp[i] = dp[i-2] +dp[i-1]
func numDecodings(s string) int {
	count := 0
	var subDivide func(s string)

	subDivide = func(sub string) {
		if len(sub) == 0 {
			count++
			return
		}
		if sub[0] == '0' {
			return
		}
		subDivide(sub[1:])

		if len(sub) <= 1 || sub[0]-'0' > 2 {
			return
		}

		if sub[0]-'0' == 2 && sub[1]-'0' > 6 {
			return
		}

		subDivide(sub[2:])
	}
	subDivide(s)
	return count
}

func numDecodings2(s string) int {
	if s[0] == '0' {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	dp := make([]int, len(s)+1)
	dp[0] = 1
	dp[1] = 1
	for index := 2; index <= len(s); index++ {
		cur := s[index-1]
		pre := s[index-2]
		if cur >= '1' && cur <= '9' {
			dp[index] = dp[index-1]
		}
		if pre == '1' || pre == '2' && cur <= '6' {
			dp[index] += dp[index-2]
		}
	}
	return dp[len(dp)-1]
}
