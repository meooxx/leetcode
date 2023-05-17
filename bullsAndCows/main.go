package main

import "fmt"

func getHint0(secret string, guess string) string {
	mp := make(map[byte]int)
	aCount := 0
	bCount := 0
	for i := range secret {
		mp[secret[i]]++
		if secret[i] == guess[i] {
			aCount++
			mp[guess[i]]--
		}
	}
	for i := range secret {
		if secret[i] != guess[i] && mp[guess[i]] > 0 {
			bCount++
			mp[guess[i]]--
		}
	}
	return fmt.Sprintf("%dA%dB", aCount, bCount)
}

func getHint(secret string, guess string) string {
	nums := [10]int{}
	aCount := 0
	bCount := 0
	for i := range secret {
		if secret[i] == guess[i] {
			aCount++
		} else {
			i1 := int(secret[i] - '0')
			i2 := int(guess[i] - '0')
			// case:  12	2 1   secret i1
			//        12	1 2   guess  i2
			//          	| |
			//          	|_|____ nums[i1] < 0
			//          	 	|_ nums[i2] > 0
			if nums[i1] < 0 {
				bCount++
			}
			if nums[i2] > 0 {
				bCount++
			}
			nums[i1]++
			nums[i2]--

		}
	}

	return fmt.Sprintf("%dA%dB", aCount, bCount)
}
