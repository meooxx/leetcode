package main	

import (
	"math"
	"fmt"
)

/**
 * 1 the given n was reordered the digits in any order
 * 123 can be 132, 321
 * 
 * shift left the 1 from 0 -> 32 (1, 2,4...MaxInt32)
 * 		one of reorder digits == shitf value
 * n is power of 2 
 * 	
 * 10^3 + 10^2 + 10^1 == 10 ^1 + 10^2 + 10^3 
 * == 10^3 + 10^1 + 10^2
 * 
 */

func main(){
	fmt.Println(reorderedPowerOf2(125))
	//fmt.Println(count(111))
}

func count(n int) float64 {
	var r float64
	for n>0 {
		r += math.Pow10(n%10)
		n= n/10
	}
	return r
	
}

func reorderedPowerOf2(n int) bool {
	result := 0
	match:= count(n)
	var i uint
	
	for i = 0; result < math.MaxInt32; i++ {
		if result =1<<i; count(result) == match {
			return true
		}
	}
	return false
}