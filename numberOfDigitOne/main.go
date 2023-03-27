package main

func countDigitOne(n int) int {
	//  121 one's digital '1'
	// 	121/10 = 12, 121 % 10 = 1
	//  12 * 1 + 1 % 1 + 1
	//  1 11 21 31 41 51 61 71 81 91 101 111 121

	//  121 10's digital '1'
	//  ten's digital
	// 	121 % 100 == 21
	//   ==>  1 * 10 + 10
	//  10  ~ 19 (10, 11, 12,....1x(19))
	//  ^     ^
	//  110 ~ 119
	
	// 121 hun's digital '1'
	// 100 ~ 121 0 + 121 % 10 + 1
	
	// d = q%10
	// if d == 0 => count += q * digital
	// if d == 1 => count += q * digital + d % digital + 1
	// if d > 1 => count += q * digital + digital
	q := n
	digital := 1
	count := 0
	for q > 0 {
		d := q % 10
		q = q / 10
		if d == 0 {
			count += digital * q
		} else if d == 1 {
			count += digital*q + n%digital + 1
			// d > 1
		} else {
			count += digital*q + digital
		}
		digital *= 10
	}
	return count
}
