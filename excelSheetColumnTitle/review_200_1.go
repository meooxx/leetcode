package main

func convertToTitle(columnNumber int) string {
	result := ""
	for columnNumber != 0 {
		b := byte((columnNumber-1)%26 + 'A')
		columnNumber = (columnNumber - 1) / 26
		result = string(b) + result
	}
	return result
}
