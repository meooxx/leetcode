package main

import (
	"fmt"
)

func main() {
	fmt.Println(generateParenthese(3))

}

func generateParenthese(n int) []string {

	result := []string{}
	genPar(&result, "", n, n)
	return result
}

// n == 2
//
//
//		 (( => (())
//  ( => () => ()()    
// 

func genPar(result *[]string, s string, left, right int) {
	if left == 0 && right == 0 {
		*result = append(*result, s)
	}
	if left > 0 {
		genPar(result, s+"(", left-1, right)
	}
	if right > left {
		genPar(result, s+")", left, right-1)
	}

}
