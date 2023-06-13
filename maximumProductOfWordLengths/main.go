package main

func maxProduct(words []string) int {
	val := make([]int, len(words))

	for i := range words {
		for j := range words[i] {
			val[i] |= 1 << int(words[i][j]-'a')
		}
	}

	maxVal := 0
	for i := range words {
		for j := i + 1; j < len(words); j++ {
			if val[i]&val[j] == 0 {
				product := len(words[i]) * len(words[j])
				if product > maxVal {
					maxVal = product
				}
			}
		}
	}
	return maxVal

}
