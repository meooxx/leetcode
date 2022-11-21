package main

import "fmt"

func main() {
	fmt.Println(fractionToDecimal(1, 4))
	fmt.Println(fractionToDecimal(8, 4))
	fmt.Println(fractionToDecimal(4, 333))
	fmt.Println(fractionToDecimal(1, 3))
	fmt.Println(fractionToDecimal(1, 6))

	fmt.Println(fractionToDecimal(9, 4))
	fmt.Println(fractionToDecimal(22, 7))
	fmt.Println(fractionToDecimal(500, 10))
	fmt.Println(fractionToDecimal(-50, 8))
	fmt.Println(fractionToDecimal(-50, 80))

}

// case: numerator 0
// case: 0.12323 ==> 0.1(23)
// case 4/2 ==> 2
// case -1/2 ==> -0.5
// 1/4 =>> 0.25
func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	neg := false
	if numerator < 0 {
		neg = !neg
		numerator = -numerator
	}
	if denominator < 0 {
		neg = !neg
		denominator = -denominator
	}

	d := 0
	f := ""

	n := 0
	// n1Str := fmt.Sprint(numerator)
	// for i := 0; i < len(n1Str); i++ {
	// 	n = int(n1Str[i]-'0') + n*10
	// 	if n >= denominator {
	// 		d = d*10 + n/denominator
	// 		n = n % denominator
	// 	} else {
	// 		d *= 10
	// 	}
	// }
	// aboving code is equal below
	d = numerator / denominator
	n = numerator % denominator
	repeating := map[int]int{}
	n *= 10
	index := 1
	for n != 0 && repeating[n] == 0 {
		repeating[n] = index
		// 1 / 6
		// 10/6,  1 4
		// 40/6   6 4
		if n >= denominator {
			f += fmt.Sprint(n / denominator)
			n = n % denominator
		} else {
			f += "0"
		}
		index++
		n *= 10
	}
	anwser := ""
	if neg {
		anwser += "-"
	}
	anwser += fmt.Sprint(d)
	if f != "" {
		anwser += "."
	}
	if n != 0 {
		return anwser + fmt.Sprintf("%s(%s)", f[:repeating[n]-1], f[repeating[n]-1:])
	}
	return anwser + f
}
