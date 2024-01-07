package main

import "fmt"

func getHint(secret string, guess string) string {
	nums := [10]int{}

	Acount := 0
	Bcount := 0

	for i := range secret {
		if secret[i] == guess[i] {
			Acount++
		} else {
			i1 := int(secret[i] - '0')
			i2 := int(guess[i] - '0')
			// guess 中有过 i2
			if nums[i1] < 0 {
				Bcount++
			}
			// secret中存在i2
			if nums[i2] > 0 {
				Bcount++
			}
			nums[i1]++
			nums[i2]--
		}
	}
	return fmt.Sprintf("%dA%dB", Acount, Bcount)

}
