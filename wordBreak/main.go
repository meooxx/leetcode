package main

import "fmt"

func main() {
	fmt.Println(wordBreak("leetcode", []string{
		"leet", "code",
	}))
	fmt.Println(wordBreak("applepenapple", []string{
		"apple", "pen",
	}))
	fmt.Println(wordBreak("catsandog", []string{
		"cats", "dog", "sand", "and", "cat",
	}))
	fmt.Println(wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", []string{
		"a", "aa", "aaaa", "aaaaa", "aaaaaaa",
	}))
	fmt.Println(wordBreakDp("leetcode", []string{
		"leet", "code",
	}))
	fmt.Println(wordBreakDp("applepenapple", []string{
		"apple", "pen",
	}))
	fmt.Println(wordBreakDp("catsandog", []string{
		"cats", "dog", "sand", "and", "cat",
	}))
	fmt.Println(wordBreakDp("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", []string{
		"a", "aa", "aaaa", "aaaaa", "aaaaaaa",
	}))
	fmt.Println(wordBreakDp("aaaaaaa", []string{
		"aaa", "aaaa",
	}))
}

func wordBreak(s string, wordDict []string) bool {
	mp := map[string]bool{}
	for _, w := range wordDict {
		mp[w] = true
	}
	cache := map[string]bool{}
	var nextWord func(s string) bool
	nextWord = func(s string) bool {
		if s == "" {
			return true
		}
		if v, ok := cache[s]; ok {
			return v
		}
		for i := 1; i <= len(s); i++ {
			word := s[:i]
			if mp[word] {
				if nextWord(s[i:]) {
					cache[s[i:]] = true
					return true
				}
			}
		}
		cache[s] = false
		return false
	}
	return nextWord(s)
}

func wordBreakDp(s string, wordDict []string) bool {
	wordmp := map[string]bool{}
	for _, w := range wordDict {
		wordmp[w] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			word := s[i:j]
			if dp[i] && wordmp[word] {
				dp[j] = true

			}
		}
	}
	return dp[len(s)]

}
