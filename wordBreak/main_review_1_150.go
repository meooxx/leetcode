package main

func wordBreak(s string, wordDict []string) bool {
	wordMap := map[string]bool{}
	for i := range wordDict {
		wordMap[wordDict[i]] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := range s {
		for j := i + 1; j <= len(s); j++ {
			word := s[i:j]
			if wordMap[word] && dp[i] {
				dp[j] = true
			}
		}
	}
	return dp[len(s)]
}