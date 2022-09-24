package main

import "fmt"

func main() {

	a := 1 +'0' - '0'
	fmt.Println(a, int('0' + 1))
	fmt.Println("✅")
	pass := []string{
		"2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789",
	}
	for _, v := range pass {
		fmt.Println((isNumber(v)))
	}
	fmt.Println("❌")
	fail := []string{
		"abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53",
	}
	for _, v := range fail {
		fmt.Println((isNumber(v)))
	}
	fmt.Println("----others-----")
	// // "--6"
	fmt.Println("T", (isNumber("-0.1")))
	fmt.Println("F", (isNumber("1e")))
	fmt.Println("T",isNumber("2e10"))
	fmt.Println("F", isNumber("e10"))
	fmt.Println("True", isNumber("2.0e10"))
	fmt.Println("True",(isNumber("2.e10")))

}

func isNumber2(s string) bool {

	// sign     fistNumber   .   preE   e   ESign  lastNumber
	// 1          1     	 1     1    1      1    1

	var sign uint8 = 0b1000000
	firstNum := sign >> 1
	point := sign >> 2
	preE := sign >> 3
	EBit := sign >> 4
	ESign := sign >> 5
	lastNum := sign >> 6
	var n uint8 = 0

	for _, s := range s {
		switch true {
		case s == '+' || s == '-':
			{
				if n != 0 && (n&EBit == 0 || n&ESign == ESign || n&lastNum == lastNum) {
					return false
				}
				if n&EBit == EBit {
					n |= ESign
				} else {
					n |= sign
				}

			}
		case s >= '0' && s <= '9':

			if n&EBit == EBit {
				n |= lastNum
			} else if n&point == point {
				n |= preE
				n |= lastNum
			} else {
				n |= firstNum
				n |= lastNum
				n |= preE
			}

		case s == '.':
			if n&point == point || n&EBit == EBit {
				return false
			}
			n |= point
			n &= ^preE
			n &= ^lastNum
		case s == 'e' || s == 'E':
			{
				if n&EBit == EBit || n&preE == 0 && n&firstNum == 0 {
					return false
				}
				n |= EBit
				n &= ^lastNum
			}
		default:
			return false
		}

	}
	//  lastNum
	if n&EBit == EBit {
		return n&lastNum == lastNum

	}

	// firstNum || lastNum
	return firstNum&n == firstNum || n&lastNum == lastNum

}

//  0-9 number, afterENumber = true
// + at index 0 or which following E
// . 1 是否已经存在'.'  2 set afterENumber false
// e 1 是否已经存在 'e' 2 set afterENumber false

func isNumber(s string) bool {
	//  +-        0-9        .   0-9|''  e     +-     0-9
	// sign     fistNumber   .   preE    e    ESign   lastNumber
	// 1          1     	 1    1      1      1     1

	seenE := false
	seenPoint := false
	seenNumber := false
	seenNumberAfterE := false

	for i := range s {
		if s[i] >= '0' && s[i] <= '9' {
			seenNumber = true
			seenNumberAfterE = true
		} else if s[i] == '+' || s[i] == '-' {
			if !(i == 0 || s[i-1] == 'e') {
				return false
			}
		} else if s[i] == '.' {
			if seenPoint || seenE {
				return false
			}
			seenNumberAfterE = false
		} else if s[i] == 'e' {
			if seenE || !seenNumber {
				return false
			}
			seenE = true
			seenNumberAfterE = false
		} else {
			return false
		}

	}

	return seenNumber && seenNumberAfterE

}
