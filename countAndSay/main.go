package main

import (
	// "strconv"
	"fmt"
)

func main() {
	r := countAndSay(5)
	fmt.Println(r)

	fmt.Println(countAndSayLoop(5))

}

// 1
// 11
// 21
// 1211
// 111221
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	str := countAndSay(n - 1)
	count := 0
	s := ""
	for i := 0; i < len(str)-1; i++ {
		count++
		if str[i] != str[i+1] {
			s += fmt.Sprint(count) + string(str[i])
			count = 0
		}
	}
	s += fmt.Sprint(count+1) + string(str[len(str)-1])

	return s
}

func countAndSayLoop(n int) string {
	if n == 1 {
		return "1"
	}
	s := "1"
	var r string
	for i := 1; i < n; i++ {
		r = ""
		var b []byte
		count := 0
		for j := 0; j < len(s)-1; j++ {
			count++
			if s[j] != s[j+1] {
				b = append(b, byte(count+'0'), s[j])
				count = 0
			}
		}
		b = append(b, byte(count+1+'0'), s[len(s)-1])
		r = string(b)
		s = r
	}
	return r
}
