package main

func main() {}

// 由注释得将t反序计算一样的
func numDistinct(s string, t string) int {
	dp := make([]int, len(t)+1)
	// 默认就是 0
	// for i := range dp {
	// 	dp[i] = 0
	// }
	dp[0] = 1
	for sIndex := 0; sIndex < len(s); sIndex++ {
		for tIndex := len(t) - 1; tIndex >= 0; tIndex-- {
			if s[sIndex] == t[tIndex] {
				dp[tIndex+1] += dp[tIndex]
			}
		}
	}
	return dp[len(t)]
}

//     r a b    b    b i t
//   1 1 1 1    1    1 1 1
// r 0 1     |
// a 0 0     v
// b 0      2(3,4)
// b 0      1(4,4) ?(4,5)
//            使用 s[5]的 b 则从s[0:5]中找一个'b', 即sequence 'rb', 则等于dp(3,4) == 2
//            不使用 s[5] b 则从s[0:5]中找俩个'b'即 'rabb' 等于dp(4,4) == 1
func numDistinct1(s string, t string) int {
	dp := make([]int, len(t)+1)
	for index := range dp {
		dp[index] = 0
	}
	for sIndex := 0; sIndex < len(s); sIndex++ {
		pre := 1
		for tIndex := 0; tIndex < len(t); tIndex++ {
			temp := dp[tIndex+1]
			if s[sIndex] == t[tIndex] {
				dp[tIndex+1] = dp[tIndex+1] + pre
			}
			pre = temp
		}
	}
	return dp[len(t)]
}
