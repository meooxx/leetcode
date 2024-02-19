package main

import "fmt"

func main() {

	fmt.Println(countDigitOne(14))
}


// e.g.
//  3 2 2 0     d > 1 [3100, 3199] + [2100, 2199]+...
//                      partial of d > 1
//  3 1 2 0     d == 1 [3100, 3120] +  [2100, 2199] + [1100, 1199]...

func countDigitOne(n int) int {

	base := 1
	anwser := 0
	q := n
	for q > 0 {
		digit := q % 10
		q /= 10
		anwser += q * base
		if digit == 1 {
			anwser += n%base + 1
		}
		if digit > 1 {
			anwser += base
		}
		base *= 10
	}

	return anwser

}
