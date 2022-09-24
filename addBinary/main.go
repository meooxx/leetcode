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

func addBinary(a string, b string) string {

	sizeA := len(a)
	sizeB := len(b)
	result := make([]byte, sizeA+sizeB)
	index := sizeA + sizeB - 1
	rest := 0
	sizeA--
	sizeB--
	for sizeA >= 0 && sizeB >= 0 {
		sum := int(a[sizeA]-'0'+b[sizeB]-'0') + rest
		rest = sum / 2
		result[index] = byte(sum%2 + '0')
		index--
		sizeA--
		sizeB--
	}

	for sizeA >= 0 {
		result[index] = a[sizeA]
		index--
		sizeA--
	}

	for sizeB >= 0 {
		result[index] = b[sizeB]
		index--
		sizeB--
	}

	if rest == 1 {
		return "1" + string(result)
	}
	return string(result)

}
