package main

import (
	"fmt"
)

func main() {
	fmt.Println(minDistance("abc", "dba"))
	fmt.Println(minDistance("intenton", "execution"))
	fmt.Println(minDistance("a", "1"))
}

// 1  abc -> ac 如果 c ==c 等于 ab->a
// 2  abd -> ac    d !=c
// 如果已知替换ab->ac, 求abd->ac 加一步删除d即可,   i-1, j
// 已知ab ->a, 求abd -> ac  加一次替换d->c 即可,   i-1, j-1
// abd -> a, 求abd -> ac   加一次步添加一个c即可  i, j-1
func minDistance(word1 string, word2 string) int {
	r := make([]int, len(word2)+1)

	for index := range r {
		r[index] = index
	}
	for i := 1; i <= len(word1); i++ {
		leftTop := r[0]
		r[0] = i
		for j := 1; j <= len(word2); j++ {
			tmp := r[j]
			if word1[i-1] == word2[j-1] {
				r[j] = leftTop
			} else {
				r[j] = min(min(leftTop, r[j-1]), r[j]) + 1
			}
			leftTop = tmp

		}
	}
	return r[len(r)-1]
}

func minDistance1(word1 string, word2 string) int {
	if word1 == "" {
		return len(word2)
	}
	if word2 == "" {
		return len(word1)
	}
	//      a  b   c    i-->
	//   0  1  2   3
	// d 1  1           j
	// b 2  2
	// a 3  2

	dp := make([][]int, len(word1)+1)
	for index := range dp {
		dp[index] = make([]int, len(word2)+1)
		dp[index][0] = index
	}
	for index := range dp[0] {
		dp[0][index] = index
	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(min(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1]) + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}
