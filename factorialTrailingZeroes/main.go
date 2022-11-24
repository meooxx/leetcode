package main

import "fmt"

func trailingZeroes1(n int) int {
	m := 1
	for n > 1 {
		m *= n
		n--
	}
	s := fmt.Sprint(m)
	zeroCount := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '0' {
			zeroCount++
		}
	}
	return zeroCount
}

// https://www.purplemath.com/modules/factzero.htm
// zero is comes 10, 10 = 10 , 100 = 10*10,...
// 10 == 5*2, and 2 is more than 5, so we can only consider 5
// but 25 == 5*5, every 25 has extra factor 
// 	e.g. 5 15 10 15 25(5 5) = 5 15 10 5(replace 25) 5(extra)
// 	
// n / 5 + n/5^2 + n/5^3
// n/5 + n/5/5 + n/5/5/5
// == n/5 => m => m/5 => g => g/5 == n/125
//  		n/5	       n/25         
func trailingZeroes(n int) int {
	sum := 0
	for n > 5 {
		sum += n / 5
		n = n / 5
	}
	return sum
}