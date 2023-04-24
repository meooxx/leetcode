package main

import "strings"

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	digit1 := []string{
		"", "One", "Two", "Three", "Four", "Five", "Six", "Seven",
		"Eight", "Nine",
	}
	digit10 := []string{
		"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen",
		"Nineteen",
	}
	digit20 := []string{
		"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety",
	}
	digit1000 := []string{
		"", "Thousand", "Million", "Billion",
	}
	encode := func(num int) string {
		result := []string{}
		if num >= 100 {
			index := num / 100
			result = append(result, digit1[index])
			result = append(result, "Hundred")
			// X Hundred
			num %= 100
		}
		if num >= 10 {
			if num < 20 {
				result = append(result, digit10[num%10])
			} else {
				result = append(result, digit20[num/10])
				result = append(result, digit1[num%10])
				// 23
				// twenty three
			}
		} else {
			result = append(result, digit1[num])
		}
		return strings.TrimSpace(strings.Join(result, " "))
	}
	thousandIndex := 0
	result := ""
	for num > 0 {
		d := num % 1000
		encodeDecimal := encode(d)
		if encodeDecimal != "" {
			result = " " + digit1000[thousandIndex] + " " + result
		}
		result = encodeDecimal + result

		num /= 1000
		thousandIndex++
	}
	return strings.TrimSpace(result)
}
