package main

//    1    1        0   			1       0       0
//  0-9   (1-9) * (10-1)     (1-9)*(10-1)*(10-2)
//  if n > 10,  P(n) == P(10)

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 0
	}
	result := 10
	base := 9
	factor := 9
	for k := 2; k <= n && k <= 10; k++ {
		base *= factor
		result += base
		factor--
	}
	return result
}
