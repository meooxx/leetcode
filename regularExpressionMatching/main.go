package main

import "fmt"

func main() {
	// fmt.Println(isMatch("abbbc", "ab*c"))
	// fmt.Println(isMatch("a", ".*"))
	// fmt.Println(isMatch("avbbb", "avbbb"))

	fmt.Println(isMatch("aab", "c*a*b"))

}

//         0  1   2    3
//        ""  1   2   3    4
//            a   b   *    c
// 0 ""   T
// 1 a
// 2 b
// 3 b
//   b
//   c

// 1, If p.charAt(j) == s.charAt(i) :  dp[i][j] = dp[i-1][j-1];
// 2, If p.charAt(j) == '.' : dp[i][j] = dp[i-1][j-1];
// 3, If p.charAt(j) == '*':
//    here are two sub conditions:
//                1   if p.charAt(j-1) != s.charAt(i) : dp[i][j] = dp[i][j-2]  //in this case, a* only counts as empty
//                2   if p.charAt(j-1) == s.charAt(i) or p.charAt(i-1) == '.':
//                               dp[i][j] = dp[i-1][j]    //in this case, a* counts as multiple a
//                            or dp[i][j] = dp[i][j-1]   // in this case, a* counts as single a
//                            or dp[i][j] = dp[i][j-2]   // in this case, a* counts as empty

func isMatch3(s string, p string) bool {
	m := make([][]bool, len(s)+1)
	for index := range m {
		m[index] = make([]bool, len(p)+1)
	}
	m[0][0] = true
	// case: d vs  c*d
	for i := 0; i < len(p); i++ {
		if p[i] == '*' {
			// +1-2 from 0
			// a*a*
			m[0][i+1] = m[0][i-1]
		}
	}
	// ab*c   
	// abbc   
	//  a-a T,  a-ab F,  a-ab* T, a-ab*c F,
	//  b-a F, b-ab T(ab-ab), b-ab* T, b-ab*c F
	//  b-aF, b-ab F(abb-ab)...
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(p); j++ {
			if s[i] == p[j] || p[j] == '.' {
				m[i+1][j+1] = m[i][j]
			}
			if p[j] == '*' {
				if p[j-1] == s[i] || p[j-1] == '.' {
					m[i+1][j+1] = m[i+1][j-1] || m[i+1][j] || m[i][j+1]
				} else {
					m[i+1][j+1] = m[i+1][j-1]
				}
			}
		}
	}

	return m[len(s)][len(p)]

}

func isMatch(s string, p string) bool {
    dp := make([][]bool, len(p) + 1)
    for i:= 0;i<=len(p);i++{
        dp[i] = make([]bool, len(s)+1)
    }
	dp[0][0] = true

	for i:=1;i<len(p);i++{
		if p[i] == '*' {
			dp[i+1][0] = dp[i-1][0]
		}
	}
    
    for pIndex := 0;pIndex<len(p);pIndex++ {
        for index:=0;index<len(s);index++{
            if s[index] == p[pIndex] || p[pIndex] == '.' {
                dp[pIndex+1][index+1] = dp[pIndex][index]    
            }else if p[pIndex] == '*' {
				// aa a*
				// abc ab*c
				if p[pIndex-1] == '.' || s[index] == p[pIndex-1] {
					dp[pIndex+1][index+1] = dp[pIndex][index+1] || dp[pIndex+1][index] || dp[pIndex-1][index+1]
				}else {
					dp[pIndex+1][index+1] = dp[pIndex-1][index+1]
				}
            }
            
        }
    }
    return dp[len(p)][len(s)]
    
}
