package main

func partition(s string) [][]string {
	result := [][]string{}
	if s == "" {
		return result
	}
	mpCache := map[string]bool{}
	isPalindrome := func(a string) bool {
		if len(a) <= 1 {
			return true
		}
		if v, ok := mpCache[a]; ok {
			return v
		}
		for left, right := 0, len(a)-1; left < right; {
			if a[left] != a[right] {
				mpCache[a] = false
				return false
			}
			left++
			right--
		}
		mpCache[a] = true
		return true
	}

	var subPartition func(s string, start, end int, current []string)
	subPartition = func(s string, start, end int, current []string) {
		if start > end {
			item := make([]string, len(current))
			copy(item, current)
			result = append(result, item)
			return
		}
		for i := start; i <= end; i++ {
			if !isPalindrome(s[start : i+1]) {
				continue
			}
			newCurrent := append(current, s[start:i+1])
			subPartition(s, i+1, end, newCurrent)
		}
	}
	subPartition(s, 0, len(s)-1, []string{})
	return result
}
