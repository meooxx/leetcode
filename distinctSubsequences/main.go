package main

import "fmt"

func main() {
	fmt.Println(numDistinct("rabbbit", "rabbit"))
	fmt.Println(numDistinct("babgbag", "bag"))
	fmt.Println(numDistinct("aabb", "ab"))

}
func numDistinct(s string, t string) int {
	dp := make([]int, len(t)+1)
	dp[0] = 1
	// 由2可知, 如果从g开始, 就不再需要变量pre了
	// 因为从后面往前 t[tIndex] 覆盖后不再使用了
	// 方法2从前往后 t[tIndex], t[tIndex-1]获取到的值不对了
	// tagbg
	// tag
	//       t a g g
	//     1 1 1 1 1
	//   t 0            0  tIndex
	//   a 0 0          1
	///  g 0 0          2
	for sIndex := 1; sIndex <= len(s); sIndex++ {
		for tIndex := len(t); tIndex >= 1; tIndex-- {
			if s[sIndex-1] == t[tIndex-1] {
				dp[tIndex] += dp[tIndex-1]
			}
		}
	}
	return dp[len(t)]
}

func numDistinct2(s string, t string) int {
	dp := make([]int, len(t)+1)
	dp[0] = 1
	// https://leetcode.com/problems/distinct-subsequences/discuss/572192/Understand-DP-through-question-115-explanation-with-pictures
	// rabbbit
	// rabi
	//       r a b b b i t
	//     1 1 1 1 1 1 1 1
 	//   r 0 1 1 1 1 1       0 tIndex
	//   a 0 0 1 1 2 1       1
	///  b 0 0 0 1 2 4       2  
	//   b 0 0 0 0 1 _       3 _ ? 因为 b == b, 如果这个b不使用,需要s去匹配rabb的数量 n(4,4) == 1
 	//   i           			如果使用b使用从s 去匹配rab即 n(3,4)最终结果即 n(3,4) + n(4,4)
	//   t
	for sIndex := 1; sIndex <= len(s); sIndex++ {
		pre := 1
		for tIndex := 1; tIndex <= len(t); tIndex++ {
			temp := dp[tIndex]
			if s[sIndex-1] == t[tIndex-1] {
				dp[tIndex] = dp[tIndex] + pre
			}
			pre = temp
		}
	}
	return dp[len(t)]
}
