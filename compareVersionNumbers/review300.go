package main

func compareVersion0(version1 string, version2 string) int {
	l1, l2 := 0, 0

	for l1 < len(version1) || l2 < len(version2) {
		num1 := 0
		for l1 < len(version1) && version1[l1] != '.' {
			num1 = num1*10 + int(version1[l1]-'0')
			l1++
		}
		num2 := 0
		for l2 < len(version2) && version2[l2] != '.' {
			num2 = num2*10 + int(version2[l2]-'0')
			l2++
		}
		if num1 > num2 {
			return 1
		} else if num1 < num2 {
			return -1
		}

	}
	return 0

}
func compareVersion(version1 string, version2 string) int {
	l1, l2 := 0, 0

	for l1 < len(version1) || l2 < len(version2) {
		num1 := 0
		for l1 < len(version1) && version1[l1] != '.' {
			num1 = num1*10 + int(version1[l1]-'0')
			l1++
		}
		num2 := 0
		for l2 < len(version2) && version2[l2] != '.' {
			num2 = num2*10 + int(version2[l2]-'0')
			l2++
		}
		if num1 > num2 {
			return 1
		} else if num1 < num2 {
			return -1
		}
		l1++
		l2++

	}
	return 0

}
