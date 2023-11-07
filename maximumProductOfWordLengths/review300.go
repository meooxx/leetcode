package main

func maxProduct(words []string) int {
	bits := make([]int, len(words))

	for i := range words {
		for j := range words[i] {
			bits[i] |= 1 << int(words[i][j]-'a')
		}
	}
	maxProduct := 0
	for i := range words {
		for j := i + 1; j < len(words); j++ {
			if bits[i]&bits[j] == 0 {
				if len(words[i])*len(words[j]) > maxProduct {
					maxProduct = len(words[i]) * len(words[j])
				}
			}

		}
	}
	return maxProduct
}
