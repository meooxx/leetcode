package main

import (
	"strconv"
	"strings"
	"fmt"
)


func main() {
	fmt.Println(readBinaryWatch(2))

}

func readBinaryWatch(n int) []string {
	res := []string{}
	for i := 0; i<12; i++ {
		for j := 0; j < 60; j++ {
			t:= int64(i<<6 + j)
			b := strconv.FormatInt(t,2)
			if strings.Count(b  , "1") == n {
				res = append(res, fmt.Sprintf("%d:%02.2d", i, j))
			}
		}
	}


	return res
}