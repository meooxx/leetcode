package main

import "fmt"

func main() {
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("1010", "1011"))
}

func addBinary2(a string, b string) string {

	l1 := len(a) - 1
	l2 := len(b) - 1

	carr := 0
	sum := 0
	r := []byte{}
	for l1 >= 0 || l2 >= 0 {
		sum = carr
		if l1 >= 0 {
			sum += int(a[l1] - '0')
			l1--
		}
		if l2 >= 0 {
			sum += int(b[l2] - '0')
			l2--
		}
		r = append(r, byte(sum%2+'0'))
		carr = sum / 2

	}
	if carr == 1 {
		r = append(r, byte('1'))
	}
	for i, j := 0, len(r)-1; i < j; {
		r[i], r[j] = r[j], r[i]
		i++
		j--
	}
	return string(r)
}

func addBinary(a, b string) string {
	lenA := len(a)
	lenB := len(b)

	answer := make([]byte, lenA+lenB)
	reminder := 0
	for lenA >= 0 || lenB >= 0 {
		sum := reminder
		index := lenA + lenB - 1
		if lenA >= 0 {
			sum += int(a[lenA] - '0')
			lenA--
		}
		if lenB >= 0 {
			sum += int(b[lenB] - '0')
			lenB--
		}
		answer[index] = byte(sum%2 + '0')
		reminder = sum / 2
	}
	if reminder > 0 {
		answer[lenA+lenB-1] = '1'
	}
	return string(answer)

}
