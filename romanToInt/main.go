package main

import (
	"fmt"
)

func main() {

	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("XV"))
	fmt.Println(romanToInt("LI"))
	fmt.Println(romanToInt("LX"))
	fmt.Println(romanToInt("XL"))
	fmt.Println(romanToInt("MCMXCIV"))

}

// solution1:
// case IV, I < V , 5-1
// case VI, 5+1
// solution2:

// index 0
// case VII
// try look up m[0:2], find nothing 
// try m[0:0+1], V
// case IVI 
// try m[0:2], IV
func romanToInt(s string) int {
	m := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}
	num := 0
	for index := 0; index < len(s); index++ {
		if index+1 < len(s) {
			if n, ok := m[s[index:index+1+1]]; ok {
				num += n
				index += 1
				continue
			}
		}
		num += m[s[index:index+1]]
	}
	return num
}
