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
	result := ""
	fMap := map[int]int{}
	d := numerator / denominator
	n := (numerator % denominator) * 10
	floatStr := ""
	index := 1
	for n > 0 && fMap[n] == 0 {
		fMap[n] = index
		if n > denominator {
			floatStr += fmt.Sprint(n / denominator)
			n = (n % denominator) * 10
		} else {
			floatStr += "0"
			n *= 10
		}
		index++
	}
	if neg {
		result = "-"
	}
	result += fmt.Sprint(d)
	if floatStr != "" {
		result += "."
	}
	if n > 0 {
		result += fmt.Sprintf("%s(%s)", floatStr[:fMap[n]-1], floatStr[fMap[n]-1:])
	} else {
		result += floatStr

	}
	return result
}
