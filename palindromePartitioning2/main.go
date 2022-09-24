package main

import "fmt"

func main() {
	fmt.Println(minCut("abbaa"))
	fmt.Println(minCut("apjesgpsxoeiokmqmfgvjslcjukbqxpsobyhjpbgdfruqdkeiszrlmtwgfxyfostpqczidfljwfbbrflkgdvtytbgqalguewnhvvmcgxboycffopmtmhtfizxkmeftcucxpobxmelmjtuzigsxnncxpaibgpuijwhankxbplpyejxmrrjgeoevqozwdtgospohznkoyzocjlracchjqnggbfeebmuvbicbvmpuleywrpzwsihivnrwtxcukwplgtobhgxukwrdlszfaiqxwjvrgxnsveedxseeyeykarqnjrtlaliyudpacctzizcftjlunlgnfwcqqxcqikocqffsjyurzwysfjmswvhbrmshjuzsgpwyubtfbnwajuvrfhlccvfwhxfqthkcwhatktymgxostjlztwdxritygbrbibdgkezvzajizxasjnrcjwzdfvdnwwqeyumkamhzoqhnqjfzwzbixclcxqrtniznemxeahfozp"))
	fmt.Println(minCut1("abbaa"))
	fmt.Println(minCut1("apjesgpsxoeiokmqmfgvjslcjukbqxpsobyhjpbgdfruqdkeiszrlmtwgfxyfostpqczidfljwfbbrflkgdvtytbgqalguewnhvvmcgxboycffopmtmhtfizxkmeftcucxpobxmelmjtuzigsxnncxpaibgpuijwhankxbplpyejxmrrjgeoevqozwdtgospohznkoyzocjlracchjqnggbfeebmuvbicbvmpuleywrpzwsihivnrwtxcukwplgtobhgxukwrdlszfaiqxwjvrgxnsveedxseeyeykarqnjrtlaliyudpacctzizcftjlunlgnfwcqqxcqikocqffsjyurzwysfjmswvhbrmshjuzsgpwyubtfbnwajuvrfhlccvfwhxfqthkcwhatktymgxostjlztwdxritygbrbibdgkezvzajizxasjnrcjwzdfvdnwwqeyumkamhzoqhnqjfzwzbixclcxqrtniznemxeahfozp"))
}
func minCut(s string) int {

	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}
	cuts := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		minCut := i
		for j := 0; j <= i; j++ {
			if s[i] == s[j] && (i-j <= 2 || dp[j+1][i-1]) {
				dp[j][i] = true
				if j == 0 {
					minCut = 0
				} else {
					minCut = min(minCut, cuts[j-1]+1)
				}

			}
		}
		cuts[i] = minCut
	}
	return cuts[len(s)-1]
}

func minCut1(s string) int {
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
	// center of a palinrome str
	for i := 0; i < len(s); i++ {
		//0 1 2 3 4 5 cuts
		//  a b b a
		//  a b c b a
		for left, right := i, i; left >= 0 && right < len(s) && s[left] == s[right]; {
			cuts[right+1] = min(cuts[right+1], 1+cuts[left])
			left--
			right++
		}
		for left, right := i, i+1; left >= 0 && right < len(s) && s[left] == s[right]; {
			cuts[right+1] = min(cuts[right+1], 1+cuts[left])

			left--
			right++
		}
	}
	return cuts[len(cuts)-1]
}

// not pass
func minCut2(s string) int {
	result := 0
	if len(s) <= 1 {
		return result
	}

	mp := map[string]bool{}
	isPalindrome := func(s string) bool {
		if len(s) <= 1 {
			return true
		}
		if v, ok := mp[s]; ok {
			return v
		}
		for left, right := 0, len(s)-1; left < right; {
			if s[left] != s[right] {
				mp[s] = false
				return false
			}
			left++
			right--
		}
		mp[s] = true
		return true
	}
	if isPalindrome(s) {
		return result
	}
	queue := []string{s}
	for len(queue) > 0 {
		qSize := len(queue)
		result++
		for i := 0; i < qSize; i++ {
			str := queue[0]
			queue = queue[1:]
			// a ab
			for i := 1; i <= len(str); i++ {
				leftIsPal := isPalindrome(str[:i])
				rightIsPal := isPalindrome(str[i:])
				if leftIsPal && rightIsPal {
					return result
				} else if leftIsPal {
					queue = append(queue, str[i:])
				}
			}
		}
	}
	return result
}
