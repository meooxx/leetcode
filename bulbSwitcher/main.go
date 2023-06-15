package main

import "math"

// i-th bulb will be switched k times
// like 6-th, Round: 1 on, 2 off, 3 on, 6 off
// these nums 1, 2, 3 , 6 are all factors of 6
//
//	so if k is odd, then the bulb is on
//	else k is even, then the bulb is off
//
// only perfect square number has odd nums of factors
// like 4(1,2,4), 9(1,3,9),
//
//	we need to find out how many perfect square
//	that is no more than n
//	and the num of perfect square is equal square root  of n
func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n)))
}
