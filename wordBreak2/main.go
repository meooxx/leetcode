package main

import (
	"fmt"
)

func main() {
	fmt.Println(wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
}

func wordBreak(s string, wordDict []string) []string {

	wordmp := map[string]bool{}
	for _, w := range wordDict {
		wordmp[w] = true
	}

	dp := make([][]string, len(s)+1)
	dp[0] = []string{}

	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			word := s[i:j]
			if dp[i] != nil && wordmp[word] {
				if len(dp[i]) == 0 {
					dp[j] = []string{word}
				} else {
					for _, preword := range dp[i] {
						str := preword + " " + word
						dp[j] = append(dp[j], str)
					}
				}
			}
		}
	}

	return dp[len(s)]
}
