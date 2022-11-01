package main

func wordBreak(s string, wordDict []string) []string {
	wordmp := map[string]bool{}
	for i := range wordDict {
		wordmp[wordDict[i]] = true
	}
	dp := make([][]string, len(s)+1)
	dp[0] = []string{}
	// [cat, cats] -> ["cat sand", "cats and"]
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			word := s[i:j]
			if dp[i] != nil && wordmp[word] {
				if len(dp[i]) == 0 {
					dp[j] = []string{word}
					continue
				}
				for k := range dp[i] {
					fragment := dp[i][k] + " " + word
					dp[j] = append(dp[j], fragment)
				}

			}
		}
	}
	return dp[len(s)]
}
