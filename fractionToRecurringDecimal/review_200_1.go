package main

import "fmt"

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	neg := false
	if numerator < 0 {
		neg = true
		numerator = -numerator
	}
	if denominator < 0 {
		neg = !neg
		denominator = -denominator
	}
	d := numerator / denominator
	n := (numerator % denominator) * 10
	floatNum := ""
	// why index start from 1?
	// CUZ' 0 is stand for 'not existing'
	// and 0 is default value of int
	index := 1
	// num -> index
	fMap := map[int]int{}
	for n > 0 && fMap[n] == 0 {
		fMap[n] = index
		if n >= denominator {
			floatNum += fmt.Sprint(n / denominator)
			n = n % denominator * 10
		} else {
			floatNum += "0"
			n = n * 10
		}
		index++
	}

	result := ""
	if neg {
		result += "-"
	}
	result += fmt.Sprint(d)
	if floatNum != "" {
		result += "."
	}
	fmt.Println(fMap[n])
	if n != 0 {
		result += fmt.Sprintf("%s(%s)", floatNum[:fMap[n]-1], floatNum[fMap[n]-1:])
	} else {
		result += floatNum
	}
	return result

}
