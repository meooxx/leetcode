package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myPow(2, -3))
	fmt.Println(myPow(2, 3))
	fmt.Println(myPow(2, 2))
}

// 3 ^ 11
// 11 == ob1011

//                       0b(1011)
//                           3^2 =m   m^2 =m*m = n     n*n
//                         3*3^2*   (0*3^4) *     3 ^8
// 3^11= 3* 3^2 * 3^8 = 3^(1+2+8)
// 3*2^0  * 3^2 * 3^4 * 3^8 * 3^16....
// 转化成 (x|1) * x^2 * x^4 ... x^(2^n)
func myPow(x float64, n int) float64 {
	if x == 0 || x == 1 {
		return x
	}
	times := int(math.Abs(float64(n)))

	var sum float64 = 1
	for times > 0 {
		if times&1 == 1 {
			sum *= x
		}
		times >>= 1
		x *= x

	}

	if n > 0 {
		return sum
	}
	return 1 / sum
}
