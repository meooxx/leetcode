package main

func titleToNumber(columnTitle string) int {
	anwser := 0
	right := len(columnTitle) - 1
	base := 1
	for ; right >= 0; right-- {
		n := int(columnTitle[right]-'A') + 1
		anwser += n * base
		base *= 26
	}
	return anwser
}
