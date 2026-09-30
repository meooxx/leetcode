func isNumber(s string) bool {
	//  +-        0-9        .   0-9|''  e     +-     0-9
	// sign     fistNumber   .   preE    e    ESign   lastNumber
	// 1          1     	 1    1      1      1     1

	seenE := false
	seenPoint := false
	seenNumber := false
	seenNumberAfterE := false
    s = strings.ToLower(s)
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
			seenPoint = true
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
