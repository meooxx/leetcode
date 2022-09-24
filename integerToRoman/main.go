package main

import (
	"fmt"
)

func main() {

	n := 1994
	fmt.Println(intToRoman(n))
	fmt.Println(intToRoman(2))
	fmt.Println(intToRoman(4))
	fmt.Println(intToRoman(5))
	fmt.Println(intToRoman(7))
	fmt.Println(intToRoman(10))
	fmt.Println(intToRoman(20))
	fmt.Println(intToRoman(40))
	fmt.Println(intToRoman(50))
	fmt.Println(intToRoman(70))
	fmt.Println(intToRoman(100))
}

func intToRoman(num int) string {
	m := map[int]string{
		1:    "I",
		2:    "II",
		3:    "III",
		4:    "IV",
		5:    "V",
		6:    "VI",
		7:    "VII",
		8:    "VIII",
		9:    "IX",
		10:   "X",
		20:   "XX",
		30:   "XXX",
		40:   "XL",
		50:   "L",
		60:   "LX",
		70:   "LXX",
		80:   "LXXX",
		90:   "XC",
		100:  "C",
		200:  "CC",
		300:  "CCC",
		400:  "CD",
		500:  "D",
		600:  "DC",
		700:  "DCC",
		800:  "DCCC",
		900:  "CM",
		1000: "M",
		2000: "MM",
		3000: "MMM",
	}
	s := ""
	base := 1
	for num > 0 {
		key := num % 10
		v := m[key*base]
		s = v + s
		num = num / 10
		base = base * 10
	}
	return s

}
