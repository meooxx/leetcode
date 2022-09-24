package main

import (
	"fmt"
)

// Expand around center
func longestPalindromeExpand(s string) string {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		// aba
		//  |
		// left,right from b
		left1, right1 := expandAroundCenter(s, i, i)
		// abba
		//  ||
		// left,right from b and b
		left2, right2 := expandAroundCenter(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	fmt.Println(start, end)
	return s[start : end+1]

}

// 如果 abbbba   bbbb是回文, 如果a==a, 则abbbba一定是回文
// 以为 i(left) < j(right), 只填写T右上边的值

// i\j  0     1      2     3
//      a     b      b     a
// 0 a  T   t(1)    t(2)  t(4)
// 1 b       T      t(3)  t(5)
// 2 b         	    T     t(6)
// 3 a         		  	  T

// 且 i, j关系如下
// j == 3
// i ==0, dp[0][3] = T/F
// i ==1, dp[1][3] = T/F
// i ==2, dp[2][3] = T/F

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	dp := make([][]bool, len(s))
	// initalize 2D slice
	// set dp[i][i] true
	for i := range dp {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	maxStr := s[0:1]
	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			if s[i] == s[j] {
				// j -i + 1 <= 3
				// case aba, a==a, true
				if j-i <= 2 {
					dp[i][j] = true
				} else {
					// abba
					// if bb true, then abba true
					dp[i][j] = dp[i+1][j-1]
				}
				if dp[i][j] && j-i+1 > len(maxStr) {
					maxStr = s[i : j+1]
				}
			}

		}

	}
	return maxStr
}

func expandAroundCenter(s string, left, right int) (int, int) {
	// bab
	//  |
	// a == a left-1 = 0, right+1=2
	// bab
	// | |
	// b == b left -1 = -1, right + 1 =3
	//  bab
	// |   | 结束 -1, 3 return 0,2
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left -= 1
		right += 1
	}
	// left, right 总是在有效范围外 + 1
	// aba 最后一个有效 0,2满足条件, 又继续循环一次 left -1,  right 3
	return left + 1, right - 1
}

func main() {
	s := longestPalindrome("aba")
	s1 := longestPalindrome("abba")
	fmt.Println(s, s1)
}
