package main

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
