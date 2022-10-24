package main

import "fmt"

func main() {

	// fmt.Println(isInterleave("1", "2", "12"))

	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbcbcac"))

	fmt.Println(isInterleave("aabcc", "", "aadbbcbcac"))
	fmt.Println(isInterleave("", "aabcc", "aadbbcbcac"))
	fmt.Println(isInterleave("ab", "bc", "bbac"))
	fmt.Println("isInterleave3", isInterleave3("aabcc", "dbbca", "aadbbcbcac"))
	fmt.Println("isInterleave3", isInterleave3("aabcc", "", "aadbbcbcac"))
	fmt.Println("isInterleave3", isInterleave3("", "aabcc", "aadbbcbcac"))
	fmt.Println("isInterleave3", isInterleave3("ab", "bc", "bbac"))
}


//  dp[i][j] 表示 s1[:i]  和 s2[:j]  的插入结果
// abc  de    
// ||   
// ab   de c   此时 s1[i] == s3[i+j] 要看 s1第i前一位的情况dp[i-1][j]
// 如果 s2[j] == s3[i+j] 要看j前面一位的情况, 所以查看dp[i][j-1]
// 又因为外层循环是i 所以可以一位数组 dp[0:j]
// 新循环中(i)dp[j] 表示dp[i-1][j]
// j loop中, dp[j-1] 表示dp[i][j-1]

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	// space O(min(s1, s2))
	if len(s2) > len(s1) {
		s1, s2 = s2, s1
	}
	dp := make([]bool, len(s2)+1)
	dp[0] = true
	for index := range s2 {
		if s2[index] == s3[index] {
			dp[index+1] = dp[index]
		}
	}
	for i := 1; i <= len(s1); i++ {
		dp[0] = dp[0] && s1[i-1] == s3[i-1]
		for j := 1; j <= len(s2); j++ {
			dp[j] = dp[j] && s1[i-1] == s3[i+j-1] || dp[j-1] && s2[j-1] == s3[i+j-1]
		}
	}
	return dp[len(s2)]
}

func isInterleave2(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	if len(s2) > len(s1) {
		s1, s2 = s2, s1
	}
	dp := make([][]bool, len(s1)+1)

	for index := range dp {
		dp[index] = make([]bool, len(s2)+1)
	}
	// ""+"" == ""
	dp[0][0] = true
	for index := 1; index <= len(s1); index++ {
		if index-1 < len(s2) && s3[index-1] == s2[index-1] {
			dp[0][index] = dp[0][index-1]
		}
		if index-1 < len(s1) && s3[index-1] == s1[index-1] {
			dp[index][0] = dp[index-1][0]
		}
	}

	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s3[i+j-1] && dp[i-1][j] {
				dp[i][j] = true
			}
			if s2[j-1] == s3[i+j-1] && dp[i][j-1] {
				dp[i][j] = true
			}

		}
	}
	return dp[len(s1)][len(s2)]
}

// 递归尝试(未验证所有测试)
// abc   de   abcde
// a          a
// ab         ab
// abc        abc
//      d     abcd
//      de    abcde 
func isInterleave3(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	var matchStr func(s1, s2, s3 string, s1Index, s2Index, s3Index int) bool

	matchStr = func(s1, s2, s3 string, s1Index, s2Index, s3Index int) bool {
		if s3Index == len(s3) {
			return true
		}
		if s1Index == len(s1) {
			for s2Index < len(s2) && s2[s2Index] == s3[s3Index] {
				s2Index++
				s3Index++
			}
			return s3Index == len(s3)
		}
		if s2Index == len(s2) {
			for s1Index < len(s1) && s1[s1Index] == s3[s3Index] {
				s1Index++
				s3Index++
			}
			return s3Index == len(s3)
		}

		if s1[s1Index] == s3[s3Index] {
			if matchStr(s1, s2, s3, s1Index+1, s2Index, s3Index+1) {
				return true
			}
		}
		if s2[s2Index] == s3[s3Index] {
			if matchStr(s1, s2, s3, s1Index, s2Index+1, s3Index+1) {
				return true

			}
		}

		return false
	}

	return matchStr(s1, s2, s3, 0, 0, 0)

}
