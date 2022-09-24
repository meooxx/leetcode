package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(divide(5,
		-8))
	fmt.Println(divide(-math.MinInt32,
		1))

}

// 15 / 2      0
// 2^3         3
// 15-2^3 / 2
// 2^2         3+2
// 7-2^2 / 2
// 2^1        3+2+1
func divide(dividend, divisor int) int {
	if dividend == 0 || divisor == 0 {
		return 0
	}

	negative := true
	if dividend > 0 && divisor > 0 || dividend < 0 && divisor < 0 {
		negative = false
	}

	dividend = int(math.Abs(float64(dividend)))
	divisor = int(math.Abs(float64(divisor)))

	n := 0
	for divisor <= dividend {
		i := 1
		d := divisor
		for d<<1 <= dividend {
			d = d << 1
			i = i << 1

		}
		dividend = dividend - d
		n += i
	}

	if negative {
		if -n < math.MinInt32 {
			return math.MinInt32
		}
		return -n
	}
	if n > math.MaxInt32 {
		return math.MaxInt32
	}
	return n

}
