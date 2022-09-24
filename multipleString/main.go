package main

import "fmt"

func main() {
	fmt.Println(multiply("123", "456"))
}

// 1 2
// 1 2
//     i+j,i+j+1
//      04(2*2)
//      02(2*1)
//     02 (1*2)
//    01  (1*1)
//    0144  result
func multiply(num1, num2 string) string {
	result := make([]byte, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			sum := result[i+j+1] + (num1[i]-'0')*(num2[j]-'0')
			result[i+j+1] = sum % 10
			// i+j > 10  也不怕
			// 下次循环变成i+j+1会处理
			// 最终后一次相乘的 i+j 不可能大于10 如 99 * 99
			//    9 * 9 
			//   81
			// i+j
			result[i+j] += sum / 10
		}
	}
	for i := range result {
		result[i] += '0'
	}

	for i := 0; i < len(result); i++ {
		if result[i] != '0' {
			return string(result[i:])
		}
	}
	return "0"

}
